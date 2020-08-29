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

func selectByID(rowID int) customer {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	defer db.Close()
	stmt, err := db.Prepare("select id,name,email,status from customers where id=$1")
	if err != nil {
		log.Println(err)
		return customer{}
	}
	var id int
	var name, email, status string
	row := stmt.QueryRow(rowID)
	err = row.Scan(&id, &name, &email, &status)
	if err != nil {
		log.Println(err)
		return customer{}
	}
	log.Println(rowID)
	log.Println(id, name, status)
	return customer{
		ID:     id,
		Name:   name,
		Email:  email,
		Status: status,
	}
}

func selectAll() []customer {
	var cuss []customer
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	defer db.Close()
	stmt, err := db.Prepare("select id,name,email,status from customers")
	if err != nil {
		log.Fatal(err)
		return cuss
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
		return cuss
	}
	for rows.Next() {
		var id int
		var name, email, status string
		err = rows.Scan(&id, &name, &email, &status)
		if err != nil {
			log.Fatal(err)
			return cuss
		}
		cuss = append(cuss, customer{
			ID:     id,
			Name:   name,
			Email:  email,
			Status: status,
		})
	}
	return cuss
}

//GetCustomerByIDHandler is function select where id
func GetCustomerByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, selectByID(id))
}

//GetAllCustomerHandler is function select where id
func GetAllCustomerHandler(c *gin.Context) {
	cus := selectAll()
	if cus == nil {
		cus = append(cus, customer{})
	}
	c.JSON(http.StatusOK, cus)
}
