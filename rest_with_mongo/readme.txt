//install mongodb then make sure its running (use ps aux | grep mongod)
sudo service mongod start
//to create json events use the example go file(example_marshall.go) 
//to build, cd to each directory and do the go install in each. then in root folder do go build main.go
then run it ./main
it should connect to the database and be ready for input.

send this (the json file was created from running example_marshall.go) (./example_marshall > new.json)
curl -d "@new.json" -X POST http://localhost:8181/events

check this by opening browser at
http://localhost:8181//events

try a delete event with 
curl -i -X DELETE http://localhost:8181/events/opera%20aida

//TODO try more things with it like searching etc.

