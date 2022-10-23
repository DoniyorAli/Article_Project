package inmemory_test

import (
	"UacademyGo/Article/models"
	"UacademyGo/Article/storage/inmemory"
	"errors"
	"testing"
)

func DataSet(stg *inmemory.InMemory) error {
	err := stg.AddAuthor("52c9c9bd-2e68-401f-8ade-31dc7e7d6eea", models.CreateModelAuthor{
		Firstname: "John",
		Lastname: "Doe",
	})
	if err != nil {
		return err
	}

	err = stg.AddNewArticle("8048264e-21a4-4ad5-b1d4-baf809210db6", models.CreateModelArticle{
		Content: models.Content{
			Title: "Sherlock Holmes",
			Body: "Sherlock Holmes is a fictional detective genre",
		},
		AuthorID: "52c9c9bd-2e68-401f-8ade-31dc7e7d6eea",
	})
	if err != nil {
		return err
	}

	err = stg.AddNewArticle("6c5dc295-e3fc-4970-8841-7505630d796d", models.CreateModelArticle{
		Content: models.Content{
			Title: "1",
			Body: "This is my Lorem pages",
		},
		AuthorID: "52c9c9bd-2e68-401f-8ade-31dc7e7d6eea",
	})
	if err != nil {
		return err
	}

	err = stg.AddNewArticle("d0c6557c-eaea-4e63-8c80-71caf73c7d89", models.CreateModelArticle{
		Content: models.Content{
			Title: "2",
			Body: "This is my Lorem pages",
		},
		AuthorID: "52c9c9bd-2e68-401f-8ade-31dc7e7d6eea",
	})
	if err != nil {
		return err
	}

	err = stg.AddNewArticle("23ae0860-84a2-4b9e-b119-918ad85b9f03", models.CreateModelArticle{
		Content: models.Content{
			Title: "3",
			Body: "This is my Lorem pages",
		},
		AuthorID: "52c9c9bd-2e68-401f-8ade-31dc7e7d6eea",
	})
	if err != nil {
		return err
	}

	err = stg.AddNewArticle("d8aed20c-aca3-4c1f-875f-ce823cd0495a", models.CreateModelArticle{
		Content: models.Content{
			Title: "4",
			Body: "This is my Lorem pages",
		},
		AuthorID: "52c9c9bd-2e68-401f-8ade-31dc7e7d6eea",
	})
	if err != nil {
		return err
	}

	err = stg.AddNewArticle("3c4f62cf-b6ac-49f7-95bd-ce776ea58b34", models.CreateModelArticle{
		Content: models.Content{
			Title: "5",
			Body: "This is my Lorem pages",
		},
		AuthorID: "52c9c9bd-2e68-401f-8ade-31dc7e7d6eea",
	})
	if err != nil {
		return err
	}

	err = stg.AddNewArticle("0828afee-86bf-4a58-8c85-b8f1a7fe304b", models.CreateModelArticle{
		Content: models.Content{
			Title: "6",
			Body: "This is my Lorem pages",
		},
		AuthorID: "52c9c9bd-2e68-401f-8ade-31dc7e7d6eea",
	})
	if err != nil {
		return err
	}

	err = stg.AddNewArticle("44d50ef7-349c-49a0-a715-bd624f1fb397", models.CreateModelArticle{
		Content: models.Content{
			Title: "7",
			Body: "This is my Lorem pages",
		},
		AuthorID: "52c9c9bd-2e68-401f-8ade-31dc7e7d6eea",
	})
	if err != nil {
		return err
	}

	err = stg.AddNewArticle("e914f3a6-8316-4e68-9241-8218e673bc14", models.CreateModelArticle{
		Content: models.Content{
			Title: "8",
			Body: "This is my Lorem pages",
		},
		AuthorID: "52c9c9bd-2e68-401f-8ade-31dc7e7d6eea",
	})
	if err != nil {
		return err
	}

	err = stg.AddNewArticle("595a44d4-a382-489c-9bd8-c638dce27c45", models.CreateModelArticle{
		Content: models.Content{
			Title: "9",
			Body: "This is my Lorem pages",
		},
		AuthorID: "52c9c9bd-2e68-401f-8ade-31dc7e7d6eea",
	})
	if err != nil {
		return err
	}

	err = stg.AddNewArticle("a525fcd5-7ab6-417d-89dd-90799b46d32e", models.CreateModelArticle{
		Content: models.Content{
			Title: "10",
			Body: "This is my Lorem pages",
		},
		AuthorID: "52c9c9bd-2e68-401f-8ade-31dc7e7d6eea",
	})
	if err != nil {
		return err
	}

	err = stg.AddNewArticle("653a1fc3-65f6-42e0-8b18-b044fba31c7f", models.CreateModelArticle{
		Content: models.Content{
			Title: "11",
			Body: "This is my Lorem pages",
		},
		AuthorID: "52c9c9bd-2e68-401f-8ade-31dc7e7d6eea",
	})

	if err != nil {
		return err
	}
	return nil
}

func TestAddNewArticle(t *testing.T) {
	var err error
	IM := inmemory.InMemory{
		DB: &inmemory.DataB{},
	}

	DataSet(&IM)

	errorAuthorNotFound := errors.New("author not found")

	authorID := "52c9c9bd-2e68-401f-8ade-31dc7e7d6eea"

	notFoundAuthorID := "761372d3-7a34-4ad4-8ab2-7a97e41fa40b"

	content := models.Content{
		Title: "Sherlock Holmes",
		Body:  "Sherlock Holmes is a fictional detective genre",
	}

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
					t.Errorf("IM.AddNewArrticle() got error: %+v", err)
				}
				article, err := IM.GetArticleById(box.id)
				if err != nil {
					t.Errorf("IM.AddNewArrticle() got error: %+v", err)
				}

				if box.wantResult.Content != article.Content {
					t.Errorf("IM.AddNewArrticle() expected: %+v, but got: %+v", box.wantResult.Content, article.Content)
				}
			} else {
				if box.wantError.Error() != err.Error() {
					t.Errorf("IM.AddNewArrticle() expected error: %+v, but got error: %+v", box.wantError, err)
				}
			}
		})
	}
	t.Log("Test has been finished")
}

//* =========================================================================

func TestGetArticleById(t *testing.T) {
	IM := inmemory.InMemory{
		DB: &inmemory.DataB{},
	}

	DataSet(&IM)

	errorArticleNotFound := errors.New("article not found")

	articleID := "8048264e-21a4-4ad5-b1d4-baf809210db6"
	
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
				Content: models.Content{
					Title: "Sherlock Holmes",
					Body: "Sherlock Holmes is a fictional detective genre",
				},
			},
		},
		{
			name: "fail: article not found",
			id:   "ee06e62c-d353-404a-87f6-a279ceedf89f",
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
					t.Errorf("IM.GetArticleById() got error: %+v", err)
				}

				if box.wantResult.Content != article.Content {
					t.Errorf("IM.GetArticleById() expected: %+v, but got: %+v", box.wantResult.Content, article.Content)
				}
			} else {
				if box.wantError.Error() != err.Error() {
					t.Errorf("IM.GetArticleById() expected error: %+v, but got error: %+v", box.wantError, err)
				}
			}
		})
	}
	t.Log("Test has been finished")
}

//* =========================================================================

func TestGetArticleList(t *testing.T) {
	IM := inmemory.InMemory{
		DB: &inmemory.DataB{},
	}

	DataSet(&IM)

	var tests = []struct {
		name       string
		offset int
		limit int
		search string
		wantError  error
		numberOfResult int
	}{
		{
			name: "success default",
			offset: 0,
			limit: 10,
			search: "",
			wantError: nil,
			numberOfResult: 10,
		},
		{
			name: "success limit",
			offset: 0,
			limit: 5,
			search: "",
			wantError: nil,
			numberOfResult: 5,
		},
		{
			name: "success offset",
			offset: 2,
			limit: 10,
			search: "",
			wantError: nil,
			numberOfResult: 10,
		},
		{
			name: "success mix offset limit ",
			offset: 2,
			limit: 5,
			search: "",
			wantError: nil,
			numberOfResult: 5,
		},
//*Searching============================
		{
			name: "search Lorem",
			offset: 0,
			limit: 10,
			search: "Lorem",
			wantError: nil,
			numberOfResult: 10,
		},
		{
			name: "succes",
			offset: 0,
			limit: 10,
			search: "5",
			wantError: nil,
			numberOfResult: 1,
		},
	}

	for _, box := range tests {
		t.Run(box.name, func(t *testing.T) {
			articleList, err := IM.GetArticleList(box.offset, box.limit, box.search)

			if box.wantError == nil {
				if err != nil {
					t.Errorf("IM.GetArticleList() got error: %+v", err)
				}

				if box.numberOfResult != len(articleList) {
					t.Errorf("IM.GetArticleList() expected: %d objects, but got objects: %+v", box.numberOfResult, len(articleList))
				}
			} else {
				if box.wantError.Error() != err.Error() {
					t.Errorf("IM.GetArticleList() expected error: %+v, but got error: %+v", box.wantError, err)
				}
			}
		})
	}
	t.Log("Test has been finished")
}

//* =========================================================================

func TestUpdateArticle(t *testing.T) {
	IM := inmemory.InMemory{
		DB: &inmemory.DataB{},
	}

	DataSet(&IM)


	var tests = []struct {
		name       string
		id         string
		data       models.UpdateArticleResponse
		wantError  error
		wantResult models.GetByIDArticleModel
	}{
		{
			name: "Successfully",
			data: models.UpdateArticleResponse{
				ID:   "8048264e-21a4-4ad5-b1d4-baf809210db6",
				Content:  models.Content{
					Title: "a",
					Body: "b",
				},
			},
			wantError: nil,
			wantResult: models.GetByIDArticleModel{
				Content: models.Content{
					Title: "a",
					Body: "b",
				},
			},
		},
		{
			name: "fail",
			data: models.UpdateArticleResponse{
				ID:   "7e6242cd-1467-45e6-96c2-57c2be96facc",
				Content:  models.Content{},
			},
			wantError: errors.New("article not found"),
			wantResult: models.GetByIDArticleModel{
				Content: models.Content{},
			},
		},
	}


	for _, box := range tests {
		t.Run(box.name, func(t *testing.T) {
			err := IM.UpdateArticle(box.data)

			if box.wantError == nil {
				if err != nil {
					t.Errorf("IM.UpdateArticle() got error: %+v", err)
				}

				article,err := IM.GetArticleById(box.data.ID)
				if err != nil {
					t.Errorf("IM.UpdateArticle() got unexpected error: %+v", err)
				}

				if box.wantResult.Content != article.Content {
					t.Errorf("IM.UpdateArticle() expected: %+v objects, but got objects: %+v", box.wantResult.Content, article.Content)
				}
			} else {
				if box.wantError.Error() != err.Error() {
					t.Errorf("IM.UpdateArticle() expected error: %+v, but got error: %+v", box.wantError, err)
				}
			}
		})
	}
	t.Log("Test has been finished")
}