package main

import (
	company "companymanager/src/companyManager"
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
)

//Service port
var port = flag.String("port", ":8080", "")

//DATABASE FLAGS
var dbHost = flag.String("db.host", "", "Database host address")
var dbName = flag.String("db.name", "", "Database table name")
var dbPort = flag.String("db.port", "", "Database port")
var dbUsr = flag.String("db.usr", "", "Database username")
var dbPw = flag.String("db.pw", "", "Database password")

//KAFKA FLAGS
var kafkaBrokers = flag.String("kafka.brokers", "", "Ip address of kafka broker")
var kafkaTopic = flag.String("kafka.topic", "", "Topic of data")

func main() {
	cm := company.CompanyManager{}
	flag.Parse()
	brokers := strings.Split(*kafkaBrokers, ",")
	err := cm.Init(*port, *dbHost, *dbPort, *dbName, *dbUsr, *dbPw, brokers, *kafkaTopic)
	if err != nil {
		log.Fatal("[ERROR] CompanyManager init failed, ", err)
	}
	go cm.Run()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	<-done
	//After signal is received, gracefully shutdown application
	cm.Stop()
}
