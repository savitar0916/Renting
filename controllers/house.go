package controllers

import (
	houseModel "Renting/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetAllHouses _
func GetAllHouses(c *gin.Context) {
	result, err := houseModel.GetAllHouses()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "The database's reading is failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{

		"data": result,
	})
}

//GetHouse _
func GetHouse(c *gin.Context) {

	houseID, errID := primitive.ObjectIDFromHex(c.Param("houseid"))
	if errID != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "Get Id failed",
		})
	}

	result, err := houseModel.GetHouse(houseID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    2,
		"message": result,
	})
}

//AddHouse _
func AddHouse(c *gin.Context) {
}

//UpdateHouse _
func UpdateHouse(c *gin.Context) {
	var house houseModel.House

	houseID, err := primitive.ObjectIDFromHex(c.Param("houseid"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "Get Id failed",
		})

	}
	price := c.PostForm("price")
	fmt.Println(c.PostForm("price"))
	changeprice, err := strconv.ParseInt(price, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "Change Type failed",
		})
	}
	//fmt.Println(changeprice)
	house.Price = changeprice
	update := house.UpdateHouse(houseID)

	if update != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    2,
			"message": update,
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"code":    3,
		"message": "Update price success",
	})
}

//DeleteHouse _
func DeleteHouse(c *gin.Context) {
	var house houseModel.House
	houseID, err := primitive.ObjectIDFromHex(c.Param("houseid"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err,
		})
	}
	err = house.DeleteHouse(houseID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    2,
		"message": "delete success",
	})

}
