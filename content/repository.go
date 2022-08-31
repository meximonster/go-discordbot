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

func AddImages(table string, name string, image string) error {
	q := fmt.Sprintf(`UPADTE %s SET images = images || '%s'::jsonb WHERE name = %s`, table, image, name)
	dbC.MustExec(q, table, image, name)
	return nil
}

func Store(table string, name string, images []byte, discordID string) error {
	q := fmt.Sprintf(`INSERT INTO %s (alias,images,discord_id) VALUES ($1,$2,$3)`, table)
	dbC.MustExec(q, name, images, discordID)
	return nil
}
