package models

import (
	"context"
	"log"
	"strconv"
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
			username , idadminrule,name,statuslogin,  
			COALESCE(lastlogin,""), COALESCE(joindate,"")  
			FROM ` + config.DB_tbl_mst_admin + ` 
			ORDER BY username ASC 
		`
	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		no++
		var (
			idadminrule_db                       int
			username_db, name_db, statuslogin_db string
			lastlogin_db, joindate_db            string
		)

		err = row.Scan(
			&username_db, &idadminrule_db, &name_db, &statuslogin_db,
			&lastlogin_db, &joindate_db)
		helpers.ErrorCheck(err)

		obj.Admin_username = username_db
		obj.Admin_idadminrule = idadminrule_db
		obj.Admin_name = name_db
		obj.Admin_statuslogin = statuslogin_db
		obj.Admin_lastlogin = lastlogin_db
		obj.Admin_joindate = joindate_db
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
func Save_domain(admin, username, password, name, status, sData string, idrecord, idadminrule int) (helpers.Response, error) {
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
					username , password, idadminrule, name, statuslogin
				) values (
					$1, $2, $3, $4, $5,
				)
			`
			field_column := config.DB_tbl_mst_admin + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_column)
			flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_mst_admin, "INSERT",
				tglnow.Format("YY")+strconv.Itoa(idrecord_counter), username, idadminrule, name,
				status)

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
		sql_update := `
				UPDATE 
				` + config.DB_tbl_mst_domain + `  
				SET nmdomain =?, statusdomain=?, tipedomain=?,
				updatedomain=?, updatedatedomain=? 
				WHERE iddomain =? 
			`

		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_mst_domain, "UPDATE",
			nmdomain, status, tipe,
			admin,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idrecord)

		if flag_update {
			flag = true
			msg = "Succes"
			log.Println(msg_update)
		} else {
			log.Println(msg_update)
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	}

	return res, nil
}
