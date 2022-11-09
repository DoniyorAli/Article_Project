package postgres

import (
	"UacademyGo/Article/models"
	"errors"
)

func (stg Postgres) AddAuthor(id string, box models.CreateModelAuthor) error {
	var err error
	if len(id) > 36 {
		return errors.New("ID was entered wrong! exc: ID must be entered with 36 or big from 36 elements")
	}

	_, err = stg.homeDB.Exec(`INSERT INTO author 
	(
		id,
		firstname,
		lastname,
	) VALUES (
		$1,
		$2,
		$3,
	)`,
		id,
		box.Firstname,
		box.Lastname,
	)
	if err != nil {
		return err
	}
	return nil
}

//*=========================================================================
func (stg Postgres) GetAuthorById(id string) (models.Author, error) {
	var author models.Author
	err := stg.homeDB.QueryRow(`SELECT 
		id,
		firstname,
		lastname,
		created_at,
		updated_at,
		deleted_at
    FROM author WHERE id = $1 AND delete_at NULL`, id).Scan(
		&author.ID,
		&author.Firstname,
		&author.Lastname,
		&author.CreateAt,
		&author.UpdateAt,
		&author.DeletedAt,
	)
	if err != nil {
		return author, err
	}

	return author, nil
}

//*=========================================================================
func (stg Postgres) GetAuthorList(offset, limit int, search string) ([]models.Author, error) {
	var res []models.Author
	rows, err := stg.homeDB.Queryx(`SELECT 
		id,
		firstname,
		lastname,
		created_at,
		updated_at,
		deleted_at 
		FROM author
		WHERE ((firstname ILIKE '%' || $1 || '%') or (lastname ILIKE '%' || $1 || '%') ) AND deleted_at IS NULL
		LIMIT $2
		OFFSET $3
	`,
		search,
		limit,
		offset,
	)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var author models.Author
		err := rows.Scan(
			&author.ID,
			&author.Firstname,
			&author.Lastname,
			&author.CreateAt,
			&author.UpdateAt,
			&author.DeletedAt,
		)
		if err != nil {
			return res, err
		}

		res = append(res, author)

	}

	return res, err
}

//*=========================================================================
func (stg Postgres) UpdateAuthor(box models.UpdateAuthorResponse) error {
	// var temp models.Author
	// for i, v := range inM.DB.InMemoryAuthorData {
	// 	if v.ID == box.ID {
	// 		temp = v
	// 		temp.Firstname = box.Firstname
	// 		temp.Lastname = box.Lastname
	// 		t := time.Now()
	// 		temp.UpdateAt = &t
	// 		inM.DB.InMemoryAuthorData[i] = temp
		
	// 		return nil
	// 	}
	// }
	return errors.New("author not found")
}

//*=========================================================================
func (stg Postgres) DeleteAuthor(id string) error {
	// for i, v := range inM.DB.InMemoryAuthorData {
	// 	if v.ID == id {
	// 		// InMemoryArticleData = remove(InMemoryArticleData, i)
	// 		inM.DB.InMemoryAuthorData = append(inM.DB.InMemoryAuthorData[:i], inM.DB.InMemoryAuthorData[i+1:]...)
	// 		return nil
	// 	}
	// }
	return errors.New("deletion failed")
}