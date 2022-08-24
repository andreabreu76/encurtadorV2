package entitie

import (
	"context"
	"github.com/andreabreu76/encurtadorV2/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type Link struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	Type          string             `bson:"type" json:"type"`
	Target        string             `bson:"url" json:"url"`
	TempTarget    string             `bson:"tempurl" json:"tempurl"`
	IDCampaign    int64              `bson:"idCampaign" json:"idCampaign"`
	IDCampaignV2  int64              `bson:"idCampaignV2" json:"idCampaignV2"`
	IDCreative    string             `bson:"idCreative" json:"idCreative"`
	CreativeType  string             `bson:"creativeType" json:"creativeType"`
	Source        string             `bson:"source" json:"source"`
	IDSite        int64              `bson:"idSite" json:"idSite"`
	TempRedirect  bool               `bson:"tempredirect" json:"tempredirect"`
	HitCount      int64              `bson:"hitCount" json:"hitCount"`
	DateCreated   time.Time          `bson:"dateCreated" json:"dateCreated"`
	DateModified  time.Time          `bson:"dateModified" json:"dateModified"`
	OriginEnv     string             `bson:"originEnv" json:"originEnv"`
	IDClient      int64              `bson:"idClient" json:"idClient"`
	IDAffiliate   int64              `bson:"idAffiliate" json:"idAffiliate"`
	ComissionType int64              `bson:"comissionType" json:"comissionType"`
	Status        int64              `bson:"status" json:"status"`
	BlockParams   int64              `bson:"blockParams" json:"blockParams"`
}

func CreateNewLink(link Link) (primitive.ObjectID, error) {
	session, err := config.ConnectMongo()
	if err != nil {
		log.Fatalf("CreateNewLink :: Error connecting to mongo: %v", err)
		return primitive.NilObjectID, err
	}

	collection := session.Database("encurtador").Collection("links")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, link)
	if err != nil {
		log.Printf("CreateNewLink :: Error inserting link: %v", err)
		return primitive.NilObjectID, err
	}

	link.ID = result.InsertedID.(primitive.ObjectID)
	return link.ID, nil
}
