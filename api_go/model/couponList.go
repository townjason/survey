package model

import (
	"api/database/mysql"
	"api/util/log"
	"time"
)

type CouponList struct {
	Id                 int64     `table:"id"`
	StoreId            int64     `table:"store_id"`
	Name               string    `table:"name"`
	Content            string    `table:"content"`
	ContentPrecautions string    `table:"content_precautions"`
	CouponTypeId       int64     `table:"coupon_type_id"`
	GetCouponLimitType int64     `table:"get_coupon_limit_type"`
	GetLimitAmount     int64     `table:"get_limit_amount"`
	IsGift             int64     `table:"is_gift"`
	GiftLimitType      int64     `table:"gift_limit_type"`
	GiftHourLimit      int64     `table:"gift_hour_limit"`
	GiftPeriodLimit    time.Time `table:"gift_period_limit"`
	GiftFrequencyLimit int64     `table:"gift_frequency_limit"`
	SendMax            int64     `table:"send_max"`
	TitleImage         string    `table:"title_image"`
	ContentImage       string    `table:"content_image"`
	ShowArea           int64     `table:"show_area"`
	CreatedAt          time.Time `table:"created_at"`
	StartAt            time.Time `table:"start_at"`
	EndAt              time.Time `table:"end_at"`
	ConvertLimitType   int64     `table:"convert_limit_type"`
	ConvertPeriodLimit time.Time `table:"convert_period_limit"`
	ConvertDateLimit   int64     `table:"convert_date_limit"`
	IsDiscountShare    bool      `table:"is_discount_share"`
	Status             int64     `table:"status"`
	AdminId            int64     `table:"admin_id"`
	DeadlineNotice     int64     `table:"deadline_notice"`
	PublishEnable      int64     `table:"publish_enable"`
	UseSet             int64     `table:"use_set"`
}

func (model *CouponList) SetId(id int64) *CouponList {
	model.Id = id
	return model
}

func (model *CouponList) GetCouponById() *CouponList {
	log.Error(mysql.Model(model).
		Where("id", "=", model.Id).
		Select([]string{"name"}).
		Find().
		Scan(&model.Name))

	return model
}

//獲取優惠券已經發幾張了
func (model *CouponList) GetSendCouponCountById(couponId int64) int64 {
	var userCoupon CouponUser
	var count int64
	log.Error(mysql.Model(&userCoupon).
		Where("status", "!=", 2).
		Where("coupon_list_id", "=", couponId).
		Select([]string{"count(*)"}).
		Find().Scan(&count))

	return count
}

//取得剩餘張數
func (model *CouponList) GetResidue(couponId int64) int64 {
	log.Error(mysql.Model(model).
		Where("id", "=", couponId).
		Select([]string{"id", "name", "content", "content_precautions", "end_at", "convert_limit_type", "convert_period_limit", "convert_date_limit", "content_image", "send_max"}).
		Find().
		Scan(&model.Id, &model.Name, &model.Content, &model.ContentPrecautions, &model.EndAt, &model.ConvertLimitType, &model.ConvertPeriodLimit, &model.ConvertDateLimit, &model.ContentImage, &model.SendMax))

	sendCount := model.GetSendCouponCountById(couponId)
	residue := model.SendMax - sendCount
	if residue < 0 {
		residue = 0
	}
	return residue
}
