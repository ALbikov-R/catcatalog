package service

import "carcat/internal/models"

type store interface {
	Create(models.Catalog) error
	Read(fil models.Filter) ([]models.Catalog, error)
	Update(models.Patch) error
	Delete(id string) error
}
