package spingo

import (
	"os"
)

type tokens struct {
	uri          string
	user         string
	port         string
	host         string
	password     string
	authDatabase string
	database     string
	collection   string
}

func (t *tokens) init() {
	t.uri = os.Getenv("MongoURI")
	t.port = os.Getenv("MongoPort")
	t.user = os.Getenv("MongoUser")
	t.password = os.Getenv("MongoPassword")
	t.authDatabase = os.Getenv("MongoAuthDatabase")
}
