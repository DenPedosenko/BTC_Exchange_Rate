package main

import (
	"bufio"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type apiResponse struct {
	Time         string  `json:"time"`
	AssetIDBase  string  `json:"asset_id_base"`
	AssetIDQuote string  `json:"asset_id_quote"`
	Rate         float64 `json:"rate"`
}

func getRate(c *gin.Context) {
	var rate float64 = getCurrentBTCToUAHRate()
	c.IndentedJSON(http.StatusOK, rate)
}

func postEmail(c *gin.Context) {
	request := c.Request
	writter := c.Writer
	headerContentType := request.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		writter.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	request.ParseForm()

	newEmail := request.FormValue("email")
	if isEmailAlreadyExists(newEmail) {
		writter.WriteHeader(http.StatusConflict)
		return
	}

	file, err := openFile(os.O_WRONLY)
	if err != nil {
		writter.WriteHeader(http.StatusInternalServerError)
		return
	}
	file.WriteString(newEmail + "\n")
	defer file.Close()
	writter.WriteHeader(http.StatusOK)
}

func getEmails(c *gin.Context) {
	file, err := openFile(os.O_RDONLY)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	var emails []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		emails = append(emails, scanner.Text())
	}
	c.IndentedJSON(http.StatusOK, emails)
}

func sendEmails(c *gin.Context) {
	file, err := openFile(os.O_RDONLY)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	var emails []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		emails = append(emails, scanner.Text())
	}

	err = send(emails)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
