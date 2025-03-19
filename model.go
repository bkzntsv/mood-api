package main

// SentimentRequest represents the input for sentiment analysis
type SentimentRequest struct {
	Text string `json:"text" binding:"required,min=1"` // Text cannot be empty
}

// SentimentResponse contains the sentiment analysis result
type SentimentResponse struct {
	Sentiment string    `json:"sentiment"` // One of: positive, negative, neutral
	Score     float64  `json:"score"`      // Confidence score between 0 and 1
}