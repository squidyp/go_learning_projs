package main

type UserProfile struct{
	UserId string
	FirstName string
	LastName string
	Email string
	Zip int16
	ContractorId string
	CustomerId string
}
type ContractorProfile struct{
	ContractorId string
	FirstName string //same as user
	LastName string //same as user
	Email string 
	Zip int16 //same as user
	CompanyId string
	Rating int8
	BankingDetails string //only visible to owner
	Licenses string
	Specialties string

}
type CustomerProfile struct{
	CustomerId string
	FirstName string //same as user
	LastName string //same as user
	Email string
	Zip int16 //same as user
	Rating int8
	BankingDetails string //only visible to owner
}
type Jobs struct{
	JobId string
	CreatedBy string // userId of creator
	CustomerId string // bound to a customer
	JobType string
	Location string
	Description string
	ScheduledDate string
	Status string
	ContractorId string
	StartDateTime string
	EndDateTime string
	PaymentCode string
}
type Company struct{
	CompanyId string
	CompanyRating int8 //average of employee ratings
	HqZip int16
	Insured bool
}
type Ledger struct {
	EntryId string //from payment system
	DateReceived string
	PaymentCode string
	CustomerId string
	AmountReceived float32
	CardFee float32
	ProcessFee float32
	AmountReleased float32
	DateReleased string
	ContractorId string
}
type JobType struct{
	JobTypeId string
	JobTypeName string
}