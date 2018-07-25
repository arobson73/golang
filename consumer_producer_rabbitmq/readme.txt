//setup rabbit mq broker with docker first
docker run --detach \
--name rabbitmq \
-p 5672:5672 \
-p 15672:15672 \
rabbitmq:3-management

//login to the management console localhost:15672
//user and password are guest
