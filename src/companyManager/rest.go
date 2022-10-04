package company

import (
	"companymanager/src/auth"
	"companymanager/src/kafka"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (cm *CompanyManager) GetCompany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		data, err := cm.ReadCompany(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if data == (Company{}) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		resp, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

func (cm *CompanyManager) PatchCompany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing id"))
			return
		}
		token := r.Header.Get("Token")

		err := auth.VerifyToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token unauthorized"))
			return
		}
		var company Company
		err = json.NewDecoder(r.Body).Decode(&company)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		data, err := cm.UpdateCompany(company, id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if data == (Company{}) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		resp, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//if everything is successful, publish to the kafka
		cm.producer.SendMessage(kafka.KafkaMessage{
			ID:        id,
			Data:      data,
			Operation: "update",
		})
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

func (cm *CompanyManager) InsertCompany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")

		err := auth.VerifyToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token unauthorized"))
			return
		}
		var company Company
		err = json.NewDecoder(r.Body).Decode(&company)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if company.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing name"))
			return
		}
		if company.EmployeesNum == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing employees number "))
			return
		}
		if company.Type == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing company type "))
			return
		}
		data, err := cm.Writecompany(company)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if data == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		company.ID = data
		//if everything is successful, publish to the kafka
		cm.producer.SendMessage(kafka.KafkaMessage{
			ID:        data,
			Data:      company,
			Operation: "create",
		})
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(data))
	}
}

func (cm *CompanyManager) DeleteCompany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		token := r.Header.Get("Token")

		err := auth.VerifyToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token unauthorized"))
			return
		}
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = cm.DeleteCompanyDB(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//if everything is successful, publish to the kafka
		cm.producer.SendMessage(kafka.KafkaMessage{
			ID:        id,
			Data:      nil,
			Operation: "delete",
		})
		w.WriteHeader(http.StatusOK)
	}
}
