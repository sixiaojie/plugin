package main

import "plugin/src"

func main(){
	s := src.Config("conf/config.ini")
	e := src.LogFormat{200,"ok"}
	l,_ := src.Log(s.Logfile)
	e.Error(l,"hello")
}

