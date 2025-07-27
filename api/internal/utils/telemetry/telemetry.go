package telemetry

import (
	"context"

	"github.com/ajohnston1219/eatme/api/internal/utils/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.uber.org/zap"
)

func InitTracer(ctx context.Context, serviceName string) func(context.Context) error {
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithEndpoint("jaeger:4317"), otlptracegrpc.WithInsecure())
	if err != nil {
		logger.Logger(ctx).Error("failed to create OTLP trace exporter", zap.Error(err))
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		logger.Logger(ctx).Error("failed to create OTEL resource", zap.Error(err))
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)

	return tp.Shutdown
}
