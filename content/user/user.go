package user

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/meximonster/go-discordbot/image"
)

var (
	dbC   *sqlx.DB
	table = "users"
)

type User struct {
	Alias              string `db:"alias"`
	Images             []byte `db:"images"`
	LastImageURLServed string
}

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

func (u *User) Type() string {
	return "user"
}

func (u *User) GetName() string {
	return u.Alias
}

func (u *User) RandomImage() (image.Image, error) {
	img, err := image.Random(u.Images, u.LastImageURLServed)
	if err != nil {
		return image.Image{}, err
	}
	u.LastImageURLServed = img.Url
	return img, nil
}

func (u *User) Store() error {
	q := fmt.Sprintf(`INSERT INTO %s (alias) VALUES ($1)`, table)
	dbC.MustExec(q, u.Alias)
	return nil
}

func (u *User) AddImage(text string, url string) error {
	img, err := image.Validate(table, text, url)
	if err != nil {
		return err
	}
	all, err := image.Add(u.Images, img)
	if err != nil {
		return err
	}
	u.Images = all
	q := fmt.Sprintf(`UPDATE %s SET images = images || '%s'::jsonb WHERE alias = '%s'`, table, string(img), u.Alias)
	dbC.MustExec(q)
	return nil
}

func GetAll() ([]*User, error) {
	users := []*User{}
	err := dbC.Select(&users, `SELECT alias, images FROM users`)
	if err != nil {
		return nil, err
	}
	return users, nil
}
