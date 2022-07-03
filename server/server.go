package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"time"
	"timestamp-service/internal"
)

var validPeriods = []string{"1h", "1d", "1mo", "1y"}

func InitServer() *gin.Engine {
	server := gin.Default()
	server.GET("/ptlist", HandlePtlist)
	server.UnescapePathValues = false
	return server
}

func checkIfPeriodIsValid(period string) bool {
	for _, v := range validPeriods {
		if period == v {
			return true
		}
	}
	return false
}

func checkQueryParams(query map[string]string) (time.Time, time.Time, time.Location, error) {
	start, err := internal.ParseString(query["t1"])
	if err != nil {
		return time.Time{}, time.Time{}, time.Location{}, &ErrInvalidT1Timestamp
	}
	end, err := internal.ParseString(query["t2"])
	if err != nil {
		return time.Time{}, time.Time{}, time.Location{}, &ErrInvalidT2Timestamp
	}
	loc, err := time.LoadLocation(query["tz"])
	if err != nil {
		return time.Time{}, time.Time{}, time.Location{}, &ErrInvalidTimezone
	}
	return start.In(loc), end.In(loc), *loc, nil
}

func getPeriodicTaskTimestamps(startTimestamp *internal.PeriodicTask, timestamps []*internal.PeriodicTask) []string {
	for i, _ := range timestamps {
		if i == 0 {
			timestamps[i] = startTimestamp
		} else {
			timestamps[i] = timestamps[i-1].Next()
		}
	}
	returnArr := make([]string, len(timestamps))
	for i, _ := range returnArr {
		returnArr[i] = internal.DateToString(timestamps[i].InvocationPoint.In(time.UTC))
	}
	return returnArr
}

func HandlePtlist(c *gin.Context) {
	params := []string{"period", "tz", "t1", "t2"}
	queryParams := make(map[string]string, len(params))
	for _, v := range params {
		value, ok := c.GetQuery(v)
		if ok {
			queryParams[v] = value
		} else {
			c.JSON(http.StatusBadRequest, ResponseError{
				Status:      http.StatusBadRequest,
				Description: fmt.Sprintf("%s wasn't specified in the url. Make sure you specify all the necessary parameters", v),
			})
			return
		}
	}
	if !checkIfPeriodIsValid(queryParams["period"]) {
		c.JSON(http.StatusConflict, ErrUnsupportedPeriod)
		return
	}
	period := queryParams["period"]
	start, end, loc, err := checkQueryParams(queryParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if end.Before(start) {
		c.JSON(http.StatusBadRequest, ErrStartDateGreater)
		return
	}
	years, months, _, _, _, _ := internal.Diff(start, end)
	var timestamps []*internal.PeriodicTask
	var startTimestamp internal.PeriodicTask
	switch period {
	case "1h":
		timestamps = make([]*internal.PeriodicTask, int(math.Round(end.Sub(start).Hours())*math.Pow(10, 0)))
		startTimestamp = internal.PeriodicTask{
			TaskPeriod:      internal.Hourly,
			Timezone:        loc,
			InvocationPoint: start.Round(time.Hour),
		}
		break
	case "1d":
		timestamps = make([]*internal.PeriodicTask, int(math.Round((end.Sub(start).Hours()/24)*math.Pow(10, 0))))
		startTimestamp = internal.PeriodicTask{
			TaskPeriod:      internal.Daily,
			Timezone:        loc,
			InvocationPoint: time.Date(start.Year(), start.Month(), start.Day(), start.Round(time.Hour).Hour(), 0, 0, 0, start.Location()),
		}
		break
	case "1mo":
		timestamps = make([]*internal.PeriodicTask, years*12+months)
		startTimestamp = internal.PeriodicTask{
			TaskPeriod:      internal.Monthly,
			Timezone:        loc,
			InvocationPoint: time.Date(start.Year(), start.Month(), start.Day(), start.Hour(), 0, 0, 0, start.Location()),
		}
		break
	case "1y":
		timestamps = make([]*internal.PeriodicTask, years)
		startTimestamp = internal.PeriodicTask{
			TaskPeriod:      internal.Yearly,
			Timezone:        loc,
			InvocationPoint: time.Date(start.Year(), start.Month(), start.Day(), start.Hour(), 0, 0, 0, start.Location()),
		}
		break
	default:
		break
	}
	c.JSON(http.StatusOK, getPeriodicTaskTimestamps(&startTimestamp, timestamps))
}
