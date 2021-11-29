package fcm

import (
	"api/model"
	"api/util/log"
	"encoding/json"
	"github.com/NaySoftware/go-fcm"
	"time"
)

type FCM struct {
	appId 		 int64
	storeIdList  []int64
	adminIdList  []int64
	userId       int64
	storeId      int64
	fcmServerKey string
	token        []string
	title        string
	body         string
	data         map[string]interface{}
}

func (f *FCM) SetAppId(appId int64) *FCM {
	f.appId = appId
	return f
}

func (f *FCM) SetStoreIdList(storeIdList []int64) *FCM {
	f.storeIdList = storeIdList
	return f
}

func (f *FCM) SetAdminIdList(adminIdList []int64) *FCM {
	f.adminIdList = adminIdList
	return f
}


func (f *FCM) SetStoreId(storeId int64) *FCM {
	f.storeId = storeId
	return f
}

func (f *FCM) SetUserId(userId int64) *FCM {
	f.userId = userId
	return f
}

func (f *FCM) SetFcmServerKey(fcmServerKey string) *FCM {
	f.fcmServerKey = fcmServerKey
	return f
}

func (f *FCM) SetTokens(token []string) *FCM {
	f.token = token
	return f
}

func (f *FCM) SetTitle(title string) *FCM {
	f.title = title
	return f
}

func (f *FCM) SetBody(body string) *FCM {
	f.body = body
	return f
}

func (f *FCM) SetData(data map[string]interface{}) *FCM {
	f.data = data
	return f
}

func (f *FCM) SendByAdminIdList() {
	var (
		adminDeviceToken model.AdminDeviceToken
		np               fcm.NotificationPayload
	)

	tokens := adminDeviceToken.SetAdminIdList(f.adminIdList).QueryTokenByAdminIdList()

	c := fcm.NewFcmClient(f.fcmServerKey)
	c.NewFcmRegIdsMsg(tokens, f.data)

	np.Title = f.title
	np.Body = f.body
	c.SetNotificationPayload(&np)

	c.NewFcmRegIdsMsg(tokens, f.data)
	status, err := c.Send()

	if err == nil {
		status.PrintResults()
	} else {
		log.Error(err)
	}

}

func (f *FCM) SendByStoreId() {
	var (
		deviceToken model.DeviceToken
		appList     model.AppList
		np          fcm.NotificationPayload
		notice 		model.Notice
	)
	//取得app info
	getStoreAppInfo := appList.GetStoreAppInfo(f.storeId)

	tokens := deviceToken.SetAppId(getStoreAppInfo[0]["appId"].(int64)).SetUserId(f.userId).QueryToken()

	c := fcm.NewFcmClient(getStoreAppInfo[0]["fcmServerKey"].(string))
	c.NewFcmRegIdsMsg(tokens, f.data)

	np.Title = f.title
	np.Body = f.body
	c.SetNotificationPayload(&np)
	status, err := c.Send()

	if err == nil {
		status.PrintResults()
		if f.userId != 0 && status.StatusCode == 200 && status.Success > 0 {
			dataEncode, _ := json.Marshal(f.data)
			if _, err := notice.SetCreatedAt(time.Now()).SetContent(f.title).SetContentBody(f.body).SetFcmData(string(dataEncode)).SetUserId(f.userId).SetParamId(0).SetNoticeTypeId(1).Insert(); err != nil {
				log.Error(err)
			}
		}
	} else {
		log.Error(err)
	}

}

func (f *FCM) Send() {
	var (
		deviceToken model.DeviceToken
		np 			fcm.NotificationPayload
		notice 		model.Notice
	)

	tokens := deviceToken.SetAppId(f.appId).SetUserId(f.userId).QueryToken()

	c := fcm.NewFcmClient(f.fcmServerKey)
	c.NewFcmRegIdsMsg(tokens, f.data)

	np.Title = f.title
	np.Body = f.body
	np.Sound = "default"
	np.ClickAction = "FCM_PLUGIN_ACTIVITY" //針對cordova的FCM Plugin而加
	c.SetNotificationPayload(&np)
	status, err := c.Send()

	if err == nil {
		status.PrintResults()
		if f.userId != 0 && status.StatusCode == 200 && status.Success > 0 {
			dataEncode, _ := json.Marshal(f.data)
			if _, err := notice.SetCreatedAt(time.Now()).SetContent(f.title).SetContentBody(f.body).SetFcmData(string(dataEncode)).SetUserId(f.userId).SetParamId(0).SetNoticeTypeId(1).Insert(); err != nil {
				log.Error(err)
			}
		}
	} else {
		log.Error(err)
	}

}