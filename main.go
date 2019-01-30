package main
import(
	"plugin/src"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
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
	data := src.ApolloClient(s,l)
	intervals ,_:= strconv.Atoi(s.Intervals)
	r := gin.Default()
	r.GET("/metrics",func(c *gin.Context){
		data := src.CacheConf(&data,intervals,s,l)
		if data == nil{
			c.JSON(200,gin.H{"message":"error"})
		}
		c.String(200,src.Statis(data))
		//c.String(200,"hello")
	})
	r.Run(":8081")
}
