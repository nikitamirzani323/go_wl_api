package models

import (
	"context"
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/nikitamirzani323/go_wl_api/config"
	"github.com/nikitamirzani323/go_wl_api/db"
	"github.com/nikitamirzani323/go_wl_api/helpers"
)

func CheckDB(table, field, value string) bool {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	sql_db := `SELECT 
					` + field + ` 
					FROM ` + table + ` 
					WHERE ` + field + ` = ? 
				`

	row := con.QueryRowContext(ctx, sql_db, value)
	switch e := row.Scan(&field); e {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
		flag = false
	case nil:
		flag = true
	default:
		panic(e)
	}
	return flag
}
func CheckDBTwoField(table, field_1, value_1, field_2, value_2 string) bool {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	sql_db := `SELECT 
					` + field_1 + ` 
					FROM ` + table + ` 
					WHERE ` + field_1 + ` = ? 
					AND ` + field_2 + ` = ? 
				`
	row := con.QueryRowContext(ctx, sql_db, value_1, value_2)
	switch e := row.Scan(&field_1); e {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
		flag = false
	case nil:
		flag = true
	default:
		panic(e)
	}
	return flag
}
func Get_counter(field_column string) int {
	con := db.CreateCon()
	ctx := context.Background()
	idrecord_counter := 0
	sqlcounter := `SELECT 
					counter 
					FROM ` + config.DB_tbl_counter + ` 
					WHERE nmcounter = $1 
				`
	var counter int = 0
	row := con.QueryRowContext(ctx, sqlcounter, field_column)
	switch e := row.Scan(&counter); e {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
	case nil:

	default:
		panic(e)
	}
	if counter > 0 {
		idrecord_counter = int(counter) + 1
		stmt, e := con.PrepareContext(ctx, "UPDATE "+config.DB_tbl_counter+" SET counter=? WHERE nmcounter=? ")
		helpers.ErrorCheck(e)
		res, e := stmt.ExecContext(ctx, idrecord_counter, field_column)
		helpers.ErrorCheck(e)
		a, e := res.RowsAffected()
		helpers.ErrorCheck(e)
		if a > 0 {
			log.Println("UPDATE TBL_COUNTER")
		}
	} else {
		stmt, e := con.PrepareContext(ctx, "insert into "+config.DB_tbl_counter+" (nmcounter, counter) values ($1, $2)")
		helpers.ErrorCheck(e)
		res, e := stmt.ExecContext(ctx, field_column, 1)
		helpers.ErrorCheck(e)
		id, e := res.RowsAffected()
		helpers.ErrorCheck(e)
		log.Println("Insert id", id)
		log.Println("NEW TBL_COUNTER")
		idrecord_counter = 1
	}
	return idrecord_counter
}
func Get_mappingdatabase(company string) (string, string, string, string) {
	tbl_mst_company := "db_tot_" + strings.ToLower(company)
	tbl_trx_keluarantogel := "db_tot_" + strings.ToLower(company) + ".tbl_trx_keluarantogel"
	tbl_trx_keluarantogel_detail := "db_tot_" + strings.ToLower(company) + ".tbl_trx_keluarantogel_detail"
	tbl_trx_keluarantogel_member := "db_tot_" + strings.ToLower(company) + ".tbl_trx_keluarantogel_member"

	return tbl_trx_keluarantogel, tbl_trx_keluarantogel_detail, tbl_trx_keluarantogel_member, tbl_mst_company
}

func GenerateNewTable(company string) (string, string, string, string) {
	tbl_trx_keluaran := `tbl_trx_keluarantogel (
		idtrxkeluaran bigint(20) NOT NULL,
		idcomppasaran int(11) NOT NULL,
		idcompany varchar(10) NOT NULL,
		yearmonth varchar(20) NOT NULL,
		keluaranperiode int(11) NOT NULL,
		datekeluaran date NOT NULL,
		keluarantogel varchar(4) NOT NULL,
		prize2 varchar(4) NOT NULL,
		prize3 varchar(4) NOT NULL,
		total_member int(11) NOT NULL,
		total_bet double NOT NULL,
		total_outstanding double NOT NULL,
		total_win double NOT NULL,
		total_lose double NOT NULL,
		total_referal double NOT NULL,
		total_buangan double NOT NULL,
		total_cancel double NOT NULL,
		winlose double NOT NULL,
		revisi tinyint(1) NOT NULL,
		noterevisi varchar(150) NOT NULL,
		createkeluaran varchar(70) NOT NULL,
		createdatekeluaran datetime NOT NULL,
		updatekeluaran varchar(70) NOT NULL,
		updatedatekeluaran datetime DEFAULT NULL,
		PRIMARY KEY (idtrxkeluaran)
	)`
	tbl_trx_keluaran_detail := `tbl_trx_keluarantogel_detail (
	  idtrxkeluarandetail bigint(20) NOT NULL,
	  idtrxkeluaran bigint(20) NOT NULL,
	  idcompany varchar(10) NOT NULL,
	  datetimedetail datetime NOT NULL,
	  ipaddress varchar(30) NOT NULL,
	  username varchar(50) NOT NULL,
	  typegame varchar(50) NOT NULL,
	  nomortogel varchar(50) NOT NULL,
	  posisitogel varchar(20) NOT NULL,
	  bet double NOT NULL,
	  diskon double NOT NULL,
	  win double NOT NULL,
	  winhasil int(11) NOT NULL,
	  cancelbet int(11) NOT NULL,
	  kei double NOT NULL,
	  upline varchar(50) NOT NULL,
	  upline_ref double NOT NULL,
	  type_ref char(2) NOT NULL,
	  browsertogel varchar(50) NOT NULL,
	  devicetogel varchar(50) NOT NULL,
	  statuskeluarandetail varchar(10) CHARACTER SET latin1 COLLATE latin1_swedish_ci DEFAULT NULL,
	  senddata varchar(20) NOT NULL,
	  senddatacreatedate datetime NOT NULL,
	  updatedata varchar(20) NOT NULL,
	  updatedatacreatedate datetime NOT NULL,
	  createkeluarandetail varchar(70) NOT NULL,
	  createdatekeluarandetail datetime NOT NULL,
	  updatekeluarandetail varchar(70) NOT NULL,
	  updatedatekeluarandetail datetime NOT NULL,
	  PRIMARY KEY (idtrxkeluarandetail)
	)`

	tbl_trx_keluaran_member := `tbl_trx_keluarantogel_member (
	  idkeluaranmember bigint(20) NOT NULL,
	  idtrxkeluaran bigint(20) NOT NULL,
	  idcompany varchar(10) NOT NULL,
	  username varchar(50) NOT NULL,
	  totalbet double NOT NULL,
	  totalbayar double NOT NULL,
	  totalreferal double NOT NULL,
	  totalwin double NOT NULL,
	  totalcancel double NOT NULL,
	  createkeluaranmember varchar(70) NOT NULL,
	  createdatekeluaranmember datetime NOT NULL,
	  updatekeluaranmember varchar(70) NOT NULL,
	  updatedatekeluaranmember datetime DEFAULT NULL,
	  PRIMARY KEY (idkeluaranmember)
	)`

	view_tbl_keluaran := "CREATE OR REPLACE  ALGORITHM = UNDEFINED VIEW db_tot_" + strings.ToLower(company) + ".client_view_invoice_" + company + " AS " +
		" select " +
		"	A.idtrxkeluarandetail AS idtrxkeluarandetail, " +
		"	A.idtrxkeluaran AS idtrxkeluaran, " +
		"	A.idcompany AS idcompany, " +
		"	C.idpasarantogel AS idpasarantogel, " +
		"	D.nmpasarantogel AS nmpasarantogel, " +
		"	A.datetimedetail AS datetimedetail, " +
		"	A.ipaddress AS ipaddress, " +
		"	A.devicetogel AS devicetogel, " +
		"	A.username AS username, " +
		"	A.typegame AS typegame, " +
		"	A.nomortogel AS nomortogel, " +
		"	A.bet AS bet, " +
		"	A.diskon AS diskon, " +
		"	A.kei AS kei, " +
		"	A.win AS win, " +
		"	A.winhasil AS winhasil, " +
		"	A.statuskeluarandetail AS statuskeluarandetail, " +
		"	B.idcomppasaran AS idcomppasaran, " +
		"	B.keluaranperiode AS keluaranperiode, " +
		"	B.datekeluaran AS datekeluaran, " +
		"	B.keluarantogel AS keluarantogel " +
		"from " +
		"	(((db_tot_" + strings.ToLower(company) + ".tbl_trx_keluarantogel_detail A " +
		"join db_tot_" + strings.ToLower(company) + ".tbl_trx_keluarantogel B on " +
		"	((B.idtrxkeluaran = A.idtrxkeluaran))) " +
		"join db_tot.tbl_mst_company_game_pasaran C on " +
		"	((C.idcomppasaran = B.idcomppasaran))) " +
		"join db_tot.tbl_mst_pasaran_togel D on " +
		"	((D.idpasarantogel = C.idpasarantogel)))"

	return tbl_trx_keluaran, tbl_trx_keluaran_detail, tbl_trx_keluaran_member, view_tbl_keluaran
}

func CreateNewCompanyDB(dbName, company string, db *sql.DB) (string, error) {
	_, err := db.Exec("CREATE DATABASE  IF NOT EXISTS " + dbName)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + dbName)
	if err != nil {
		panic(err)
	}

	tbl_kel, tbl_kel_det, tbl_kel_member, view_keluaran := GenerateNewTable(company)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS db_tot_" + strings.ToLower(company) + "." + tbl_kel)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS db_tot_" + strings.ToLower(company) + "." + tbl_kel_det)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE db_tot_" + strings.ToLower(company) + "." + tbl_kel_member)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(view_keluaran)
	if err != nil {
		panic(err)
	}
	return "ok", nil
}
func Exec_SQL(sql, table, action string, args ...interface{}) (bool, string) {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	msg := ""
	stmt_exec, e_exec := con.PrepareContext(ctx, sql)
	helpers.ErrorCheck(e_exec)
	defer stmt_exec.Close()
	rec_exec, e_exec := stmt_exec.ExecContext(ctx, args...)

	helpers.ErrorCheck(e_exec)
	exec, e := rec_exec.RowsAffected()
	helpers.ErrorCheck(e)
	if exec > 0 {
		flag = true
		msg = "Data " + table + " Berhasil di " + action
	} else {
		msg = "Data " + table + " Failed di " + action
	}
	return flag, msg
}
