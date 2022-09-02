package content

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var dbC *sqlx.DB

func NewDB(db *sqlx.DB) {
	dbC = db
}

func CloseDB() error {
	err := dbC.Close()
	if err != nil {
		return err
	}
	return nil
}

func Store(table string, name string, images []byte) error {
	q := fmt.Sprintf(`INSERT INTO %s (alias,images) VALUES ($1,$2)`, table)
	dbC.MustExec(q, name, images)
	return nil
}

func StoreImage(table string, name string, image string) error {
	q := fmt.Sprintf(`UPDATE %s SET images = images || '%s'::jsonb WHERE name = %s`, table, image, name)
	dbC.MustExec(q, table, image, name)
	return nil
}
