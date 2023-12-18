package piiredactionprocessor

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

type piiRedactionProcessor struct {
	config *Config
}

func newPiiRedactionProcessor(ctx context.Context, config *Config, logger *zap.Logger) (*piiRedactionProcessor, error) {
	return &piiRedactionProcessor{
		config: config,
	}, nil
}

func (processor *piiRedactionProcessor) processLogs(ctx context.Context, batch plog.Logs) (plog.Logs, error) {
	resourcesLogs := batch.ResourceLogs()
	for i := 0; i < resourcesLogs.Len(); i++ {
		resource := resourcesLogs.At(i)
		scopeLogs := resource.ScopeLogs()
		for j := 0; j < scopeLogs.Len(); j++ {
			scope := scopeLogs.At(j)
			logRecords := scope.LogRecords()
			for k := 0; k < logRecords.Len(); k++ {
				record := logRecords.At(k)
				processor.redactPii(ctx, record)
			}
		}
	}
	return batch, nil
}

func (processor *piiRedactionProcessor) redactPii(ctx context.Context, record plog.LogRecord) {
	// record.Body().SetStr(strings.Repeat("*", len(record.Body().AsString())))
	inputString := record.Body().AsString()
	requestData, err := json.Marshal(inputString)
	if err != nil {
		panic(err)
	}
	response, err := http.Post(processor.config.Endpoint, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var entities []Entity
	err = json.Unmarshal(responseBody, &entities)
	if err != nil {
		panic(err)
	}
	for _, entity := range entities {
		mask := strings.Repeat("*", len(inputString[entity.Start:entity.End]))
		inputString = inputString[:entity.Start] + mask + inputString[entity.End:]
	}
	record.Body().SetStr(inputString)
}
