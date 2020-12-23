package migratefunctions

import (
	"fmt"

	"gorm.io/gorm"
)

//CreateTables : creamos tablas si existen
func DropTables(db *gorm.DB) {
	//Imprimimos la base de datos actual
	fmt.Println(db.Migrator().CurrentDatabase())
	//Comenzamos a tirar tablas
	db.Migrator().DropTable(link)
	db.Migrator().DropTable(user)
}
