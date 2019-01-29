package src

import(
	"os/exec"
	"strings"
	"bytes"
)

func Statistics() string{
	base := ","
	funclist := map[string]string{"hostname":"hostname","system_v":"uname -r"}
	for k,v := range funclist{
		base += k+":"+"\""+Cmd(v)+"\","
	}
	return base
}


func hostname()(string){
	return Cmd("hostname")
}

func system_v() string{
	return Cmd("uname -r")
}

func Cmd(shell string) string{
    cmd := exec.Command("/bin/bash","-c",shell)
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil{
    	return "unknown"
	}
	return strings.TrimSpace(out.String())
}