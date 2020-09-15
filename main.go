package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Notifications struct {
	ID     string `json:"id"`
	Notif  string `json:"notification"`
	UserID string `json:"user_id"`
}

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/timesheet")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	//Insert data to database
	// insert, err := db.Query("INSERT INTO notifications (notification, user_id) VALUES ('jajang', '4')")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	// fmt.Println("sussces insert")

	// Execute the query
	results, err := db.Query("SELECT * FROM notifications")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var notifications Notifications
		// for each row, scan the result into our tag composite object
		err = results.Scan(&notifications.ID, &notifications.Notif, &notifications.UserID)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the notification's Name attribute
		fmt.Println(notifications.Notif)
		fmt.Println(notifications.UserID)
	}

}
