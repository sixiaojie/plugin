package src

import (
	"net/http"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
	"fmt"
)


func ApolloClient(c *ServiceConfg, l *logrus.Logger)(map[string]string){
	e := LogFormat{}
	url := "http://" + c.Server + "/configfiles/json/" + c.Appid + "/" + c.ClusterName + "/" + c.NamespaceName
	resp,err := http.Get(url)
	if err != nil{
		e.Code = 500
		e.Status = "failed"
		e.Error(l,err.Error())
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		e.Code = 501
		e.Status = "http body read failed"
		e.Error(l,err.Error())
	}else {
		data,err := JsonUnMarshal(body)
		//fmt.Println(data)
		if err != nil{
			e.Code = 502
			e.Status = "json unmarshal http body failed"
			e.Error(l,"json body data is unmarsha failed")
		}else{
			a,err := JsonUnMarshal([]byte(data["monitor"]))
			if err != nil{
				e.Code = 500
				e.Status = "key monitor json failed"
				e.Error(l,"json  monitor data is unmarsha failed")
			}else {
				return a
			}
		}
	}
	return nil
}

func JsonUnMarshal(data []byte)(map[string]string,error){
	temp := make(map[string]string)
	err := json.Unmarshal(data,&temp)
	if err != nil{
		fmt.Println(err.Error())
		return nil,err
	}
	return temp,nil
}