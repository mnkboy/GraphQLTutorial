package authentication

import (
	"fmt"
	"gqlGenTutorial/dataaccess/repositories/userrepositories"
	"gqlGenTutorial/models/usermodel"

	"gorm.io/gorm"
)

//Authenticate :
func Authenticate(model usermodel.UserModel, db *gorm.DB) (string, error) {

	models, _ := userrepositories.RetrieveUser(&model, db)

	if len(*models) > 1 || len(*models) == 0 {
		return "", fmt.Errorf("Wrong user name or password")
	}

	token, err := GenerateToken(model.UserName)

	if err != nil {
		return "", err
	}

	return token, nil
}
