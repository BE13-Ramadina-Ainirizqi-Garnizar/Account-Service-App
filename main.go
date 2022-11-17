package main

import (
	"github.com/BE13-Ramadina-Ainirizqi-Garnizar/Account-Service-App/config"
	"github.com/BE13-Ramadina-Ainirizqi-Garnizar/Account-Service-App/halaman"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConnection := config.InitToDB()

	defer dbConnection.Close()

	halaman.HalUtama(dbConnection)

}
