package dao

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"roomies/code/shared"
	"time"
)

type DummyGroceryListDAO struct {
	GroceryListDAO
	dummyList *shared.GroceryList
}

func GetNewDummyGroceryListDAO() *DummyGroceryListDAO {
	l := shared.GetNewGroceryList([]*shared.GroceryListItem{
		{
			ItemID:        "1111",
			GroupID:       "group1",
			ListItem:      "eggs",
			AddedByID:     "user",
			Recurring:     false,
			RecurringDate: time.Time{},
			Timestamp:     time.Time{},
		}, {
			ItemID:        "2222",
			GroupID:       "group1",
			ListItem:      "milk",
			AddedByID:     "user",
			Recurring:     false,
			RecurringDate: time.Time{},
			Timestamp:     time.Time{},
		}, {
			ItemID:        "3333",
			GroupID:       "group1",
			ListItem:      "flour",
			AddedByID:     "user",
			Recurring:     false,
			RecurringDate: time.Time{},
			Timestamp:     time.Time{},
		}})

	return &DummyGroceryListDAO{
		dummyList: l,
	}
}

func setUpDummyData(db *sql.DB) (sql.Result, error) {
	result, err := db.Exec("INSERT INTO GroceryList (listID, itemID, groupID, listItem, addedByID, recurring, recurringDate, timestamp) VALUES (?,?,?,?,?,?,?,?)",
		shared.GroceryListItem{
			ListID:        "1",
			ItemID:        "12",
			GroupID:       "123",
			ListItem:      "TEST",
			AddedByID:     "me",
			Recurring:     false,
			RecurringDate: time.Time{},
			Timestamp:     time.Time{},
		})
	if err != nil {
		return nil, err
	}
	return result, err
}

func (d *DummyGroceryListDAO) GetList(listID string) *shared.GroceryList {
	cfg := mysql.Config{
		Net:    "tcp",
		User:   "jared@smarty.com",
		Passwd: "GeneticNetwork1!",
		Addr:   "127.0.0.1:3306",
		DBName: "Roomies",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(err)
	}

	//result, err := setUpDummyData(db)
	//if err != nil {
	//	panic(err)
	//}
	//if result == nil {
	//	panic(err)
	//}

	rows, err := db.Query("SELECT * FROM GroceryList WHERE listID = ?", listID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var list []*shared.GroceryListItem

	for rows.Next() {
		var item *shared.GroceryListItem
		if err := rows.Scan(
			&item.ItemID, &item.ListItem, &item.Timestamp,
			&item.AddedByID, &item.GroupID, &item.Recurring,
			&item.RecurringDate); err != nil {
			return nil
		}
		list = append(list, item)
	}

	return shared.GetNewGroceryList(list)
	//return d.dummyList
}

func (d *DummyGroceryListDAO) GetListItem(listID, itemID string) *shared.GroceryListItem {
	return d.dummyList.GetItem(itemID)
}

func (d *DummyGroceryListDAO) CreateList() {
	fmt.Println("CreateList not implemented for dummy")
}
func (d *DummyGroceryListDAO) AppendListItem(listID string, item *shared.GroceryListItem) {
	fmt.Println("AppendListItem not implemented for dummy")
}
func (d *DummyGroceryListDAO) DeleteListItem(listID, itemID string) {
	fmt.Println("Don't delete the dummy list")
}
