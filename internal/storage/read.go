package storage

import (
	"carcat/internal/models"
	"fmt"
)

func (r *Repo) Read(fil models.Filter) ([]models.Catalog, error) {
	query, args := queryrow(fil)
	fmt.Println(query, args)
	rows, err := r.storage.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Catalog
	for rows.Next() {
		var c models.Catalog
		err := rows.Scan(&c.RegNum, &c.Mark, &c.Model, &c.Year, &c.Owner.Name, &c.Owner.Surname, &c.Owner.Patronymic)
		if err != nil {
			return nil, err
		}
		result = append(result, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil

}
func queryrow(fil models.Filter) (string, []interface{}) {
	query := "SELECT reg_num, mark, model, year, owner_name, owner_surname, owner_patronymic FROM catalog WHERE 1=1"
	args := []interface{}{}
	argCount := 1

	if fil.Regnum != "" {
		query += fmt.Sprintf(" AND reg_num = $%d", argCount)
		args = append(args, fil.Regnum)
		argCount++
	}
	if fil.Mark != "" {
		query += fmt.Sprintf(" AND mark = $%d", argCount)
		args = append(args, fil.Mark)
		argCount++
	}
	if fil.Model != "" {
		query += fmt.Sprintf(" AND model = $%d", argCount)
		args = append(args, fil.Model)
		argCount++
	}
	if fil.Year != 0 {
		query += fmt.Sprintf(" AND year = $%d", argCount)
		args = append(args, fil.Year)
		argCount++
	}
	if fil.Lyear != 0 && fil.Tyear != 0 {
		query += fmt.Sprintf(" AND year >= $%d AND year <=$%d", argCount, argCount+1)
		args = append(args, fil.Lyear, fil.Tyear)
		argCount += 2
	}
	if fil.Name != "" {
		query += fmt.Sprintf(" AND owner_name = $%d", argCount)
		args = append(args, fil.Name)
		argCount++
	}
	if fil.Surname != "" {
		query += fmt.Sprintf(" AND owner_surname = $%d", argCount)
		args = append(args, fil.Surname)
		argCount++
	}
	if fil.Patronymic != "" {
		query += fmt.Sprintf(" AND owner_patronymic = $%d", argCount)
		args = append(args, fil.Patronymic)
		argCount++
	}
	return query, args
}
