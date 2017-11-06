# GolangRestfulApiExample

This personal project shows how to write a simple Restful Microservice API written in Go (Golang).

### Prerequisites

This project uses the Go language to build the Miscroservice. The database connection is made with MongoDB.
So both Go and MongoDB must be installed.
The configuration is not externalized, so the connection with the database will be done int he same machine as the database server is running (localhost).

### Installing

To run this project just run the main file and run the MongoDB server in another terminal.

Start the microservice:
```
go run main.go
```

Start the database server:
```
mongod
```

Exemple of the requests that can be done with the microservice.
Create, Read, Update and Delete.

POST
- The POST request will insert a new user in the database, all of the fields must have an value associated or it will not be accepted by the service.
- The socialNumber must be unique within all of the users.

Body:
```
{
	"name": "JP",
	"age": "26",
	"phones": ["32511234", "999991991"],
	"socialNumber": "88877766600",
	"address":{
		"streetName": "Rua Principal",
		"streetNumber": "123",
		"zipCode": "18000000",
		"country": "Brasil",
		"state": "São Paulo"
	}
}
```

GET
- To recover and see the information about a user, the GET request will only need a socialNumber associated with the user.
- Each user has its own and unique socialNumber.

Header:
```
socialNumber: 88877766600
```

PATCH
- To update or change some data from a user, just create a PATCH request.
- The data send within the request must have at leat one completed field in the body and the socialNumber in the header to be accepted by the service.

Header:
```
socialNumber: 88877766600
```

Body:
```
{
	"name": "João Pedro",
	"age": "27",
	"phones": ["999991991"],
	"socialNumber": "88877766600",
	"address":{
		"streetName": "Rua Matriz",
		"streetNumber": "456",
		"zipCode": "18111111",
		"country": "Peru",
		"state": "Lima"
	}
}
```

DELETE
- To delete all of the selected user`s data the DELETE request must be sent with only a socialNumber associated with the user in the header.
- Each user has its own and unique socialNumber.

Header:
```
socialNumber: 88877766600
```

## Running the tests

Unit and Integration Tests are being written within each service.
A script is being made that the test and its verification with the database must pass it all.
