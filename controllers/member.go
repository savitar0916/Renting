package controllers

import (
	memberModel "Renting/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	// for BSON ObjectID
)

// AddMember _ Create New Member
func AddMember(c *gin.Context) {
	var member memberModel.Member
	member.Account = c.PostForm("account")
	member.Name = c.PostForm("name")
	Password := c.PostForm("password")
	hash, hasherr := HashPassword(Password)
	if hasherr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "Hash Password failed",
		})
		return
	}
	member.Password = hash
	/*member.Password = c.PostForm("password")*/
	member.Phone = c.PostForm("phone")
	member.Sex = c.PostForm("sex")

	err := member.AddMember()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "Add new member failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    2,
		"message": "Add new member success",
	})
}

// Login _
func Login(c *gin.Context) {
	var member memberModel.Member
	member.Account = c.PostForm("account")
	member.Password = c.PostForm("password")
	//test1, _ := HashPassword("123456")
	data, err := member.Login()
	if err != nil {

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "Get data failed",
			//"Account": member.Account,
		})
		return
	}
	result := CheckPasswordHash(member.Password, data.Password)
	if result != true {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "Login failed",
		})
		return
	}
	//test2, _ := HashPassword("123456")
	c.JSON(http.StatusOK, gin.H{
		"code": 2,
		"data": result,
	})
	//fmt.Println(test1)
	//fmt.Println(test2)
}

// HashPassword _
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash _
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
