package logger

import (
	"log"
	"os"
)

func getLogger() (logger *log.Logger) {
	file, err := os.OpenFile("storage/logs/app.log", os.O_APPEND | os.O_RDWR | os.O_CREATE, 0777)
	if err != nil {
		log.Fatalln("日志驱动[文件]初始化失败：日志文件创建失败", err)
	}
	//defer file.Close()
	logger = log.New(file, "", log.LstdFlags | log.Lshortfile)
	return
}

func Info(value ...interface{})  {
	args := []interface{}{"Log.Info"}
	args = append(args, value...)
	log.Println(args)
	getLogger().Println(args)
}

func Error(value ...interface{})  {
	args := []interface{}{"Log.Error"}
	args = append(args, value...)
	log.Println(args)
	getLogger().Fatalln(args)
}