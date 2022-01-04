package entities

type Model_admin struct {
	Admin_username        string `json:"admin_username"`
	Admin_idadminrule     string `json:"admin_idadminrule"`
	Admin_name            string `json:"admin_name"`
	Admin_statuslogin     string `json:"admin_statuslogin"`
	Admin_lastlogin       string `json:"admin_lastlogin"`
	Admin_joindate        string `json:"admin_joindate"`
	Admin_ipaddres        string `json:"admin_ipaddres"`
	Admin_timezone        string `json:"admin_timezone"`
	Admin_createadmin     string `json:"admin_createadmin"`
	Admin_createdateadmin string `json:"admin_createdateadmin"`
	Admin_updateadmin     string `json:"admin_updateadmin"`
	Admin_updatedateadmin string `json:"admin_updatedateadmin"`
}

type Controller_admin struct {
	Master string `json:"master" validate:"required"`
}
type Controller_adminsave struct {
	Admin_username        string `json:"admin_username" validate:"required"`
	Admin_idadminrule     string `json:"admin_idadminrule" validate:"required"`
	Admin_name            string `json:"admin_name" validate:"required"`
	Admin_statuslogin     string `json:"admin_statuslogin" validate:"required"`
	Admin_lastlogin       string `json:"admin_lastlogin"`
	Admin_joindate        string `json:"admin_joindate"`
	Admin_ipaddres        string `json:"admin_ipaddres"`
	Admin_timezone        string `json:"admin_timezone"`
	Admin_createadmin     string `json:"admin_createadmin"`
	Admin_createdateadmin string `json:"admin_createdateadmin"`
	Admin_updateadmin     string `json:"admin_updateadmin"`
	Admin_updatedateadmin string `json:"admin_updatedateadmin"`
}
