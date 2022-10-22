package inmemory

import "UacademyGo/Article/models"

type InMemory struct {
	DB *DataB
}

type DataB struct {
	InMemoryArticleData []models.Article
	InMemoryAuthorData []models.Author
}