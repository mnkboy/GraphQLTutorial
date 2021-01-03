package userrepositories

import (
	"gqlGenTutorial/models/usermodel"

	"gorm.io/gorm"
)

// //UserRepository :
// type UserRepository struct {
// }

//CreateUser es una funcion para crear un elemento en la base de datos
func CreateUser(model *usermodel.UserModel, db *gorm.DB) (bool, error) {

	er := db.Create(model).Error

	if er != nil {
		return false, er
	}

	return true, nil
}

//RetrieveUser es una funcion para crear un elemento en la base de datos
func RetrieveUser(model *usermodel.UserModel, db *gorm.DB) (*[]usermodel.UserModel, error) {

	models := []usermodel.UserModel{}

	db = db.Where("user_name", model.UserName)

	er := db.Find(&models).Error

	if er != nil {
		return &[]usermodel.UserModel{}, er
	}

	return &models, nil
}
