package noti

import (
	"context"
	"fmt"
	_ "fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson"
	"mindfit.noti/graph/model"
	"mindfit.noti/mongo/maindb"
)

type Idonly struct {
	Userid string `json:"userid"`
}

type ReturnNotiFromDB struct {
	Userid   string `json:"userid"`
	Notilist []*model.NotiSchema
}

var ctx = context.TODO()

func Getnoti(idonly Idonly) []*model.NotiSchema {

	res, err := maindb.Getcol().Find(ctx, idonly)
	if err != nil {
		log.Fatal(err)
	}

	var msg1 []*ReturnNotiFromDB
	if err = res.All(ctx, &msg1); err != nil {
		panic(err)
	}

	var msg []*model.NotiSchema = msg1[0].Notilist

	return msg

}

func Updatenoti(notiin ReturnNotiFromDB) bool {

	fmt.Println(notiin)

	res, err := maindb.Getcol().UpdateOne(ctx, bson.M{"userid": notiin.Userid}, bson.M{"$set": bson.M{"notilist": notiin.Notilist}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	msg := true

	return msg

}
