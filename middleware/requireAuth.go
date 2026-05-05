package middleware

import (
	"auth-go-test/initializers"
	"auth-go-test/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// it will give us the access to the req/res/next 
func RequireAuth(c *gin.Context){

	//get the cookie/token

	tokenString , err:=c.Cookie("Authorization")
	if err != nil{
		c.AbortWithStatus(http.StatusUnauthorized) 
		return
	}


	//decode/validate the token
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
	
	return []byte(os.Getenv("SECRET")), nil
}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))


if err != nil {
	c.AbortWithStatus(http.StatusUnauthorized)
	return
}

if claims, ok := token.Claims.(jwt.MapClaims); ok {


	//check the expiry date
	if float64(time.Now().Unix()) > claims["exp"].(float64){
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//find the user with token sub 
    var user models.User 
	initializers.DB.First(&user, claims["sub"])
	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	//attach request
	c.Set("user",user) // user is available in the request context for the next handler 

    //continue
	
} else {
	c.AbortWithStatus(http.StatusUnauthorized) 
	return
}
	c.Next()
	
}