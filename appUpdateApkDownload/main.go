package main

import (
	"fmt"
	"net/http"
)

const (
	port         = "63999"
	urlBasicPath = "/appUpdate/download/"

	//URL路徑
	urlPathName1 = "1" // 教育訓練語音Service
	urlPathName2 = "2" // 教育訓練
	urlPathName3 = "3" // 熱像偵測
	urlPathName4 = "4" // EnvironmentSysteem
	urlPathName5 = "5" // 教育訓練語音Service
	urlPathName6 = "6" // 自動點擊

	// 檔案路徑與下載檔名
	filePath1 = "apk/1/教育訓練語音Service.apk" // 教育訓練語音Service
	fileName1 = "教育訓練語音Service.apk"

	filePath2 = "apk/2/教育訓練.apk" // 教育訓練
	fileName2 = "教育訓練.apk"

	filePath3 = "apk/3/熱像偵測.apk" // 熱像偵測
	fileName3 = "熱像偵測.apk"

	filePath4 = "apk/4/EnvironmentSysteem.apk" // EnvironmentSysteem
	fileName4 = "EnvironmentSysteem.apk"

	filePath5 = "apk/5/launcher_Fii.apk" // 教育訓練語音Service
	fileName5 = "launcher_Fii.apk"

	filePath6 = "apk/6/Auto-click_Service.apk" // 自動點擊
	fileName6 = "Auto-click_Service.apk"
)

func main() {

	// 檔案一下載路徑
	http.HandleFunc(urlBasicPath+urlPathName1, downloadFile1)
	http.HandleFunc(urlBasicPath+urlPathName2, downloadFile2)
	http.HandleFunc(urlBasicPath+urlPathName3, downloadFile3)
	http.HandleFunc(urlBasicPath+urlPathName4, downloadFile4)
	http.HandleFunc(urlBasicPath+urlPathName5, downloadFile5)
	http.HandleFunc(urlBasicPath+urlPathName6, downloadFile6)

	// 開啟Port
	http.ListenAndServe(":"+port, nil)

	// 提示
	fmt.Println("已經開啟Port:" + port + " 提供APK下載服務")

}

// 檔案一下載
func downloadFile1(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName1) // 下載檔名
	http.ServeFile(w, r, filePath1)                                          // 檔案路徑
}

// 檔案二下載
func downloadFile2(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName2) // 下載檔名
	http.ServeFile(w, r, filePath2)                                          // 檔案路徑
}

// 檔案三下載
func downloadFile3(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName3) // 下載檔名
	http.ServeFile(w, r, filePath3)                                          // 檔案路徑
}

// 檔案四下載
func downloadFile4(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName4) // 下載檔名
	http.ServeFile(w, r, filePath4)                                          // 檔案路徑
}

// 檔案五下載
func downloadFile5(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName5) // 下載檔名
	http.ServeFile(w, r, filePath5)                                          // 檔案路徑
}

// 檔案六下載
func downloadFile6(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName6) // 下載檔名
	http.ServeFile(w, r, filePath6)                                          // 檔案路徑
}
