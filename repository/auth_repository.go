package repository

import (
	"app-ecommerce-server/data/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthRepository interface {
	SignUp(c *gin.Context, db *gorm.DB, user *entity.User) *entity.User
	CheckEmail(c *gin.Context, db *gorm.DB, email string) bool
	CheckUsername(c *gin.Context, db *gorm.DB, username string) bool
}
