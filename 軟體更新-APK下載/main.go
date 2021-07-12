package main

import (
	"fmt"
	"net/http"
)

const (
	port         = "63999"
	urlBasicPath = "/appUpdate/download/"

	//URL路徑
	urlPathName1  = "1"  // 教育訓練語音Service
	urlPathName2  = "2"  // 教育訓練
	urlPathName3  = "3"  // 熱像偵測
	urlPathName4  = "4"  // EnvironmentSysteem
	urlPathName5  = "5"  // 教育訓練語音Service
	urlPathName6  = "6"  // 自動點擊
	urlPathName7  = "7"  // expertglass (專家系統眼鏡端)
	urlPathName8  = "8"  // expertpad (專家系統平板端)
	urlPathName9  = "9"  //
	urlPathName10 = "10" //
	urlPathName11 = "11" //
	urlPathName12 = "12" //
	urlPathName13 = "13" //
	urlPathName14 = "14" //
	urlPathName15 = "15" //

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

	filePath7 = "apk/7/expertglass.apk" // expertglass (專家系統眼鏡端)
	fileName7 = "expertglass.apk"

	filePath8 = "apk/8/expertpad.apk" // expertpad (專家系統平板端)
	fileName8 = "expertpad.apk"

	filePath9 = "apk/9/test.apk" //
	fileName9 = ""

	filePath10 = "apk/10/test.apk" //
	fileName10 = ""

	filePath11 = "apk/11/test.apk" //
	fileName11 = ""

	filePath12 = "apk/12/test.apk" //
	fileName12 = ""

	filePath13 = "apk/13/test.apk" //
	fileName13 = ""

	filePath14 = "apk/14/test.apk" //
	fileName14 = ""

	filePath15 = "apk/15/test.apk" //
	fileName15 = ""
)

func main() {

	// 檔案一下載路徑
	http.HandleFunc(urlBasicPath+urlPathName1, downloadFile1)
	http.HandleFunc(urlBasicPath+urlPathName2, downloadFile2)
	http.HandleFunc(urlBasicPath+urlPathName3, downloadFile3)
	http.HandleFunc(urlBasicPath+urlPathName4, downloadFile4)
	http.HandleFunc(urlBasicPath+urlPathName5, downloadFile5)
	http.HandleFunc(urlBasicPath+urlPathName6, downloadFile6)
	http.HandleFunc(urlBasicPath+urlPathName7, downloadFile7)
	http.HandleFunc(urlBasicPath+urlPathName8, downloadFile8)
	http.HandleFunc(urlBasicPath+urlPathName9, downloadFile9)
	http.HandleFunc(urlBasicPath+urlPathName10, downloadFile10)
	http.HandleFunc(urlBasicPath+urlPathName11, downloadFile11)
	http.HandleFunc(urlBasicPath+urlPathName12, downloadFile12)
	http.HandleFunc(urlBasicPath+urlPathName13, downloadFile13)
	http.HandleFunc(urlBasicPath+urlPathName14, downloadFile14)
	http.HandleFunc(urlBasicPath+urlPathName15, downloadFile15)

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

// 檔案七下載
func downloadFile7(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName7) // 下載檔名
	http.ServeFile(w, r, filePath7)                                          // 檔案路徑
}

// 檔案八下載
func downloadFile8(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName8) // 下載檔名
	http.ServeFile(w, r, filePath8)                                          // 檔案路徑
}

// 檔案9
func downloadFile9(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName9) // 下載檔名
	http.ServeFile(w, r, filePath8)                                          // 檔案路徑
}

// 檔案10
func downloadFile10(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName10) // 下載檔名
	http.ServeFile(w, r, filePath8)                                           // 檔案路徑
}

// 檔案11
func downloadFile11(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName11) // 下載檔名
	http.ServeFile(w, r, filePath8)                                           // 檔案路徑
}

// 檔案12
func downloadFile12(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName12) // 下載檔名
	http.ServeFile(w, r, filePath8)                                           // 檔案路徑
}

// 檔案13
func downloadFile13(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName13) // 下載檔名
	http.ServeFile(w, r, filePath8)                                           // 檔案路徑
}

// 檔案14
func downloadFile14(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName14) // 下載檔名
	http.ServeFile(w, r, filePath8)                                           // 檔案路徑
}

// 檔案15
func downloadFile15(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName15) // 下載檔名
	http.ServeFile(w, r, filePath8)                                           // 檔案路徑
}
