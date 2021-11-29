package model

import (
	"api/database/mysql"
	"api/util/log"
	"database/sql"
	"time"
)

type AppList struct {
	Id                     int64     `table:"id" `
	Name                   string    `table:"name"`
	AppName                string    `table:"app_name" `
	IsJoinOin              bool      `table:"is_join_oin"`
	FcmServerKey           string    `table:"fcm_server_key"`
	HasPosApi              bool      `table:"has_pos_api"`
	BackgroundImg          string    `table:"background_img"`
	HomeIconSelect         string    `table:"home_icon_select"`
	HomeIconUnselect       string    `table:"home_icon_unselect"`
	GiftIconSelect         string    `table:"gift_icon_select"`
	GiftIconUnselect       string    `table:"gift_icon_unselect"`
	QrIcon                 string    `table:"qr_icon"`
	AboutIconSelect        string    `table:"about_icon_select"`
	AboutIconUnselect      string    `table:"about_icon_unselect"`
	PersonalIconSelect     string    `table:"personal_icon_select"`
	PersonalIconUnselect   string    `table:"personal_icon_unselect"`
	FooterColor            string    `table:"footer_color"`
	StartImg               string    `table:"start_img"`
	UpdatedFooterIconAt    time.Time `table:"updated_footer_icon_at"`
	UpdatedFooterColorAt   time.Time `table:"updated_footer_color_at"`
	UpdatedBackgroundImgAt time.Time `table:"updated_background_img_at"`
	UpdatedStartImgAt      time.Time `table:"updated_start_img_at"`
	IsShowBackgroundImg    bool      `table:"is_show_background_img"`
	IsShowStartImg         bool      `table:"is_show_start_img"`
}

func (model *AppList) SetId(id int64) *AppList {
	model.Id = id
	return model
}

func (model *AppList) QueryOne() *AppList {
	log.Error(mysql.Model(model).
		Where("id", "=", model.Id).
		Select([]string{"fcm_server_key"}).
		Find().
		Scan(&model.FcmServerKey))

	return model
}

func (model *AppList) GetStoreAppInfo(storeId int64) []map[string]interface{} {
	var dataList = make([]map[string]interface{}, 0)
	mysql.Model(model).
		Where("store.id", "=", storeId).
		InnerJoin("brand", "brand.app_list_id", "=", "app_list.id").
		InnerJoin("store", "store.brand_id", "=", "brand.id").
		Select([]string{"app_list.id", "app_list.name", "app_list.fcm_server_key"}).
		Get(func(rows *sql.Rows) (isBreak bool) {
			log.Error(rows.Scan(&model.Id, &model.Name, &model.FcmServerKey))
			data := map[string]interface{}{
				"appId":       	model.Id,
				"name":    		model.Name,
				"fcmServerKey": model.FcmServerKey,
			}
			dataList = append(dataList, data)
			return
		})
	return dataList
}