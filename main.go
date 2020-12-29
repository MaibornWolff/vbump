package main

import (
	"log"
	"net/http"
	"time"

	"github.com/maibornwolff/vbump/adapter"
	"github.com/prometheus/client_golang/prometheus"
	logrus "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	numberOfDeployments = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "vbump_deployments_total",
			Help: "Number of deployments tracked by vbump",
		},
		[]string{"project"},
	)
)

func init() {
	prometheus.MustRegister(numberOfDeployments)
}

func main() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	w := logger.Writer()
	defer w.Close()

	log.SetOutput(w)

	listenAddr := kingpin.Flag("listen", "Address to listen on.").Short('l').Default(":8080").String()
	datadir := kingpin.Flag("datadir", "Directory path for storing version files (must exist).").Short('d').Required().String()

	kingpin.Parse()
	logger.Info("Server is starting...")

	fileProvider := adapter.New(*datadir)
	version := NewVersion(fileProvider)
	handler := NewHandler(version, logger)
	router := handler.GetRouter()

	server := &http.Server{
		Addr:         *listenAddr,
		Handler:      router,
		ErrorLog:     log.New(w, "", 0),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	logger.Infof("Server is ready to handle requests at %v", *listenAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %v: %v\n", *listenAddr, err)
	}
}
