package integrations

import (
	"github.com/ffelipelimao/excel-go/internal/domain"
)

const _gamePath = "games"

type GameIntegrationImpl struct {
}

func NewIntegration() GameIntegration {
	return GameIntegrationImpl{}
}

func (i GameIntegrationImpl) Get() ([]domain.Game, error) {

	return []domain.Game{
		{
			ID:    "1",
			Name:  "Batman Arkham City",
			Price: "4.99",
			Thumb: "https://s3.batman.arkham.jpg",
		},
		{
			ID:    "1",
			Name:  "Batman Arkham Asylum",
			Price: "3.99",
			Thumb: "https://s3.batman.arkham.jpg",
		},
		{
			ID:    "1",
			Name:  "Batman Arkham Knight",
			Price: "2.99",
			Thumb: "https://s3.batman.arkham.jpg",
		},
	}, nil
}
