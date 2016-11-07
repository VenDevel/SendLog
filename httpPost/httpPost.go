package httpPost

import (
	"SendLog/config"
	"SendLog/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type HttpSend struct {
	Url         string
	ContentType string
	Body        string
}

var HttpInstance *HttpSend = nil

func init() {

	xmlContent := new(models.Configxml)
	err := config.LoadConfig("./SendLogConfig.xml", xmlContent)
	if err != nil {
		log.Println("func init config.LoadConfig error:", err.Error())
		return
	}
	HttpInstance = &HttpSend{
		Url:         xmlContent.Url,
		ContentType: xmlContent.ContentType,
	}
	log.Println("xml config", xmlContent)
	log.Println("func init HttpInstance", HttpInstance)
}

func (this *HttpSend) HttpSendPost(code int64, logname string, data interface{}) {

	httpBody := &models.LogInfo{
		Code:    code,
		LogName: logname,
		Log:     data,
	}
	jsonbyte, err := json.Marshal(httpBody)
	if err != nil {
		log.Println(err.Error())
		panic(1)
	}
	go httpPost(this.Url, string(jsonbyte))

}

func testgo(url string, jsonstr string) {
	log.Println("testgo----------------------------")
	log.Println(url)
	log.Println(jsonstr)
}

func httpPost(url string, jsonstr string) {
	log.Println("go httpPost")
	log.Println("go httpPost url:", url)
	log.Println("go httpPost jsonStr:", jsonstr)

	resp, err := http.Post(url, "application/json", strings.NewReader(jsonstr))
	if err != nil {
		log.Println("go httpPost http.Post error:", err)
		panic(4)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Println("go httpPost ReadAll resp.Body error", err.Error())
		panic(5)
	}

	log.Println(string(body))
}

func GetInstance() *HttpSend {
	if HttpInstance == nil {
		HttpInstance = &HttpSend{}
		xmlContent := new(models.Configxml)
		err := config.LoadConfig("./conf/SendLogConfig.xml", xmlContent)
		if err != nil {
			log.Println(err.Error())
			panic(1)
		}
		HttpInstance.Url = xmlContent.Url
		HttpInstance.ContentType = xmlContent.ContentType
	}
	log.Println("get instance")
	log.Println(HttpInstance)
	return HttpInstance
}
