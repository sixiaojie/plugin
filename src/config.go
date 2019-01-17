package src

import (
	"io/ioutil"
	"encoding/json"
)

type ServiceConfg struct {
	Appid  string
	Server string
	ClusterName string
	NamespaceName string
	Logfile string
}



func Config(path string)(*ServiceConfg){
	s := &ServiceConfg{}
	data,err := ioutil.ReadFile(path)
	if err != nil{
		panic(err.Error())
	}
	err = json.Unmarshal(data,s)
	if err != nil{
		panic(err)
	}
	return s
}






