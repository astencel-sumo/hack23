package piiredactionprocessor

import (
	"context"
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

func (p *piiRedactionProcessor) processLogs(ctx context.Context, batch plog.Logs) (plog.Logs, error) {
	resourcesLogs := batch.ResourceLogs()
	for i := 0; i < resourcesLogs.Len(); i++ {
		resource := resourcesLogs.At(i)
		scopeLogs := resource.ScopeLogs()
		for j := 0; j < scopeLogs.Len(); j++ {
			scope := scopeLogs.At(j)
			logRecords := scope.LogRecords()
			for k := 0; k < logRecords.Len(); k++ {
				record := logRecords.At(k)
				redactPii(ctx, record)
			}
		}
	}
	return batch, nil
}

func redactPii(ctx context.Context, record plog.LogRecord) {
	record.Body().SetStr(strings.Repeat("*", len(record.Body().AsString())))
}
