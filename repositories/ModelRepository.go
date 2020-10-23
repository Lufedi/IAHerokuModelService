package repositories

import (
	"ModelsService/client"
	"ModelsService/model"
)

type ModelRepository struct {
}

func (repo ModelRepository) Insert(model *model.Model) error {
	client.Db.Create(model)
	return nil
}
