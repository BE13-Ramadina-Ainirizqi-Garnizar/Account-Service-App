package halaman

import (
	"database/sql"
	"fmt"
	"os"
)

func HalUtama(db *sql.DB) {
	fmt.Println("Selamat Datang !")
	fmt.Println("Silakan Pilih Menu :\n1. Register \n2. Login \n0. Keluar Program")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		HalRegister(db)

	case 2:
		HalLogin(db)

	case 0:
		fmt.Println("Terima kasih")
		os.Exit(1)
	}
}
