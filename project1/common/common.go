package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"project/project1/config"
	"project/project1/models"
	"sync"
)

// 进行幅值操作

var Template models.HtmlTemplate

func LoadTemplate() {
	//
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		//耗时 扔到协程来进行
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}
func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}

func Success(w http.ResponseWriter, data interface{}) {
	// 处理前端返回的东西 返回不同的结果  返回的是对应的json数据

	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, _ := json.Marshal(result)
	// 转化成json 之后高数 前端您输入的是json格式的
	w.Header().Set("Contet-Type", "application/json")

	_, err := w.Write(resultJson)
	if err != nil {
		// 打印日志
		log.Println(err)
	}
}

func Error(w http.ResponseWriter, err error) {
	// 处理前端返回的东西 返回不同的结果  返回的是对应的json数据

	var result models.Result
	result.Code = -999
	result.Error = err.Error()
	resultJson, _ := json.Marshal(result)
	// 转化成json 之后高数 前端您输入的是json格式的
	w.Header().Set("Contet-Type", "application/json")

	_, err = w.Write(resultJson)
	if err != nil {
		// 打印日志
		log.Println(err)
	}
}
