package src

import (
	"os/exec"
	"bytes"
	"strings"
)

func Statis(data map[string]string)(string){
	res := ""
	out := make(chan string)
	for k,v := range(data){
		a,_ := Command(k,v,out)
		res += k+ " "+a + "\n"
	}

	return res
}

func Command(k,v string,o chan string)(string,error){
	cmd := exec.Command("/bin/bash","-c",v)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil{
		return "0",err
	}else {
		return strings.TrimSpace(out.String()),nil
	}
}

