package halaman

import (
	"database/sql"
	"fmt"

	"github.com/BE13-Ramadina-Ainirizqi-Garnizar/Account-Service-App/controller"
)

func HalLogin(db *sql.DB) {
	var NoTelp string
	var Password string
	fmt.Println("Masukkan No. Telp : ")
	fmt.Scanln(&NoTelp)
	fmt.Println("Masukkan Password : ")
	fmt.Scanln(&Password)

	_, errLog := controller.Login(db, NoTelp, Password)
	if errLog != nil {
		fmt.Println(errLog)
		HalUtama(db)
	}
	HalMainMenu(db, NoTelp)
}
