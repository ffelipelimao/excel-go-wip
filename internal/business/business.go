package business

import (
	"bytes"
	"mime/multipart"

	"github.com/ffelipelimao/excel-go/internal/domain"
)

//go:generate mockgen -source=./business.go -destination=../../mocks/business.go -package=mocks
type GameService interface {
	List() ([]domain.Game, error)
	ListExcel() (*bytes.Buffer, error)
	ReadExcel(file multipart.File) ([]domain.Game, error)
}
