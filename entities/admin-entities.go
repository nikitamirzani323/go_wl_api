package entities

type Model_admin struct {
	Admin_username    string `json:"admin_username"`
	Admin_password    string `json:"admin_password"`
	Admin_idadminrule int    `json:"admin_idadminrule"`
	Admin_name        string `json:"admin_name"`
	Admin_statuslogin string `json:"admin_statuslogin"`
	Admin_lastlogin   string `json:"admin_lastlogin"`
	Admin_joindate    string `json:"admin_joindate"`
	Admin_ipaddres    string `json:"admin_ipaddres"`
}

type Controller_admin struct {
	Admin_username    string `json:"admin_username" validate:"required"`
	Admin_password    string `json:"admin_password" validate:"required"`
	Admin_idadminrule int    `json:"admin_idadminrule" validate:"required"`
	Admin_name        string `json:"admin_name" validate:"required"`
	Admin_statuslogin string `json:"admin_statuslogin" validate:"required"`
	Admin_lastlogin   string `json:"admin_lastlogin" validate:"required"`
	Admin_joindate    string `json:"admin_joindate" validate:"required"`
	Admin_ipaddres    string `json:"admin_ipaddres" validate:"required"`
}
