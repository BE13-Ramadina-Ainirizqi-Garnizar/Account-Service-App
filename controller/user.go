package controller

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BE13-Ramadina-Ainirizqi-Garnizar/Account-Service-App/entity"
)

func Register(db *sql.DB, newUser entity.User) (sql.Result, error) {
	if newUser.Username == "" || newUser.Password == "" || newUser.Nama == "" || newUser.Gender == "" || newUser.NoTelp == "" || newUser.Email == "" {
		log.Fatal("Input tidak boleh kosong")
	}

	var query = "insert into users(username, pass_word, nama, gender, no_telp, email) value (?,?,?,?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error preparing", errPrepare.Error())
	}

	result, errExec := statement.Exec(newUser.Username, newUser.Password, newUser.Nama, newUser.Gender, newUser.NoTelp, newUser.Email)
	if errExec != nil {
		log.Fatal("error executing", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("Register Berhasil")
		} else {
			fmt.Println("Register Gagal")
		}
	}
	return result, nil
}

func Login(db *sql.DB, NoTelp string, Password string) (string, error) {

	if NoTelp == "" || Password == "" {
		log.Fatal("No. HP & Password harus diisi")

	}

	query := fmt.Sprintf("select pass_word from users where no_telp = %s", NoTelp)
	result, err := db.Query(query)
	if err != nil {
		log.Fatal("error", err.Error())
	}

	var pass string
	for result.Next() {
		errScan := result.Scan(&pass)
		if errScan != nil {
			log.Fatal("error scan", errScan.Error())
		}
	}

	if pass != "" {
		if pass == Password {
			fmt.Println("Login Berhasil")
			return "", nil
		} else {
			log.Fatal("Password Salah")
			return "", err
		}
	} else {
		log.Fatal("Nomor Tidak Terdaftar")
		return "", err
	}
}

func ReadAccount(db *sql.DB, NoTelp string) ([]entity.User, error) {
	query := fmt.Sprintf("select username, nama, gender, email from users where no_telp = %s", NoTelp)
	result, err := db.Query(query)
	if err != nil {
		log.Fatal("error", err.Error())
	}

	var dataUser []entity.User
	for result.Next() {
		var userrow entity.User
		errScan := result.Scan(&userrow.Username, &userrow.Nama, &userrow.Gender, &userrow.Email)
		if errScan != nil {
			log.Fatal("error scan", errScan.Error())
		}

		dataUser = append(dataUser, userrow)
	}
	return dataUser, nil
}

func UpdateAcc(db *sql.DB, update entity.User, NoTelp string) (sql.Result, error) {
	query := fmt.Sprintf("Update users set username = ?, pass_word = ?, nama = ?, gender = ?, no_telp = ?, email = ? where no_telp = %s", NoTelp)
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error preparing", errPrepare.Error())
	}

	result, errExec := statement.Exec(update.Username, update.Password, update.Nama, update.Gender, update.NoTelp, update.Email)
	if errExec != nil {
		log.Fatal("error execute", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("Update Berhasil !")
		} else {
			fmt.Println("Update Gagal.")
		}
	}
	return result, nil
}

func Delete(db *sql.DB, NoTelp string) (sql.Result, error) {
	var query = "delete from users where no_telp = ?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error preparing", errPrepare.Error())
	}

	result, errExec := statement.Exec(NoTelp)
	if errExec != nil {
		log.Fatal("error executing", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("Delete Berhasil !")
		} else {
			fmt.Println("Delete Gagal.")
		}
	}
	return result, nil
}

func ProfilTeman(db *sql.DB, HpTeman string) ([]entity.User, error) {
	query := fmt.Sprintf("select username, nama, gender, email from users where no_telp=%s", HpTeman)
	result, err := db.Query(query)
	if err != nil {
		log.Fatal("error", err.Error())
	}

	var dataTeman []entity.User
	for result.Next() {
		var userrow entity.User
		errScan := result.Scan(&userrow.Username, &userrow.Nama, &userrow.Gender, &userrow.Email)
		if errScan != nil {
			log.Fatal("error scan", errScan.Error())
		}
		dataTeman = append(dataTeman, userrow)
	}
	return dataTeman, nil
}
