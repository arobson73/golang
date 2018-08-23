package main

import (
	"flag"
	"fmt"
	"time"

	"andy/booking/eventService/listner"
	"andy/booking/eventService/rest"
	"andy/booking/lib/configuration"
	"andy/booking/lib/msgqueue"
	msgqueue_amqp "andy/booking/lib/msgqueue/amqp"
	"andy/booking/lib/msgqueue/kafka"
	"andy/booking/lib/persistence/dblayer"

	"github.com/Shopify/sarama"

	"log"

	"github.com/streadway/amqp"
)

func init() {

	//log.SetOutput(ioutil.Discard)

}

func main() {
	var eventEmitter msgqueue.EventEmitter
	var eventListener msgqueue.EventListener

	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	switch config.MessageBrokerType {
	case "amqp":
		log.Println("EventService: messagebrokertype=amqp")

		conn, err := amqp.Dial(config.AMQPMessageBroker)
		if err != nil {

			time.Sleep(10 * time.Second)
			conn, err = amqp.Dial(config.AMQPMessageBroker)
		}

		if err != nil {
			panic(err)
		}
		eventListener, err = msgqueue_amqp.NewAMQPEventListener(conn, "events", "user")

		eventEmitter, err = msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
		if err != nil {
			panic(err)
		}
	case "kafka":
		log.Println("EventService: messagebrokertype=kafka")
		conf := sarama.NewConfig()
		conf.Producer.Return.Successes = true
		conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
		if err != nil {
			panic(err)
		}

		eventEmitter, err = kafka.NewKafkaEventEmitter(conn)
		if err != nil {
			panic(err)
		}
	default:
		panic("Bad message broker type: " + config.MessageBrokerType)
	}

	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	processor := listener.BookingProcessor{eventListener, dbhandler}
	go processor.ProcessEvents()

	fmt.Println("Serving API")
	//RESTful API start
	err := rest.ServeAPI(config.RestfulEndpoint, dbhandler, eventEmitter)
	if err != nil {
		panic(err)
	}
}
