package pipelines

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewHTTPMetricsExporter(t *testing.T) {
	const endpoint = "https://otlp-gateway-prod-us-east-0.grafana.net/otlp"
	exp, err := newHTTPMetricsExporter(endpoint, false, map[string]string{})
	require.NoError(t, err)
	require.NotNil(t, exp)

	t.Cleanup(func() {
		require.NoError(t, exp.Shutdown(context.Background()))
	})
}
