package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/christle/vbump/adapter"
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
