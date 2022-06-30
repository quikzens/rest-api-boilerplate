package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Insert new document to a collection
func InsertOne(coll *mongo.Collection, newData any) error {
	_, err := coll.InsertOne(context.TODO(), newData)
	if err != nil {
		return err
	}

	return nil
}

// Update existing document in a collection
func UpdateOne(coll *mongo.Collection, id string, updatedData any) error {
	_, err := coll.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": updatedData},
	)
	if err != nil {
		return err
	}

	return nil
}

// Find all documents in a collection, sort by `created_at` field descending
func FindAll(coll *mongo.Collection, results any) error {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := coll.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return err
	}

	if err = cursor.All(context.TODO(), results); err != nil {
		return err
	}

	return nil
}

// Find several documents in a collection, sort by `created_at` field descending
func FindMany(coll *mongo.Collection, filter any, results any) error {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := coll.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return err
	}
	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(), results)
	if err != nil {
		return err
	}

	return nil
}

// Find one document in a collection
func FindOne(coll *mongo.Collection, filter any, result any) error {
	err := coll.FindOne(context.TODO(), filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

// Delete one document in a collection
func DeleteOne(coll *mongo.Collection, id string) error {
	_, err := coll.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

// Delete several documents in a collection
func DeleteMany(coll *mongo.Collection, filter any) error {
	_, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

// Delete all documents in a collection
func DeleteAll(coll *mongo.Collection) error {
	_, err := coll.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	return nil
}
