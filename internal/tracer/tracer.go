package tracer

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"io"
	"ova-conference-api/internal/configs"
)

func InitTracer(serviceName string, endpoint *configs.JaegerConfiguration) io.Closer {
	cfg := jaegerConfig.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: fmt.Sprintf("%s:%v", endpoint.Host, endpoint.Port),
		},
	}

	jLogger := jaegerLog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jLogger), jaegerConfig.Metrics(jMetricsFactory))
	if err != nil {
		log.Error().Err(err).Msg("Could not initialize jaeger tracer -> NewTracer")
		return nil
	}
	opentracing.SetGlobalTracer(tracer)
	return closer
}
