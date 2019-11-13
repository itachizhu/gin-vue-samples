package main

import (
	"github.com/itachizhu/gin-vue-samples/chapter06/logger"
	"io"
	"log"
	"os"
)

func main() {
	log.Println("这是日志")

	// 设置前缀
	log.SetPrefix("TRACE: ")
	// 设置需要输出的额外信息
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.Println("这是另一行日志")

	file, err := os.OpenFile("messages.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	// 输出到控制台和文件
	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.Println("日志会打印在控制台和文件中")

	logg := logger.New("mymessages.log")
	logg.Info("使用自定制的日志记录器打印日志。")

	zap := logger.NewZap("zap.log")
	zap.Info("这个是集成zap的日志方式。")
}
