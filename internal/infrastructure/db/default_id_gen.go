package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type DefaultIDGenerator struct {
}

func (gen DefaultIDGenerator) Generate() string {
	return primitive.NewObjectID().Hex()
}

func NewIdGenerator() DefaultIDGenerator {
	return DefaultIDGenerator{}
}
