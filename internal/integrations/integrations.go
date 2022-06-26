package integrations

import "github.com/ffelipelimao/excel-go/internal/domain"

//go:generate mockgen -source=./integrations.go -destination=../../mocks/integrations.go -package=mocks
type GameIntegration interface {
	Get() ([]domain.Game, error)
}
