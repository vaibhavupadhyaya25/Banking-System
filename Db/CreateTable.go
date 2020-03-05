package Db

//CreateTables ...
func CreateTables() {
	d := Connect()
	//CreateAccountDetailsTable(d)
	//CreateTableTransaction(d)
	CreateCustomerTable(d)
}
