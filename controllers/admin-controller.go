package controllers

import (
	"log"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/go_wl_api/entities"
	"github.com/nikitamirzani323/go_wl_api/helpers"
	"github.com/nikitamirzani323/go_wl_api/models"
)

const Fieldadmin_home_redis = "LISTADMIN"

func Adminhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_admin)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	var obj entities.Model_admin
	var arraobj []entities.Model_admin
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldadmin_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		admin_username, _ := jsonparser.GetString(value, "admin_username")
		admin_password, _ := jsonparser.GetString(value, "admin_password")
		admin_idadminrule, _ := jsonparser.GetInt(value, "admin_idadminrule")
		admin_name, _ := jsonparser.GetString(value, "admin_name")
		admin_statuslogin, _ := jsonparser.GetString(value, "admin_statuslogin")
		admin_lastlogin, _ := jsonparser.GetString(value, "admin_lastlogin")
		admin_joindate, _ := jsonparser.GetString(value, "admin_joindate")
		admin_ipaddres, _ := jsonparser.GetString(value, "admin_ipaddres")

		obj.Admin_username = admin_username
		obj.Admin_password = admin_password
		obj.Admin_idadminrule = int(admin_idadminrule)
		obj.Admin_name = admin_name
		obj.Admin_statuslogin = admin_statuslogin
		obj.Admin_lastlogin = admin_lastlogin
		obj.Admin_joindate = admin_joindate
		obj.Admin_ipaddres = admin_ipaddres
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_admin()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldadmin_home_redis, result, 60*time.Minute)
		log.Println("DOMAIN MYSQL")
		return c.JSON(result)
	} else {
		log.Println("DOMAIN CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
