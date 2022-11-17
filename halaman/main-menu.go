package halaman

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/BE13-Ramadina-Ainirizqi-Garnizar/Account-Service-App/controller"
	"github.com/BE13-Ramadina-Ainirizqi-Garnizar/Account-Service-App/entity"
)

func HalMainMenu(db *sql.DB, NoTelp string) {

	fmt.Println("Silakan Pilih Menu :\n1. Lihat Profil \n2. Update Profil \n3. Delete Profil \n4. Lihat Profil Teman \n5. Cek Saldo \n6. Transfer Saldo \n7. Riwayat Transfer \n8. Top Up \n9. History Top up \n0. Keluar Program")
	var pilihan2 int
	fmt.Scanln(&pilihan2)

	switch pilihan2 {
	case 1:
		dataUser, errReadAcc := controller.ReadAccount(db, NoTelp)
		if errReadAcc != nil {
			fmt.Println("Error Read Account")
			HalMainMenu(db, NoTelp)
		}

		for _, value := range dataUser {
			fmt.Printf("Username: %s, Nama: %s, Gender: %s, Email: %s\n\n\n", value.Username, value.Nama, value.Gender, value.Email)
			var kembali int
			fmt.Println("Masukkan 0 untuk kembali ke main menu: ")
			fmt.Scanln(&kembali)

			switch kembali {
			default:
				HalMainMenu(db, NoTelp)
			}

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

		_, errUpdate := controller.UpdateAcc(db, updateUser, NoTelp)
		if errUpdate != nil {
			fmt.Println(errUpdate.Error())
			HalMainMenu(db, NoTelp)
		} else {
			HalMainMenu(db, NoTelp)
		}

	case 3:
		var pilihan3 int
		fmt.Println("Apakah anda yakin akan menghapus akun? \n1.Ya \n2.Tidak")
		fmt.Scanln(&pilihan3)

		switch pilihan3 {
		case 1:
			_, errDelete := controller.Delete(db, NoTelp)
			if errDelete != nil {
				fmt.Println(errDelete.Error())
				HalMainMenu(db, NoTelp)
			} else {
				HalUtama(db)
			}

		case 2:
			HalMainMenu(db, NoTelp)
		}

	case 4:
		var HpTeman string
		fmt.Println("Masukan No. HP Teman: ")
		fmt.Scanln(&HpTeman)

		dataTeman, errReadTeman := controller.ProfilTeman(db, HpTeman)
		if errReadTeman != nil {
			fmt.Println("error read teman", errReadTeman.Error())
			HalMainMenu(db, NoTelp)
		} else {

			for _, value := range dataTeman {
				fmt.Printf("Username: %s Nama: %s Gender: %s Email: %s\n\n\n", value.Username, value.Nama, value.Gender, value.Email)

			}
			var kembali int
			fmt.Println("Masukkan 0 untuk kembali ke main menu: ")
			fmt.Scanln(&kembali)

			switch kembali {
			default:
				HalMainMenu(db, NoTelp)
			}
		}

	case 5:
		dataSaldo, errReadSaldo := controller.CekSaldo(db, NoTelp)
		if errReadSaldo != nil {
			fmt.Println(errReadSaldo.Error())
			HalMainMenu(db, NoTelp)
		} else {
			fmt.Printf("Rp. %d\n\n\n", dataSaldo)
			var kembali int
			fmt.Println("Masukkan 0 untuk kembali ke main menu: ")
			fmt.Scanln(&kembali)

			switch kembali {
			default:
				HalMainMenu(db, NoTelp)
			}

		}

	case 6:
		var NoTeman string
		var Nominal int
		fmt.Println("Masukkan No. Telp tujuan transfer: ")
		fmt.Scanln(&NoTeman)
		fmt.Println("Masukkan Nominal Transfer: ")
		fmt.Scanln(&Nominal)

		_, errTransfer := controller.TransferDana(db, NoTelp, NoTeman, Nominal)
		if errTransfer != nil {
			fmt.Println(errTransfer.Error())
			HalMainMenu(db, NoTelp)
		} else {
			_, errTambah := controller.SaldoBertambah(db, NoTeman, Nominal)
			if errTambah != nil {
				fmt.Println(errTambah.Error())
				HalMainMenu(db, NoTelp)
			}

			_, errKurang := controller.SaldoBerkurang(db, NoTelp, Nominal)
			if errKurang != nil {
				fmt.Println(errKurang.Error())
				HalMainMenu(db, NoTelp)
			}

			var kembali int
			fmt.Println("Masukkan 0 untuk kembali ke main menu: ")
			fmt.Scanln(&kembali)

			switch kembali {
			default:
				HalMainMenu(db, NoTelp)
			}
		}

	case 7:
		dataRiwayat, errHistoryTransfer := controller.HistoryTransfer(db, NoTelp)
		if errHistoryTransfer != nil {
			HalMainMenu(db, NoTelp)
		} else {
			for _, value := range dataRiwayat {
				fmt.Printf("ID Transaksi = %d Nama Pengirim = %s Nama Penerima = %s Nominal Transfer = Rp. %d Created at = %s\n", value.IDTransfer, value.NamaPengirim, value.NamaPenerima, value.SaldoTransfer, value.CreatedAt)

			}
			var kembali int
			fmt.Println("Masukkan 0 untuk kembali ke main menu: ")
			fmt.Scanln(&kembali)

			switch kembali {
			default:
				HalMainMenu(db, NoTelp)
			}
		}
	case 8:
		var Nominal int
		fmt.Println("Masukan Nominal: ")
		fmt.Scanln(&Nominal)
		_, errTopUp := controller.TopUp(db, Nominal, NoTelp)
		if errTopUp != nil {
			fmt.Println(errTopUp.Error())
			HalMainMenu(db, NoTelp)
		} else {
			controller.SaldoBertambah(db, NoTelp, Nominal)
			HalMainMenu(db, NoTelp)
		}

	case 9:
		dataRiwayat, errDataRiwayat := controller.HistoryTopUp(db, NoTelp)
		if errDataRiwayat != nil {
			fmt.Println("Tidak ada data untuk ditampilkan")
			HalMainMenu(db, NoTelp)
		} else {

			for _, value := range dataRiwayat {
				fmt.Printf("ID Transaksi = %d Nama = %s  Nominal Top up = Rp. %d Created at = %s\n", value.IDTopUp, value.Nama, value.SaldoTopUp, value.CreatedAt)
			}
			var kembali int
			fmt.Println("Masukkan 0 untuk kembali ke main menu: ")
			fmt.Scanln(&kembali)

			switch kembali {
			default:
				HalMainMenu(db, NoTelp)
			}
		}
	case 0:
		fmt.Println("Terima kasih")
		os.Exit(1)
	}
}
