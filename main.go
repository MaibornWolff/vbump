package main

import (
	"log"
	"maibornwolff/vbump/service"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"maibornwolff/vbump/adapter"
)

var (
	numberOfBumps = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "vbump_bumps_total",
			Help: "Number of bumps tracked by vbump, labelled with projectname and semVer element",
		},
		[]string{"project", "element"},
	)
)

func init() {
	prometheus.MustRegister(numberOfBumps)
}

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	loggerWriter := logger.Writer()
	log.SetOutput(loggerWriter)
	defer loggerWriter.Close()

	listenAddr := kingpin.Flag("listen", "Address to listen on.").Short('l').Default(":8080").String()
	dataDir := kingpin.Flag("datadir", "Directory path for storing versionManager files (must exist).").Short('d').Required().String()
	kingpin.Parse()

	logger.Info("Server is starting...")

	fileProvider := adapter.NewFileProvider(*dataDir)
	versionManager := service.NewVersionManager(fileProvider)
	handler := NewHandler(versionManager, logger)
	router := handler.GetRouter()

	server := &http.Server{
		Addr:         *listenAddr,
		Handler:      router,
		ErrorLog:     log.New(loggerWriter, "", 0),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	logger.Infof("Server is ready to handle requests at %v", *listenAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %v: %v\n", *listenAddr, err)
	}
}
