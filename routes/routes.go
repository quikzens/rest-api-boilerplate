package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quikzens/rest-api-boilerplate/domain/users"
)

func Set(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/users/register", users.UserRegister)
	api.POST("/users/login", users.UserLogin)
	api.GET("/users/auth", users.VerifyAuth, users.UserCheckAuth)
}
