package domain

//Business Side logic
type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// Secondary Port
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
