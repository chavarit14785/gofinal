package customers

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func (cus customer) insert() (customer, error) {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return customer{}, err
	}
	row := db.QueryRow("insert into customers (name, email, status) values ($1,$2,$3) returning id", cus.Name, cus.Email, cus.Status)
	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Println(err)
		return customer{}, err
	}
	return customer{
		ID:     id,
		Name:   cus.Name,
		Email:  cus.Email,
		Status: cus.Status,
	}, nil
}

//CreateTable is init function for customers
func CreateTable() error {
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
		return err
	}
	return nil
}

//PostCreateCustomerHandler is handler function
func PostCreateCustomerHandler(c *gin.Context) {
	var cus customer
	err := c.ShouldBindJSON(&cus)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	r, err := cus.insert()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusCreated, r)
}
