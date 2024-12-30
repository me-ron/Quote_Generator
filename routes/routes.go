package routes

import (
    "github.com/gin-gonic/gin"
    "quote-generator-backend/controllers"
)

func SetupRoutes(r *gin.Engine, qc *controllers.QuoteController, uc *controllers.UserController) {
    r.GET("/quotes/:category", qc.GetQuotesByCategory)
    r.GET("/quotes/random", qc.GetRandomQuotes)
    r.POST("/users/:user_id/favorites", uc.AddFavorite)
    r.GET("/users/:user_id/favorites", uc.GetFavorites)
}