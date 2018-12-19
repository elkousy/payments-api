package errors

import (
	"context"
	"net/http"

	logger "github.com/elkousy/payments-api/utility/logger"
	kithttp "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"
)

// LoggingErrorEncoder wraps GoKit's DefaultErrorEncoder to provide, on top of it, logging into stderr
func LoggingErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	logger.LogStdErr.Error("err", zap.Error(err),
		zap.Any("http.url", ctx.Value(kithttp.ContextKeyRequestURI)),
		zap.Any("http.path", ctx.Value(kithttp.ContextKeyRequestPath)),
		zap.Any("http.method", ctx.Value(kithttp.ContextKeyRequestMethod)),
		zap.Any("http.user_agent", ctx.Value(kithttp.ContextKeyRequestUserAgent)),
	)
	kithttp.DefaultErrorEncoder(ctx, err, w)
}
