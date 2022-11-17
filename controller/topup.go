package controller

import (
	"database/sql"
	"fmt"

	"github.com/BE13-Ramadina-Ainirizqi-Garnizar/Account-Service-App/entity"
)

func TopUp(db *sql.DB, Nominal int, NoTelp string) (sql.Result, error) {
	query := fmt.Sprintf("select id from users where no_telp=%s", NoTelp)
	result := db.QueryRow(query)

	var id int
	errScan := result.Scan(&id)
	if errScan != nil {
		fmt.Println("error scan", errScan.Error())
		return nil, errScan
	}

	query2 := "insert into top_up(user_id, saldo_top_up) values(?,?)"
	statement, errPrepare := db.Prepare(query2)
	if errPrepare != nil {
		fmt.Println("error prepare", errPrepare.Error())
		return nil, errPrepare

	}
	result2, errExec := statement.Exec(id, Nominal)
	if errExec != nil {
		fmt.Println("error exec", errExec.Error())
		return nil, errExec
	} else {
		row, _ := result2.RowsAffected()
		if row > 0 {
			fmt.Println("Top Up Berhasil")

		} else {
			fmt.Println("Top Up Gagal")
		}
	}
	return result2, nil

}
func HistoryTopUp(db *sql.DB, NoTelp string) ([]entity.TopUp, error) {
	query := fmt.Sprintf("select id from users where no_telp = %s", NoTelp)
	result := db.QueryRow(query)

	var id int
	errScan1 := result.Scan(&id)
	if errScan1 != nil {
		fmt.Println("Nomor Telpon tidak terdaftar")
		return nil, errScan1
	}
	query2 := fmt.Sprintf("select t.id, s.nama, t.saldo_top_up, t.created_at from top_up t join users s on s.id = t.user_id where t.user_id = %d ", id)
	result2, errSelect := db.Query(query2)
	if errSelect != nil {
		fmt.Println("error select", errSelect.Error())
		return nil, errSelect
	}
	var dataHistory []entity.TopUp
	for result2.Next() {
		var userrows entity.TopUp
		errScan2 := result2.Scan(&userrows.IDTopUp, &userrows.Nama, &userrows.SaldoTopUp, &userrows.CreatedAt)
		if errScan2 != nil {
			fmt.Println("error scan", errScan2.Error())
			return nil, errScan2
		}

		dataHistory = append(dataHistory, userrows)
	}

	return dataHistory, nil
}
