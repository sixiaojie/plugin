package main

import "net/http"

func main(){
	for {
		go func() {
			get()
		}()
	}

}

func get(){
	http.Get("http://127.0.0.1:8080/metrics")
}
