package customers

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func update(id int, cus customer) customer {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	defer db.Close()
	stmt, err := db.Prepare("update customers set name = $2 ,email = $3 ,status = $4 where id=$1")
	if err != nil {
		log.Println(err)
		return customer{}
	}
	if _, err := stmt.Exec(id, cus.Name, cus.Email, cus.Status); err != nil {
		log.Println(err)
		return customer{}
	}
	cus.ID = id
	return cus
}

//PutUpdateCustomerHandler is handler function
func PutUpdateCustomerHandler(c *gin.Context) {
	var cus customer
	err := c.ShouldBindJSON(&cus)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	r := update(id, cus)
	c.JSON(http.StatusOK, r)
}
