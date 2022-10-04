package company

type CompanyType int

const (
	Corporation CompanyType = iota
	NonProfit
	Cooperative
	SoleProprietorship
)

type Company struct {
	ID           string
	Name         string
	Description  string
	EmployeesNum int32
	Registred    bool
	Type         CompanyType
}
