package model

import (
	. "api/database/mysql"
	"api/util/log"
	"database/sql"
	"time"
)


type DeviceToken struct {
	Id         int64      `table:"id"`
	AppId      int64      `table:"app_id"`
	UserId     int64      `table:"user_id"`
	Token 	   string     `table:"token"`
	Os 		   string     `table:"os"`
	CreatedAt  time.Time  `table:"created_at"`
}

func (model *DeviceToken) SetAppId(appId int64) *DeviceToken {
	model.AppId = appId
	return model
}

func (model *DeviceToken) SetUserId(userId int64) *DeviceToken {
	model.UserId = userId
	return model
}

func (model *DeviceToken) SetToken(token string) *DeviceToken {
	model.Token = token
	return model
}

func (model *DeviceToken) SetOs(os string) *DeviceToken {
	model.Os = os
	return model
}

func (model *DeviceToken) SetCreatedAt(createdAt time.Time) *DeviceToken {
	model.CreatedAt = createdAt
	return model
}

func (model *DeviceToken) QueryOne() *DeviceToken {
	table := Model(model)

	if model.UserId != 0 {
		table.Where("user_id", "=", model.UserId)
	}

	log.Error(table.
		Select([]string{"id"}).
		Where("app_id", "=", model.AppId).
		Where("token", "like", model.Token).
		Where("os", "like", model.Os).
		Find().Scan(&model.Id))
	return model
}

func (model *DeviceToken) QueryToken() []string {
	var data []string
	Model(model).
		Select([]string{"token"}).
		Where("app_id", "=", model.AppId).
		Where("user_id", "=", model.UserId).
		Get(func(row *sql.Rows) (isBreak bool) {
		log.Error(row.Scan(&model.Token))
		data = append(data, model.Token)
		return
	})
	return data
}

func (model *DeviceToken) Insert() {
	model.CreatedAt = time.Now()
	_, err := Model(model).Insert()
	log.Error(err)
}

func (model *DeviceToken) Update(columns []string) {
	log.Error(Model(model).
		Where("app_id", "=", model.AppId).
		Where("token", "like", model.Token).
		Where("os", "like", model.Os).
		Update(columns))
}