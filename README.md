# CompanyManager
## Preparation <br />
Run make <br />
docker-compose is used to run kafka and postgres <br />


## Run commad <br />
### flags: <br />
    -port    -port on which service runs 
    -db.host -hostname for postgres 
    -db.name -database name 
    -db.usr  -postgres user
    -db.pw   -postgres password 
    -kafka.broker -kafka brokers seperated wiht "," 
    -kafka.topic  -kafka topic that should receive updates
    -secretKey    -secret key that is used for jwt authorization
### Local example: <br />
- docker build --tag cm .
- docker run cm<br />
- run without docker: go run ./src/cmd/cm/ -db.host=localhost -db.port=:5111 -db.name=company -db.usr=company -db.pw=company -kafka.brokers=localhost:9092 -kafka.topic=company -secretKey=

### REST
- GET - /company/{id} - retrieve specified company
- POST - /company/ + body with specified fields + token in header - post new company
- PATCH - /company/{id} + body with specified fields + token in header - patch specified company
- DELETE - /company/{id} + token in header - delete specifeid company



