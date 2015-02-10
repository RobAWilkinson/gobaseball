# Go Baseball

## MVP for accessing a database of Baseball players
Dependencies:  

* gin
* go mysql

### Getting started
* create the proper gopath directory: `mkdir -p $GOPATH/src/bitbucket.org/Robawilkinson/`
* clone this repo in there.
* open up mysql `$ mysql -uroot`
* create the database, call it "baseball" `#=: CREATE DATABASE baseball`
* Use the database: `USE baseball`
* And load the file: `source baseball.sql`
* build baseball `$ go build sql.go`
* And run it `$ ./sql`

