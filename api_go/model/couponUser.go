package model

import (
	"time"
)

type CouponUser struct {
	Id            int64     `table:"id"`
	UserId        int64     `table:"user_id"`
	CouponListId  int64     `table:"coupon_list_id"`
	Status        int64     `table:"status"`
	GiftTypeId    int64     `table:"gift_type_id"`
	CouponStartAt time.Time `table:"coupon_start_at"`
	CreatedAt     time.Time `table:"created_at"`
	UpdatedAt     time.Time `table:"updated_at"`
	GiftUserId    int64     `table:"gift_user_id"`
	FromUserId    int64     `table:"from_user_id"`
	Note          string    `table:"note"`
	AdminsId      int64     `table:"admins_id"`
}
