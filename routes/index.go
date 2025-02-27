package routes

import (

	"evendor.com/go/controllers/usercontroller"
	"evendor.com/go/middleware"
	"github.com/gin-gonic/gin"

)



func InitRoutes(app *gin.Engine) {

	route := gin.Default()
	route.GET("api/users",middleware.RequireAuth, usercontroller.Index)
	route.POST("api/users/signup", usercontroller.SignUp)
	route.POST("api/users/login", usercontroller.SignIn)
	route.POST("api/users/changepassword",middleware.RequireAuth, usercontroller.ChangePassword)
	route.GET("api/users/show/:id",middleware.RequireAuth,  usercontroller.Show)
	route.GET("api/users/show", middleware.RequireAuth,  usercontroller.Show)
	route.POST("api/users/update/:id",middleware.RequireAuth, usercontroller.Update)

	route.GET("api/validate", middleware.RequireAuth, usercontroller.Verify)

	route.Run(":8000")

}
