package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
	"gopkg.in/alecthomas/kingpin.v2"
	"maibornwolff/vbump/adapter"
	"maibornwolff/vbump/service"
)

var (
	numberOfBumps = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "vbump_bumps_total",
			Help: "Number of bumps tracked by vbump, labelled with project name and semVer element",
		},
		[]string{"project", "element"},
	)
)

func init() {
	prometheus.MustRegister(numberOfBumps)
}

func main() {
	listenAddr := kingpin.Flag("listen", "Address to listen on.").Short('l').Default(":8080").String()
	dataDir := kingpin.Flag("datadir", "Directory path for storing version files (must exist).").Short('d').Required().String()
	kingpin.Parse()

	log.Info().Msg("Server is starting...")

	fileProvider := adapter.NewFileProvider(*dataDir)
	versionManager := service.NewVersionManager(fileProvider)
	handler := NewHandler(versionManager)
	router := handler.GetRouter()

	server := &http.Server{
		Addr:         *listenAddr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Info().Str("listenAddr", *listenAddr).Msg("Server is ready to handle requests")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Str("listenAddr", *listenAddr).Err(err).Msg("Failed to listen")
	}
}
