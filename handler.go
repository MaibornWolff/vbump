package main

import (
	"net/http"

	"maibornwolff/vbump/service"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

// Handler for handling http routes
type Handler struct {
	versionManager *service.VersionManager
}

// NewHandler constructs a new handler
func NewHandler(versionManager *service.VersionManager) *Handler {
	return &Handler{
		versionManager: versionManager,
	}
}

// LoggerMiddleware logs the last error
func (handler *Handler) LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err != nil {
			log.Error().Err(err).Msg("")
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
	r.POST("/transient/minor/:version", handler.OnTransientMinor)
	r.POST("/transient/patch/:version", handler.OnTransientPatch)
	r.POST("/version/:project/:version", handler.OnSetVersion)
	r.GET("/version/:project", handler.OnGetVersion)
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
	log.Info().Str("version", version.String()).Str("project", project).Msg("Bumped major version")
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
	log.Info().Str("version", version.String()).Str("project", project).Msg("Bumped minor version")
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
	log.Info().Str("version", version.String()).Str("project", project).Msg("Bumped patch version")
	context.String(http.StatusOK, "%s", version.String())
}

// OnSetVersion is a handler for setting the version for a given project
func (handler *Handler) OnSetVersion(context *gin.Context) {
	project := context.Param("project")
	version := context.Param("version")
	_, err := handler.versionManager.SetVersion(project, version)
	if err != nil {
		_ = context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Info().Str("version", version).Str("project", project).Msg("Set version explicitly")
	context.String(http.StatusOK, "%s", version)
}

// OnGetVersion is a handler for getting the version for a given project
func (handler *Handler) OnGetVersion(context *gin.Context) {
	project := context.Param("project")
	version, err := handler.versionManager.GetVersion(project)
	if err != nil {
		_ = context.AbortWithError(http.StatusNotFound, err)
		return
	}

	log.Info().Str("project", project).Msg("Got version")
	context.String(http.StatusOK, "%s", version.String())
}

// OnTransientPatch is a handler for a transient patch bump
func (handler *Handler) OnTransientPatch(context *gin.Context) {
	version := context.Param("version")
	bumpedVersion, err := handler.versionManager.BumpTransientPatch(version)
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Info().Str("version", bumpedVersion.String()).Msg("Bumped patch version transiently")
	context.String(http.StatusOK, "%s", bumpedVersion.String())
}

// OnTransientMinor is a handler for a transient minor bump
func (handler *Handler) OnTransientMinor(context *gin.Context) {
	version := context.Param("version")
	bumpedVersion, err := handler.versionManager.BumpTransientMinor(version)
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Info().Str("version", bumpedVersion.String()).Msg("Bumped minor version transiently")
	context.String(http.StatusOK, "%s", bumpedVersion.String())
}
