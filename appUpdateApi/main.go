package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// apps
// type Apps1 struct {
// 	AppName       string `json:"appName"`
// 	Id            int    `json:"id"`
// 	AppButtonName string `json:"appButtonName"`
// 	// IsMainMenu       bool   `json:"isMainMenu"`
// 	AppState         int    `json:"appState"`
// 	MultiLanguageKey string `json:"multiLanguageKey"`
// 	PackageName      string `json:"packageName"`
// 	//Download_path string `json:"download_path"`
// }

// type Apps2 struct {
// 	IsPlus bool `json:"isPlus"`
// }

// 最新apk與檔案路徑
type Apk_Version_And_File_Path struct {
	ProjectName  string `json:"projectName"`
	AppName      string `json:"appName"`
	LastVersion  string `json:"lastVersion"`
	PackageName  string `json:"packageName"`
	DownloadPath string `json:"downloadPath"`
}

func main() {

	router := mux.NewRouter() // 新路由
	// router.HandleFunc(`/software/update/apps`, launcherAppsHandler)
	router.HandleFunc(`/appUpdate/allLastVersionApks`, allLastVersionApksAPIHandler)

	apiServerPointer := &http.Server{
		Addr:           ":5004",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	} // 設定伺服器

	log.Fatal(apiServerPointer.ListenAndServe())

}

// 列出所有apk狀態
func allLastVersionApksAPIHandler(w http.ResponseWriter, r *http.Request) {

	// 連接Mongodb
	mongoClientPointer, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(`mongodb://localhost:27017`)) // 連接預設主機

	// 若連接錯誤
	if nil != err {
		fmt.Fprintf(w, "連MongoDB Server錯誤") // 寫入回應
		return
	}

	// 選資料庫與collection
	cursor, err := mongoClientPointer.Database(`Apps_Update`).Collection(`appsInformation`).Find(context.TODO(), bson.M{}) //找全部
	//Find(context.TODO(), bson.M{"time": bson.M{`$gt`: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local), `$lt`: time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local)}}) //時間要大於某時間 並且小於某時間

	if nil != err {
		fmt.Fprintf(w, "連MongoDB.Collection錯誤") // 寫入回應
		return
	}

	// 取得DB結果
	var results []Apk_Version_And_File_Path

	for cursor.Next(context.TODO()) { // 針對每一紀錄

		var data Apk_Version_And_File_Path

		err = cursor.Decode(&data) // 解析紀錄

		if nil != err {
			fmt.Fprintf(w, "no data") // 寫入回應
			return
		}

		results = append(results, data) // 儲存紀錄

	}

	jsonBytes, err := json.Marshal(results) // 轉成JSON

	if nil != err {
		fmt.Fprintf(w, "轉成JSON格式錯誤") // 寫入回應
		return
	}

	fmt.Fprintf(w, "%s", string(jsonBytes)) // 寫入回應

}

// func launcherAppsHandler(w http.ResponseWriter, r *http.Request) {

// 	mongoClientPointer, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(`mongodb://localhost:27017`)) // 連接預設主機

// 	if nil != err {
// 		fmt.Fprintf(w, "連線MongoDB Server錯誤") // 寫入回應
// 		return
// 	}

// 	cursor, err := mongoClientPointer.
// 		Database(`Launcher`).
// 		Collection(`apps`).
// 		//Find(context.TODO(), bson.M{"time": bson.M{`$gt`: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local), `$lt`: time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local)}}) //時間要大於某時間 並且小於某時間
// 		Find(context.TODO(), bson.M{}) //時間要大於某時間 並且小於某時間

// 	if nil != err {
// 		fmt.Fprintf(w, "連線MongoDB.Collection錯誤") // 寫入回應
// 		return
// 	}

// 	// 取得DB結果
// 	// var results1 []Apps1
// 	// var results2 []Apps2

// 	var results []interface{} // 以interface盛裝兩種形狀的json物件

// 	for cursor.Next(context.TODO()) { // 針對每一紀錄

// 		var data1 Apps1
// 		var data2 Apps2

// 		err = cursor.Decode(&data1) // 以Apps1物件形狀來解碼

// 		if nil != err {
// 			fmt.Fprintf(w, "no data1") // 印出err
// 			return
// 		}

// 		err = cursor.Decode(&data2) // 以Apps2物件形狀來解碼

// 		if nil != err {
// 			fmt.Fprintf(w, "no data2") // 印出err
// 			return
// 		}

// 		//若 Apps1
// 		if `` != data1.AppName {
// 			results = append(results, data1) // 儲存紀錄
// 		} else {
// 			results = append(results, data2) // 儲存紀錄
// 		}

// 		// results1 = append(results1, data1) // 儲存紀錄

// 	}

// 	jsonBytes, err := json.Marshal(results) // 轉成JSON

// 	if nil != err {
// 		fmt.Fprintf(w, "轉成JSON格式錯誤") // 寫入回應
// 		return
// 	}

// 	fmt.Fprintf(w, "%s", string(jsonBytes)) // 寫入回應

// }
