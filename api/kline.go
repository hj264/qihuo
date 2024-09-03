package api

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pt/cons"
	"strconv"
	"strings"
	"time"
)

type KlineData struct {
	Symbols    string `json:"symbols" binding:"required"`
	Resolution string `json:"resolution" binding:"required"`
	From       string `json:"from" binding:"required"`
	To         string `json:"to" binding:"required"`
}

type KlineResponseData struct {
	T   []interface{} `json:"t"`
	C   []interface{} `json:"c"`
	O   []interface{} `json:"o"`
	H   []interface{} `json:"h"`
	L   []interface{} `json:"l"`
	V   []interface{} `json:"v"`
	Vo  []interface{} `json:"vo"`
	Vac []interface{} `json:"vac"`
	S   string        `json:"s"`
}

func GetKline(context *gin.Context) {
	var fdata KlineData

	// 返回数据格式
	var data Data

	// 解析请求参数
	if err := context.ShouldBindJSON(&fdata); err != nil {
		data.Error = err.Error()
		data.Code = 400
		context.JSON(http.StatusBadRequest, fdata)
		return
	}

	// 判断参数
	if fdata.Symbols == "" {
		data.Error = "Symbols 参数不能为空"
		data.Code = 400
		context.JSON(http.StatusBadRequest, fdata)
		return
	}
	if fdata.Resolution == "" {
		data.Error = "Resolution 参数不能为空"
		data.Code = 400
		context.JSON(http.StatusBadRequest, fdata)
		return
	}
	if fdata.From == "" {
		data.Error = "From 参数不能为空"
		data.Code = 400
		context.JSON(http.StatusBadRequest, fdata)
		return
	}
	if fdata.To == "" {
		data.Error = "To 参数不能为空"
		data.Code = 400
		context.JSON(http.StatusBadRequest, fdata)
		return
	}
	varietyId := cons.GetVariety(fdata.Symbols)
	if varietyId == 0 {
		data.Error = "无该品种数据！"
		data.Code = 400
		context.JSON(http.StatusBadRequest, data)
		return
	}

	// 时间戳
	timestamp := time.Now().Unix()

	// md5
	str := strconv.FormatInt(timestamp, 10) // Conv
	strByte := []byte(str)
	has := md5.Sum(strByte)
	md5str := fmt.Sprintf("%x", has)
	url := " https://tvc4.investing.com/%s/%d/1/1/8/history?symbol=%d&resolution=%s&from=%s&to=%s"
	url = fmt.Sprintf(url, md5str, timestamp, varietyId, fdata.Resolution, fdata.From, fdata.To)

	res := RequestUrl(url, "body")
	// Remove <body> and </body>
	output := strings.Replace(res, "<body>", "", -1)
	output = strings.Replace(output, "</body>", "", -1)

	// 解析json
	var response_ KlineResponseData
	err := json.Unmarshal([]byte(output), &response_)
	if err != nil {
		data.Error = err.Error() + "JOSN 解析错误"
		data.Code = 400
		context.JSON(http.StatusBadRequest, data)
		return
	}

	data.Data = response_
	context.JSON(http.StatusBadRequest, data)
	return
}
