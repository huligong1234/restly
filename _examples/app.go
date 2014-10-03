// app.go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/huligong1234/restly"
	"net/http"
	"net/url"
)

type App struct {
	Id      int    `json:"appId"`
	AppCode string `json:"appCode"`
	AppName string `json:"appName"`
}

type Result struct {
	Rows      []App  `json:"rows"`
	Total     int    `json:"total"`
	Success   bool   `json:"success"`
	ResultMsg string `json:resultMsg`
}

func (app App) Get(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	app01 := App{Id: 1, AppCode: "1001", AppName: "Angry Birds"}
	app02 := App{2, "1002", "Temple Run"}
	app03 := App{3, "1003", "Baldur's Gate series"}
	app04 := App{4, "1004", "Bastion"}
	apps := []App{app01, app02}
	apps = append(apps, app03, app04)
	result := Result{Rows: apps, Total: len(apps), Success: true}

	strResult, _ := json.Marshal(result)
	fmt.Println(string(strResult))
	return 200, result, http.Header{"Content-type": {"application/json"}}
}

func main() {
	app := new(App)

	api := restly.NewAPI()
	api.AddResource(app, "/apps")
	api.Start(8080)
}
