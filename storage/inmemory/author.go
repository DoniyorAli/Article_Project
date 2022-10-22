package inmemory

import (
	"UacademyGo/Article/models"
	"errors"
	"time"
)

func (inM InMemory) AddAuthor(id string, box models.CreateModelAuthor) error {
	var author models.Author
	author.ID = id
	author.Firstname = box.Firstname
	author.Lastname = box.Lastname
	author.CreateAt = time.Now()
	inM.DB.InMemoryAuthorData = append(inM.DB.InMemoryAuthorData, author)
	return nil
}

//*=========================================================================
func (inM InMemory) GetAuthorById(id string) (models.Author, error) {
	var result models.Author
	for _, v := range inM.DB.InMemoryAuthorData {
		if v.ID == id {
			result = v
			return result, nil
		}
	}
	return result, errors.New("author not found")
}

//*=========================================================================
func (inM InMemory) GetAuthorList() (dataset []models.Author, err error) {
	dataset = inM.DB.InMemoryAuthorData
	return dataset, err
}

//*=========================================================================
func (inM InMemory) UpdateAuthor(box models.UpdateAuthorResponse) error {
	var temp models.Author
	for i, v := range inM.DB.InMemoryAuthorData {
		if v.ID == box.ID {
			temp = v
			temp.Firstname = box.Firstname
			temp.Lastname = box.Lastname
			t := time.Now()
			temp.UpdateAt = &t
			inM.DB.InMemoryAuthorData[i] = temp
		
			return nil
		}
	}
	return errors.New("author not found")
}

//*=========================================================================
func (inM InMemory) DeleteAuthor(id string) error {
	for i, v := range inM.DB.InMemoryAuthorData {
		if v.ID == id {
			// InMemoryArticleData = remove(InMemoryArticleData, i)
			inM.DB.InMemoryAuthorData = append(inM.DB.InMemoryAuthorData[:i], inM.DB.InMemoryAuthorData[i+1:]...)
			return nil
		}
	}
	return errors.New("deletion failed")
}