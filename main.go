package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

const (
	host     = "belajar.mysql.database.azure.com"
	database = "guest"
	user     = "steve@belajar"
	password = "EmbromSkolkov25"
)

var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type fruit struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Quantity int    `json:"quantity"`
}

var fruits = []fruit{}
var db = sql.DB{}

func update() {
	rows, err := db.Exec("UPDATE inventory SET quantity = ? WHERE name = ?", 250, "banana")
	checkError(err)
	rowCount, err := rows.RowsAffected()
	fmt.Printf("Updated %d row(s) of data.\n", rowCount)
	fmt.Println("Done.")
}

func read(c *gin.Context) {
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
	c.IndentedJSON(http.StatusOK, fruits)
}
func delete(c *gin.Context) {
	idP := c.Param("id")
	id, _ := strconv.Atoi(idP)

	rows, err := db.Exec("DELETE FROM inventory WHERE id = ?", id)
	checkError(err)
	rowCount, err := rows.RowsAffected()
	fmt.Printf("Deleted %d row(s) of data.\n", rowCount)
	fmt.Println("Done.")
	//c.IndentedJSON(http.StatusOK, "orange")
}
func insert(c *gin.Context) {
	titleP := c.Query("title")
	fmt.Sprintf(titleP)
	idP := c.Query("id")
	quantityP := c.Query("quantity")
	title := titleP
	id, _ := strconv.Atoi(idP)
	quantity, _ := strconv.Atoi(quantityP)

	var query = fmt.Sprintf("INSERT INTO inventory(id, quantity, name) VALUES(%v,%v,%v);", id, quantity, title)
	fmt.Println(query)
	var rows, err = db.Prepare(query)
	checkError(err)
	rowCount, err := rows.Query()
	fmt.Printf("Inserted %v row(s) of data.\n", rowCount)
	fmt.Printf("Inserted data.\n")
	fmt.Println("  Done. ")

	fruits = append(fruits, fruit{ID: id, Quantity: quantity, Title: title})

	c.String(200, title)
}

func getfruits(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fruits)
}
func getFruitById(c *gin.Context) {
	id := c.Param("id")
	getid, _ := strconv.Atoi(id)
	fruit2, err := searchFruitById(getid)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, fruit2)
}

func searchFruitById(id int) (*fruit, error) {
	for i, b := range fruits {
		if b.ID == id {
			return &fruits[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func main() {
	AZdb, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer AZdb.Close()
	err = AZdb.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")
	db = *AZdb
	router := gin.Default()
	router.GET("/fruits/:id", getFruitById)
	router.GET("/read", read)
	router.GET("/", getfruits)
	router.POST("/insert", insert)
	router.DELETE("/delete/:id", delete)
	router.Run("localhost:5000")

}
