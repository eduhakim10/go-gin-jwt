package usercontroller

import (
	"fmt"
	"net/http"
	"time"

	"evendor.com/go/initializers"
	"evendor.com/go/models"
	"evendor.com/go/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("secret-key")
var (
	jwtKey = []byte("your_secret_key") // Define jwtKey as a []byte
)

type Claims struct {
	UserID uint `json:"user_id"`
}
type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Age  uint
}

func SignUp(c *gin.Context) {

	//	get email/pass from body
	//UserReq := new(request.UserRequest)

	var body struct {
		Nama     string
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed cuy"})

		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed cuy"})
	}
	hashedPasswordString := string(hashedPassword)

	// fmt.Println("Hashed password:", hashedPasswordString)
	// return
	//create User
	user := new(models.User)
	user.Nama = &body.Nama
	user.Email = &body.Email
	user.Password = &hashedPasswordString

	errDb := initializers.DB.Table("users").Create(&user).Error
	if errDb != nil {

		c.JSON(500, gin.H{
			"Message": "Cant Create data",
		})

		return

	}
	c.JSON(200, gin.H{
		"Message": "Data saved",
	})

}

func SignIn(c *gin.Context) {

	// get email password reques body
	var body struct {
		Email    string
		Password string
	}
	// check email
	if c.Bind(&body) != nil {
		c.JSON(500, gin.H{
			"Message": "need body request",
		})
		return
	}
	var user models.User
	// var user = User{emai: 10};

	result := initializers.DB.Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			//	fmt.Println("Email does not exist!")
			c.JSON(500, gin.H{
				"Message": "Email not found",
			})
			return
			// Handle case when the email doesn't exist
		} else {
			//log.Fatal(result.Error)
			c.JSON(500, gin.H{
				"Message": "Hushno",
			})
		}
	}

	hashedPasswordString := string(*user.Password)
	errz := bcrypt.CompareHashAndPassword([]byte(hashedPasswordString), []byte(body.Password))

	if errz != nil {
		//fmt.Println("Password does not match")
		c.JSON(200, gin.H{
			//	"message": hashedPasswordString,
			"error": " password does match",
		})
		return
	}

	tokenString, err := generateToken(1) // Example user ID
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		c.JSON(200, gin.H{
			"Message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"token": tokenString,
	})

}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Token not provided", http.StatusUnauthorized)
		return
	}

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	// Check for errors in parsing the token
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Validate token and claims
	if !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		w.Write([]byte(fmt.Sprintf("Welcome, %s!", username)))
	} else {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
	}
}
func Verify(c *gin.Context) {

	c.JSON(200, gin.H{
		"Message": "I logged in",
	})

}
func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func Profile(w http.ResponseWriter, r *http.Request) {

}
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}

	fmt.Fprint(w, "Welcome to the the protected area")

}

func generateToken(userID uint) (string, error) {

	//var secretKey = "r434rwre"
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = 123
	claims["username"] = "example_user"
	claims["exp"] = time.Now().Add(time.Hour).Unix() // Token expires in 1 hour
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	fmt.Println("Generated token:", tokenString) // Print the generated token for debugging
	return tokenString, nil
	//return token.SignedString(secretKey)
}
func Index(c *gin.Context) {
	//	fmt.Println("connect cuy")
	var users []models.User
	//c.String(200, "Ini controller fung index  \n")
	initializers.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}
func Show(c *gin.Context) {
	//	fmt.Println("connect cuy")
	var users []models.User

	//var users []User
	id := c.Param("id")

	// If ID is provided, retrieve user by ID
	if id != "" {
		var user []models.User
		result := initializers.DB.First(&user, id)
		if result.Error != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(200, user)
		return
	}
	//c.String(200, "Ini controller fung show \n")
	initializers.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}
func Store(c *gin.Context) {

	UserReq := new(request.UserRequest)
	if errReq := c.ShouldBind(&UserReq); errReq != nil {
		c.JSON(400, gin.H{
			"Message": errReq.Error(),
		})

	}

	// //	fmt.Println("connect cuy")
	user := new(models.User)
	user.Nama = &UserReq.Nama
	user.Email = &UserReq.Email
	user.Password = &UserReq.Password

	errDb := initializers.DB.Table("users").Create(&user).Error
	if errDb != nil {

		c.JSON(500, gin.H{
			"Message": "Cant Create data",
		})

		return

	}
	c.JSON(200, gin.H{
		"Message": "Data saved",
	})
}
func Update(c *gin.Context) {

	var user models.User
	//user := new(models.User)

	var requestBody struct {
		Nama  string `json:"nama"`
		Email string `json:"email"`
	}

	ID := c.Param("id")

	//UserReq := new(request.UserRequest)

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := initializers.DB.First(&user, ID)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	user.Nama = &requestBody.Nama
	user.Email = &requestBody.Email

	if err := initializers.DB.Model(&User{}).Where("id = ?", ID).Updates(map[string]interface{}{"nama": requestBody.Nama, "email": requestBody.Email}).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(200, gin.H{
		"Message": "Data saved",
	})
}

// $2a$10$urLlAvWcB0mOknMBozPrWOD4jZRyFENEb3ywqSjrpRoq5PN0j/yiK
func ChangePassword(c *gin.Context) {

	// get email password reques body
	var body struct {
		ID          uint
		OldPassword string
		NewPassword string
	}

	if c.Bind(&body) != nil {
		c.JSON(500, gin.H{
			"Message": "need body request",
		})
		return
	}
	var user models.User
	// var user = User{emai: 10};

	result := initializers.DB.Where("id = ?", body.ID).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			//	fmt.Println("Email does not exist!")
			c.JSON(500, gin.H{
				"Message": "User not found",
			})
			return
			// Handle case when the email doesn't exist
		}
	}

	hashedPasswordString := string(*user.Password)
	errz := bcrypt.CompareHashAndPassword([]byte(hashedPasswordString), []byte(body.OldPassword))

	if errz != nil {
		//fmt.Println("Password does not match")
		c.JSON(200, gin.H{
			//	"message": hashedPasswordString,
			"error": "Old password does match",
		})
		return
	}

	NewhashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed cuy"})
	}
	NewhashedPasswordString := string(NewhashedPassword)

	// fmt.Println("Hashed password:", hashedPasswordString)
	// return
	//create User

	//user.Nama = &requestBody.Nama
	//user.Password = NewhashedPasswordString

	if err := initializers.DB.Model(&User{}).Where("id = ?", body.ID).Updates(map[string]interface{}{"password": NewhashedPasswordString}).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(200, gin.H{
		"Message": "Password Changed",
	})

}
