package main

import (
	"maibornwolff/vbump/service"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"
	"maibornwolff/vbump/adapter"
	"maibornwolff/vbump/model"
)

func TestBumbMajor(t *testing.T) {
	Ω := NewGomegaWithT(t)

	fileProvider := adapter.NewMock(model.NewVersion(1, 0, 0), "p1")
	versionManager := service.NewVersionManager(fileProvider)
	handler := NewHandler(versionManager)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/major/p1", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Body.String()).To(Equal("2.0.0"))
}

func TestBumbMinor(t *testing.T) {
	Ω := NewGomegaWithT(t)

	fileProvider := adapter.NewMock(model.NewVersion(1, 0, 0), "p1")
	versionManager := service.NewVersionManager(fileProvider)
	handler := NewHandler(versionManager)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/minor/p1", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Body.String()).To(Equal("1.1.0"))
}

func TestBumbPatch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	fileProvider := adapter.NewMock(model.NewVersion(1, 0, 0), "p1")
	versionManager := service.NewVersionManager(fileProvider)
	handler := NewHandler(versionManager)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/patch/p1", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Body.String()).To(Equal("1.0.1"))
}

func TestNumberOfVbumpMetric(t *testing.T) {
	Ω := NewGomegaWithT(t)

	fileProvider := adapter.NewMock(model.NewVersion(1, 0, 0), "init")
	versionManager := service.NewVersionManager(fileProvider)
	handler := NewHandler(versionManager)
	router := handler.GetRouter()
	res := httptest.NewRecorder()
	metrics, _ := http.NewRequest("GET", "/metrics", nil)

	// Test for prom1
	patchp1, _ := http.NewRequest("POST", "/patch/prom1", nil)
	minorp1, _ := http.NewRequest("POST", "/minor/prom1", nil)
	majorp1, _ := http.NewRequest("POST", "/major/prom1", nil)

	router.ServeHTTP(res, patchp1)
	router.ServeHTTP(res, minorp1)
	router.ServeHTTP(res, majorp1)
	router.ServeHTTP(res, metrics)

	Ω.Expect(res.Body.String()).To(ContainSubstring("vbump_bumps_total{element=\"patch\",project=\"prom1\"} 1"))
	Ω.Expect(res.Body.String()).To(ContainSubstring("vbump_bumps_total{element=\"minor\",project=\"prom1\"} 1"))
	Ω.Expect(res.Body.String()).To(ContainSubstring("vbump_bumps_total{element=\"major\",project=\"prom1\"} 1"))

	// Test for prom2
	patchp2, _ := http.NewRequest("POST", "/patch/prom2", nil)
	minorp2, _ := http.NewRequest("POST", "/minor/prom2", nil)
	majorp2, _ := http.NewRequest("POST", "/major/prom2", nil)

	router.ServeHTTP(res, patchp2)
	router.ServeHTTP(res, minorp2)
	router.ServeHTTP(res, majorp2)
	router.ServeHTTP(res, metrics)

	Ω.Expect(res.Body.String()).To(ContainSubstring("vbump_bumps_total{element=\"patch\",project=\"prom2\"} 1"))
	Ω.Expect(res.Body.String()).To(ContainSubstring("vbump_bumps_total{element=\"minor\",project=\"prom2\"} 1"))
	Ω.Expect(res.Body.String()).To(ContainSubstring("vbump_bumps_total{element=\"major\",project=\"prom2\"} 1"))
}

func TestBumbTransientPatch(t *testing.T) {
	Ω := NewGomegaWithT(t)

	fileProvider := adapter.NewMock(model.NewVersion(1, 0, 0), "p1")
	versionManager := service.NewVersionManager(fileProvider)
	handler := NewHandler(versionManager)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/transient/patch/1.0", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Body.String()).To(Equal("1.0.1"))
}

func TestBumbTransientMinor(t *testing.T) {
	Ω := NewGomegaWithT(t)

	fileProvider := adapter.NewMock(model.NewVersion(1, 0, 0), "p1")
	versionManager := service.NewVersionManager(fileProvider)
	handler := NewHandler(versionManager)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/transient/minor/1.0", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Body.String()).To(Equal("1.1"))
}

func TestSetVersion(t *testing.T) {
	Ω := NewGomegaWithT(t)

	fileProvider := adapter.NewMock(model.NewVersion(1, 0, 0), "p1")
	versionManager := service.NewVersionManager(fileProvider)
	handler := NewHandler(versionManager)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/version/p1/3.1.2", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Body.String()).To(Equal("3.1.2"))
}

func TestGetVersionWithHandler(t *testing.T) {
	Ω := NewGomegaWithT(t)

	fileProvider := adapter.NewMock(model.NewVersion(1, 0, 0), "p1")
	versionManager := service.NewVersionManager(fileProvider)
	handler := NewHandler(versionManager)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/version/p1", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Body.String()).To(Equal(model.NewVersion(1, 0, 0).String()))
}

func TestSetVersionWithValidationError(t *testing.T) {
	Ω := NewGomegaWithT(t)

	fileProvider := adapter.NewMock(model.NewVersion(1, 0, 0), "p1")
	versionManager := service.NewVersionManager(fileProvider)
	handler := NewHandler(versionManager)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/version/p1/3.1.2.", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Code).To(Equal(400))
}

func TestBumpMajorWithInvalidPath(t *testing.T) {
	Ω := NewGomegaWithT(t)

	fileProvider := adapter.NewFileProvider("dirnotexist")
	versionManager := service.NewVersionManager(fileProvider)
	handler := NewHandler(versionManager)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/major/p1", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Code).To(Equal(500))
	Ω.Expect(res.Body.String()).To(Equal(""))
}

func TestBumpMinorWithInvalidPath(t *testing.T) {
	Ω := NewGomegaWithT(t)

	fileProvider := adapter.NewFileProvider("dirnotexist")
	versionManager := service.NewVersionManager(fileProvider)
	handler := NewHandler(versionManager)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/minor/p1", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Code).To(Equal(500))
	Ω.Expect(res.Body.String()).To(Equal(""))
}

func TestBumpPatchWithInvalidPath(t *testing.T) {
	Ω := NewGomegaWithT(t)

	fileProvider := adapter.NewFileProvider("dirnotexist")
	versionManager := service.NewVersionManager(fileProvider)
	handler := NewHandler(versionManager)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/patch/p1", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Code).To(Equal(500))
	Ω.Expect(res.Body.String()).To(Equal(""))
}

func TestGetHealth(t *testing.T) {
	Ω := NewGomegaWithT(t)

	handler := NewHandler(nil)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Body.String()).To(Equal("hello from vbump!"))
}
