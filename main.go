package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Animal struct {
	Id           int    `json:"id" gorm:"column:id;"`
	NameVN       string `json:"nameVN" gorm:"column:nameVN;"`
	NameEN       string `json:"nameEN" gorm:"column:nameEN;"`
	Size         string `json:"size" gorm:"column:size;"`
	Color        string `json:"color" gorm:"column:color;"`
	Behavior     string `json:"behavior" gorm:"column:behavior;"`
	Longevity    string `json:"longevity" gorm:"column:longevity;"`
	Temperature  string `json:"temperature" gorm:"column:temperature;"`
	PH           string `json:"pH" gorm:"column:pH;"`
	Food         string `json:"food" gorm:"column:food;"`
	Reproduction string `json:"reproduction" gorm:"column:reproduction;"`
}

func (Animal) TableName() string { return "animals" }

func main() {
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/aquarium_list_animal?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(db)

	r := gin.Default()
	// GET /v1/animals
	v1 := r.Group("/v1")
	{
		animals := v1.Group("/animals")
		{
			animals.GET("", GetAllAnimals(db))
			animals.GET("/:id")
		}
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
func GetAllAnimals(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var result [] Animal

		if err := db.Find(&result).Error; err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}
