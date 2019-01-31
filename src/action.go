package src

import (
	"os/exec"
	"bytes"
	"strings"
	"time"
	"fmt"
)

func Statis(data map[string]string)(string){
	channel := make(chan map[string]string)
	for k,v := range(data){
		k = strings.Replace(k,"}",base,-1)
		k = strings.Trim(k,",") + "}"
		 go Command(k,v,channel)
		//res += k+ " "+a + "\n"
	}
	res := ""
	go func(){
		for {
			a := <- channel
			for k,v := range(a){
				res += k + " "+v + "\n"
			}
		}
	}()
	time.Sleep(5e8)
	return res
}

func Command(k,v string,in chan map[string]string){
	result := make(map[string]string)
	cmd := exec.Command("/bin/bash","-c",v)
	//cmd := exec.Command("/bin/bash","-c","ls")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil{
		fmt.Println(err.Error())
	}else {
		result[k] = strings.TrimSpace(out.String())
	}
	in <- result
}

