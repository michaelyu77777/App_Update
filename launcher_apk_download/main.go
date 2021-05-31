package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/download/apk", downloadFile)
	fmt.Println("Server run on port 5001 for apk download")
	http.ListenAndServe(":5001", nil)
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	file := "delog201102.apk"

	// 告知檔名
	w.Header().Set("Content-Disposition", "attachment; filename="+file)

	http.ServeFile(w, r, file)
}
