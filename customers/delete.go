package customers

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func deleteByID(id int) {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	_, err = db.Exec("delete from customer where id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}
}

//DeleteCustomerByIDHandler is handler function
func DeleteCustomerByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	deleteByID(id)
	c.JSON(http.StatusOK, gin.H{
		"message": "customer deleted", // cast it to string before showing
	})
}
