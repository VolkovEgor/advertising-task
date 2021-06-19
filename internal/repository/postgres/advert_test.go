package postgres

import (
	"fmt"
	"strings"
	"testing"

	"github.com/VolkovEgor/advertising-task/internal/model"

	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestAdvertPg_GetAll(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewAdvertPg(db)

	type args struct {
		page      int
		sortField string
		order     string
	}
	type mockBehavior func(args args)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    []*model.Advert
		wantErr bool
	}{
		{
			name: "Ok with sorting",
			input: args{
				page:      1,
				sortField: "price",
				order:     "desc",
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
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"a.id", "a.title",
					"a.photos", "a.price"}).
					AddRow(1, "First Advert", "link1", 10000).
					AddRow(2, "Second Advert", "link2", 30000)
				mock.ExpectQuery("SELECT (.+) FROM adverts").
					WithArgs().WillReturnRows(rows)
			},
		},
		{
			name: "Ok without sorting",
			input: args{
				page:      1,
				sortField: "",
				order:     "",
			},
			want: []*model.Advert{
				{
					Id:        1,
					Title:     "First Ad",
					MainPhoto: "link1",
					Price:     10000,
				},
				{
					Id:        2,
					Title:     "Second Ad",
					MainPhoto: "link2",
					Price:     30000,
				},
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"a.id", "a.title",
					"a.photos", "a.price"}).
					AddRow(1, "First Ad", "link1", 10000).
					AddRow(2, "Second Ad", "link2", 30000)
				mock.ExpectQuery("SELECT (.+) FROM adverts").
					WithArgs().
					WillReturnRows(rows)
			},
		},
		{
			name: "Internal error",
			input: args{
				page:      1,
				sortField: "",
				order:     "",
			},
			wantErr: true,
			mock: func(args args) {
				mock.ExpectQuery("SELECT (.+) FROM adverts").
					WithArgs().
					WillReturnError(fmt.Errorf("Some error"))
			},
		},
		{
			name: "Wrong input data",
			input: args{
				page:      1,
				sortField: "time",
				order:     "desc",
			},
			wantErr: true,
			mock: func(args args) {
				mock.ExpectQuery("SELECT (.+) FROM adverts").
					WithArgs().
					WillReturnError(fmt.Errorf("Sort field does not exist"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input)

			got, err := r.GetAll(tt.input.page, tt.input.sortField, tt.input.order)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestAdvertPg_GetById(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewAdvertPg(db)

	type args struct {
		advertId int
		fields   bool
	}
	type mockBehavior func(args args)

	tests := []struct {
		name    string
		mock    mockBehavior
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
				Title:       "Test Advert",
				Description: "Some Description",
				Photos:      []string{"link1", "link2", "link3"},
				Price:       10000,
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"a.id", "a.title",
					"a.description", "a.photos", "a.price"}).
					AddRow(1, "Test Advert", "Some Description",
						"{link1,link2,link3}", 10000)
				mock.ExpectQuery("SELECT (.+) FROM adverts").
					WithArgs(args.advertId).WillReturnRows(rows)
			},
		},
		{
			name: "Ok without getting all fields",
			input: args{
				advertId: 1,
				fields:   false,
			},
			want: &model.DetailedAdvert{
				Id:     1,
				Title:  "Test Advert",
				Photos: []string{"link1"},
				Price:  10000,
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"a.id", "a.title",
					"a.photos", "a.price"}).
					AddRow(1, "Test Advert", "link1", 10000)
				mock.ExpectQuery("SELECT (.+) FROM adverts").
					WithArgs(args.advertId).WillReturnRows(rows)
			},
		},
		{
			name: "Wrong data",
			input: args{
				advertId: 1,
				fields:   false,
			},
			wantErr: true,
			mock: func(args args) {

				mock.ExpectQuery("SELECT (.+) FROM adverts").
					WithArgs(args.advertId).WillReturnError(fmt.Errorf("Error not positive advert id"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input)

			got, err := r.GetById(tt.input.advertId, tt.input.fields)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestAdvertPg_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewAdvertPg(db)

	type args struct {
		advert *model.DetailedAdvert
	}
	type mockBehavior func(args args, id int)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				advert: &model.DetailedAdvert{
					Title:        "New Test Advert",
					Description:  "Some Description",
					Photos:       []string{"link1", "link2", "link3"},
					Price:        10000,
					CreationDate: 1,
				},
			},
			want: 1,
			mock: func(args args, id int) {
				photos := strings.Join(args.advert.Photos, `", "`)
				photos = `{"` + photos + `"}`

				input := args.advert
				adRows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO adverts").
					WithArgs(input.Title, input.Description, photos,
						input.Price, input.CreationDate).
					WillReturnRows(adRows)
			},
		},
		{
			name: "Wrong data",
			input: args{
				advert: &model.DetailedAdvert{
					Title:        "New Test Advert",
					Description:  "Some Description",
					Photos:       []string{"link1", "link2", "link3"},
					Price:        -10000,
					CreationDate: 1,
				},
			},
			wantErr: true,
			mock: func(args args, id int) {
				input := args.advert
				mock.ExpectQuery("INSERT INTO adverts").
					WithArgs(input.Title, input.Description, input.Photos,
						input.Price, input.CreationDate).
					WillReturnError(fmt.Errorf("Error not positive number"))
			},
		},
		{
			name: "Failed insert",
			input: args{
				advert: &model.DetailedAdvert{
					Title:        "New Test Advert",
					Description:  "Some Description",
					Photos:       []string{"link1", "link2", "link3"},
					Price:        10000,
					CreationDate: 1,
				},
			},
			wantErr: true,
			mock: func(args args, id int) {
				input := args.advert
				mock.ExpectQuery("INSERT INTO adverts").
					WithArgs(input.Title, input.Description, input.Photos,
						input.Price, input.CreationDate).
					WillReturnError(fmt.Errorf("Some error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input, tt.want)

			got, err := r.Create(tt.input.advert)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
