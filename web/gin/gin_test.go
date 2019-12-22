package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Example() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, world!"})
	})
	err := r.Run(":8000")
	if err != nil {
		panic(err)
	}
}
