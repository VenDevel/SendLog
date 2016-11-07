package main

import (
	"SendLog/httpPost"
	"SendLog/models"
	"fmt"
	"time"
)

func main() {

	httpPost.GetInstance().HttpSendPost(1, "111", "aasdfasd")

	//xmlContent := new(models.Configxml)

	// err := config.LoadConfig("./conf/SendLogConfig.xml", xmlContent)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	panic(1)
	// }
	// fmt.Println(xmlContent)
	time.Sleep(10 * time.Second)
	fmt.Println(models.LogInfo{Code: 1111})

}
