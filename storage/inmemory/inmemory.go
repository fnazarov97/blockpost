package inmemory

import "article/models"

type Inmemory struct {
	Db *DB
}

type DB struct {
	InMemoryAuthorData  []models.Author
	InMemoryArticleData []models.Article
}
