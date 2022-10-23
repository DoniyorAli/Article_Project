package inmemory_test

import (
	"UacademyGo/Article/models"
	"UacademyGo/Article/storage/inmemory"
	"errors"
	"testing"
)

func TestAddNewArticle(t *testing.T) {
	var err error
	IM := inmemory.InMemory{
		DB: &inmemory.DataB{},
	}

	errorAuthorNotFound := errors.New("author not found")

	authorID := "099cd64c-e2f9-4a94-8901-a49beb07357b"
	authorData := models.CreateModelAuthor{
		Firstname: "Antoniyo",
		Lastname:  "Banderos",
	}

	notFoundAuthorID := "761372d3-7a34-4ad4-8ab2-7a97e41fa40b"

	content := models.Content{
		Title: "Sherlock Holmes",
		Body:  "Sherlock Holmes is a fictional detective genre",
	}

	err = IM.AddAuthor(authorID, authorData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

//* ===========================================================================================

	var tests = []struct {
		name       string
		id         string
		data       models.CreateModelArticle
		wantError  error
		wantResult models.GetByIDArticleModel
	}{
		{
			name: "Successfully",
			id:   "eec91f7b-8a6e-4f3c-8df8-48f313803fc6",
			data: models.CreateModelArticle{
				Content:  content,
				AuthorID: authorID,
			},
			wantError: nil,
			wantResult: models.GetByIDArticleModel{
				Content: content,
			},
		},
		{
			name: "fail",
			id:   "c98f19ff-2ee0-4af4-9758-c794661759f8",
			data: models.CreateModelArticle{
				Content:  content,
				AuthorID: notFoundAuthorID,
			},
			wantError:  errorAuthorNotFound,
			wantResult: models.GetByIDArticleModel{},
		},
	}

	for _, box := range tests {
		t.Run(box.name, func(t *testing.T) {
			err = IM.AddNewArticle(box.id, box.data)

			if box.wantError == nil {
				if err != nil {
					t.Errorf("IM.AddNewArrticle() got error: %v", err)
				}
				article, err := IM.GetArticleById(box.id)
				if err != nil {
					t.Errorf("IM.AddNewArrticle() got error: %v", err)
				}

				if box.wantResult.Content != article.Content {
					t.Errorf("IM.AddNewArrticle() expected: %v, but got: %v", box.wantResult.Content, article.Content)
				}
			} else {
				if box.wantError.Error() != err.Error() {
					t.Errorf("IM.AddNewArrticle() expected error: %v, but got error: %v", box.wantError, err)
				}
			}
		})
	}
	t.Log("Test has been finished")
}

//* ===========================================================================================

func TestGetArticleById(t *testing.T) {
	var err error
	IM := inmemory.InMemory{
		DB: &inmemory.DataB{},
	}

	errorArticleNotFound := errors.New("article not found")

	authorID := "099cd64c-e2f9-4a94-8901-a49beb07357b"
	authorData := models.CreateModelAuthor{
		Firstname: "Antoniyo",
		Lastname:  "Banderos",
	}

	err = IM.AddAuthor(authorID, authorData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	articleID := "a8ad6df8-9455-4840-aace-83b44ccf47bc"
	articleData := models.CreateModelArticle{
		Content:   models.Content{
			Title: "Sherlock Holmes",
			Body:  "Sherlock Holmes is a fictional detective genre",
		},
		AuthorID: authorID,
	}

	err = IM.AddNewArticle(articleID, articleData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
//* ===========================================================================================

	var tests = []struct {
		name       string
		id         string
		mockFunc 	func()
		wantError  error
		wantResult models.GetByIDArticleModel
	}{
		{
			name: "Successfully",
			id:   articleID,
			mockFunc: func ()  {
				
			},
			wantError: nil,
			wantResult:  models.GetByIDArticleModel{
				Content: articleData.Content,
			},
		},
		{
			name: "fail: article not found",
			id:   "de89270e-4a1e-4142-a70f-843ebb35fefc",
			mockFunc: func ()  {
				
			},
			wantError: errorArticleNotFound,
			wantResult:  models.GetByIDArticleModel{},
		},
	}

	for _, box := range tests {
		t.Run(box.name, func(t *testing.T) {
			article, err := IM.GetArticleById(box.id)

			if box.wantError == nil {
				if err != nil {
					t.Errorf("IM.GetArticleById() got error: %v", err)
				}

				if box.wantResult.Content != article.Content {
					t.Errorf("IM.GetArticleById() expected: %v, but got: %v", box.wantResult.Content, article.Content)
				}
			} else {
				if box.wantError.Error() != err.Error() {
					t.Errorf("IM.GetArticleById() expected error: %v, but got error: %v", box.wantError, err)
				}
			}
		})
	}
	t.Log("Test has been finished")
}
