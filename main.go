package main
import(
	"client/src"
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
	src.ApolloClient(s,l)
}
