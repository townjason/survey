package model

import (
	. "api/database/mysql"
	"api/util/log"
	"database/sql"
	"encoding/json"
	"time"
)

type Questionnaire struct {
	Id                        int64     `table:"id"`
	StoreId                   int64     `table:"store_id"`
	Name                      string    `table:"name"`
	Content                   string    `table:"content"`
	IsOpen                    bool      `table:"is_open"`
	BuildInItem               string    `table:"built_in_item"`
	ImagePath                 string    `table:"image_path"`
	PushNotificationSettingId int64     `table:"push_notification_setting_id"`
	IsRepeatWrite             bool      `table:"is_repeat_write"`
	IsPush                    bool      `table:"is_push"`
	JoinPushTime              int64     `table:"join_push_time"`
	IsGift                    bool      `table:"is_gift"`
	Coupon                    string    `table:"coupon"`
	IsLimitTime               bool      `table:"is_limit_time"`
	LimitTime                 int       `table:"limit_time"`
	CongratulationText        string    `table:"congratulation_text"`
	CreatedAt                 time.Time `table:"created_at"`
	UpdatedAt                 time.Time `table:"updated_at"`
	AdminId                   int64     `table:"admin_id"`
	Limit                     int
	Offset                    int
}

type BuildInItem struct {
	Type      string `json:"type"`
	IsShow    bool   `json:"isShow"`
	IsRequire bool   `json:"isRequire"`
}

func (model *Questionnaire) SetId(id int64) *Questionnaire {
	model.Id = id
	return model
}

func (model *Questionnaire) QueryOne() *Questionnaire {
	log.Error(Model(model).Select([]string{"store_id", "name", "content", "is_open", "built_in_item", "image_path", "is_repeat_write", "is_push",
		"join_push_time", "is_gift", "coupon", "is_limit_time", "limit_time", "congratulation_text", "admin_id"}).
		Where("id", "=", model.Id).
		Find().
		Scan(&model.StoreId, &model.Name, &model.Content, &model.IsOpen, &model.BuildInItem, &model.ImagePath, &model.IsRepeatWrite, &model.IsPush,
			&model.JoinPushTime, &model.IsGift, &model.Coupon, &model.IsLimitTime, &model.LimitTime, &model.CongratulationText, &model.AdminId))
	return model
}

func (model *Questionnaire) QueryAll() ([]map[string]interface{}, int) {
	var dataList = make([]map[string]interface{}, 0)

	table := Model(model)
	tableCount := Model(model)

	if model.StoreId != 0 {
		table.Where("questionnaire.store_id", "=", model.StoreId)
		tableCount.Where("questionnaire.store_id", "=", model.StoreId)
	}

	table.Select([]string{"id", "name", "is_open", "created_at"}).
		LeftJoin("admins", "admins.id", "=", "questionnaire.admin_id").
		Get(func(rows *sql.Rows) (isBreak bool) {
			log.Error(rows.Scan(&model.Id, &model.Name, &model.IsOpen, &model.CreatedAt))
			data := map[string]interface{}{
				"id":        model.Id,
				"name":      model.Name,
				"isOpen":    model.IsOpen,
				"createdAt": model.CreatedAt.Format("2006-01-02 15:04:05"),
			}
			dataList = append(dataList, data)
			return
		})

	count := tableCount.Count() //取得目前資料筆數
	return dataList, count
}

func (model *Questionnaire) GetRank(questionnaireId int64) int {
	var count = 1
	var rank = 0

	table := Model(model)

	if model.StoreId != 0 {
		table.Where("questionnaire.store_id", "=", model.StoreId)
	}

	table.Select([]string{"id"}).
		Get(func(rows *sql.Rows) (isBreak bool) {
			log.Error(rows.Scan(&model.Id))
			if model.Id == questionnaireId {
				rank = count
			}
			count++
			return
		})

	return rank
}

func (model *Questionnaire) Update(columns []string) {
	table := Model(model)

	if model.Id != 0 {
		table.Where("id", "=", model.Id)
	}

	log.Error(table.Update(columns))
}

func (model *Questionnaire) GetRequire(buildInItemType, buildInItem string) bool {
	var buildInItemList []BuildInItem

	log.Error(json.Unmarshal([]byte(buildInItem), &buildInItemList))

	for _, buildInItem := range buildInItemList {
		if buildInItem.Type == buildInItemType {
			return buildInItem.IsRequire
		}
	}

	return false
}
