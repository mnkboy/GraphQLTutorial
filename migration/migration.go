package main

import (
	"gqlGenTutorial/connection"
	"gqlGenTutorial/migration/migratefunctions"
	"gqlGenTutorial/settings"
)

func main() {
	// Pedimos una conexion a la base de datos POSTGRES
	db := connection.OpenConnection(settings.PostgresDB)

	migratefunctions.DropTables(db)
	migratefunctions.CreateTables(db)
	migratefunctions.AddForeignKey(db)

	// migratefunctions.LoadData(db)
}
