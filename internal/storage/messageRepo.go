package storage

import (
	"database/sql"
	"fmt"
)

type Repo struct {
	storage *sql.DB
}
type ErrRegNumExists struct {
	RegNum string
}

func (e *ErrRegNumExists) Error() string {
	return fmt.Sprintf("element reg_num is already existes: %s", e.RegNum)
}

type ErrDeleteNoEffect struct {
	RegNum string
}

func (e *ErrDeleteNoEffect) Error() string {
	return fmt.Sprintf("element %s does not exsites in database", e.RegNum)
}
