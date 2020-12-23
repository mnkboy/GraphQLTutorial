package usermodel

type UserModel struct {
	IDUser string `gorm:"primary_key;type:uuid;default:uuid_generate_v4();"`
	Name   string `gorm:"type:varchar(150)";`
}

//Cambiamos el nombre de la tabla a singular
func (User UserModel) TableName() string {
	return "user"
}
