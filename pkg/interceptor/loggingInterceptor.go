package interceptor

import (
	"log"
	"os"
)

type LoggingInterceptor struct {
}

func (l LoggingInterceptor) execute(request string) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Printf("request for %s sent", request)
}
