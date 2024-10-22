package main

import (
	"context"
	"fmt"
	"log/slog"
	"logging_sample/loggers"
	"net/http"
	"os"
)

func requestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "host", r.Host)
		ctx = context.WithValue(ctx, "method", r.Method)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func fearHandler(w http.ResponseWriter, r *http.Request) {

	traceId := "1234567890"
	spanId := "9999999999"

	ctx := context.WithValue(r.Context(), "trace_id", traceId)
	ctx = context.WithValue(ctx, "span_id", spanId)
	r = r.WithContext(ctx)

	// default 설정을 해놨기 때문에 slog 로 호출해도 contextHandler 가 젹용된 채 나옴
	slog.InfoContext(r.Context(), "get / 로그 테스트")

	host := r.Context().Value("host").(string) // 컨텍스트에서 값 추출
	method := r.Context().Value("method").(string)
	w.Write([]byte(fmt.Sprintf("ContextHandler 테스트 중 ( host: %s, method : %s )", host, method)))
}

func main() {

	// server on 용 slog
	simpleJsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	mainLogger := slog.New(simpleJsonHandler)
	mainLogger.Info("server on ..")

	// default slog
	opts := slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: loggers.ReplaceOption,
	}

	jsonHandler := slog.NewJSONHandler(os.Stdout, &opts)
	ctxHandler := loggers.ContextHandler{
		Handler: jsonHandler,
	}
	logger := slog.New(&ctxHandler)
	slog.SetDefault(logger) // default 설정

	http.HandleFunc("/fear", func(w http.ResponseWriter, r *http.Request) {
		requestMiddleware(http.HandlerFunc(fearHandler)).ServeHTTP(w, r)
	})

	http.ListenAndServe(":7000", nil)
}
