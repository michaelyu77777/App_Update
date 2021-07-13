package main

import (
	"fmt"
	"net/http"
)

const (
	port         = "63999"
	urlBasicPath = "/appUpdate/download/"

	//URL路徑
	urlPathName1  = "1"  //
	urlPathName2  = "2"  //
	urlPathName3  = "3"  //
	urlPathName4  = "4"  //
	urlPathName5  = "5"  //
	urlPathName6  = "6"  //
	urlPathName7  = "7"  //
	urlPathName8  = "8"  //
	urlPathName9  = "9"  //
	urlPathName10 = "10" //
	urlPathName11 = "11" //
	urlPathName12 = "12" //
	urlPathName13 = "13" //
	urlPathName14 = "14" //
	urlPathName15 = "15" //

	// 檔案路徑與下載檔名
	filePath1 = "apk/1/camera.apk" //
	fileName1 = "camera.apk"

	filePath2 = "apk/2/album.apk" //
	fileName2 = "album.apk"

	filePath3 = "apk/3/webBrowser.apk" //
	fileName3 = "webBrowser.apk"

	filePath4 = "apk/4/throne.apk" //
	fileName4 = "throne.apk"

	filePath5 = "apk/5/leapsyStore.apk" //
	fileName5 = "leapsyStore.apk"

	filePath6 = "apk/6/intelligenceSystem.apk" //
	fileName6 = "intelligenceSystem.apk"

	filePath7 = "apk/7/environment.apk" //
	fileName7 = "environment.apk"

	filePath8 = "apk/8/palace.apk" //
	fileName8 = "palace.apk"

	filePath9 = "apk/9/FH1.apk" //
	fileName9 = "FH1.apk"

	filePath10 = "apk/10/settings.apk" //
	fileName10 = "settings.apk"

	filePath11 = "apk/11/programs.apk" //
	fileName11 = "programs.apk"

	filePath12 = "apk/12/expert.apk" //
	fileName12 = "expert.apk"

	filePath13 = "apk/13/faceRecognize.apk" //
	fileName13 = "faceRecognize.apk"

	filePath14 = "apk/14/posture.apk" //
	fileName14 = "posture.apk"

	filePath15 = "apk/15/gesture.apk" //
	fileName15 = "gesture.apk"
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

	// 提示
	fmt.Println("開啟Port:" + port + " 提供APK下載服務")

	// 開啟Port
	http.ListenAndServe(":"+port, nil)

}

// 檔案一下載
func downloadFile1(w http.ResponseWriter, r *http.Request) {

	//print
	fmt.Println("收到HOST ", r.Host, " 要求路徑 ", r.URL, "下載檔案", fileName1)

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
