package utils

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB
var MongoDB *mongo.Client
var EmailCode = make(map[string]time.Time)

const YesCode = 0
const NoCode = 1
