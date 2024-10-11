package metrics

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/Conflux-Chain/go-conflux-util/http/middlewares"
	"github.com/Conflux-Chain/go-conflux-util/metrics"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type ResponseWriter struct {
	writer http.ResponseWriter
	code   int
	// cache response data
	buf *bytes.Buffer
}

func newResponseWriter(writer http.ResponseWriter, code int) *ResponseWriter {
	var buf bytes.Buffer
	return &ResponseWriter{
		writer: writer,
		code:   code,
		buf:    &buf,
	}
}

func (w *ResponseWriter) Write(bs []byte) (int, error) {
	w.buf.Write(bs)
	return w.writer.Write(bs)
}

func (w *ResponseWriter) Header() http.Header {
	return w.writer.Header()
}

func (w *ResponseWriter) WriteHeader(code int) {
	w.code = code
	w.writer.WriteHeader(code)
}

func Metrics() middlewares.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			wrappedW := newResponseWriter(w, http.StatusOK)
			next.ServeHTTP(wrappedW, r)

			if len(wrappedW.buf.Bytes()) > 0 {
				urlType, ok := r.Context().Value(CtxKeyURLType).(string)
				if !ok {
					return
				}

				var resp api.BusinessError
				err := json.Unmarshal(wrappedW.buf.Bytes(), &resp)
				if err != nil {
					return
				}

				UpdateDuration(urlType, wrappedW.code, resp.Code, start)
				logrus.WithFields(logrus.Fields{
					"Path":         r.URL.Path,
					"urlType":      urlType,
					"http status":  wrappedW.code,
					"resp code":    resp.Code,
					"resp message": resp.Message,
				}).Debug("report API metrics")
			}
		})
	}
}

func UpdateDuration(url string, status, code int, start time.Time) {
	var isSuccess, isInternalError bool
	if isSuccess = status == http.StatusOK && code == api.ErrCodeSuccess; !isSuccess {
		isInternalError = status == api.ErrCodeInternal
	}

	// Overall rate statistics
	metrics.GetOrRegisterTimeWindowPercentageDefault("openapi/rate/success").Mark(isSuccess)
	metrics.GetOrRegisterTimeWindowPercentageDefault("openapi/rate/internalErr").Mark(isInternalError)
	metrics.GetOrRegisterTimeWindowPercentageDefault("openapi/rate/nonInternalErr").Mark(!isSuccess && !isInternalError)

	// API rate statistics
	path := url[5:]
	metrics.GetOrRegisterTimeWindowPercentageDefault("openapi/rate/success/%v", path).Mark(isSuccess)
	metrics.GetOrRegisterTimeWindowPercentageDefault("openapi/rate/internalErr/%v", path).Mark(isInternalError)
	metrics.GetOrRegisterTimeWindowPercentageDefault("openapi/rate/nonInternalErr/%v", path).Mark(!isSuccess && !isInternalError)

	// Update QPS & Latency
	metrics.GetOrRegisterTimer("openapi/duration/all").UpdateSince(start)
	metrics.GetOrRegisterTimer("openapi/duration/%v", path).UpdateSince(start)
}

var CtxKeyURLType = middlewares.CtxKey("X-URL-TYPE")

func URLType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path

		urlType, _ := ReplaceStrByRegex(url, "/accounts/[a-zA-Z0-9]+", "/accounts/address")
		urlType, _ = ReplaceStrByRegex(urlType, "/txs/[a-zA-Z0-9]+", "/txs/detail")
		if i := strings.Index(urlType, "/txs/detail"); i != -1 {
			urlType = urlType[0:(i + len("/txs/detail"))]
		}

		ctx := context.WithValue(r.Context(), CtxKeyURLType, urlType)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ReplaceStrByRegex(str, rule, replace string) (string, error) {
	reg, err := regexp.Compile(rule)
	if reg == nil || err != nil {
		return "", errors.WithMessagef(err, "Compile regex rule %v", rule)
	}
	return reg.ReplaceAllString(str, replace), nil
}
