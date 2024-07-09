package repository

import (
	"errors"
	"go-inertia/app/models"
)

func GetProviders() []models.Provider {
	providers := make([]models.Provider, 3)
	providers[0] = models.Provider{ID: "1", Name: "Claro"}
	providers[1] = models.Provider{ID: "2", Name: "Vivo"}
	providers[2] = models.Provider{ID: "3", Name: "Oi"}

	return providers
}

func FindProvider(id string) (*models.Provider, error) {
	providers := GetProviders()

	for _, provider := range providers {
		if provider.ID == id {
			return &provider, nil
		}
	}

	return nil, errors.New("Provider not found")
}
