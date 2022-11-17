package controller

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BE13-Ramadina-Ainirizqi-Garnizar/Account-Service-App/config"
	"github.com/BE13-Ramadina-Ainirizqi-Garnizar/Account-Service-App/entity"
)

func TransferDana(db *sql.DB, NoTelp string, TelpTeman string, Nominal int) (sql.Result, error) {
	dbConnection := config.InitToDB()

	defer dbConnection.Close()

	saldoNow := CekSaldo(dbConnection, NoTelp)
	if saldoNow < Nominal {
		log.Fatal("Dana anda tidak mencukupi.")
	}

	query := fmt.Sprintf("select id from users where no_telp = %s", NoTelp)
	result := db.QueryRow(query)

	var id int
	errScan := result.Scan(&id)
	if errScan != nil {
		log.Fatal("error scan", errScan.Error())
	}

	queryTeman := fmt.Sprintf("select id from users where no_telp = %s", TelpTeman)
	resultTeman := db.QueryRow(queryTeman)

	var idTeman int
	errScan3 := resultTeman.Scan(&idTeman)
	if errScan3 != nil {
		log.Fatal("Nomor yang anda tuju tidak terdaftar")
	}

	queryTransfer := "insert into transfer(user_id_pengirim, user_id_penerima, saldo_transfer) values(?,?,?)"
	statement, errPrepare := db.Prepare(queryTransfer)
	if errPrepare != nil {
		log.Fatal("error preparing", errPrepare.Error())
	}

	result2, errExec := statement.Exec(id, idTeman, Nominal)
	if errExec != nil {
		log.Fatal("error execute", errExec.Error())
	} else {
		row, _ := result2.RowsAffected()
		if row > 0 {
			fmt.Println("Transfer Berhasil !")
		} else {
			fmt.Println("Transfer Gagal.")
		}
	}
	return result2, nil
}

func HistoryTransfer(db *sql.DB, NoTelp string) ([]entity.Transfer, error) {
	query := fmt.Sprintf("select id from users where no_telp = %s", NoTelp)
	result := db.QueryRow(query)

	var id int
	errScan1 := result.Scan(&id)
	if errScan1 != nil {
		log.Fatal("Nomor Telpon tidak terdaftar")
	}

	query2 := fmt.Sprintf("select t.id, s.nama as pengirim, r.nama as penerima, t.saldo_transfer, t.created_at from transfer t join users s on s.id = t.user_id_pengirim join users r on r.id = t.user_id_penerima where t.user_id_pengirim = %d or t.user_id_penerima = %d;", id, id)
	result2, errSelect := db.Query(query2)
	if errSelect != nil {
		log.Fatal("error select", errSelect.Error())
	}

	var dataHistory []entity.Transfer
	for result2.Next() {
		var userrows entity.Transfer
		errScan2 := result2.Scan(&userrows.IDTransfer, &userrows.NamaPengirim, &userrows.NamaPenerima, &userrows.SaldoTransfer, &userrows.CreatedAt)
		if errScan2 != nil {
			log.Fatal("error scan", errScan2.Error())
		}

		dataHistory = append(dataHistory, userrows)
	}

	return dataHistory, nil
}
