package controllers

import (
	"auth-go-test/initializers"
	"auth-go-test/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

//reigistering the user
func SignUp(c *gin.Context ){ // it will take the request and response

	 //get the email/pass from req body 
	 var body struct{
		Email string
		Password string 
	 }
	
	if c.Bind(&body) != nil { // if there is an err in reading the body we send a bad request response
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}
	// hashed the pass 

  hash , err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// create user 
	user:=models.User{Email: body.Email,Password: string(hash)} 
	result:= initializers.DB.Create(&user) // we are creating a table by refering to the User model

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	// respond with user 

    c.JSON(http.StatusOK, gin.H{})
}

//login the user 
 func Login(c *gin.Context){
	 //GET THE EMAIL/PASS FROM REQ BODY 
	 var body struct{
		Email string
		Password string 
	 }
	
	if c.Bind(&body) != nil { // if there is an err in reading the body we send a bad request response
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}
	// CHECK IF USER EXISTS IN THE DB 
	var user models.User
	initializers.DB.First(&user, "email= ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	//COMPARE THE PASSWORD
	err:=bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	//GENERATING JWT TOKEN
	
	// Create a new token object, specifying signing method and the claims
// you would like it to contain.
token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"sub": user.ID,
	"exp": time.Now().Add(24 * time.Hour * 30).Unix(),
})
// Sign and get the complete encoded token as a string using the secret
tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
if err!=nil{
	c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
}
//send it back
c.SetSameSite(http.SameSiteLaxMode)//to prevent cross site request forgery attacks
c.SetCookie("Authorization",tokenString, 60 * 60 * 24 * 30, "", "", false, true) // 30 days duration

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
 }

 //validate user

 func Validate(c *gin.Context){

	c.JSON(http.StatusOK,gin.H{
		"message": "You are logged in",
	})
	
 }	