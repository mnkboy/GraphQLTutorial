package migratefunctions

import "gorm.io/gorm"

//AddForeignKey :
func AddForeignKey(db *gorm.DB) {
	db.Model(link).Association("user")
}
