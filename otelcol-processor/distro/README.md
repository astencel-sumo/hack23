# Otelcol distro containing the PII redaction processor

- Install Otelcol Builder: https://github.com/open-telemetry/opentelemetry-collector/tree/main/cmd/builder#installation
- Build the otelcol-pii distro with `opentelemetry-collector-builder --config=./manifest.yaml`
- Run the collector and send data to Sumo Logic: `SUMOLOGIC_INSTALLATION_TOKEN=<your-installation-token> ./out/otelcol-pii --config=./config-sumologic.yaml`
- To run the collector locally and display data in the collector's console output without sending to Sumo Logic: `./out/otelcol-pii --config=./config-console.yaml`
