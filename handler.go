package main

import (
	"maibornwolff/vbump/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

// Handler for handling http routes
type Handler struct {
	versionManager *service.VersionManager
	logger         *log.Logger
}

// NewHandler constructs a new handler
func NewHandler(version *service.VersionManager, logger *log.Logger) *Handler {
	if logger == nil {
		logger = log.New()
	}

	return &Handler{
		versionManager: version,
		logger:         logger,
	}
}

// LoggerMiddleware logs the last error
func (handler *Handler) LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err != nil {
			handler.logger.Error(err)
		}
	}
}

// GetRouter configures all routes
func (handler *Handler) GetRouter() *gin.Engine {
	r := gin.New()
	r.Use(handler.LoggerMiddleware())
	gin.SetMode(gin.ReleaseMode)

	r.POST("/major/:project", handler.OnMajor)
	r.POST("/minor/:project", handler.OnMinor)
	r.POST("/patch/:project", handler.OnPatch)
	r.POST("/transient/minor/:versionManager", handler.OnTransientMinor)
	r.POST("/transient/patch/:versionManager", handler.OnTransientPatch)
	r.POST("/versionManager/:project/:versionManager", handler.OnSetVersion)
	r.GET("/versionManager/:project", handler.OnGetVersion)
	r.GET("/", handler.OnHealth)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return r
}

// OnHealth is a handler for a health check
func (handler *Handler) OnHealth(context *gin.Context) {
	context.String(http.StatusOK, "hello from vbump!")
}

// OnMajor is a handler for bumping the major part for a given project
func (handler *Handler) OnMajor(context *gin.Context) {
	project := context.Param("project")
	version, err := handler.versionManager.BumpMajor(project)
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	numberOfBumps.With(prometheus.Labels{"project": project, "element": "major"}).Inc()
	handler.logger.Infof("bump major versionManager to %v on project %v", version, project)
	context.String(http.StatusOK, "%s", version.String())
}

// OnMinor is a handler for bumping the minor part for a given project
func (handler *Handler) OnMinor(context *gin.Context) {
	project := context.Param("project")
	version, err := handler.versionManager.BumpMinor(project)
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	numberOfBumps.With(prometheus.Labels{"project": project, "element": "minor"}).Inc()
	handler.logger.Infof("bump minor versionManager to %v on project %v", version, project)
	context.String(http.StatusOK, "%s", version.String())
}

// OnPatch is a handler for bumping the patch part for a given project
func (handler *Handler) OnPatch(context *gin.Context) {
	project := context.Param("project")
	version, err := handler.versionManager.BumpPatch(project)
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	numberOfBumps.With(prometheus.Labels{"project": project, "element": "patch"}).Inc()
	handler.logger.Infof("bump patch versionManager to %v on project %v", version, project)
	context.String(http.StatusOK, "%s", version.String())
}

// OnSetVersion is a handler for setting the versionManager for a given project
func (handler *Handler) OnSetVersion(context *gin.Context) {
	project := context.Param("project")
	version := context.Param("versionManager")
	_, err := handler.versionManager.SetVersion(project, version)
	if err != nil {
		_ = context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	handler.logger.Infof("set versionManager explicitly to %v on project %v", version, project)
	context.String(http.StatusOK, "%s", version)
}

// OnGetVersion is a handler for getting the versionManager for a given project
func (handler *Handler) OnGetVersion(context *gin.Context) {
	project := context.Param("project")
	version, err := handler.versionManager.GetVersion(project)
	if err != nil {
		_ = context.AbortWithError(http.StatusNotFound, err)
		return
	}

	handler.logger.Infof("get versionManager from project %v", project)
	context.String(http.StatusOK, "%s", version.String())
}

// OnTransientPatch is a handler for a transient patch bump
func (handler *Handler) OnTransientPatch(context *gin.Context) {
	version := context.Param("versionManager")
	bumpedVersion, err := handler.versionManager.BumpTransientPatch(version)
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	handler.logger.Infof("bump transient patch versionManager to %v", bumpedVersion)
	context.String(http.StatusOK, "%s", bumpedVersion.String())
}

// OnTransientMinor is a handler for a transient minor bump
func (handler *Handler) OnTransientMinor(context *gin.Context) {
	version := context.Param("versionManager")
	bumpedVersion, err := handler.versionManager.BumpTransientMinor(version)
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	handler.logger.Infof("bump transient minor versionManager to %v", bumpedVersion)
	context.String(http.StatusOK, "%s", bumpedVersion.String())
}
