package appevendor

import (
	"evendor.com/go/initializers"
	"evendor.com/go/routes"
	"github.com/gin-gonic/gin"
)

func EvendorApp() {

	app := gin.Default()
	initializers.NowDB()
	routes.InitRoutes(app)

}
