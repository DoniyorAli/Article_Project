package inmemory

import (
	"UacademyGo/Article/models"
	"errors"
	"strings"
	"time"
)

//*=========================================================================
func (inM InMemory) AddNewArticle(id string, box models.CreateModelArticle) error {
	var article models.Article
	article.ID = id
	article.Content = box.Content
	article.AuthorID = box.AuthorID
	article.CreateAt = time.Now()
	inM.DB.InMemoryArticleData = append(inM.DB.InMemoryArticleData, article)
	return nil
}
//*=========================================================================
func (inM InMemory) GetArticleById(id string) (models.GetByIDArticleModel, error) {
	var result models.GetByIDArticleModel
	for _, v := range inM.DB.InMemoryArticleData {
		if v.ID == id && v.DeletedAt == nil {
			author, err := inM.GetAuthorById(v.AuthorID)
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
func (inM InMemory) GetArticleList(offset, limit int, search string) (dataset []models.Article, err error) {
	throw := 0
	count := 0

	for _, v := range inM.DB.InMemoryArticleData {
		if v.DeletedAt == nil && (strings.Contains(v.Title, search) || strings.Contains(v.Body, search)) {
			if offset <= throw {
				count++
				dataset = append(dataset, v)	//! dataset ---> malumotlar to'plami
			}
			
			if count >= limit {
				break
			}
			throw++
		}
	}
	return dataset, err
}
//*=========================================================================
func (inM InMemory) UpdateArticle(box models.UpdateArticleResponse) error {
	var temp models.Article
	for i, v := range inM.DB.InMemoryArticleData {
		if v.ID == box.ID && v.DeletedAt == nil {
			temp = v
			temp.Content = box.Content
			t := time.Now()
			temp.UpdateAt = &t
			inM.DB.InMemoryArticleData[i] = temp
		
			return nil
		}
	}
	return errors.New("article not found")
}
//*=========================================================================
func (inM InMemory) DeleteArticle(id string) error {
	for i, v := range inM.DB.InMemoryArticleData {
		if v.ID == id {
			if v.DeletedAt != nil {
				return errors.New("article already deleted")
			}
			t := time.Now()
			v.DeletedAt = &t
			inM.DB.InMemoryArticleData[i] = v
			return nil
		}
	}
	return errors.New("article not found")
}







// // ! Delete Article with remove function
// func remove(slice []models.Article, s int) []models.Article {
// 	return append(slice[:s], slice[s+1:]...)
// }


