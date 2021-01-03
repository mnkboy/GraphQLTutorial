package connection

import (
	"encoding/xml"
	"fmt"
	settingsModels "gqlGenTutorial/models/settingsmodels"
	"gqlGenTutorial/settings"
	"io/ioutil"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

//OpenConnection : abrimos conexion a la base de datos
func OpenConnection(database string) *gorm.DB {
	//Capturamos el path original
	originalPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	//Imprimimos el path original
	fmt.Println(originalPath)

	//Nos cambiamos a settings
	os.Chdir(settings.SettingsPath)

	//Imprimimos el path de settings
	fmt.Println(settings.SettingsPath)

	current, err := os.Getwd()
	fmt.Println(current)

	//Abrimos archivo
	xmlFile, err := os.Open(settings.BdSettingsFile)

	//Verificamos que no existan errores
	if err != nil {
		panic(err)
	}

	fmt.Println("Archivo " + settings.BdSettingsFile + " abierto exitosamente")

	//Por defecto siempre cerramos el archivo
	defer xmlFile.Close()

	//Leemos en un byte array el contenido del archivo
	dbSettingsByteArray, _ := ioutil.ReadAll(xmlFile)

	//Declaramos variables
	DBs := settingsModels.DBSettingsModel{}

	//volcamos el contenido del byte array en las estructuras
	xml.Unmarshal(dbSettingsByteArray, &DBs)

	//Creamos un item de tipo DB
	DBItem := &settingsModels.DBModel{}

	//Imprimimos las configuraciones
	for _, dbItem := range DBs.DataBase {
		if dbItem.Name == database {
			DBItem = &dbItem
			break
		}
	}

	//Creamos la cadena de conexion
	strcon := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%s",
		DBItem.Server,
		DBItem.Port,
		DBItem.User,
		DBItem.Database,
		DBItem.Password,
		DBItem.SslMode,
	)

	//Abrimos una conexion a la base de datos
	db, err := gorm.Open(postgres.Open(strcon), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	//Si el error es distinto a nil tiramos el error
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Conexion abierta exitosamente a base dedatos: " + DBItem.Name)
	}

	//Creamos la extension para uuid
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	//Regresamos al directorio original
	os.Chdir(originalPath)

	//Devolvemos la conexion abierta a la base de datos
	return db
}
