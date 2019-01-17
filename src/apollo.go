package src

import (
	"net/http"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type item map[string]string

type MonitorConfig struct {
	Item item
} 

func ApolloClient(c *ServiceConfg, l *logrus.Logger){
	v := &MonitorConfig{}
	e := LogFormat{}
	url := "http://" + c.Server + "/configfiles/json/" + c.Appid + "/" + c.ClusterName + "/" + c.NamespaceName
	fmt.Println(url)
	resp,err := http.Get(url)
	if err != nil{
		e.Code = 500
		e.Status = "failed"
		e.Error(l,err.Error())
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(string(body))
		err := json.Unmarshal(body,v)
		if err != nil{
			fmt.Println(err.Error())
		}else{
			fmt.Println(v)
		}
	}
}
