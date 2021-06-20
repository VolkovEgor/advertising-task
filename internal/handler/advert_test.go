package handler

import (
	"bytes"
	"errors"
	"fmt"
	"net/http/httptest"
	"strconv"
	"testing"

	. "github.com/VolkovEgor/advertising-task/internal/error"
	"github.com/VolkovEgor/advertising-task/internal/model"
	"github.com/VolkovEgor/advertising-task/internal/service"
	mock_service "github.com/VolkovEgor/advertising-task/internal/service/mocks"
	"github.com/labstack/echo/v4"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAdvertHandler_GetAll(t *testing.T) {
	type args struct {
		page int
		sort string
	}
	type mockBehavior func(r *mock_service.MockAdvert, args args)

	tests := []struct {
		name                 string
		inputBody            string
		input                args
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			input: args{
				page: 1,
				sort: "price_asc",
			},
			mockBehavior: func(r *mock_service.MockAdvert, args args) {
				data := []*model.Advert{
					{
						Id:        1,
						Title:     "First Advert",
						MainPhoto: "link1",
						Price:     30000,
					},
					{
						Id:        2,
						Title:     "Second Advert",
						MainPhoto: "link2",
						Price:     10000,
					},
				}
				r.EXPECT().GetAll(args.page, args.sort).Return(data, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `[{"id":1,"title":"First Advert","main_photo":"link1","price":30000},` +
				`{"id":2,"title":"Second Advert","main_photo":"link2","price":10000}]` + "\n",
		},
		{
			name: "Wrong sort parameter",
			input: args{
				page: 1,
				sort: "price_ascc",
			},
			mockBehavior: func(r *mock_service.MockAdvert, args args) {
				r.EXPECT().GetAll(args.page, args.sort).Return(nil, ErrWrongSortParams)
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"` + ErrWrongSortParams.Error() + `"}` + "\n",
		},
		{
			name: "Wrong page number",
			input: args{
				page: -1,
				sort: "price_asc",
			},
			mockBehavior:         func(r *mock_service.MockAdvert, args args) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"` + ErrWrongPageNumber.Error() + `"}` + "\n",
		},
		{
			name: "Service error",
			input: args{
				page: 1,
				sort: "price_asc",
			},
			mockBehavior: func(r *mock_service.MockAdvert, args args) {
				r.EXPECT().GetAll(args.page, args.sort).Return(nil, errors.New("Some error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"Some error"}` + "\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAdvert(c)
			test.mockBehavior(repo, test.input)

			service := &service.Service{Advert: repo}
			handler := Handler{service}

			app := echo.New()
			handler.Init(app)

			w := httptest.NewRecorder()
			target := fmt.Sprintf("/api/adverts?page=%d&sort=%s",
				test.input.page, test.input.sort)
			req := httptest.NewRequest("GET", target, nil)

			app.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestAdvertHandler_GetById(t *testing.T) {
	type args struct {
		advertId int
		fields   string
	}
	type mockBehavior func(r *mock_service.MockAdvert, args args)

	tests := []struct {
		name                 string
		inputBody            string
		input                args
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			input: args{
				advertId: 1,
				fields:   "true",
			},
			mockBehavior: func(r *mock_service.MockAdvert, args args) {
				data := &model.DetailedAdvert{
					Id:          1,
					Title:       "First Advert",
					Description: "Some Description",
					Photos:      []string{"link1", "link2", "link3"},
					Price:       30000,
				}
				boolFields, _ := strconv.ParseBool(args.fields)
				r.EXPECT().GetById(args.advertId, boolFields).Return(data, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{"id":1,"title":"First Advert","description":"Some Description",` +
				`"photos":["link1","link2","link3"],"price":30000}` + "\n",
		},
		{
			name: "Wrong advert id",
			input: args{
				advertId: -1,
				fields:   "true",
			},
			mockBehavior:         func(r *mock_service.MockAdvert, args args) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"` + ErrWrongAdvertId.Error() + `"}` + "\n",
		},
		{
			name: "Wrong fields",
			input: args{
				advertId: 1,
				fields:   "truee",
			},
			mockBehavior:         func(r *mock_service.MockAdvert, args args) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"` + ErrWrongFieldsParam.Error() + `"}` + "\n",
		},
		{
			name: "Service error",
			input: args{
				advertId: 1,
				fields:   "true",
			},
			mockBehavior: func(r *mock_service.MockAdvert, args args) {
				boolFields, _ := strconv.ParseBool(args.fields)
				r.EXPECT().GetById(args.advertId, boolFields).Return(nil, errors.New("Some error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"Some error"}` + "\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAdvert(c)
			test.mockBehavior(repo, test.input)

			service := &service.Service{Advert: repo}
			handler := Handler{service}

			app := echo.New()
			handler.Init(app)

			w := httptest.NewRecorder()
			target := fmt.Sprintf("/api/adverts/%d?fields=%s",
				test.input.advertId, test.input.fields)
			req := httptest.NewRequest("GET", target, nil)

			app.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestAdvertHandler_Create(t *testing.T) {
	type args struct {
		advert *model.DetailedAdvert
	}
	type mockBehavior func(r *mock_service.MockAdvert, args args)

	tests := []struct {
		name                 string
		inputBody            string
		input                args
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			inputBody: `{"title": "Test Advert", "Description": "Some Description",` +
				`"photos": ["link1", "link2", "link3"], "price": 10000}`,
			input: args{
				&model.DetailedAdvert{
					Title:       "Test Advert",
					Description: "Some Description",
					Photos:      []string{"link1", "link2", "link3"},
					Price:       10000,
				},
			},
			mockBehavior: func(r *mock_service.MockAdvert, args args) {
				r.EXPECT().Create(args.advert).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"advert_id":1}` + "\n",
		},
		{
			name: "Wrong data",
			inputBody: `{"title": "", "Description": "Some Description",` +
				`"photos": ["link1", "link2", "link3"], "price": 10000}`,
			input: args{
				&model.DetailedAdvert{
					Title:       "",
					Description: "Some Description",
					Photos:      []string{"link1", "link2", "link3"},
					Price:       10000,
				},
			},
			mockBehavior: func(r *mock_service.MockAdvert, args args) {
				r.EXPECT().Create(args.advert).Return(0, ErrWrongTitle)
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"` + ErrWrongTitle.Error() + `"}` + "\n",
		},
		{
			name: "Service error",
			inputBody: `{"title": "Test Advert", "Description": "Some Description",` +
				`"photos": ["link1", "link2", "link3"], "price": 10000}`,
			input: args{
				&model.DetailedAdvert{
					Title:       "Test Advert",
					Description: "Some Description",
					Photos:      []string{"link1", "link2", "link3"},
					Price:       10000,
				},
			},
			mockBehavior: func(r *mock_service.MockAdvert, args args) {
				r.EXPECT().Create(args.advert).Return(0, errors.New("Some error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"Some error"}` + "\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAdvert(c)
			test.mockBehavior(repo, test.input)

			service := &service.Service{Advert: repo}
			handler := Handler{service}

			app := echo.New()
			handler.Init(app)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/adverts",
				bytes.NewBufferString(test.inputBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			app.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
