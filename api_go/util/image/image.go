package image

import (
	"api/config"
	"api/util/hit"
)

func ReturnPhotoPath(path string) string{
	//判斷圖片的路徑是否為空字串，是的話則回傳空字串
	return  hit.If(path == "", "", config.ServerInfo.FileHost + path).(string)
}
