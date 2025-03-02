package appevendor

import (
	"fmt"
	"os"
	"strings"
	"evendor.com/go/initializers"
	"evendor.com/go/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
	"github.com/joho/godotenv"
	"log"
)

func EvendorApp() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")

	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	initializers.NowDB()
	routes.InitRoutes(app)

	fmt.Println("Server running on :8000")
	app.Run(":8000") // Pastikan pakai `app`, bukan `route`

}
