package storage

import (
	"UacademyGo/Article/models"
	"errors"
	"time"
)

var InMemoryArticleData []models.Article
//*=========================================================================
func AddNewArticle(id string, box models.CreateModelArticle) error {
	var article models.Article
	article.ID = id
	article.Content = box.Content
	article.AuthorID = box.AuthorID
	article.CreateAt = time.Now()
	InMemoryArticleData = append(InMemoryArticleData, article)
	return nil
}
//*=========================================================================
func GetArticleById(id string) (models.GetByIDArticleModel, error) {
	var result models.GetByIDArticleModel
	for _, v := range InMemoryArticleData {
		if v.ID == id {
			author, err := GetAuthorById(v.AuthorID)
			if err != nil {
				return result, err
			}

			result.ID = v.ID
			result.Author = author
			result.Content = v.Content
			result.CreateAt = v.CreateAt
			result.UpdateAt = v.UpdateAt
			result.DeletedAt = v.DeletedAt
			return result, nil
		}
	}
	return result, errors.New("article not found")
}
//*=========================================================================
func GetArticleList() (dataset []models.Article, err error) {
	dataset = InMemoryArticleData	//! dataset ---> malumotlar to'plami
	return dataset, err
}
//*=========================================================================
func UpdateArticle(box models.UpdateArticleResponse) error {
	var temp models.Article
	for i, v := range InMemoryArticleData {
		if v.ID == box.ID {
			temp = v
			temp.Content = box.Content
			t := time.Now()
			temp.UpdateAt = &t
			InMemoryArticleData[i] = temp
		
			return nil
		}
	}
	return errors.New("article not found")
}
//*=========================================================================
func DeleteArticle(id string) error {
	for i, v := range InMemoryArticleData {
		if v.ID == id {
			// InMemoryArticleData = remove(InMemoryArticleData, i)
			InMemoryArticleData = append(InMemoryArticleData[:i], InMemoryArticleData[i+1:]...)
			return nil
		}
	}
	return errors.New("deletion failed")
}







// // ! Delete Article with remove function
// func remove(slice []models.Article, s int) []models.Article {
// 	return append(slice[:s], slice[s+1:]...)
// }


