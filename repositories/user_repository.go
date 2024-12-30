package repositories

import (
	"context"
	"quote-generator-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
    Collection *mongo.Collection
}

func (r *UserRepository) AddFavorite(userID primitive.ObjectID, quoteID primitive.ObjectID) error {
    // Use context.TODO() instead of timeout context
    filter := bson.M{"_id": userID}
    update := bson.M{"$addToSet": bson.M{"favorites": quoteID}}
    _, err := r.Collection.UpdateOne(context.TODO(), filter, update)

    return err
}

func (r *UserRepository) GetFavorites(userID primitive.ObjectID) ([]models.Quote, error) {
    var user models.User
    var quotes []models.Quote

    if err := r.Collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user); err != nil {
        return nil, err
    }

    quotesCollection := r.Collection.Database().Collection("quotes")
    cursor, err := quotesCollection.Find(context.TODO(), bson.M{"_id": bson.M{"$in": user.Favorites}})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var quote models.Quote
        if err := cursor.Decode(&quote); err != nil {
            return nil, err
        }
        quotes = append(quotes, quote)
    }

    return quotes, nil
}
