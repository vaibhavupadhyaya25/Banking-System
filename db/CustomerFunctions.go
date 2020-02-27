package db

import (
	"log"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

//to update the key or something else
func (pi *Customerinfo) UpdateCustomer(db *pg.DB) error {
	_, updateErr := db.Model(pi).Set("fname = ?fname").Where("cid= ?cid").Update()

	if updateErr != nil {
		log.Printf("Error while updating, Reason %v\n", updateErr)
		return updateErr
	}

	log.Printf("Customer data Upadted successfully for ID %d\n", pi.Cid)
	return nil
}

func (pi *Customerinfo) Updatelname(db *pg.DB) error {
	_, updateErr := db.Model(pi).Set("lname = ?lname").Where("cid= ?cid").Update()

	if updateErr != nil {
		log.Printf("Error while updating, Reason %v\n", updateErr)
		return updateErr
	}

	log.Printf("Customer data Upadted successfully for ID %d\n", pi.Cid)
	return nil
}

func (pi *Customerinfo) Updateaddress(db *pg.DB) error {
	_, updateErr := db.Model(pi).Set("address = ?address").Where("cid= ?cid").Update()

	if updateErr != nil {
		log.Printf("Error while updating, Reason %v\n", updateErr)
		return updateErr
	}

	log.Printf("Customer data Upadted successfully for ID %d\n", pi.Cid)
	return nil
}

func (pi *Customerinfo) Updatephone(db *pg.DB) error {
	_, updateErr := db.Model(pi).Set("phoneno = ?phoneno").Where("cid= ?cid").Update()

	if updateErr != nil {
		log.Printf("Error while updating, Reason %v\n", updateErr)
		return updateErr
	}

	log.Printf("Customer data Upadted successfully for ID %d\n", pi.Cid)
	return nil
}

func (pi *Customerinfo) Updateemail(db *pg.DB) error {
	_, updateErr := db.Model(pi).Set("emailid = ?emailid").Where("cid= ?cid").Update()

	if updateErr != nil {
		log.Printf("Error while updating, Reason %v\n", updateErr)
		return updateErr
	}

	log.Printf("Customer data Upadted successfully for ID %d\n", pi.Cid)
	return nil
}

func (pi *Customerinfo) GetById(db *pg.DB) Customerinfo {
	var item Customerinfo
	getErr := db.Model(&item).Where("cid = ?0", pi.Cid).Select()
	if getErr != nil {
		log.Printf("Error while querying value by id, Reason: %v\n", getErr)
	}
	log.Printf("Get by id successfull  %v\n", *pi)
	return item
}

func (pi *Customerinfo) DeleteRecord(db *pg.DB) error {
	//delete by primary key
	//db.delete(pi)
	_, deleteErr := db.Model(pi).Where("cid = ?cid").Delete()
	if deleteErr != nil {
		log.Printf("error deleteing for reason, %v\n", deleteErr)
		return deleteErr
	}

	log.Printf("Delete Successful for %s customer,\n ", pi.Fname)
	return nil
}

func (pi *Customerinfo) Save(db *pg.DB) error {
	insertErr := db.Insert(pi)
	if insertErr != nil {
		log.Printf("Error while insering new item into db, Reason: %v\n", insertErr)
		return insertErr
	}
	log.Printf("Customer details %s inserted successfully \n", pi.Cid)
	return nil
}

// func (pi *ProductItem) SaveMultiple(db *pg.DB, items []*ProductItem) error {
// 	_, insertErr := db.Model(items[0], items[1]).Insert()
// 	if insertErr != nil {
// 		log.Printf("Error while bulk inserting, resons: %v\n", insertErr)
// 		return insertErr
// 	}
// 	log.Printf("Bulk insert successful")
// 	return nil
// }

// func (pi *Customerinfo) Adding(db *pg.DB) (*Customerinfo, error) {
// 	InsertResult, insertErr := db.Model(pi).Returning("*").Insert()
// 	if insertErr != nil {
// 		log.Printf("Error while adding customer, Reason: %v\n", insertErr)
// 		return nil, insertErr
// 	}
// 	log.Printf("Customer Inserted Successfully \n")
// 	log.Printf("Customer details added to the bank %v\n", InsertResult)
// 	return pi, nil
// }

func CreateCustomerTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Customerinfo{}, opts)
	if createErr != nil {
		log.Printf("Error while creating tableItems, Reason: %v\n, createErr")
		return createErr
	}
	log.Printf("Customer Table is created\n")
	return nil
}
