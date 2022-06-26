package business

import (
	"bytes"

	"github.com/ffelipelimao/excel-go/internal/domain"
)

//TODO: Add mockgen to generate mockups
type GameService interface {
	List() ([]domain.Game, error)
	ListExcel() (*bytes.Buffer, error)
}
