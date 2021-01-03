package usermodel

type UserModel struct {
	IDUser   string `gorm:"primary_key;type:uuid;default:uuid_generate_v4();"`
	UserName string `gorm:"type:varchar(50);unique";`
	Password string `gorm:"type:varchar(60)";`
}

//Cambiamos el nombre de la tabla a singular
func (User UserModel) TableName() string {
	return "user_access"
}
