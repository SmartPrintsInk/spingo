package spingo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connection *mongo.Client

// Access default env credentials
func Access() (*mongo.Client, error) {
	token := tokens{}
	token.init()
	return connect(token)
}

// AccessFor any host
func AccessFor(host string) (*mongo.Client, error) {
	token := tokens{}
	token.init()
	token.uri = fmt.Sprintf("mongodb://%s:%s", host, token.port)
	return connect(token)
}

func connect(token tokens) (*mongo.Client, error) {
	credential := options.Credential{
		AuthMechanism: token.authMechanism,
		AuthSource:    token.authDatabase,
		Username:      token.user,
		Password:      token.password,
	}
	clientOptions := options.Client().ApplyURI(token.uri).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	connection = client
	return client, err
}

func Close() {
	if err := connection.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func Ping() {
	if err := connection.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
}
