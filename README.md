# CompanyManager
<br />
## Preparation <br />
Run make <br />
docker-compose is used to run kafka and postgres <br />

## Run commad <br />
### flags: <br />
    -port    -port on which service runs <br />
    -db.host -hostname for postgres <br />
    -db.name -database name <br />
    -db.usr  -postgres user <br />
    -db.pw   -postgres password <br />
    -kafka.broker -kafka brokers seperated wiht "," <br />
    -kafka.topic  -kafka topic that should receive updates <br /> 
    -secretKey    -secret key that is used for jwt authorization
### Local example: <br />
go run ./src/cmd/companyManagerd/ -db.host=localhost -db.port=:5111 -db.name=company -db.usr=company -db.pw=company -kafka.brokers=localhost:9092 -kafka.topic=company -secretKey=tajna <br />
