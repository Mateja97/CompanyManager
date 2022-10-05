package company

import (
	"companymanager/src/cors"
	"companymanager/src/kafka"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Producer interface {
	Init([]string, string) error
	SendMessage(interface{}) error
	Stop() error
}
type CompanyManager struct {
	producer Producer
	server   *http.Server
	db       *sql.DB
}

func (cm *CompanyManager) Init(port, dbHost, dbPort, dbName, dbUsr, dbPw string, brokers []string, topic string) error {

	connStr := fmt.Sprintf("postgresql://%s:%s@%s%s/%s?sslmode=disable", dbUsr, dbPw, dbHost, dbPort, dbName)
	var err error
	cm.db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println("[ERROR] Sql.Open")
		return err
	}
	r := mux.NewRouter()
	r.HandleFunc("/company/{id}", cm.GetCompany()).Methods("GET")
	r.HandleFunc("/company/{id}", cm.PatchCompany()).Methods("PATCH")
	r.HandleFunc("/company/", cm.InsertCompany()).Methods("POST")
	r.HandleFunc("/company/{id}", cm.DeleteCompany()).Methods("DELETE")

	cm.server = &http.Server{
		Addr:    port,
		Handler: cors.CORSEnabled(r),
	}
	p := &kafka.KafkaProducer{}
	err = p.Init(brokers, topic)
	if err != nil {
		log.Println("[ERROR] Producer init failed")
		return err
	}
	cm.producer = p
	return nil
}

func (cm *CompanyManager) Run() {
	log.Println("[INFO] Company manager has started succesfully")
	if err := cm.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println("[ERROR] Company manager server run failed")
	}

}
func (cm *CompanyManager) Stop() {
	if err := cm.server.Shutdown(context.Background()); err != nil {
		log.Println("[ERROR] Company manager server shutdown failed")
	}
	err := cm.producer.Stop()
	if err != nil {
		log.Fatalln("[ERROR] Could not close consumer gracefully:", err.Error())
	}
	err = cm.db.Close()
	if err != nil {
		log.Fatalln("[ERROR] Could not close db gracefully:", err.Error())
	}
	log.Println("[INFO] CompanyManager stopped gracefully")
}
