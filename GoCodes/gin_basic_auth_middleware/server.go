package main

/*
curl request

curl --request GET \
  --url http://localhost:8085/rest/names \
  --header 'authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=' \
  --data '{
	"username":"admin",
	"password":"pass"
}'

*/
import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	//authenticate the creds here
	return func(c *gin.Context) {

		s := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			c.JSON(http.StatusUnauthorized, "Authorization header is not correct")
			c.Abort()
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Decode error")
			c.Abort()
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			c.JSON(http.StatusUnauthorized, "Decoded string is not correct")
			c.Abort()
			return
		}

		if pair[0] != "username" || pair[1] != "password" {
			c.JSON(http.StatusUnauthorized, "Username/password is not correct")
			c.Abort()
			return
		}

	}
}

func getNames(c *gin.Context) {
	var names []string

	names = []string{"bharath", "kumar"}
	c.JSON(http.StatusOK, names)
}

func main() {
	//create a new gin router
	router := gin.New()
	//create a new router group, this will create localhost:<port>/rest
	rest := router.Group("/rest")
	//create a new group under /rest group, this will create localhost:<port>/rest/names
	names := rest.Group("/names")
	//All middlewares must be passed will come under USE
	names.Use(AuthMiddleware())

	names.GET("", getNames)

	//start the server at port 8085
	router.Run(":8085")
}
