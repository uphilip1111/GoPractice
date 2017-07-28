package main

import (
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var DB = make(map[string]string)

func init() {
	db, err := sql.Open("mysql", "root:admin@/project")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmIns, err := db.Prepare("insert into project.user_account values(?,?,?,?,?,?)")
	_, err = stmIns.Exec(1, "acc", "pwd", "name", "0800-000-000", "AuroraTech")

	defer stmIns.Close()

}

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/api/account/login/:acc/*action", func(context *gin.Context) {
		acc := context.Param("acc")
		action := context.Param("action")
		message := "acc: " + acc + " , action: " + action
		context.String(http.StatusOK, message)
	})

	r.GET("api/account/login", func(context *gin.Context) {
		acc := context.DefaultQuery("acc", "123") //DefaultQuery 沒參數(key不見) 預設用後面那個，用Query 沒參數 就是空白
		pwd := context.Query("pwd")
		msg := acc + "," + pwd
		context.String(http.StatusOK, msg)
	})

	v1 := r.Group("/v1")

	v1.GET("api/login", func(c *gin.Context) {
		c.String(http.StatusOK, "login successful")
	})
	v1.GET("api/logout", func(c *gin.Context) {
		c.String(http.StatusOK, "logout successful!")
	})
	v1.GET("api/register", func(c *gin.Context) {
		c.String(http.StatusOK, "register OK")
	})
	v1.POST("api/post", func(c *gin.Context) {
		acc := c.PostForm("acc")
		pwd := c.PostForm("pwd")
		msg := acc + pwd
		/*c.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},
			"message": msg,
		})*/
		c.String(http.StatusOK, msg)
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := DB[user]
		if ok {
			c.JSON(200, gin.H{"user": user, "value": value})
		} else {
			c.JSON(200, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			DB[user] = json.Value
			c.JSON(200, gin.H{"status": "ok"})
		}
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
