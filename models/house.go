package models

import (
	con "Renting/models/db"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
)

type House struct {
	ID      primitive.ObjectID `bson:"_id"`
	Owner   string             `json:"owner"`
	Address string             `json:"address"`
	Price   int64              `json:"price"`
}

func GetAllHouses() (data []*House, err error) {
	db := con.ConnectDB()

	collection := db.Database("Reservation").Collection("House")

	cursor, err := collection.Find(context.TODO(), bson.M{})

	// 關閉 DB 連線
	db.Disconnect(context.TODO())
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func GetHouse(houseID primitive.ObjectID) (data []*House, err error) {
	db := con.ConnectDB()

	collection := db.Database("Reservation").Collection("House")
	fmt.Println(houseID)
	result, err := collection.Find(context.TODO(), bson.M{"_id": houseID})

	db.Disconnect(context.TODO())
	err = result.All(context.TODO(), &data)
	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (house *House) UpdateHouse(houseID primitive.ObjectID) error {

	db := con.ConnectDB()

	collection := db.Database("Reservation").Collection("House")

	_, err := collection.UpdateOne(context.TODO(),
		bson.M{"_id": houseID},
		bson.D{
			{"$set", bson.M{"price": house.Price}},
		})
	db.Disconnect(context.TODO())

	if err != nil {
		return err
	}
	return nil
}
func (house *House) DeleteHouse(houseID primitive.ObjectID) error {
	db := con.ConnectDB()

	collection := db.Database("Reservation").Collection("House")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": houseID})

	db.Disconnect(context.TODO())

	if err != nil {
		return err
	}
	return nil
}
