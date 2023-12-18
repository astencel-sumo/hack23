package piiredactionprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

func NewFactory() processor.Factory {
	return processor.NewFactory(
		"pii_redaction",
		createDefaultConfig,
		processor.WithLogs(createLogsProcessor, component.StabilityLevelDevelopment),
	)
}

func createLogsProcessor(
	ctx context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	nextLogsConsumer consumer.Logs,
) (processor.Logs, error) {
	processor, err := newPiiRedactionProcessor(ctx, cfg.(*Config), set.Logger)
	if err != nil {
		return nil, err
	}
	return processorhelper.NewLogsProcessor(
		ctx,
		set,
		cfg,
		nextLogsConsumer,
		processor.processLogs,
		processorhelper.WithCapabilities(consumer.Capabilities{MutatesData: true}))
}
