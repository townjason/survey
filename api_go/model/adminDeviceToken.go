package model

import (
	. "api/database/mysql"
	"api/util/log"
	"database/sql"
)

type AdminDeviceToken struct {
	Id          int64  `table:"id"`
	StoreId     int64  `table:"store_id"`
	Token       string `table:"token"`
	StoreIdList []int64
	AdminIdList []int64
}

func (model *AdminDeviceToken) SetStoreIdList(storeIdList []int64) *AdminDeviceToken {
	model.StoreIdList = storeIdList
	return model
}

func (model *AdminDeviceToken) SetAdminIdList(adminIdList []int64) *AdminDeviceToken {
	model.AdminIdList = adminIdList
	return model
}

func (model *AdminDeviceToken) QueryTokenByAdminIdList() []string {
	var data []string
	if len(model.AdminIdList) > 0{
		adminIdList := make([]interface{}, 0)

		for _, adminId := range model.AdminIdList {
			adminIdList = append(adminIdList, adminId)
		}

		Model(model).
			Select([]string{"token"}).
			WhereIn("and", "admin_device_token.admin_id", adminIdList).
			Get(func(row *sql.Rows) (isBreak bool) {
				log.Error(row.Scan(&model.Token))
				data = append(data, model.Token)
				return
			})
	}

	return data
}

