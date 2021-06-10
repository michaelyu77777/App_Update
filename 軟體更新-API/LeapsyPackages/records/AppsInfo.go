package records

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HourlyRecord - 小時紀錄
type AppsInfo struct {
	ProjectName  string //專案名稱
	AppName      string //軟體名稱
	LastVersion  string //最新版本號
	PackageName  string //封包名稱
	DownloadPath string //下載URL
}

// 包成回給前端的格式
type AppsInfoResponse struct {
	Code    string     `json:"code"`    //錯誤代碼
	Message string     `json:"message"` //錯誤訊息
	Data    []AppsInfo `json:"data"`    //查詢結果
}

// PrimitiveM - 轉成primitive.M
/*
 * @return primitive.M returnPrimitiveM 回傳結果
 */
func (appsInfo AppsInfo) PrimitiveM() (returnPrimitiveM primitive.M) {

	returnPrimitiveM = bson.M{
		`projectName`:  appsInfo.ProjectName,
		`appName`:      appsInfo.AppName,
		`lastVersion`:  appsInfo.LastVersion,
		`packageName`:  appsInfo.PackageName,
		`downloadPath`: appsInfo.DownloadPath,
	}

	return
}
