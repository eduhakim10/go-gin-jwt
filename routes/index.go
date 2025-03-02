package routes

import (
	"fmt"
	"evendor.com/go/controllers/usercontroller"
	"evendor.com/go/middleware"
	"github.com/gin-gonic/gin"

	

)



func InitRoutes(app *gin.Engine) {

	app.Use(func(c *gin.Context) {
		fmt.Println("Incoming request:", c.Request.Method, c.Request.URL.Path)
		fmt.Println("Request headers:", c.Request.Header)
		c.Next()
	})
	app.GET("api/users/",middleware.RequireAuth, usercontroller.Index)
	app.POST("api/users/signup", usercontroller.SignUp)
	app.POST("api/users/login", usercontroller.SignIn)
	app.POST("api/users/changepassword",middleware.RequireAuth, usercontroller.ChangePassword)
	app.GET("api/users/show/:id",middleware.RequireAuth,  usercontroller.Show)
	app.GET("api/users/show", middleware.RequireAuth,  usercontroller.Show)
	app.POST("api/users/update/:id",middleware.RequireAuth, usercontroller.Update)


	app.GET("api",middleware.RequireAuth, usercontroller.Verify)

	//app.GET("api/validate", middleware.RequireAuth, usercontroller.Verify)

	//app.Run(":8000")

}
