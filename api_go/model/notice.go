package model

import (
	. "api/database/mysql"
	"time"
)

type Notice struct {
	Id           int64     `table:"id"`
	ParamId      int64     `table:"param_id"`
	NoticeTypeId int64     `table:"notice_type_id"`
	UserId       int64     `table:"user_id"`
	Content      string    `table:"content"`
	ContentBody  string    `table:"content_body"`
	IsRead       bool      `table:"is_read"`
	CreatedAt    time.Time `table:"created_at"`
	FcmData      string    `table:"fcm_data"`
	Db           Database
}

const (
	NoticeTypeForCommentReply    = 1
	NoticeTypeForCommentLike     = 2
	NoticeTypeForGiftCoupon      = 3
	NoticeTypeForGiftCard        = 4
	NoticeTypeForGiftLoyaltyCard = 5
	NoticeRegistGift             = 6
)

func (model *Notice) SetId(id int64) *Notice {
	model.Id = id
	return model
}

func (model *Notice) SetNoticeTypeId(noticeType int64) *Notice {
	model.NoticeTypeId = noticeType
	return model
}

func (model *Notice) SetParamId(paramId int64) *Notice {
	model.ParamId = paramId
	return model
}

func (model *Notice) SetUserId(userId int64) *Notice {
	model.UserId = userId
	return model
}

func (model *Notice) SetContent(content string) *Notice {
	model.Content = content
	return model
}

func (model *Notice) SetContentBody(contentBody string) *Notice {
	model.ContentBody = contentBody
	return model
}

func (model *Notice) SetFcmData(fcmData string) *Notice {
	model.FcmData = fcmData
	return model
}

func (model *Notice) SetIsRead(isRead bool) *Notice {
	model.IsRead = isRead
	return model
}

func (model *Notice) SetCreatedAt(createdAt time.Time) *Notice {
	model.CreatedAt = createdAt
	return model
}

func (model *Notice) SetDb(db Database) *Notice {
	model.Db = db
	return model
}

func (model *Notice) Insert() (int64, error) {
	return Model(model).Insert()
}

func (model *Notice) TransactionCreate() (int64, error) {
	return model.Db.Model(model).Insert()
}
