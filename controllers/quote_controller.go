package controllers

import (
	"net/http"
	"quote-generator-backend/models"
	"quote-generator-backend/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuoteController struct {
    Service *services.QuoteService
}

// AddQuote adds a new quote to the database.
func (h *QuoteController) AddQuote(c *gin.Context) {
	var quote models.Quote
	if err := c.ShouldBindJSON(&quote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Set a new ObjectID for the quote if it's not set
	if quote.ID.IsZero() {
		quote.ID = primitive.NewObjectID()
	}

	err := h.Service.AddQuote(quote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add quote"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quote added successfully", "quote_id": quote.ID.Hex()})
}

func (qc *QuoteController) GetQuotesByCategory(c *gin.Context) {
    category := c.Param("category")
    quotes, err := qc.Service.GetQuotesByCategory(category)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, quotes)
}

func (qc *QuoteController) GetRandomQuotes(c *gin.Context) {
    quotes, err := qc.Service.GetRandomQuotes(5) // Default limit to 5
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, quotes)
}
