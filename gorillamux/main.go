package main

import (
	"fmt"
	"time"

)

// Order represents the model for an order
// Default table name will be `orders`
type Order struct {
	// gorm.Model
	OrderID      uint      `json:"orderId" gorm:"primary_key"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignkey:OrderID"`
}
type Item struct {
	// gorm.Model
	LineItemID  uint   `json:"lineItemId" gorm:"primary_key"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"-"`
}

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("CREATE DATABASE orders_db")
	db.Exec("USE orders_db")

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&Order{}, &Item{})
}


dataSourceName := "your_username:your_password@tcp(localhost:3306)/your_database_name?parseTime=True"
/* Commenting out, so that order_db is not used, and your_database_name is used instead
db.Exec("CREATE DATABASE orders_db")
db.Exec("USE orders_db")
*/

func main() {
    router := mux.NewRouter()
    // Create
    router.HandleFunc("/orders", createOrder).Methods("POST")
    // Read
    router.HandleFunc("/orders/{orderId}", getOrder).Methods("GET")
    // Read-all
    router.HandleFunc("/orders", getOrders).Methods("GET")
    // Update
    router.HandleFunc("/orders/{orderId}", updateOrder).Methods("PUT")
    // Delete
    router.HandleFunc("/orders/{orderId}", deleteOrder).Methods("DELETE")
    // Initialize db connection
    initDB()

    log.Fatal(http.ListenAndServe(":8080", router))
}
