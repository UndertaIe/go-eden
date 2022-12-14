package cache

import (
	"time"

	"github.com/UndertaIe/go-eden/logger"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var cacheTracerName = "github.com/UndertaIe/go-eden/cache/otel"
var deleteSpanName = "cache.delete"
var cacheSpanName = "cache.decorator"

func DeleteCacheWithTracing(c *gin.Context, cacheName ...string) error {
	_, span := otel.GetTracerProvider().Tracer(cacheTracerName).Start(c.Request.Context(), deleteSpanName, trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()
	return DeleteCache(c)
}

func CachePageWithTracing(store Cache, expire time.Duration, log logger.Log, handle gin.HandlerFunc) gin.HandlerFunc {
	cb := CachePage(store, expire, log, handle)
	return func(c *gin.Context) {
		_, span := otel.GetTracerProvider().Tracer(cacheTracerName).Start(c.Request.Context(), cacheSpanName, trace.WithSpanKind(trace.SpanKindClient))
		defer span.End()

		cb(c)
	}
}

func DeleteCachePageWithTracing(handle gin.HandlerFunc, cacheName ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		handle(c)
		_, span := otel.GetTracerProvider().Tracer(cacheTracerName).Start(c.Request.Context(), deleteSpanName, trace.WithSpanKind(trace.SpanKindClient))
		defer span.End()
		DeleteCache(c, cacheName...)
	}
}
