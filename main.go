package main

import (
	"github.com/joho/godotenv"
	"github.com/qalexandre/imersaofsfc2-simulator/infra/kafka"

	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {

	producer := kafka.NewKafkaProducer()
	kafka.Publish("oi", "readtest", producer)

	for {
		_ = 1
	}
	//route := route2.Route{
	//	ID:       "1",
	//	ClientId: "1",
	//}
	//
	//route.LoadPositions()
	//stringJson, _ := route.ExportJsonPositions()
	//fmt.Println(stringJson[0])
}
