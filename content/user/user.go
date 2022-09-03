package user

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/meximonster/go-discordbot/image"
)

var (
	dbC   *sqlx.DB
	table = "users"
)

type User struct {
	Name               string
	Images             []image.Image
	LastImageURLServed string
}

func NewDB(db *sqlx.DB) {
	dbC = db
}

func (u *User) Type() string {
	return "user"
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) RandomImage() (image.Image, error) {
	img, err := image.RandomImage(u.Images, u.LastImageURLServed)
	if err != nil {
		return image.Image{}, err
	}
	u.LastImageURLServed = img.Url
	return img, nil
}

func (u *User) Store() error {
	images, err := json.Marshal(u.Images)
	if err != nil {
		return err
	}
	q := fmt.Sprintf(`INSERT INTO %s (alias,images) VALUES ($1,$2)`, table)
	dbC.MustExec(q, u.Name, images)
	return nil
}

func (u *User) AddImage(text string, url string) error {
	img, err := image.ValidateImage(table, text, url)
	if err != nil {
		return err
	}
	q := fmt.Sprintf(`UPDATE %s SET images = images || '%s'::jsonb WHERE alias = %s`, table, string(img), u.Name)
	dbC.MustExec(q)
	return nil
}

func GetAll() []User {
	users := []User{}
	dbC.Select(&users, `SELECT * FROM users`)
	return users
}
