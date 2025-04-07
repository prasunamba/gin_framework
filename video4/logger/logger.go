package logger

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Formatlogs(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s,%s,%v,%s,%d,%v,%v,%v\n",
		param.Method,
		param.Path,
		param.Request.Proto,
		param.ClientIP,
		param.StatusCode,
		param.TimeStamp,
		param.Request,
		param.Keys,
	)
}

type locallogparams struct {
	Method     string
	Path       string
	ClientIP   string
	StatusCode int
	TimeStamp  time.Time
}

// Ideal for monitoring & analysis
// JSON format logging
func Jsonformatlogs(params gin.LogFormatterParams) string {
	param := locallogparams{
		Method:     params.Method,
		Path:       params.Path,
		ClientIP:   params.ClientIP,
		StatusCode: params.StatusCode,
		TimeStamp:  params.TimeStamp,
	}
	by, err := json.Marshal(param)
	if err != nil {
		fmt.Printf("log error: %v\n", err)
	}
	return string(by)
}
