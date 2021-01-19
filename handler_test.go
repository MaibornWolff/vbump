package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"maibornwolff/vbump/adapter"

	. "github.com/onsi/gomega"
)

func Test_Bumb_Major(t *testing.T) {
	Ω := NewGomegaWithT(t)
	fileProvider := adapter.NewMock("1.0.0", "p1")
	version := NewVersion(fileProvider)
	handler := NewHandler(version, nil)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/major/p1", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Body.String()).To(Equal("2.0.0"))
}

func Test_Bumb_Minor(t *testing.T) {
	Ω := NewGomegaWithT(t)
	fileProvider := adapter.NewMock("1.0.0", "p1")
	version := NewVersion(fileProvider)
	handler := NewHandler(version, nil)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/minor/p1", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Body.String()).To(Equal("1.1.0"))
}

func Test_Bumb_Patch(t *testing.T) {
	Ω := NewGomegaWithT(t)
	fileProvider := adapter.NewMock("1.0.0", "p1")
	version := NewVersion(fileProvider)
	handler := NewHandler(version, nil)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/patch/p1", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Body.String()).To(Equal("1.0.1"))
}

func Test_Number_Of_Vbump_Metric(t *testing.T) {
	Ω := NewGomegaWithT(t)
	fileProvider := adapter.NewMock("init", "init")
	version := NewVersion(fileProvider)
	handler := NewHandler(version, nil)
	router := handler.GetRouter()
	res := httptest.NewRecorder()
	metrics, _ := http.NewRequest("GET", "/metrics", nil)

	// test for prom1
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

	// test for prom2
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
func Test_Bumb_Transient_Patch(t *testing.T) {
	Ω := NewGomegaWithT(t)
	fileProvider := adapter.NewMock("1.0.0", "p1")
	version := NewVersion(fileProvider)
	handler := NewHandler(version, nil)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/patch/transient/1.0", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Body.String()).To(Equal("1.0.1"))
}

func Test_Bumb_Transient_Minor(t *testing.T) {
	Ω := NewGomegaWithT(t)
	fileProvider := adapter.NewMock("1.0.0", "p1")
	version := NewVersion(fileProvider)
	handler := NewHandler(version, nil)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/minor/transient/1.0", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Body.String()).To(Equal("1.1"))
}

func Test_Set_Version(t *testing.T) {
	Ω := NewGomegaWithT(t)
	fileProvider := adapter.NewMock("1.0.0", "p1")
	version := NewVersion(fileProvider)
	handler := NewHandler(version, nil)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/version/p1/3.1.2", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Body.String()).To(Equal("3.1.2"))
}

func Test_Get_Version_With_Handler(t *testing.T) {
	Ω := NewGomegaWithT(t)
	fileProvider := adapter.NewMock("1.0.0", "p1")
	version := NewVersion(fileProvider)
	handler := NewHandler(version, nil)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/version/p1", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Body.String()).To(Equal("1.0.0"))
}

func Test_Set_Version_With_Validation_Error(t *testing.T) {
	Ω := NewGomegaWithT(t)
	fileProvider := adapter.NewMock("1.0.0.", "p1")
	version := NewVersion(fileProvider)
	handler := NewHandler(version, nil)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/version/p1/3.1.2.", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Code).To(Equal(422))
	Ω.Expect(res.Body.String()).To(Equal("3.1.2. is not a valid version"))
}

func Test_Bump_Major_With_Invalid_Path(t *testing.T) {
	Ω := NewGomegaWithT(t)
	fileProvider := adapter.New("dirnotexist")
	version := NewVersion(fileProvider)
	handler := NewHandler(version, nil)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/major/p1", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Code).To(Equal(500))
	Ω.Expect(res.Body.String()).To(Equal(""))
}

func Test_Bump_Minor_With_Invalid_Path(t *testing.T) {
	Ω := NewGomegaWithT(t)
	fileProvider := adapter.New("dirnotexist")
	version := NewVersion(fileProvider)
	handler := NewHandler(version, nil)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/minor/p1", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Code).To(Equal(500))
	Ω.Expect(res.Body.String()).To(Equal(""))
}

func Test_Bump_Patch_With_Invalid_Path(t *testing.T) {
	Ω := NewGomegaWithT(t)
	fileProvider := adapter.New("dirnotexist")
	version := NewVersion(fileProvider)
	handler := NewHandler(version, nil)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/patch/p1", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Code).To(Equal(500))
	Ω.Expect(res.Body.String()).To(Equal(""))
}

func Test_Get_Health(t *testing.T) {
	Ω := NewGomegaWithT(t)
	handler := NewHandler(nil, nil)
	router := handler.GetRouter()
	res := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(res, req)

	Ω.Expect(res.Body.String()).To(Equal("hello from vbump!"))
}
