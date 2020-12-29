package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

//Handler for handling http routes
type Handler struct {
	version *Version
	logger  *log.Logger
}

//NewHandler constructs a new handler
func NewHandler(version *Version, logger *log.Logger) *Handler {
	if logger == nil {
		logger = log.New()
	}

	return &Handler{
		version: version,
		logger:  logger,
	}
}

//GetRouter configure all routes
func (handler *Handler) GetRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/major/{project}", handler.OnMajor).Methods("POST")
	r.HandleFunc("/minor/{project}", handler.OnMinor).Methods("POST")
	r.HandleFunc("/patch/{project}", handler.OnPatch).Methods("POST")
	r.HandleFunc("/minor/transient/{version}", handler.OnTransientMinor).Methods("POST")
	r.HandleFunc("/patch/transient/{version}", handler.OnTransientPatch).Methods("POST")
	r.HandleFunc("/version/{project}/{version}", handler.OnSetVersion).Methods("POST")
	r.HandleFunc("/version/{project}", handler.OnGetVersion).Methods("GET")
	r.HandleFunc("/", handler.OnHealth)
	r.Handle("/metrics", promhttp.Handler())

	return r
}

//OnHealth is a handler for a health check
func (handler *Handler) OnHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello from vbump!"))
}

//OnMajor is a handler for bumping the major part for a given project
func (handler *Handler) OnMajor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numberOfDeployments.With(prometheus.Labels{"project": vars["project"], "element": "major"}).Inc()
	version, err := handler.version.BumpMajor(vars["project"])
	if err != nil {
		handler.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	handler.logger.Infof("bump major version to %v on project %v", version, vars["project"])
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(version))
}

//OnMinor is a handler for bumping the minor part for a given project
func (handler *Handler) OnMinor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numberOfDeployments.With(prometheus.Labels{"project": vars["project"], "element": "minor"}).Inc()
	version, err := handler.version.BumpMinor(vars["project"])
	if err != nil {
		handler.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	handler.logger.Infof("bump minor version to %v on project %v", version, vars["project"])
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(version))
}

//OnPatch is a handler for bump the patch part for a given project
func (handler *Handler) OnPatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numberOfDeployments.With(prometheus.Labels{"project": vars["project"], "element": "patch"}).Inc()
	version, err := handler.version.BumpPatch(vars["project"])
	if err != nil {
		handler.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	handler.logger.Infof("bump patch version to %v on project %v", version, vars["project"])
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(version))
}

//OnSetVersion is a handler for setting the version for a given project
func (handler *Handler) OnSetVersion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	version, err := handler.version.SetVersion(vars["project"], vars["version"])
	if err != nil {
		handler.logger.Error(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}
	handler.logger.Infof("set version explicitly to %v on project %v", version, vars["project"])

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(version))
}

//OnGetVersion is a handler for getting the version for a given project
func (handler *Handler) OnGetVersion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	version, err := handler.version.GetVersion(vars["project"])
	if err != nil {
		handler.logger.Error(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	handler.logger.Infof("get version from project %v", vars["project"])

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(version))
}

//OnTransientPatch is a handler for a transient patch bumps
func (handler *Handler) OnTransientPatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	version, err := handler.version.BumpTransientPatch(vars["version"])
	if err != nil {
		handler.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	handler.logger.Infof("bump transient patch version to %v", version)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(version))
}

//OnTransientMinor is a handler for a transient minor bumps
func (handler *Handler) OnTransientMinor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	version, err := handler.version.BumpTransientMinor(vars["version"])
	if err != nil {
		handler.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	handler.logger.Infof("bump transient minor version to %v", version)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(version))
}
