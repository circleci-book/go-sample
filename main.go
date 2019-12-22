package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Product describes product
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/products", fetchAllProducts)

	return r
}

func setupDB() (*gorm.DB, error) {
	// user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local
	database, err := gorm.Open("mysql", os.Getenv("DATABASE_URL")+"?charset=utf8&parseTime=True&loc=Local")

	if err == nil {
		database.AutoMigrate(&Product{})
		return database, err
	}

	return nil, err
}

func dropDB() {
	db.DropTable(&Product{})
}

var db *gorm.DB

func main() {
	var err error

	db, err = setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := setupRouter()
	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// fetchAllProducts fetch all products
func fetchAllProducts(c *gin.Context) {
	var products []Product

	db.Find(&products)

	if len(products) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No product found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": products})
}
