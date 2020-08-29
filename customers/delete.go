package customers

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func deleteByID(id int) error {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	_, err = db.Exec("delete from customers where id = $1", id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//DeleteCustomerByIDHandler is handler function
func DeleteCustomerByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	err = deleteByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "customer deleted",
	})
}
