package main

import "net/http"
import "fmt"
import "../notificator"
import "log"
import "os"

var notify *notificator.Notificator

func main(){
	
	argsWithProg := os.Args[1]

	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     os.Args[1],
	})
	for{
		resp, err := http.Get(argsWithProg)
		if err != nil {
			notify.Push("DOWN or ERROR", err.Error(), "/Users/henriqueferreira/Documents/GoLangDEV/up-monitor/icon.png", notificator.UR_CRITICAL)
			fmt.Println(err.Error());
		} else {
			fmt.Println(string(resp.StatusCode) + resp.Status)
		}

		if err != nil {
			log.Fatal(err)
		}
	}
}