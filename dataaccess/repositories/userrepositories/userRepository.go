package userrepositories

import (
	"gqlGenTutorial/models/usermodel"

	"gorm.io/gorm"
)

//UserRepository :
type UserRepository struct {
}

//CreateUser es una funcion para crear un elemento en la base de datos
func (repository *UserRepository) CreateUser(model *usermodel.UserModel, db *gorm.DB) (bool, string) {

	er := db.Create(model).Error

	if er != nil {
		return false, "error"
	}

	return true, "success"
}
