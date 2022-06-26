package business

import (
	"bytes"

	"github.com/ffelipelimao/excel-go/internal/domain"
)

//go:generate mockgen -source=./business.go -destination=../../mocks/business.go -package=mocks
type GameService interface {
	List() ([]domain.Game, error)
	ListExcel() (*bytes.Buffer, error)
}
