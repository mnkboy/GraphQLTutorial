package linkmodel

import "gqlGenTutorial/models/usermodel"

type LinkModel struct {
	IDLink  string               `gorm:"primary_key;type:uuid;default:uuid_generate_v4();"`
	Title   string               `gorm:"type:varchar(150)";`
	Address string               `gorm:"type:varchar(150)";`
	UserID  *string              `gorm:"type:uuid;null;column:id_user;"`
	User    *usermodel.UserModel `gorm:"foreignKey:UserID;association_foreignkey:id_user;"`
}

//Cambiamos el nombre de la tabla a singular
func (Link LinkModel) TableName() string {
	return "link"
}
