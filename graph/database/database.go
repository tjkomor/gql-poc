package database

import (
	"context"
	"log"
	"time"

	"gql-poc/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	return &DB {
		client: client,
	}
}

func (db *DB) Save(input *model.NewPizza) *model.Pizza {
	collection := db.client.Database("food").Collection("pizzas")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	 res, err := collection.InsertOne(ctx, input)
	 if err != nil {
		 log.Fatal(err)
	 }

	 return &model.Pizza {
		 ID: res.InsertedID.(primitive.ObjectID).Hex(),
		 Toppings: input.Toppings,
		 Price: input.Price,
	 }
}

func (db *DB) FindByID(ID string) *model.Pizza {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("food").Collection("pizzas")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	pizza := model.Pizza{}
	res.Decode(&pizza)
	return &pizza
}

func (db *DB) All() []*model.Pizza {
	collection := db.client.Database("food").Collection("pizzas")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var pizzas []*model.Pizza
	for cur.Next(ctx) {
		var pizza *model.Pizza
		err := cur.Decode(&pizza)
		if err != nil {
			log.Fatal(err)
		}
		pizzas = append(pizzas, pizza)
	}
	return pizzas
}