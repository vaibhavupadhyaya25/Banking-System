package db

//CreateTables ...
func CreateTables() {
	d := Connect()
	CreateAccountDetailsTable(d)
	//CreateTableTransaction(d)
	//CreateCustomerTable(d)
}
