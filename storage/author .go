package storage

import (
	"UacademyGo/Article/models"
	"errors"
	"time"
)

var InMemoryAuthorData []models.Author

func AddAuthor(id string, box models.CreateModelAuthor) error {
	var author models.Author
	author.ID = id
	author.Firstname = box.Firstname
	author.Lastname = box.Lastname
	author.CreateAt = time.Now()
	InMemoryAuthorData = append(InMemoryAuthorData, author)
	return nil
}

//*=========================================================================
func GetAuthorById(id string) (models.Author, error) {
	var result models.Author
	for _, v := range InMemoryAuthorData {
		if v.ID == id {
			result = v
			return result, nil
		}
	}
	return result, errors.New("author not found")
}

//*=========================================================================
func GetAuthorList() (dataset []models.Author, err error) {
	dataset = InMemoryAuthorData
	return dataset, err
}

//*=========================================================================
func UpdateAuthor(box models.UpdateAuthorResponse) error {
	var temp models.Author
	for i, v := range InMemoryAuthorData {
		if v.ID == box.ID {
			temp = v
			temp.Firstname = box.Firstname
			temp.Lastname = box.Lastname
			t := time.Now()
			temp.UpdateAt = &t
			InMemoryAuthorData[i] = temp
		
			return nil
		}
	}
	return errors.New("author not found")
}

//*=========================================================================
func DeleteAuthor(id string) error {
	for i, v := range InMemoryAuthorData {
		if v.ID == id {
			// InMemoryArticleData = remove(InMemoryArticleData, i)
			InMemoryAuthorData = append(InMemoryAuthorData[:i], InMemoryAuthorData[i+1:]...)
			return nil
		}
	}
	return errors.New("deletion failed")
}