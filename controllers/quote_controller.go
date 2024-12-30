package controllers

import (
    "net/http"
    "quote-generator-backend/services"

    "github.com/gin-gonic/gin"
)

type QuoteController struct {
    Service *services.QuoteService
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
