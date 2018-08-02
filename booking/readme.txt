#use the simpleEvent.go to build some json input files. 
#use the bash scripts to send GET/POST messages.
#use the builder to build project (see other directory (simple builder)

first make sure you have rabbitmq setup (this is for ubuntu via docker)
docker run --detach \
--name rabbitmq \
-p 5672:5672 \
-p 15672:15672 \
rabbitmq:3-management

make sure mongo is setup/installed.

in the /eventService folder run the main in a terminal . this will use rest at localhost:8181 mongo at port default(27017).
For this i started mongo prior to running this main via sudo service mongod start (this uses 27017) and gets its config from /etc/mongo.conf

in another terminal do mongod --dbpath ~/go/src/andy/booking/bookingservice/ --port 27018
you now have 2 instances of mongo running

in another terminal run the /bookingservice main (./main -conf=config.json) .Check this config it sets up mongo to use port 27018
and the rest is localhost:8182

open 2 more terminals one for (both for mongo cli)  
mongo --port 27017
and another
mongo --port 27018

now use the bash script to create an even (newEvent). and another to create a new user (newUser). check the mongo cli
both these scripts talk to endpoint localhost:8181. these events are written to mongo db (27017) and via amqp its replicated
to the other mongo db at 27018. 
now send a makeBooking . this talks to end point localhost:8182, and updates the mongo db at 27018 only.




