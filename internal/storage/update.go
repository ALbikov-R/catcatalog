package storage

import (
	"carcat/internal/models"
	"fmt"
	"strings"
)

func (r *Repo) Update(patch models.Patch) error {
	var setParts []string
	args := []interface{}{}
	argCount := 0

	if patch.Mark != "" {
		argCount++
		setParts = append(setParts, fmt.Sprintf("mark = $%d", argCount))
		args = append(args, patch.Mark)
	}
	if patch.Model != "" {
		argCount++
		setParts = append(setParts, fmt.Sprintf("model = $%d", argCount))
		args = append(args, patch.Model)
	}
	if patch.Year != 0 {
		argCount++
		setParts = append(setParts, fmt.Sprintf("year = $%d", argCount))
		args = append(args, patch.Year)
	}
	if patch.Name != "" {
		argCount++
		setParts = append(setParts, fmt.Sprintf("owner_name = $%d", argCount))
		args = append(args, patch.Name)
	}
	if patch.Surname != "" {
		argCount++
		setParts = append(setParts, fmt.Sprintf("owner_surname = $%d", argCount))
		args = append(args, patch.Surname)
	}
	if patch.Patronymic != "" {
		argCount++
		setParts = append(setParts, fmt.Sprintf("owner_patronymic = $%d", argCount))
		args = append(args, patch.Patronymic)
	}

	if len(setParts) == 0 {
		return fmt.Errorf("no fields to update")
	}

	argCount++
	args = append(args, patch.RegNum)

	sqlQuery := fmt.Sprintf("UPDATE catalog SET %s WHERE reg_num = $%d", strings.Join(setParts, ", "), argCount)
	fmt.Println(sqlQuery, args)
	_, err := r.storage.Exec(sqlQuery, args...)
	if err != nil {
		return err
	}

	return nil
}
