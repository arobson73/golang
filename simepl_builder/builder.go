package main

import (
	"log"
	"os/exec"
)

//note had to combine commands to get them to work (proably is better way!)
//		"cd /home/febe/go/src/andy/builder/test && go build main.go",
func main() {
	installPaths := []string{

		"cd /home/febe/go/src/andy/booking/lib/configuration && go install",
		"cd /home/febe/go/src/andy/booking/lib/helper/amqp && go install",

		"cd /home/febe/go/src/andy/booking/lib/helper/kafka && go install",
		"cd /home/febe/go/src/andy/booking/lib/msgqueue/amqp && go install",
		"cd /home/febe/go/src/andy/booking/lib/msgqueue/builder && go install",
		"cd /home/febe/go/src/andy/booking/lib/msgqueue/kafka && go install",
		"cd /home/febe/go/src/andy/booking/lib/msgqueue && go install",

		"cd /home/febe/go/src/andy/booking/lib/persistence/mongolayer && go install",
		"cd /home/febe/go/src/andy/booking/lib/persistence/dblayer && go install",
		"cd /home/febe/go/src/andy/booking/lib/persistence && go install",

		"cd /home/febe/go/src/andy/booking/bookingservice/listener && go install",
		"cd /home/febe/go/src/andy/booking/bookingservice/rest && go install",
		"cd /home/febe/go/src/andy/booking/contracts && go install",
		"cd /home/febe/go/src/andy/booking/eventService/rest && go install",
	}

	buildPaths := []string{
		"cd /home/febe/go/src/andy/booking/bookingservice && go build main.go",
		"cd /home/febe/go/src/andy/booking/eventService && go build main.go",
	}
	for _, p := range installPaths {
		//	out, err := exec.Command("/bin/bash", "-c", p).Output()
		cmd := exec.Command("/bin/bash", "-c", p)
		err := cmd.Run()
		if err != nil {
			log.Println("Error running cd with path")
			continue
		}
	}

}
