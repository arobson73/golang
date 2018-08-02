package builder

import (
	"errors"
	"log"
	"os"

	"andy/booking/lib/msgqueue"
	"andy/booking/lib/msgqueue/amqp"
	"andy/booking/lib/msgqueue/kafka"
)

func init() {

	//log.SetOutput(ioutil.Discard)

}

func NewEventListenerFromEnvironment() (msgqueue.EventListener, error) {
	var listener msgqueue.EventListener
	var err error

	if url := os.Getenv("AMQP_URL"); url != "" {
		log.Printf("connecting to AMQP broker at %s", url)

		listener, err = amqp.NewAMQPEventListenerFromEnvironment()
		if err != nil {
			return nil, err
		}
	} else if brokers := os.Getenv("KAFKA_BROKERS"); brokers != "" {
		log.Printf("connecting to Kafka brokers at %s", brokers)

		listener, err = kafka.NewKafkaEventListenerFromEnvironment()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Neither AMQP_URL nor KAFKA_BROKERS specified")
	}
	log.Println("NewEventListenerFromEnvironment() - builder")

	return listener, nil
}
