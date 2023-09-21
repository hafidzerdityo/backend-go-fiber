package datastore

import "hafidzresttemplate.com/data"

func(d *DatastoreSetup) CheckUser()(queryResult []data.Book, err error){
	queryResult = []data.Book{
			{
				ID:     "1",
				Title:  "Knight in Shining Armor",
				Author: "Hafidz Erdityo",
			},
			{
				ID:     "2",
				Title:  "The Great Gatsby",
				Author: "F. Scott Fitzgerald",
			},
			{
				ID:     "3",
				Title:  "To Kill a Mockingbird",
				Author: "Harper Lee",
			},
			{
				ID:     "4",
				Title:  "1984",
				Author: "George Orwell",
			},
			{
				ID:     "5",
				Title:  "Pride and Prejudice",
			},
			{
				ID:     "6",
				Title:  "The Catcher in the Rye",
				Author: "J.D. Salinger",
			},
		}

	return
}

func(d *DatastoreSetup) CreateUser()(queryResult []map[string]interface{}, err error){
	queryResult  = []map[string]interface{}{
		{"hello":1},
		{"hello":2},
		{"hello":23},
	}

	return
}