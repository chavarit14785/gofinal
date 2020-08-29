package customers

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (cus customer) update(id int) (customer, error) {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return customer{}, err
	}
	stmt, err := db.Prepare("update customers set name = $2 ,email = $3 ,status = $4 where id=$1")
	if err != nil {
		log.Println(err)
		return customer{}, err
	}
	if _, err := stmt.Exec(id, cus.Name, cus.Email, cus.Status); err != nil {
		log.Println(err)
		return customer{}, err
	}
	cus.ID = id
	return cus, nil
}

//PutUpdateCustomerHandler is handler function
func PutUpdateCustomerHandler(c *gin.Context) {
	var cus customer
	err := c.ShouldBindJSON(&cus)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	r, err := cus.update(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, r)
}
