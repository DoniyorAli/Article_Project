package postgres

import (
	"UacademyGo/Article/models"
	"errors"
)

func (stg Postgres) AddAuthor(id string, box models.CreateModelAuthor) error {
	_, err := stg.homeDB.Exec(`INSERT INTO author 
	(
		id,
		firstname,
		lastname,
		middlename
	) VALUES (
		$1,
		$2,
		$3
		$4
	)`,
		id,
		box.Firstname,
		box.Lastname,
		box.Middlename,
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
		middlename,
		created_at,
		updated_at,
		deleted_at
    FROM author WHERE id = $1 AND deleted_at IS NULL`, id).Scan(
		&author.ID,
		&author.Firstname,
		&author.Lastname,
		&author.Middlename,
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
		middlename,
		created_at,
		updated_at,
		deleted_at 
		FROM author
		WHERE ((firstname ILIKE '%' || $1 || '%') or (lastname ILIKE '%' || $1 || '%') or (middlename ILIKE '%' || $1 || '%')) AND deleted_at IS NULL
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
			&author.Middlename,
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

	res, err := stg.homeDB.NamedExec("UPDATE author  SET firstname=:f, lastname=:l, middlename=:m, updated_at=now() WHERE deleted_at IS NULL AND id=:id", map[string]interface{}{
		"id": box.ID,
		"f":  box.Firstname,
		"l":  box.Lastname,
		"m":  box.Middlename,
	})
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect > 0 {
		return nil
	}
	return errors.New("author not found")
}

//*=========================================================================
func (stg Postgres) DeleteAuthor(id string) error {
	res, err := stg.homeDB.Exec("UPDATE author SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL", id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect > 0 {
		return nil
	}
	return errors.New("author not found")
}