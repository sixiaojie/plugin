package src

import(
	"os/exec"
	"strings"
	"bytes"
)

type Callback func() string

func funclist(callback Callback) string{
	return callback()
}

func hostname()(string){
	cmd := exec.Command("/bin/bash","-c","hostname")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil{
		return "unknown"
	}
	return strings.TrimSpace(out.String())
}

func Gather() {

}