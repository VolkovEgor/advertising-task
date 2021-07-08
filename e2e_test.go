package test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/VolkovEgor/advertising-task/internal/handler"
	"github.com/VolkovEgor/advertising-task/internal/repository"
	"github.com/VolkovEgor/advertising-task/internal/service"
	"github.com/VolkovEgor/advertising-task/test"

	"github.com/labstack/echo/v4"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_E2E_App(t *testing.T) {
	const prefix = "./"
	db, err := test.PrepareTestDatabase(prefix, false)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	app := echo.New()
	handler.Init(app)

	// Create first advert
	Convey("Given params", t, func() {
		const (
			expectedStatus   = http.StatusOK
			expectedAdvertId = 1
		)
		expectedBody := fmt.Sprintf(`{"advert_id":%d}`, expectedAdvertId) + "\n"

		const (
			inputTitle       = `"First Advert"`
			inputDescription = `"Description of first advert"`
			inputPrice       = 30000
		)
		inputPhotos := []string{`"link1"`, `"link2"`, `"link3"`}

		photos := `[` + strings.Join(inputPhotos, `, `) + `]`
		inputBody := fmt.Sprintf(`{"title": %s, "description": %s, "photos": %s, "price": %d}`,
			inputTitle, inputDescription, photos, inputPrice)

		Convey("When create method", func() {
			req := httptest.NewRequest(
				"POST",
				"/api/adverts",
				bytes.NewBufferString(inputBody),
			)
			req.Header.Set("Content-type", "application/json")

			w := httptest.NewRecorder()
			app.ServeHTTP(w, req)

			Convey("Then should be Ok", func() {
				So(w.Code, ShouldEqual, expectedStatus)
				So(w.Body.String(), ShouldEqual, expectedBody)
			})
		})
	})

	// Get all adverts
	Convey("Given params", t, func() {
		const (
			expectedStatus = http.StatusOK
		)
		expectedBody := `[{"id":1,"title":"First Advert","main_photo":"link1","price":30000}]` + "\n"

		const (
			inputPage = 1
			inputSort = ""
		)

		Convey("When get all method", func() {
			target := fmt.Sprintf(`/api/adverts?page=%d&sort=%s`, inputPage, inputSort)
			req := httptest.NewRequest("GET", target, nil)

			w := httptest.NewRecorder()
			app.ServeHTTP(w, req)

			Convey("Then should be Ok", func() {
				So(w.Code, ShouldEqual, expectedStatus)
				So(w.Body.String(), ShouldEqual, expectedBody)
			})
		})
	})

	// Create second advert
	Convey("Given params", t, func() {
		const (
			expectedStatus   = http.StatusOK
			expectedAdvertId = 2
		)
		expectedBody := fmt.Sprintf(`{"advert_id":%d}`, expectedAdvertId) + "\n"

		const (
			inputTitle       = `"Second Advert"`
			inputDescription = `"Description of second advert"`
			inputPrice       = 10000
		)
		inputPhotos := []string{`"link1"`, `"link2"`}

		photos := `[` + strings.Join(inputPhotos, `, `) + `]`
		inputBody := fmt.Sprintf(`{"title": %s, "description": %s, "photos": %s, "price": %d}`,
			inputTitle, inputDescription, photos, inputPrice)

		Convey("When create method", func() {
			req := httptest.NewRequest(
				"POST",
				"/api/adverts",
				bytes.NewBufferString(inputBody),
			)
			req.Header.Set("Content-type", "application/json")

			w := httptest.NewRecorder()
			app.ServeHTTP(w, req)

			Convey("Then should be Ok", func() {
				So(w.Code, ShouldEqual, expectedStatus)
				So(w.Body.String(), ShouldEqual, expectedBody)
			})
		})
	})

	// Get all adverts with sorting by price
	Convey("Given params", t, func() {
		const (
			expectedStatus = http.StatusOK
		)
		expectedBody := `[{"id":2,"title":"Second Advert","main_photo":"link1","price":10000},` +
			`{"id":1,"title":"First Advert","main_photo":"link1","price":30000}]` + "\n"

		const (
			inputPage = 1
			inputSort = "price_asc"
		)

		Convey("When get all method with sorting by price", func() {
			target := fmt.Sprintf(`/api/adverts?page=%d&sort=%s`, inputPage, inputSort)
			req := httptest.NewRequest("GET", target, nil)

			w := httptest.NewRecorder()
			app.ServeHTTP(w, req)

			Convey("Then should be Ok", func() {
				So(w.Code, ShouldEqual, expectedStatus)
				So(w.Body.String(), ShouldEqual, expectedBody)
			})
		})
	})

	// Get all adverts with sorting by creation date
	Convey("Given params", t, func() {
		const (
			expectedStatus = http.StatusOK
		)
		expectedBody := `[{"id":1,"title":"First Advert","main_photo":"link1","price":30000},` +
			`{"id":2,"title":"Second Advert","main_photo":"link1","price":10000}]` + "\n"

		const (
			inputPage = 1
			inputSort = "date_asc"
		)

		Convey("When get all method with sorting by price", func() {
			target := fmt.Sprintf(`/api/adverts?page=%d&sort=%s`, inputPage, inputSort)
			req := httptest.NewRequest("GET", target, nil)

			w := httptest.NewRecorder()
			app.ServeHTTP(w, req)

			Convey("Then should be Ok", func() {
				So(w.Code, ShouldEqual, expectedStatus)
				So(w.Body.String(), ShouldEqual, expectedBody)
			})
		})
	})

	// Get advert by id without all fields
	Convey("Given params", t, func() {
		const (
			expectedStatus = http.StatusOK
		)
		expectedBody := `{"id":2,"title":"Second Advert","photos":["link1"],"price":10000}` + "\n"

		const (
			advertId = 2
			fields   = "false"
		)

		Convey("When get by id without all fields", func() {
			target := fmt.Sprintf(`/api/adverts/%d?fields=%s`, advertId, fields)
			req := httptest.NewRequest("GET", target, nil)

			w := httptest.NewRecorder()
			app.ServeHTTP(w, req)

			Convey("Then should be Ok", func() {
				So(w.Code, ShouldEqual, expectedStatus)
				So(w.Body.String(), ShouldEqual, expectedBody)
			})
		})
	})

	// Get advert by id with all fields
	Convey("Given params", t, func() {
		const (
			expectedStatus = http.StatusOK
		)
		expectedBody := `{"id":2,"title":"Second Advert","description":"Description of second advert",` +
			`"photos":["link1","link2"],"price":10000}` + "\n"

		const (
			advertId = 2
			fields   = "true"
		)

		Convey("When get by id without all fields", func() {
			target := fmt.Sprintf(`/api/adverts/%d?fields=%s`, advertId, fields)
			req := httptest.NewRequest("GET", target, nil)

			w := httptest.NewRecorder()
			app.ServeHTTP(w, req)

			Convey("Then should be Ok", func() {
				So(w.Code, ShouldEqual, expectedStatus)
				So(w.Body.String(), ShouldEqual, expectedBody)
			})
		})
	})
}
