package types

import (
	"go.mongodb.org/mongo-driver/bson"
)

type MongoExpression struct {
	Condition bson.M
}
