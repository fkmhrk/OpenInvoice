package company

// DAO is interface
type DAO interface {
	// Gets all company
	GetList() ([]*Company, error)
	GetById(id string) (*Company, error)
	Create(name, zip, address, phone, unit string) (*Company, error)
	Update(id, name, zip, address, phone, unit string) (*Company, error)
	Delete(id string) error
}

// Company is copmany
type Company struct {
	Id      string
	Name    string
	Zip     string
	Address string
	Phone   string
	Unit    string
}
