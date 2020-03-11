package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"github.com/SamuelWillis/go_snake/api"
)

func main() {
	http.HandleFunc("/", api.Index)
	http.HandleFunc("/start", api.Start)
	http.HandleFunc("/move", api.Move)
	http.HandleFunc("/end", api.End)
	http.HandleFunc("/ping", api.Ping)

	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
	}

	// Add filename into logging messages
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("Running server on port %s...\n", port)
	http.ListenAndServe(":"+port, LoggingHandler(http.DefaultServeMux))
}

const LogFormat = `"%s %s %s %d %d" %f`

func LoggingHandler(next http.Handler) http.Handler {
	var logger = log.New(os.Stderr, "", log.LstdFlags)

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		loggingWriter := &LoggingResponseWriter{res, http.StatusOK}

		startTime := time.Now()
		next.ServeHTTP(loggingWriter, req)
		elapsedTime := time.Now().Sub(startTime)

		logger.Printf(
			LogFormat,
			req.Method, req.RequestURI, req.Proto,
			loggingWriter.statusCode, 0, elapsedTime.Seconds(),
		)

	})
}

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (res *LoggingResponseWriter) WriteHeader(code int) {
	res.statusCode = code
	res.ResponseWriter.WriteHeader(code)
}
