package main

/*
curl request
curl --request GET \
  --url http://localhost:8085/rest/names \
  --header 'authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ='  //the encoded string of usernam:password
*/
import (
	"encoding/base64"
	"encoding/json"
	//"io/ioutil"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	Names = []string{"Bharath Kumar", "Jon Snow"}
)

type Name struct {
	Name string `json:"name",omitempty`
}

func authMiddleware() gin.HandlerFunc {
	//authenticate the creds here
	return func(c *gin.Context) {

		s := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			c.JSON(http.StatusBadRequest, "Authorization header is not passed")
			c.Abort()
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Auth Decode error")
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
	c.JSON(http.StatusOK, Names)
}

func putNames(c *gin.Context) {
	var name Name
	//c.BindJSON(&name) //this can be used instead of unmarshal
	//bodyByte, err := ioutil.ReadAll(c.Request.Body)
	bodyByte, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Check the body sent")
		return
	}
	err = json.Unmarshal(bodyByte, &name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Check the body sent")
		return
	}

	Names = append(Names, name.Name)
	c.JSON(http.StatusAccepted, Names)
}

func main() {
	//create a new gin router
	router := gin.New()
	//create a new router group, this will create localhost:<port>/rest/names
	names := router.Group("/rest/names")
	//All middlewares must be passed will come under USE
	names.Use(authMiddleware())
	names.GET("", getNames)
	names.POST("", putNames)

	//start the server at port 8085
	router.Run(":8085")
}
