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




