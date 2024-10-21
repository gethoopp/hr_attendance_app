package services

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func User_data_entry(c *gin.Context) {

	// var Userdata Userdata

	if err := c.ShouldBindJSON(""); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/hr_attendance_app")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server errror"})
	}

	defer db.Close()

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(10)

	// ctx := context.Background()

	// query := "UPDATE users where rfid_tag = ? "
}

func Attendace_user(c *gin.Context) {

}
