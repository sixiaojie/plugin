package src

import (
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
)

type LogFormat struct {
	Code int
	Status string
}

func(l *LogFormat) MakeMap()( logrus.Fields){
	//obj1 := reflect.TypeOf(l)
	data := make(map[string]interface{})
	s := reflect.ValueOf(l).Elem()
	typeoff := s.Type()
	for i := 0; i < s.NumField(); i++{
		f := s.Field(i)
		data[typeoff.Field(i).Name] = f.Interface()
	}
	return data
}

func Log(path string)(*logrus.Logger,error) {
	log := logrus.New()
	f,err := os.OpenFile(path,os.O_APPEND|os.O_CREATE|os.O_WRONLY,07777)
	if err != nil{
		return log,err
	}
	defer f.Close()
	log.Out = f
	log.Formatter = &logrus.JSONFormatter{}
	return log,nil
}


func (d *LogFormat) Error(log *logrus.Logger,message string){
	log.WithFields(d.MakeMap()).Error(message)
}

func (d *LogFormat) Info(log *logrus.Logger,message string){
	log.WithFields(d.MakeMap()).Info(message)
}

func (d *LogFormat) Warn(log *logrus.Logger,message string){
	log.WithFields(d.MakeMap()).Warn(message)
}
