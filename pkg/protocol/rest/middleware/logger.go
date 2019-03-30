package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sunkuet02/go-grpc-rest/pkg/logger"
)

func AddLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		requestId := GetReqID(ctx)

		logger.Log.Info(fmt.Sprintf("Request started for request id : %s, method: %s, remote-address: %s",
			requestId, r.Method, r.RemoteAddr))

		startedTime := time.Now()

		h.ServeHTTP(w, r)

		logger.Log.Info(fmt.Sprintf("Request finished for request id : %s, elapsed : %f seconds",
			requestId, float64(time.Since(startedTime).Seconds())))
	})
}
