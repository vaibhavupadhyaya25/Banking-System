package db

//CreateTables ...
func CreateTables() {
	d := Connect()
	CreateTableTransaction(d)
	CreateAccountDetailsTable(d)
	CreateCustomerTable(d)
}
