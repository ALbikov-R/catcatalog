package storage

import (
	"carcat/internal/models"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

func (r *Repo) Create(cars models.Catalog) error {
	_, err := r.storage.Exec(`INSERT INTO catalog 
		(
			reg_num, mark, model , year, 
			owner_name, owner_surname, owner_patronymic 
		) VALUES ($1,$2,$3,$4,$5,$6,$7)
		`, cars.RegNum, cars.Mark, cars.Model, cars.Year, cars.Owner.Name, cars.Owner.Surname, cars.Owner.Patronymic)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" { // Код ошибки уникальности в PostgreSQL
				return &ErrRegNumExists{RegNum: cars.RegNum}
			} else {
				return fmt.Errorf("command 'Insert' failed: %w", err)
			}
		} else {
			return fmt.Errorf("unknown error: %w", err)
		}
	}
	return nil
}
