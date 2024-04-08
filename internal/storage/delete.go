package storage

import "fmt"

func (r *Repo) Delete(id string) error {
	res, err := r.storage.Exec("DELETE FROM catalog WHERE reg_num = $1", id)
	if err != nil {
		return fmt.Errorf("deleting from database: %w", err)
	}
	rowcount, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rowcount: %w", err)
	}
	if rowcount == 0 {
		return &ErrDeleteNoEffect{RegNum: id}
	}
	return nil
}
