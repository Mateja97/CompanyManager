package company

import (
	"database/sql"
	"log"
	"strconv"
)

const tableName = `company`

//ReadCompany returns Company from database for given id
func (cm *CompanyManager) ReadCompany(id string) (Company, error) {
	var company Company
	query := `select id,name,description,employeesnum,registered,type from ` + tableName + ` where id = '` + id + `';`

	err := cm.db.QueryRow(query).Scan(&company.ID, &company.Name, &company.Description, &company.EmployeesNum, &company.Registred, &company.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return Company{}, nil
		}
		log.Println("[ERROR] ReadCompany.QueryRow", err)
		return Company{}, err
	}
	return company, nil
}

//UpdateCompany updates company in the database
func (cm *CompanyManager) UpdateCompany(new Company, id string) (Company, error) {
	var company Company
	query := `update ` + tableName + ` set`
	if new.Name != "" {
		query += ` name ='` + new.Name + `',`
	}
	if new.Description != "" {
		query += ` description ='` + new.Description + `',`
	}
	if new.EmployeesNum != 0 {
		str := strconv.Itoa(int(new.EmployeesNum))
		query += ` employeesnum =` + str + `,`
	}
	if new.Type != 0 {
		str := strconv.Itoa(int(new.Type))
		query += ` type =` + str + `,`
	}
	query += ` registered = ` + strconv.FormatBool(new.Registred) + ` where id = '` + id + `' `
	query += `RETURNING id,name,description,employeesnum,registered,type;`
	err := cm.db.QueryRow(query).Scan(&company.ID, &company.Name, &company.Description, &company.EmployeesNum, &company.Registred, &company.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return Company{}, nil
		}
		log.Println("[ERROR] UpdateCompany.QueryRow", err)
		return Company{}, err
	}
	return company, nil
}

//Writecompany inserts company to the database if not exists
func (cm *CompanyManager) Writecompany(company Company) (string, error) {
	var id string
	query := `insert into ` + tableName + `(name,description,employeesnum,registered,type) VALUES($1,$2,$3,$4,$5) returning id;`
	err := cm.db.QueryRow(query, company.Name, company.Description, company.EmployeesNum, company.Registred, company.Type).Scan(&id)
	if err != nil {
		log.Println("[ERROR] Writecompany.QueryRow", err)
		return "", err
	}
	return id, nil
}

//DeleteCompanyDB deletes company from database
func (cm *CompanyManager) DeleteCompanyDB(id string) error {
	query := `delete from ` + tableName + ` where id = '` + id + `';`
	_, err := cm.db.Exec(query)
	if err != nil {
		log.Println("[ERROR] DeleteCompanyDB.QueryRow", err)
		return err
	}
	return nil
}
