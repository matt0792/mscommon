package util

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogRequestBody(c *gin.Context) {
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Failed to read body: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	log.Printf("Request Body: %s", string(bodyBytes))

	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	c.Next()
}
