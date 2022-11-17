package controller

import (
	"database/sql"
	"fmt"
	"log"
)

func CekSaldo(db *sql.DB, NoTelp string) (int, error) {
	query := fmt.Sprintf("select id from users where no_telp = %s", NoTelp)
	result := db.QueryRow(query)

	var id int
	errScan := result.Scan(&id)
	if errScan != nil {
		fmt.Println("error scan", errScan.Error())
		return 0, errScan
	}

	querySaldo := fmt.Sprintf("select saldo from dompet where user_id = %d", id)
	result2 := db.QueryRow(querySaldo)
	var saldo int
	errScan2 := result2.Scan(&saldo)
	if errScan2 != nil {
		fmt.Println("error scan", errScan2.Error())
		return 0, errScan2
	}

	return saldo, nil
}

func SaldoBerkurang(db *sql.DB, NoTelp string, Nominal int) (sql.Result, error) {
	query := fmt.Sprintf("select id from users where no_telp = %s", NoTelp)
	result := db.QueryRow(query)

	var id int
	errScan := result.Scan(&id)
	if errScan != nil {
		fmt.Println("error scan", errScan.Error())
		return nil, errScan
	}

	querySaldo := fmt.Sprintf("select saldo from dompet where user_id = %d", id)
	result2 := db.QueryRow(querySaldo)
	var saldo int
	errScan2 := result2.Scan(&saldo)
	if errScan2 != nil {
		fmt.Println("error scan", errScan2.Error())
		return nil, errScan2
	}

	SaldoMin := saldo - Nominal

	querySaldoMin := "update dompet set saldo = ? where user_id = ?"
	statement, errPrepare := db.Prepare(querySaldoMin)
	if errPrepare != nil {
		fmt.Println("error preparing", errPrepare.Error())
		return nil, errPrepare
	}

	result3, errExec := statement.Exec(SaldoMin, id)
	if errExec != nil {
		fmt.Println("error executing", errExec.Error())
		return nil, errExec
	}
	return result3, nil
}

func SaldoBertambah(db *sql.DB, NoTelp string, Nominal int) (sql.Result, error) {
	query := fmt.Sprintf("select id from users where no_telp = %s", NoTelp)
	result := db.QueryRow(query)

	var id int
	errScan := result.Scan(&id)
	if errScan != nil {
		log.Fatal("error scan", errScan.Error())
	}

	querySaldo := fmt.Sprintf("select saldo from dompet where user_id = %d", id)
	result2 := db.QueryRow(querySaldo)
	var saldo int
	errScan2 := result2.Scan(&saldo)
	if errScan2 != nil {
		log.Fatal("error scan", errScan2.Error())
	}

	SaldoPlus := saldo + Nominal

	querySaldoMin := "update dompet set saldo = ? where user_id = ?"
	statement, errPrepare := db.Prepare(querySaldoMin)
	if errPrepare != nil {
		log.Fatal("error preparing", errPrepare.Error())
	}

	result3, errExec := statement.Exec(SaldoPlus, id)
	if errExec != nil {
		log.Fatal("error executing", errExec.Error())
	}
	return result3, nil
}
