package util

// Response Success
type RS struct {
	Message  string `json:"message"`
	Status   bool `json:"status"`
	Data interface{} `json:"data"`
}

// Response Error
// field ---> error ----> return到首頁
// field ---> alarm ----> 傳送message到首頁即可

type RE struct {
	Field    string `json:"field"`
	Message  string `json:"message"`
}

type KafKaRS struct {
	Status      bool     `json:"status"`
	Message     string   `json:"message"`
	Residue     int64  	 `json:"residue"`
	Uuid        string   `json:"uuid"`
}