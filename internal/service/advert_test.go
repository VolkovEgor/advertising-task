package service

import (
	"errors"
	"testing"

	"github.com/VolkovEgor/advertising-task/internal/model"
	"github.com/VolkovEgor/advertising-task/internal/repository"
	mock_repositories "github.com/VolkovEgor/advertising-task/internal/repository/mocks"
	"github.com/VolkovEgor/advertising-task/test"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAdvertService_GetAll(t *testing.T) {
	type args struct {
		page int
		sort string
	}
	type mockBehavior func(r *mock_repositories.MockAdvert, args args)

	tests := []struct {
		name    string
		input   args
		mock    mockBehavior
		wantErr bool
		want    []*model.Advert
	}{
		{
			name: "Ok without sorting",
			input: args{
				page: 1,
			},
			mock: func(r *mock_repositories.MockAdvert, args args) {
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
				r.EXPECT().GetAll(args.page, "", "").Return(data, nil)
			},
			want: []*model.Advert{
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
			},
		},
		{
			name: "Ok with sorting by price",
			input: args{
				page: 1,
				sort: "price_asc",
			},
			mock: func(r *mock_repositories.MockAdvert, args args) {
				data := []*model.Advert{
					{
						Id:        2,
						Title:     "Second Advert",
						MainPhoto: "link2",
						Price:     10000,
					},
					{
						Id:        1,
						Title:     "First Advert",
						MainPhoto: "link1",
						Price:     30000,
					},
				}
				r.EXPECT().GetAll(args.page, "price", "asc").Return(data, nil)
			},
			want: []*model.Advert{
				{
					Id:        2,
					Title:     "Second Advert",
					MainPhoto: "link2",
					Price:     10000,
				},
				{
					Id:        1,
					Title:     "First Advert",
					MainPhoto: "link1",
					Price:     30000,
				},
			},
		},
		{
			name: "Ok with sorting by date",
			input: args{
				page: 1,
				sort: "date_asc",
			},
			mock: func(r *mock_repositories.MockAdvert, args args) {
				data := []*model.Advert{
					{
						Id:        1,
						Title:     "First Advert",
						MainPhoto: "link1",
						Price:     10000,
					},
					{
						Id:        2,
						Title:     "Second Advert",
						MainPhoto: "link2",
						Price:     30000,
					},
				}
				r.EXPECT().GetAll(args.page, "creation_date", "asc").Return(data, nil)
			},
			want: []*model.Advert{
				{
					Id:        1,
					Title:     "First Advert",
					MainPhoto: "link1",
					Price:     10000,
				},
				{
					Id:        2,
					Title:     "Second Advert",
					MainPhoto: "link2",
					Price:     30000,
				},
			},
		},
		{
			name: "Repo Error",
			input: args{
				page: 1,
				sort: "price_asc",
			},
			mock: func(r *mock_repositories.MockAdvert, args args) {
				r.EXPECT().GetAll(args.page, "price", "asc").Return(nil, errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "Page number greater than possible",
			input: args{
				page: 10,
				sort: "price_asc",
			},
			mock: func(r *mock_repositories.MockAdvert, args args) {
				r.EXPECT().GetAll(args.page, "price", "asc").Return(nil, nil)
			},
			wantErr: true,
		},
		{
			name: "Wrong sort parameter",
			input: args{
				page: 10,
				sort: "price_ascc",
			},
			mock:    func(r *mock_repositories.MockAdvert, args args) {},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repositories.NewMockAdvert(c)
			test.mock(repo, test.input)
			s := &AdvertService{repo: repo}

			got, err := s.GetAll(test.input.page, test.input.sort)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestAdvertService_GetById(t *testing.T) {
	const prefix = "../../"
	db, err := test.PrepareTestDatabase(prefix, true)
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer db.Close()

	type args struct {
		advertId int
		fields   bool
	}

	tests := []struct {
		name    string
		input   args
		want    *model.DetailedAdvert
		wantErr bool
	}{
		{
			name: "Ok with getting all fields",
			input: args{
				advertId: 1,
				fields:   true,
			},
			want: &model.DetailedAdvert{
				Id:          1,
				Title:       "Advert 1",
				Description: "Description 1",
				Photos:      []string{"link1", "link2", "link3"},
				Price:       10000,
			},
			wantErr: false,
		},
		{
			name: "Ok without getting all fields",
			input: args{
				advertId: 1,
				fields:   false,
			},
			want: &model.DetailedAdvert{
				Id:     1,
				Title:  "Advert 1",
				Photos: []string{"link1"},
				Price:  10000,
			},
			wantErr: false,
		},
		{
			name: "Wrong advert id",
			input: args{
				advertId: -1,
				fields:   true,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewRepository(db)
			s := &AdvertService{repo: repo}

			got, err := s.GetById(tt.input.advertId, tt.input.fields)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestAdvertService_Create(t *testing.T) {
	type args struct {
		advert *model.DetailedAdvert
	}
	type mockBehavior func(r *mock_repositories.MockAdvert, args args)

	tests := []struct {
		name    string
		input   args
		mock    mockBehavior
		wantErr bool
		want    int
	}{
		{
			name: "Ok",
			input: args{
				advert: &model.DetailedAdvert{
					Title:       "New Test Ad",
					Description: "Some Description",
					Photos:      []string{"link1", "link2", "link3"},
					Price:       10000,
				},
			},
			mock: func(r *mock_repositories.MockAdvert, args args) {
				r.EXPECT().Create(args.advert).Return(1, nil)
			},
			want: 1,
		},
		{
			name: "Repo Error",
			input: args{
				advert: &model.DetailedAdvert{
					Title:        "New Test Advert",
					Description:  "Some Description",
					Photos:       []string{"link1", "link2", "link3"},
					Price:        10000,
					CreationDate: 1,
				},
			},
			mock: func(r *mock_repositories.MockAdvert, args args) {
				r.EXPECT().Create(args.advert).Return(0, errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "Wrong title",
			input: args{
				advert: &model.DetailedAdvert{
					Title:        "",
					Description:  "Some Description",
					Photos:       []string{"link1", "link2", "link3"},
					Price:        10000,
					CreationDate: 1,
				},
			},
			mock:    func(r *mock_repositories.MockAdvert, args args) {},
			wantErr: true,
		},
		{
			name: "Wrong description",
			input: args{
				advert: &model.DetailedAdvert{
					Title:        "New Test Advert",
					Description:  "",
					Photos:       []string{"link1", "link2", "link3"},
					Price:        10000,
					CreationDate: 1,
				},
			},
			mock:    func(r *mock_repositories.MockAdvert, args args) {},
			wantErr: true,
		},
		{
			name: "Wrong photos",
			input: args{
				advert: &model.DetailedAdvert{
					Title:        "New Test Advert",
					Description:  "Some Description",
					Photos:       []string{"link1", "link2", "link3", "link4"},
					Price:        10000,
					CreationDate: 1,
				},
			},
			mock:    func(r *mock_repositories.MockAdvert, args args) {},
			wantErr: true,
		},
		{
			name: "Wrong price",
			input: args{
				advert: &model.DetailedAdvert{
					Title:        "New Test Advert",
					Description:  "Some Description",
					Photos:       []string{"link1", "link2", "link3"},
					Price:        -10000,
					CreationDate: 1,
				},
			},
			mock:    func(r *mock_repositories.MockAdvert, args args) {},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repositories.NewMockAdvert(c)
			test.mock(repo, test.input)
			s := &AdvertService{repo: repo}

			got, err := s.Create(test.input.advert)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
