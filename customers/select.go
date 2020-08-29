package customers

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func selectByID(rowID int) (customer, error) {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	defer db.Close()
	stmt, err := db.Prepare("select id,name,email,status from customers where id=$1")
	if err != nil {
		log.Println(err)
		return customer{}, err
	}
	var id int
	var name, email, status string
	row := stmt.QueryRow(rowID)
	err = row.Scan(&id, &name, &email, &status)
	if err != nil {
		log.Println(err)
		return customer{}, err
	}
	return customer{
		ID:     id,
		Name:   name,
		Email:  email,
		Status: status,
	}, nil
}

func selectAll() ([]customer, error) {
	var cuss []customer
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	defer db.Close()
	stmt, err := db.Prepare("select id,name,email,status from customers")
	if err != nil {
		log.Fatal(err)
		return cuss, err
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
		return cuss, err
	}
	for rows.Next() {
		var id int
		var name, email, status string
		err = rows.Scan(&id, &name, &email, &status)
		if err != nil {
			log.Fatal(err)
			return cuss, err
		}
		cuss = append(cuss, customer{
			ID:     id,
			Name:   name,
			Email:  email,
			Status: status,
		})
	}
	return cuss, nil
}

//GetCustomerByIDHandler is function select where id
func GetCustomerByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	r, err := selectByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, r)
}

//GetAllCustomerHandler is function select where id
func GetAllCustomerHandler(c *gin.Context) {
	cus, err := selectAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	if cus == nil {
		cus = append(cus, customer{})
	}
	c.JSON(http.StatusOK, cus)
}
