package migratefunctions

import (
	"fmt"

	"gorm.io/gorm"
)

//CreateTables : creamos tablas si existen
func CreateTables(db *gorm.DB) {
	//Imprimimos la base de datos actual
	fmt.Println(db.Migrator().CurrentDatabase())
	//Comenzamos a crear tablas
	db.Migrator().CreateTable(user)
	db.Migrator().CreateTable(link)

}
