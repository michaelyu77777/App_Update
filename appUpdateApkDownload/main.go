package main

import (
	"fmt"
	"net/http"
)

const (
	port = "5005"

	pathName1 = "1" // 教育訓練語音Service
	pathName2 = "2" // 教育訓練
	pathName3 = "3" // 熱像偵測
	pathName4 = "4" // EnvironmentSysteem
	pathName5 = "5" // 教育訓練語音Service

	fileName1 = "" // 教育訓練語音Service
	fileName2 = "" // 教育訓練
	fileName3 = "" // 熱像偵測
	fileName4 = "" // EnvironmentSysteem
	fileName5 = "" // 教育訓練語音Service
)

func main() {

	// 檔案一下載路徑
	http.HandleFunc("/appUpdate/download/"+pathName1, downloadFile1)
	http.HandleFunc("/appUpdate/download/"+pathName2, downloadFile2)
	http.HandleFunc("/appUpdate/download/"+pathName3, downloadFile3)
	http.HandleFunc("/appUpdate/download/"+pathName4, downloadFile4)
	http.HandleFunc("/appUpdate/download/"+pathName5, downloadFile5)

	// 開啟Port
	http.ListenAndServe(":"+port, nil)

	// 提示
	fmt.Println("已經開啟Port:" + port + " 提供APK下載服務")

}

// 檔案一下載
func downloadFile1(w http.ResponseWriter, r *http.Request) {

	// 檔案名稱
	// file := "delog201102.apk"

	// 進行設定
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName1)

	http.ServeFile(w, r, fileName1)
}

// 檔案一下載
func downloadFile2(w http.ResponseWriter, r *http.Request) {

	// 檔案名稱
	// file := "delog201102.apk"

	// 進行設定
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName1)

	http.ServeFile(w, r, fileName2)
}

// 檔案一下載
func downloadFile3(w http.ResponseWriter, r *http.Request) {

	// 檔案名稱
	// file := "delog201102.apk"

	// 進行設定
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName1)

	http.ServeFile(w, r, fileName3)
}

// 檔案一下載
func downloadFile4(w http.ResponseWriter, r *http.Request) {

	// 檔案名稱
	// file := "delog201102.apk"

	// 進行設定
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName1)

	http.ServeFile(w, r, fileName4)
}

// 檔案一下載
func downloadFile5(w http.ResponseWriter, r *http.Request) {

	// 檔案名稱
	// file := "delog201102.apk"

	// 進行設定
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName1)

	http.ServeFile(w, r, fileName5)
}
