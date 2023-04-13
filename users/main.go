package main

import (
    "users/api"
    "net/http"
    "os"

    "github.com/go-kit/kit/log"
)

// main is the uses entry point
func main() {

    var logger log.Logger
    logger = log.NewLogfmtLogger(os.Stderr)
    logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", "8081", "caller", log.DefaultCaller)

    // Start users api server
    r := api.NewHttpServer(api.NewService(), logger)
    logger.Log("msg", "HTTP", "addr", "8081")
    logger.Log("err", http.ListenAndServe(":8081", r))
}
