package api

import (
	"encoding/json"
	"fmt"
	cloudbypass "github.com/cloudbypass/golang-sdk"
	"github.com/gin-gonic/gin"
	"net/http"
	"pt/cons"
	"strings"
)

// Data 返回参数
type Data struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
}

type ForeignData struct {
	Symbols     string `json:"symbols" binding:"required"`
	Interval    string `json:"interval" binding:"required"`
	Pointscount string `json:"pointscount" binding:"required"`
}

type Event struct {
	X     int64   `json:"x"`
	Y     float64 `json:"y"`
	Title string  `json:"title"`
	Type  string  `json:"type"`
	Time  int64   `json:"time"`
	Text  string  `json:"text"`
}

type ForeignResponseData struct {
	Data   [][]float64 `json:"data"`
	Events []Event     `json:"events"`
}

func Foreign(context *gin.Context) {
	var fdata ForeignData

	// 返回数据格式
	var data Data

	// 解析请求参数
	if err := context.ShouldBindJSON(&fdata); err != nil {
		data.Error = err.Error()
		data.Code = 400
		context.JSON(http.StatusBadRequest, data)
		return
	}

	// 判断参数
	if fdata.Symbols == "" {
		data.Error = "Symbols 参数不能为空"
		data.Code = 400
		context.JSON(http.StatusBadRequest, data)
		return
	}
	if fdata.Pointscount == "" {
		data.Error = "Pointscount 参数不能为空"
		data.Code = 400
		context.JSON(http.StatusBadRequest, data)
		return
	}
	if fdata.Interval == "" {
		data.Error = "Interval 参数不能为空"
		data.Code = 400
		context.JSON(http.StatusBadRequest, data)
		return
	}

	varietyId := cons.GetVariety(fdata.Symbols)
	if varietyId == 0 {
		data.Error = "无该品种数据！"
		data.Code = 400
		context.JSON(http.StatusBadRequest, data)
		return
	}
	url := "https://api.investing.com/api/financialdata/%d/historical/chart/?interval=%s&pointscount=%s"
	url = fmt.Sprintf(url, varietyId, fdata.Interval, fdata.Pointscount)

	res := RequestUrl(url, "pre")
	// Remove <body> and </body>
	output := strings.Replace(res, "<pre>", "", -1)
	output = strings.Replace(output, "</pre>", "", -1)

	// 解析json
	var response_ ForeignResponseData
	err := json.Unmarshal([]byte(output), &response_)
	if err != nil {
		data.Error = err.Error() + "JOSN 解析错误"
		data.Code = 400
		context.JSON(http.StatusBadRequest, data)
		return
	}

	data.Data = response_.Data
	context.JSON(http.StatusBadRequest, data)
	return
}

func GetStockList(c *gin.Context) {
	var data Data
	data.Data = []string{
		cons.INTC,
		cons.MSFT,
		cons.ADEL,
		cons.APLH,
		cons.AJ,
		cons.AH,
		cons.HK0001,
		cons.HK0002,
		cons.JP1332,
		cons.DH6434,
		cons.DH5276,
		cons.AMZN,
		cons.AAPL,
		cons.TSLA,
		cons.GOOGL,
		cons.META,
		cons.LLY,
		cons.AVGO,
		cons.JPM,
		cons.NVDA,
		cons.GBI,
	}
	c.JSON(http.StatusBadRequest, data)
}

func RequestUrl(url string, query1 string) string {

	client := cloudbypass.New(cloudbypass.BypassConfig{
		Apikey: "5e6e62c13fc34af4861ae4684888b29d",
		Part:   "0",
		Proxy:  getProxy(),
	})

	resp, err := client.R().
		EnableTrace().
		Get(url)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	return resp.String()
}

func getProxy() string {
	proxy, err := cloudbypass.NewProxy("36719991-res:imbaqyyw")

	if err != nil {
		fmt.Println("代理提取失败", err)
		return err.Error()
	}
	return proxy.SetDynamic().String()

}
