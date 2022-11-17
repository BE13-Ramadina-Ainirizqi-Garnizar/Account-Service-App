package main

import (
	"fmt"
	"log"

	"github.com/BE13-Ramadina-Ainirizqi-Garnizar/Account-Service-App/config"
	"github.com/BE13-Ramadina-Ainirizqi-Garnizar/Account-Service-App/controller"
	"github.com/BE13-Ramadina-Ainirizqi-Garnizar/Account-Service-App/entity"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConnection := config.InitToDB()

	defer dbConnection.Close()

	fmt.Println("Selamat Datang !")
	fmt.Println("Silakan Pilih Menu :\n1. Register \n2. Login")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
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

		controller.Register(dbConnection, newUser)

	case 2:

		var NoTelp string
		var Password string
		fmt.Println("Masukkan No. Telp : ")
		fmt.Scanln(&NoTelp)
		fmt.Println("Masukkan Password : ")
		fmt.Scanln(&Password)
		controller.Login(dbConnection, NoTelp, Password)

		fmt.Println("Silakan Pilih Menu :\n1. Lihat Profil \n2. Update Profil \n3. Delete Profil \n4. Lihat Profil Teman \n5. Cek Saldo \n6. Transfer Saldo \n7. Riwayat Transfer \n8. Top Up \n9. History Top up")
		var pilihan2 int
		fmt.Scanln(&pilihan2)

		switch pilihan2 {
		case 1:
			dataUser, errReadAcc := controller.ReadAccount(dbConnection, NoTelp)
			if errReadAcc != nil {
				log.Fatal("Error Read Account")
			}

			for _, value := range dataUser {
				fmt.Printf("Username: %s, Nama: %s, Gender: %s, Email: %s", value.Username, value.Nama, value.Gender, value.Email)
			}

		case 2:
			updateUser := entity.User{}
			var pilGender int

			fmt.Println("Masukkan Username: ")
			fmt.Scanln(&updateUser.Username)
			fmt.Println("Masukkan Password: ")
			fmt.Scanln(&updateUser.Password)
			fmt.Println("Masukkan Nama: ")
			fmt.Scanln(&updateUser.Nama)
			fmt.Println("Pilih Gender: 1. Laki-Laki 2. Perempuan")
			fmt.Scanln(&pilGender)
			fmt.Println("Masukkan No. Telp: ")
			fmt.Scanln(&updateUser.NoTelp)
			fmt.Println("Masukkan E-mail: ")
			fmt.Scanln(&updateUser.Email)

			switch pilGender {
			case 1:
				updateUser.Gender = "L"
			case 2:
				updateUser.Gender = "P"
			}

			controller.UpdateAcc(dbConnection, updateUser, NoTelp)

		case 3:
			var pilihan3 int
			fmt.Println("Apakah anda yakin akan menghapus akun? \n1.Ya \n2.Tidak")
			fmt.Scanln(&pilihan3)

			switch pilihan3 {
			case 1:
				controller.Delete(dbConnection, NoTelp)

			case 2:
				main()
			}

		case 4:
			var HpTeman string
			fmt.Println("Masukan No. HP Teman: ")
			fmt.Scanln(&HpTeman)

			dataTeman, errReadTeman := controller.ProfilTeman(dbConnection, HpTeman)
			if errReadTeman != nil {
				log.Fatal("error read teman", errReadTeman.Error())
			}

			for _, value := range dataTeman {
				fmt.Printf("Username: %s Nama: %s Gender: %s Email: %s", value.Username, value.Nama, value.Gender, value.Email)
			}

		case 5:
			dataSaldo := controller.CekSaldo(dbConnection, NoTelp)

			fmt.Println("Rp.", dataSaldo)

		case 6:
			var NoTeman string
			var Nominal int
			fmt.Println("Masukkan No. Telp tujuan transfer: ")
			fmt.Scanln(&NoTeman)
			fmt.Println("Masukkan Nominal Transfer: ")
			fmt.Scanln(&Nominal)

			controller.TransferDana(dbConnection, NoTelp, NoTeman, Nominal)
			controller.SaldoBertambah(dbConnection, NoTeman, Nominal)
			controller.SaldoBerkurang(dbConnection, NoTelp, Nominal)

		case 7:
			dataRiwayat, errDataRiwayat := controller.HistoryTransfer(dbConnection, NoTelp)
			if errDataRiwayat != nil {
				log.Fatal("error get all data")

			}

			for _, value := range dataRiwayat {
				fmt.Printf("ID Transaksi = %d Nama Pengirim = %s Nama Penerima = %s Nominal Transfer = Rp. %d Created at = %s\n", value.IDTransfer, value.NamaPengirim, value.NamaPenerima, value.SaldoTransfer, value.CreatedAt)
			}
		case 8:
			var Nominal int
			fmt.Println("Masukan Nominal: ")
			fmt.Scanln(&Nominal)
			controller.TopUp(dbConnection, Nominal, NoTelp)
			controller.SaldoBertambah(dbConnection, NoTelp, Nominal)

		case 9:
			dataRiwayat, errDataRiwayat := controller.HistoryTopUp(dbConnection, NoTelp)
			if errDataRiwayat != nil {
				log.Fatal("error get all data")

			}

			for _, value := range dataRiwayat {
				fmt.Printf("ID Transaksi = %d Nama = %s  Nominal Top up = Rp. %d Created at = %s\n", value.IDTopUp, value.Nama, value.SaldoTopUp, value.CreatedAt)
			}
		}

	}

}
