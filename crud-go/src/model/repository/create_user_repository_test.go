package repository

import (
	"github.com/stretchr/testify/assert"
	"github.com/valdenidelgado/go-projects/crud-go/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"os"
	"testing"
)

func TestUserRepository_CreateUser(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("MONGODB_USER_COLLECTION", collection_name)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.Close()

	mtestDb.Run("TestUserRepository_CreateUser", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(database_name)

		repository := NewUserRepository(databaseMock)

		userDomain, err := repository.CreateUser(model.NewUserDomain(
			"John@doe.com", "test", "John", 80))

		_, errId := primitive.ObjectIDFromHex(userDomain.GetID())

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.EqualValues(t, userDomain.GetEmail(), "John@doe.com")
	})

	mtestDb.Run("TestUserRepository_CreateUser - Error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)

		repository := NewUserRepository(databaseMock)

		userDomain, err := repository.CreateUser(model.NewUserDomain(
			"John@doe.com", "test", "John", 80))

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}
