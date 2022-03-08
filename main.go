package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

const (
	host     = "belajar.mysql.database.azure.com"
	database = "guest"
	user     = "steve@belajar"
	password = "EmbromSkolkov25"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type fruit struct {
	ID    int    `json:"id"`
	Title string `json:"title"`

	Quantity int `json:"quantity"`
}

var fruits = []fruit{
	{ID: 1, Title: "In Search of Lost Time", Quantity: 2},
	{ID: 2, Title: "The Great Gatsby", Quantity: 5},
	{ID: 3, Title: "War and Peace", Quantity: 6},
}

//
//func update() {
//	rows, err := db.Exec("UPDATE inventory SET quantity = ? WHERE name = ?", 250, "banana")
//	checkError(err)
//	rowCount, err := rows.RowsAffected()
//	fmt.Printf("Updated %d row(s) of data.\n", rowCount)
//	fmt.Println("Done.")
//}
//func delete() {
//	rows, err := db.Exec("DELETE FROM inventory WHERE name = ?", "orange")
//	checkError(err)
//	rowCount, err := rows.RowsAffected()
//	fmt.Printf("Deleted %d row(s) of data.\n", rowCount)
//	fmt.Println("Done.")
//}
//func read(c *gin.Context) {
//	var (
//		id       int
//		name     string
//		quantity int
//	)
//	rows, err := db.Query("SELECT id, name, quantity from inventory;")
//	checkError(err)
//	defer rows.Close()
//	fmt.Println("Reading data:")
//	for rows.Next() {
//		err := rows.Scan(&id, &name, &quantity)
//		checkError(err)
//		fmt.Printf("Data row = (%d, %s, %d)\n", id, name, quantity)
//	}
//	err = rows.Err()
//	checkError(err)
//	fmt.Println("Done.")
//
//}

//func insert() {
//	rows, err := db.Prepare("INSERT INTO inventory(id, quantity, name) VALUES( 17, 4543, 'atea'  )")
//	checkError(err)
//	rowCount, err := rows.Query()
//	fmt.Printf("Inserted %v row(s) of data.\n", rowCount)
//	fmt.Printf("Inserted data.\n")
//	fmt.Println("  Done. ")
//	defer rows.Close()
//}

func getbook(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fruits)
}
func main() {

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")
	var (
		id       int
		name     string
		quantity int
	)
	rows, err := db.Query("SELECT id, name, quantity from inventory;")
	checkError(err)
	defer rows.Close()
	fmt.Println("Reading data:")
	for rows.Next() {
		err := rows.Scan(&id, &name, &quantity)
		checkError(err)
		fmt.Printf("Data row = (%d, %s, %d)\n", id, name, quantity)
		fruits = append(fruits, fruit{ID: id, Quantity: quantity, Title: name})
	}
	err = rows.Err()
	checkError(err)
	fmt.Println("Done.")

	router := gin.Default()
	router.GET("/read2", getbook)
	router.GET("/", getbook)
	router.Run()

}
