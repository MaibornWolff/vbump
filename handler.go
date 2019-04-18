package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	version *Version
	logger  *log.Logger
}

func NewHandler(version *Version, logger *log.Logger) *Handler {
	if logger == nil {
		logger = log.New()
	}

	return &Handler{
		version: version,
		logger:  logger,
	}
}

func (handler *Handler) GetRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/major/{project}", handler.OnMajor).Methods("POST")
	r.HandleFunc("/minor/{project}", handler.OnMinor).Methods("POST")
	r.HandleFunc("/patch/{project}", handler.OnPatch).Methods("POST")
	r.HandleFunc("/version/{project}/{version}", handler.OnSetVersion).Methods("POST")
	r.HandleFunc("/version/{project}", handler.OnGetVersion).Methods("GET")
	r.HandleFunc("/", handler.OnHealth)

	return r
}

func (handler *Handler) OnHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello from vbump!"))
}

func (handler *Handler) OnMajor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
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

func (handler *Handler) OnMinor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
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

func (handler *Handler) OnPatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
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
