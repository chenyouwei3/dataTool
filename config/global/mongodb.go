package global

import "go.mongodb.org/mongo-driver/mongo"

var (
	MongodbClient403 *mongo.Client
	Device           *mongo.Collection
)
