package integrations

import "github.com/ffelipelimao/excel-go/internal/domain"

//TODO: Add mockgen to generate mockups
type GameIntegration interface {
	Get() ([]domain.Game, error)
}
