package servers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"../logings"
	"../network"
	"../records"
	"../times"
)

// getRecordsDailyAPIHandler - 處理GET一日內小時資料網頁
/**
 * @param  *ECAPIServer eCAPIServer 環控API伺服器指標
 * @param  *gin.Context ginContextPointer  gin Context 指標
 */
func getRecordsDailyAPIHandler(eCAPIServer *ECAPIServer, ginContextPointer *gin.Context) {

	type Parameters struct {
		Year  int `uri:"year" json:"year"`
		Month int `uri:"month" json:"month"`
		Day   int `uri:"day" json:"day"`
	}

	var parameters Parameters

	bindJSONError := ginContextPointer.ShouldBindJSON(&parameters)

	bindURIError := ginContextPointer.ShouldBindUri(&parameters)

	defaultArgs :=
		append(
			network.GetAliasAddressPair(
				fmt.Sprintf(`%s:%d`,
					eCAPIServer.GetConfigValueOrPanic(`host`),
					eCAPIServer.GetConfigPositiveIntValueOrPanic(`port`),
				),
			),
			ginContextPointer.ClientIP(),
			ginContextPointer.FullPath(),
			parameters,
		)

	logings.SendLog(
		[]string{`%s %s 接受 %s 請求 %s %+v `},
		defaultArgs,
		nil,
		0,
	)

	parametersYear := parameters.Year
	parametersMonth := parameters.Month
	parametersDay := parameters.Day

	if bindJSONError == nil &&
		bindURIError == nil &&
		parametersMonth >= 1 && parametersMonth <= 12 &&
		parametersDay >= 1 && parametersDay <= 31 {

		low, upper :=
			times.
				GetDailyBounds(
					time.
						Date(
							parametersYear,
							time.Month(parametersMonth),
							parametersDay,
							0,
							0,
							0,
							0,
							time.Local,
						).
						AddDate(
							0,
							0,
							1,
						),
				) // 取得上下限日

		var (
			endDateTime time.Time // 結束時間
		)

		if now := time.Now(); upper.After(now) { // 若上限時間在現在時間之後
			endDateTime = now
		} else {
			endDateTime = upper
		}

		duration := time.Hour // 定義期間

		if hoursCount := int(times.ConvertToHourlyDateTime(endDateTime).Sub(low).Hours()); hoursCount != mongoDB.CountHourlyRecordsBetweenTimes(low, false, upper, true) { // 若缺資料

			hourlyRecordsDoneChannel := make(chan bool, 1)

			for dateTime := low.Add(duration); !dateTime.After(endDateTime); dateTime = dateTime.Add(duration) { // 針對每一小時

				thisDateTime := dateTime

				go func() {

					if mongoDB.CountHourlyRecordByTime(thisDateTime) == 0 {
						mongoDB.AggregateRepsertHourlyRecordByTime(thisDateTime) // 代添紀錄到現在時間

						if mongoDB.CountHourlyRecordByTime(thisDateTime) == 0 {
							mongoDB.RepsertHourlyRecord(records.HourlyRecord{Time: thisDateTime})
						}

					}

					hourlyRecordsDoneChannel <- true

				}()

			}

			for dateTime := low.Add(duration); !dateTime.After(endDateTime); dateTime = dateTime.Add(duration) {
				<-hourlyRecordsDoneChannel
			}

		}

		result := mongoDB.FindHourlyRecordsBetweenTimes(low, false, upper, true)

		ginContextPointer.JSON(http.StatusOK, result)

		logings.SendLog(
			[]string{`%s %s 回應 %s 請求 %s %+v : %+v `},
			append(
				defaultArgs,
				result,
			),
			nil,
			0,
		)

	} else {
		ginContextPointer.Status(http.StatusNotFound)

		logings.SendLog(
			[]string{`%s %s 拒絕回應 %s 請求 %s %+v `},
			defaultArgs,
			nil,
			0,
		)

	}

}

//取得所有apps info
func getAllAppsInfoAPIHandler(eCAPIServer *ECAPIServer, ginContextPointer *gin.Context) {
	type Parameters struct {
	}

	var parameters Parameters

	// bindJSONError := ginContextPointer.ShouldBindJSON(&parameters)

	// bindURIError := ginContextPointer.ShouldBindUri(&parameters)

	defaultArgs :=
		append(
			network.GetAliasAddressPair(
				fmt.Sprintf(`%s:%d`,
					eCAPIServer.GetConfigValueOrPanic(`host`),
					eCAPIServer.GetConfigPositiveIntValueOrPanic(`port`),
				),
			),
			ginContextPointer.ClientIP(),
			ginContextPointer.FullPath(),
			parameters,
		)

	logings.SendLog(
		[]string{`%s %s 接受 %s 請求 %s %+v `},
		defaultArgs,
		nil,
		0,
	)

	result := mongoDB.FindAllAppsInfoByProjectNameAndAppName()

	//包成回給前端的格式
	myResult := records.AppsInfoResponse{
		Code:    "200",
		Message: "",
		Data:    result,
	}

	ginContextPointer.JSON(http.StatusOK, myResult)

	logings.SendLog(
		[]string{`%s %s 回應 %s 請求 %s %+v : %+v `},
		append(
			defaultArgs,
			result,
		),
		nil,
		0,
	)
}

// 取得指定apps info
func getAppsInfoAPIHandler(eCAPIServer *ECAPIServer, ginContextPointer *gin.Context) {

	type Parameters struct {
		ProjectName string `uri:"projectName" json:"projectName"`
		AppName     string `uri:"appName" json:"appName"`
	}

	var parameters Parameters

	bindJSONError := ginContextPointer.ShouldBindJSON(&parameters)

	bindURIError := ginContextPointer.ShouldBindUri(&parameters)

	defaultArgs :=
		append(
			network.GetAliasAddressPair(
				fmt.Sprintf(`%s:%d`,
					eCAPIServer.GetConfigValueOrPanic(`host`),
					eCAPIServer.GetConfigPositiveIntValueOrPanic(`port`),
				),
			),
			ginContextPointer.ClientIP(),
			ginContextPointer.FullPath(),
			parameters,
		)

	logings.SendLog(
		[]string{`%s %s 接受 %s 請求 %s %+v `},
		defaultArgs,
		nil,
		0,
	)

	parametersProjectName := parameters.ProjectName
	parametersAppName := parameters.AppName

	fmt.Println("測試：已取得參數parametersProjectName=", parametersProjectName, "parametersAppName=", parametersAppName)
	if bindJSONError == nil && bindURIError == nil {

	}
	result := mongoDB.FindAppsInfoByProjectNameAndAppName(parametersProjectName, parametersAppName)

	// 包成回給前端的格式

	myResult := records.AppsInfoResponse{
		Code:    "200",
		Message: "",
		Data:    result,
	}

	ginContextPointer.JSON(http.StatusOK, myResult)

	logings.SendLog(
		[]string{`%s %s 回應 %s 請求 %s %+v : %+v `},
		append(
			defaultArgs,
			result,
		),
		nil,
		0,
	)
}
