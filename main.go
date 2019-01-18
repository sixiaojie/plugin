package main
import(
	"plugin/src"
	"github.com/gin-gonic/gin"
)

func main() {
	s := src.Config("conf/config.ini")
	l,err := src.Log(s.Logfile)
	if err != nil{
		panic(err.Error())
	}
	if err != nil{
		panic(err)
	}
	r := gin.Default()
	r.GET("/metrics",func(c *gin.Context){
		data := src.ApolloClient(s,l)
		if data == nil{
			c.JSON(200,gin.H{"message":"error"})
		}
		c.String(200,src.Statis(data))
	})
	r.Run()
}
