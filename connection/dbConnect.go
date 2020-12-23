package connection

import (
	"encoding/xml"
	"fmt"
	"gqlGenTutorial/models/settingsModels"
	"io/ioutil"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

//OpenConnection : abrimos conexion a la base de datos
func OpenConnection(database string) *gorm.DB {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path) // for example /home/user

	//Abrimos archivo
	xmlFile, err := os.Open(path + "/settings/BDSettings.xml")

	//Verificamos que no existan errores
	if err != nil {
		panic(err)
	}

	fmt.Println("Archivo BdSettings.xml abierto exitosamente")

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

	//Declaramos un logger
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold: time.Second,   // Slow SQL threshold
	// 		LogLevel:      logger.Silent, // Log level
	// 		Colorful:      true,          // Disable color
	// 	},
	// 	logger.Default.LogMode(logger.Silent),
	// )

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

	//Devolvemos la conexion abierta a la base de datos
	return db
}
