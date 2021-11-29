package model

import (
	"api/database/mysql"
	"api/util/log"
	"time"
)

type Store struct {
	Id                    int64     `table:"id"`
	BrandId               int64     `table:"brand_id"`
	StoreTypeId           int64     `table:"store_type_id"`
	Status                int64     `table:"status"`
	Name                  string    `table:"name"`
	ZipCode               string    `table:"zip_code"`
	Tel                   string    `table:"tel"`
	City                  string    `table:"city"`
	Township              string    `table:"township"`
	Address               string    `table:"address"`
	Description           string    `table:"description"`
	Lat                   float64   `table:"lat"`
	Lng                   float64   `table:"lng"`
	IsBooking             bool      `table:"is_booking"`
	IsWaiting             bool      `table:"is_waiting"`
	CreatedAt             time.Time `table:"created_at"`
	Facebook              string    `table:"facebook"`
	Instagram             string    `table:"instagram"`
	OfficialWebsite       string    `table:"official_website"`
	MinPrice              int64     `table:"min_price"`
	MaxPrice              int64     `table:"max_price"`
	OfficialWebsiteStatus int64     `table:"official_website_status"`
	OfficialWebsiteVerify string    `table:"official_website_verify"`
	AdminId               int64     `table:"admin_id"`
	IsHeadOffice          bool      `table:"is_head_office"`
	Code				  string	`table:"code"`
	StoreType             string
	Distance              float64
	Star                  float64
	Offset                int64
}

func (model *Store) SetId(id int64) *Store {
	model.Id = id
	return model
}

func (model *Store) QueryOne() *Store {
	log.Error(mysql.Model(model).
		Where("id", "=", model.Id).
		Select([]string{"brand_id", "name"}).Find().Scan(&model.BrandId, &model.Name))
	return model
}