package models

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/go_wl_api/config"
	"github.com/nikitamirzani323/go_wl_api/db"
	"github.com/nikitamirzani323/go_wl_api/entities"
	"github.com/nikitamirzani323/go_wl_api/helpers"
	"github.com/nleeper/goment"
)

func Fetch_admin() (helpers.Response, error) {
	var obj entities.Model_admin
	var arraobj []entities.Model_admin
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()

	var no int = 0
	sql_select := `SELECT 
			username , idadmin, name, statuslogin,  
			to_char(lastlogin,'YYYY-MM-DD HH24:mm:ss') as lastlogin, 
			to_char(joindate,'YYYY-MM-DD') as joindate, 
			ipaddress, timezone, 
			createadmin, to_char(createdateadmin,'YYYY-MM-DD HH24:mm:ss') as createdateadmin,  
			updateadmin, to_char(updatedateadmin,'YYYY-MM-DD HH24:mm:ss') as  updatedateadmin 
			FROM ` + config.DB_tbl_mst_admin + ` 
			ORDER BY username ASC 
		`
	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		no++
		var (
			username_db, idadmin_db, name_db, statuslogin_db, ipaddress_db, timezone_db                       string
			lastlogin_db, joindate_db, createadmin_db, createdateadmin_db, updateadmin_db, updatedateadmin_db string
		)

		err = row.Scan(
			&username_db, &idadmin_db, &name_db, &statuslogin_db,
			&lastlogin_db, &joindate_db, &ipaddress_db, &timezone_db,
			&createadmin_db, &createdateadmin_db, &updateadmin_db, &updatedateadmin_db)
		helpers.ErrorCheck(err)

		obj.Admin_username = username_db
		obj.Admin_idadminrule = idadmin_db
		obj.Admin_name = name_db
		obj.Admin_statuslogin = statuslogin_db
		obj.Admin_ipaddres = ipaddress_db
		obj.Admin_timezone = timezone_db
		obj.Admin_lastlogin = lastlogin_db
		obj.Admin_joindate = joindate_db
		obj.Admin_createadmin = createadmin_db
		obj.Admin_createdateadmin = createdateadmin_db
		obj.Admin_updateadmin = updateadmin_db
		obj.Admin_updatedateadmin = updatedateadmin_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()

	return res, nil
}

func Save_admin(admin, username, password, idadminrule, name, status, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(config.DB_tbl_mst_admin, "username", username)
		if !flag {
			sql_insert := `
				insert into
				` + config.DB_tbl_mst_admin + ` (
					username , password, idadmin, name, statuslogin, 
					joindate, ipaddress, timezone, createadmin, createdateadmin 
				) values (
					$1, $2, $3, $4, $5,
					$6, $7, $8, $9, $10 
				)
			`

			flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_mst_admin, "INSERT",
				username, password, idadminrule, name, status,
				tglnow.Format("YYYY-MM-DD"), "192.168.23.01", "Asia/Jakarta",
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if flag_insert {
				flag = true
				msg = "Succes"
				log.Println(msg_insert)
			} else {
				log.Println(msg_insert)
			}
		} else {
			msg = "Duplicate Entry"
		}
	} else {
		// sql_update := `
		// 		UPDATE
		// 		` + config.DB_tbl_mst_admin + `
		// 		SET password =?, statusdomain=?, tipedomain=?,
		// 		updatedomain=?, updatedatedomain=?
		// 		WHERE iddomain =?
		// 	`

		// flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_mst_admin, "UPDATE",
		// 	nmdomain, status, tipe,
		// 	admin,
		// 	tglnow.Format("YYYY-MM-DD HH:mm:ss"),
		// 	idrecord)

		// if flag_update {
		// 	flag = true
		// 	msg = "Succes"
		// 	log.Println(msg_update)
		// } else {
		// 	log.Println(msg_update)
		// }
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
