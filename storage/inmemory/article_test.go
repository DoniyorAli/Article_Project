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

	err = IM.AddAuthor("099cd64c-e2f9-4a94-8901-a49beb07357b", models.CreateModelAuthor{
		Firstname: "Antoniyo",
		Lastname: "Banderos",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
//* ===========================================================================================

	err = IM.AddNewArticle("6d45c3fa-960f-45af-a8c1-18a2abba14ca", models.CreateModelArticle{
		Content: models.Content{
			Title: "Sherlock Holmes",
			Body: "Sherlock Holmes is a fictional detective genre",
		},
		AuthorID: "099cd64c-e2f9-4a94-8901-a49beb07357b",
	})

	if err != nil {
		t.Errorf("IM.AddNewArticle() got error: %v", err)
	}

	article, err := IM.GetArticleById("6d45c3fa-960f-45af-a8c1-18a2abba14ca")
	if err != nil {
		t.Errorf("IM.AddNewArticle() got error: %v", err)
	}

	if article.Title != "Sherlock Holmes" || article.Body != "Sherlock Holmes is a fictional detective genre" {
		t.Errorf("mitmatch between data is wrong")
	}

//* ===========================================================================================

	err = IM.AddNewArticle("6d45c3fa-960f-45af-a8c1-18a2abba14ca", models.CreateModelArticle{
		Content: models.Content{
			Title: "Sherlock Holmes",
			Body: "Sherlock Holmes is a fictional detective genre",
		},
		AuthorID: "761372d3-7a34-4ad4-8ab2-7a97e41fa40b",
	})

	expectedError := errors.New("author not found")
	if err == nil {
		t.Errorf("IM.AddNewArticle() expected error but got nil")
	}else {
		if err.Error() != expectedError.Error() {
			t.Errorf("IM.AddNewArticle() expected: %v, but got error: %v", expectedError, err)
		}
	}
	t.Log("Test has been finished")
}

