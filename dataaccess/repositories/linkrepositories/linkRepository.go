package linkrepositories

import (
	"gqlGenTutorial/models/linkmodel"

	"gorm.io/gorm"
)

// //LinkRepository :
// type LinkRepository struct {
// }

//CreateLink es una funcion para crear un elemento en la base de datos
func CreateLink(model *linkmodel.LinkModel, db *gorm.DB) (bool, string) {

	er := db.Create(model).Error

	if er != nil {
		return false, "error"
	}

	return true, "success"

}
