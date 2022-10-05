FROM golang:1.18.6-alpine

ADD . /companymanager/
WORKDIR /companymanager/

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy

COPY ./src/cmd/cm/*.go ./

RUN go build ./src/cmd/cm/

EXPOSE 8080/tcp
ENTRYPOINT ["./cm","-port",":8080","-db.host","localhost","-db.port",":5111","-db.name","company","-db.usr","company","-db.pw","company","-kafka.brokers","host.docker.internal:29092","-kafka.topic","company","-secretKey","tajna"]