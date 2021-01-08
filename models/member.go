package models

import (
	con "Renting/models/db"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

// Member -
type Member struct {
	ID       primitive.ObjectID `bson:"_id"`
	Account  string             `json:"account"`
	Password string             `json:"password"`
	Name     string             `json:"name"`
	Sex      string             `json:"sex"`
	Phone    string             `json:"phone"`
}

/*type ID struct {
	ID primitive.ObjectID `bson:"_id"`
}*/

// AddMember _新增使用者
func (member *Member) AddMember() error {

	member.ID = primitive.NewObjectID()

	// DB 連線
	db := con.ConnectDB()
	// 連接到user collection(mongoDB)
	collection := db.Database("Reservation").Collection("Member")

	_, err := collection.InsertOne(context.TODO(), member)

	// 關閉 DB 連線
	db.Disconnect(context.TODO())

	if err != nil {
		return err
	}

	return nil
}

// Login _使用者登入
func (member *Member) Login() (*Member, error) {
	// DB 連線
	db := con.ConnectDB()
	// 連接到user collection(mongoDB)
	collection := db.Database("Reservation").Collection("Member")

	//fmt.Println(member.Account)

	result := new(Member)

	err := collection.FindOne(context.TODO(), bson.M{"account": member.Account}).Decode(&result)

	// 關閉 DB 連線
	db.Disconnect(context.TODO())

	if err != nil {
		//fmt.Println(result)
		fmt.Println(err)
		return result, err
	}
	//fmt.Println(result)
	return result, nil
}

/*func Getuser(memeber_id primitive.ObjectID) (data []*Member, err error) {
	// DB 連線
	db := con.ConnectDB()
	// 連接到user collection(mongoDB)
	collection := db.Database("Reservation").Collection("Member")

	cursor := collection.Find(context.TODO(), bson.M{"_id": memeber_id})
	// 關閉 DB 連線
	db.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}*/
