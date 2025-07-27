package logger

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func LoggerWithTrace(ctx context.Context) *zap.Logger {
	logger := zap.L()
	span := trace.SpanFromContext(ctx)
	sc := span.SpanContext()
	if sc.IsValid() {
		logger = logger.With(
			zap.String("trace_id", sc.TraceID().String()),
			zap.String("span_id", sc.SpanID().String()),
		)
	}
	return logger
}

type loggerKey struct{}

func WithRequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := LoggerWithTrace(ctx)
		ctx = context.WithValue(ctx, loggerKey{}, logger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Logger(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(loggerKey{}).(*zap.Logger)
	if !ok {
		logger = zap.L()
		logger.Warn("logger not found in context, defaulting to global logger")
	}
	return logger
}

func SetLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func AddAttributes(ctx context.Context, attributes map[string]interface{}) context.Context {
	logger := Logger(ctx)
	return SetLogger(ctx, logger.With(zap.Any("attributes", attributes)))
}
