package piiredactionprocessor

import (
	"context"

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
	return batch, nil
}
