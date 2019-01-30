package main
import(
	"plugin/src"
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
)

var data map[string]string


func main() {
	fmt.Println(data)
	s := src.Config("conf/config.ini")
	e := src.LogFormat{200,"ok"}
	l,err := src.Log(s.Logfile)
	e.Error(l,"hello")
	if err != nil{
		panic(err.Error())
	}
	if err != nil{
		panic(err)
	}
	intervals ,_:= strconv.Atoi(s.Intervals)
	r := gin.Default()
	go src.Changevalue(&data,s,l,intervals)
	r.GET("/metrics",func(c *gin.Context){
		if data == nil{
			c.JSON(200,gin.H{"message":"error"})
		}
		c.String(200,src.Statis(data))
		//c.String(200,"hello")
	})
	r.Run(":8081")
}


