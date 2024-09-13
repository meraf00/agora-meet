package infrastructure

import (
	"context"
	"log"

	"github.com/meraf00/agora-meet/agora-be/internal/users/domain/entities"
	"github.com/meraf00/agora-meet/agora-be/internal/users/infrastructure/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserMongoRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewUserMongoRepository(mongoUri string, dbName string, collectionName string) (*UserMongoRepository, error) {
	clientOptions := options.Client().ApplyURI(mongoUri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Unable to connext to database: %s", err)
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Database ping failed")
		return nil, err
	}

	db := client.Database(dbName)

	return &UserMongoRepository{
		db:         db,
		collection: db.Collection(collectionName),
	}, nil

}

func (r *UserMongoRepository) FindUser(ctx context.Context, id string) (*entities.User, error) {
	var user models.UserModel
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	err = r.collection.FindOne(ctx, bson.M{
		"_id": objectId,
	}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &entities.User{
		Email: user.Email,
		Id:    user.Id,
		Name:  user.Name,
	}, nil

}

func (r *UserMongoRepository) FindUsers(ctx context.Context) ([]*entities.User, error) {
	var users []*entities.User

	cur, err := r.collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var user models.UserModel
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, &entities.User{
			Email: user.Email,
			Id:    user.Id,
			Name:  user.Name,
		})
	}

	return users, nil
}

func (r *UserMongoRepository) SaveUser(ctx context.Context, user entities.User) error {
	return nil
}
