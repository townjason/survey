package model

import (
	"api/database/mysql"
	"api/util/log"
	"time"
)

type Brand struct {
	Id                    int64     `table:"id"`
	Name                  string    `table:"name"`
	AppListId             int64     `table:"app_list_id"`
	Code                  string    `table:"code"`
	Status                int64     `table:"status"`
	RoleOinSettingRoleId  int64     `table:"role_oin_setting_role_id"`
	StoreTypeId           int64     `table:"store_type_id"`
	Contact               string    `table:"contact"`
	Phone                 string    `table:"phone"`
	Email                 string    `table:"email"`
	Address               string    `table:"address"`
	CompanyTitle          string    `table:"company_title"`
	TaxIdNumber           string    `table:"tax_id_number"`
	Note                  string    `table:"note"`
	ContractStartAt       time.Time `table:"contract_start_at"`
	ContractEndAt         time.Time `table:"contract_end_at"`
	ProfessionCodeListId  int64     `table:"profession_code_list_id"`
	Facebook              string    `table:"facebook"`
	Instagram             string    `table:"instagram"`
	OfficialWebsite       string    `table:"official_website"`
	About                 string    `table:"about"`
	Image                 string    `table:"image"`
	CreatedAt             time.Time `table:"created_at"`
	UpdatedAt             time.Time `table:"updated_at"`
}

func (model *Brand) SetId(id int64) *Brand {
	model.Id = id
	return model
}

func (model *Brand) QueryOne() *Brand {
	log.Error(mysql.Model(model).
		Where("id", "=", model.Id).
		Select([]string{"name", "app_list_id", "image"}).Find().Scan(&model.Name, &model.AppListId, &model.Image))
	return model
}