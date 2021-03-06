package linkrepositories

import (
	"gqlGenTutorial/models/linkmodel"

	"gorm.io/gorm"
)

// //LinkRepository :
// type LinkRepository struct {
// }

//CreateLink es una funcion para crear un elemento en la base de datos
func CreateLink(model *linkmodel.LinkModel, db *gorm.DB) (bool, error) {

	er := db.Create(model).Error

	if er != nil {
		return false, er
	}

	return true, nil

}

//DeleteLink es una funcion para eliminar un elemento en la base de datos
func DeleteLink(model *linkmodel.LinkModel, db *gorm.DB) (bool, error) {

	er := db.Delete(model).Error

	if er != nil {
		return false, er
	}

	return true, nil
}
