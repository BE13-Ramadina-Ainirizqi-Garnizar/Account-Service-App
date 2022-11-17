package halaman

import (
	"database/sql"
	"fmt"

	"github.com/BE13-Ramadina-Ainirizqi-Garnizar/Account-Service-App/controller"
	"github.com/BE13-Ramadina-Ainirizqi-Garnizar/Account-Service-App/entity"
)

func HalRegister(db *sql.DB) {
	newUser := entity.User{}
	var pilGender int

	fmt.Println("Masukkan Username: ")
	fmt.Scanln(&newUser.Username)
	fmt.Println("Masukkan Password: ")
	fmt.Scanln(&newUser.Password)
	fmt.Println("Masukkan Nama: ")
	fmt.Scanln(&newUser.Nama)
	fmt.Println("Pilih Gender: 1. Laki-Laki 2. Perempuan")
	fmt.Scanln(&pilGender)
	fmt.Println("Masukkan No. Telp: ")
	fmt.Scanln(&newUser.NoTelp)
	fmt.Println("Masukkan E-mail: ")
	fmt.Scanln(&newUser.Email)

	switch pilGender {
	case 1:
		newUser.Gender = "L"
	case 2:
		newUser.Gender = "P"
	}

	_, errReg := controller.Register(db, newUser)
	if errReg != nil {
		fmt.Println(errReg.Error())
		HalUtama(db)
	} else {
		fmt.Println("Silakan Login.")
		HalLogin(db)
	}
}
