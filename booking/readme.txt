
#use the simpleEvent.go to build some json input files. 
#use the bash scripts to send GET/POST messages.
#use the builder to build project go files (see other directory (simple builder)

first make sure you have rabbitmq setup (this is for ubuntu via docker)
docker run --detach \
--name rabbitmq \
-p 5672:5672 \
-p 15672:15672 \
rabbitmq:3-management

make sure mongo is setup/installed.

in the /eventService folder run the main in a terminal . this will use rest at localhost:8181 mongo at port default(27017).
For this i started mongo prior to running this main via sudo service mongod start (this uses 27017) and gets its config from /etc/mongo.conf

in another terminal do mongod --dbpath ~/go/src/andy/booking/bookingservice/db --port 27018
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

#Note to dockerize this would do as follows:
#make sure the binary exe's in both bookingservice and eventservice are name to these names rather than main.

#create Dockerfile like this in each of the bookingservice and eventservice folders 
#event service docker file

FROM debian:jessie
COPY eventservice /eventservice
RUN useradd eventservice
USER eventservice
ENV LISTEN_URL=0.0.0.0:8181
EXPOSE 8181
CMD ["/eventservice"]

#build the image like this
docker image build -t myevents/eventservice .

#bookingservice docker file

FROM debian:jessie
COPY bookingservice /bookingservice
RUN useradd bookingservice
USER bookingservice
ENV LISTEN_URL=0.0.0.0:8181
EXPOSE 8181
CMD ["/bookingservice"]

#build image like this
docker image build -t myevents/bookingservice .

#the container network like this
docker network create myevents

#rabbitmq like this
docker container run -d --name rabbitmq --network myevents
rabbitmq:3-management

#events-db like this
docker container run -d --name events-db --network myevents mongo
#bookings-db like this (make sure its a seperate instance, hence 20718)
docker container run -d -p 20718:20718 --name bookings-db --network myevents mongo

#eventservice 
docker container run \
--detach \
--name events \
--network myevents \
-e AMQP_BROKER_URL=amqp://guest:guest@rabbitmq:5672/ \
-e MONGO_URL=mongodb://events-db/events \
-p 8181:8181 \
myevents/eventservice

#bookingservice
docker container run --detach --name bookings --network myevents -e AMQP_BROKER_URL=amqp://guest:guest@rabbitmq:5672/ -e MONGO_URL=mongodb://bookings-db/bookings -p 8182:8181 myevents/bookingservice

#booking_front
docker container build -t myevents/frontend .
#then
docker container run --name frontend -p 80:80 myevents/frontend	 


############
## Alternatively rather than using docker as above (whuch is good way to test its all working) can instead use the ##docker-compose which will build and run the whole lot in one go
docker-compose up -d

##########
##USAGE##
#########
#check its all running with docker ps -a
#use the bash script 
newEvent outx.json 
#to send events
#also bash sript to send new user
newUser userx.json
#then you can make booking via
makeBooking bash script.

#you can view this by going to http://localhost

#this is all skeleton code, front end needs login credentials to add new events / users. booking frontend buttton needs to use the user id to actually make a button which missing. TODO

