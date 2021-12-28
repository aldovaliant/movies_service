# movies_service
Stockbit Test no. 2

This microservice is built with Clean Architecture. 
There are 2 routes: Search & Get movie detail by ID

1. http://localhost:8081/api/v1/search?{params}

params : 
- searchword : search input
- year : year input
- type : type input (movie, series, or episode)
- pagination : pagination input

2. http://localhost:8081/api/v1/detail?{params}

params : 
- id : id of the movie

Run command:
```
go run main.go
```
The service will be running on port :8081
