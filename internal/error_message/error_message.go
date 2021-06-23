package error_message

import "errors"

var (
	ErrWrongPageNumber  = errors.New("wrong page number")
	ErrWrongSortParams  = errors.New("wrong sort params")
	ErrWrongAdvertId    = errors.New("wrong advert id")
	ErrWrongFieldsParam = errors.New("wrong fields param")
	ErrWrongTitle       = errors.New("title must contain from 1 to 200 characters")
	ErrWrongDescription = errors.New("description must contain from 1 to 1000 characters")
	ErrWrongPhotos      = errors.New("advert must contain from 1 to 3 photos")
	ErrNotPositivePrice = errors.New("price must be positive number")
)
