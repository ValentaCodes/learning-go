package main

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

// Create/Run engine
func main() {
	engine := gin.Default()
	engine.GET("analysis", getAnalysis)
	engine.Run("localhost:3001")

}

// declare struct data
type analysis struct {
	Mood           string `json:"mood"`
	SentimentScore int    `json:"sentimentScore"`
	Summary        string `json:"summary"`
	Subject        string `json:"subject"`
	Color          string `json:"color"`
}

// create slice of analysis
var aiAnalysis = []analysis{
	{Mood: "Happy", SentimentScore: 5, Summary: "I was happy today", Subject: "Personal", Color: "Blue"},
	{Mood: "Sad", SentimentScore: -10, Summary: "I was sad today", Subject: "Personal", Color: "Black"},
}

// function getting data
func getAnalysis(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, aiAnalysis)
	// fmt.Println(results)
}
