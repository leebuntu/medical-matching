package user

import "github.com/gin-gonic/gin"

func GetUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GetUserProfile",
		})
	}
}

func UpdateUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "UpdateUserProfile",
		})
	}
}

func DeleteUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DeleteUserProfile",
		})
	}
}
