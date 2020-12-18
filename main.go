package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/stavangler/venue/venue"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"
	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	NATS_SERVER string `envconfig:"NATS_SERVER" default:"http://127.0.0.1:4222" required:"true"`
	Port        string `envconfig:"PORT" default:"8080" required:"true"`
}

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	}
	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		logger.Log("msg", "failed to parse environment variable", "err", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	router := mux.NewRouter()

	var svc venue.Service
	{
		svc = venue.NewService(venue.NewStore())
		svc = venue.NewLoggingService(log.With(logger, "component", "venue"), svc)
	}

	router.PathPrefix("/").Handler(venue.MakeHTTPHandler(svc, log.With(logger, "component", "http")))

	router.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		//TODO add some start up flags etc, return 200 OK in the beginning
		w.WriteHeader(http.StatusOK)
	})

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", fmt.Sprintf(":%s", env.Port), "msg", "listening")
		errs <- http.ListenAndServe(fmt.Sprintf(":%s", env.Port), router)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
		cancel()
		logger.Log("terminated", <-errs)
	}()

	<-ctx.Done()
}
