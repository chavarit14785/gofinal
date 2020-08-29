package customers

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func insert(cus customer) customer {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	defer db.Close()
	row := db.QueryRow("insert into customers (name, email, status) values ($1,$2,$3) returning id", cus.Name, cus.Email, cus.Status)
	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Println(err)
		return customer{}
	}
	return customer{
		ID:     id,
		Name:   cus.Name,
		Email:  cus.Email,
		Status: cus.Status,
	}
}

func createTable() {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	defer db.Close()
	createTb := `
	create table if not exists customers (
		id serial primary key,
		name text,
		email text,
		status text
	);
	`
	_, err = db.Exec(createTb)

	if err != nil {
		log.Println(err)
		return
	}
}

//PostCreateCustomerHandler is handler function
func PostCreateCustomerHandler(c *gin.Context) {
	var cus customer
	err := c.ShouldBindJSON(&cus)
	if err != nil {
		return
	}
	createTable()
	r := insert(cus)
	c.JSON(http.StatusCreated, r)
}
