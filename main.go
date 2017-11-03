package main

import "net/http"
import "fmt"
import "../notificator"
import "log"

var notify *notificator.Notificator

func main(){
	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "Up-Monitor",
	})
	
	for{
		resp, err := http.Get("http://localhost:9200")
		if err != nil {
			notify.Push("DOWN or ERROR", err.Error(), "/Users/henriqueferreira/Documents/GoLangDEV/up-monitor/icon.png", notificator.UR_CRITICAL)
			break
		} else {
			fmt.Println(string(resp.StatusCode) + resp.Status)
		}

		if err != nil {
			log.Fatal(err)
		}
	}
}