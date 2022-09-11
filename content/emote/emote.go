package emote

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/meximonster/go-discordbot/image"
)

var (
	dbC   *sqlx.DB
	table = "emotes"
)

type Emote struct {
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

func (e *Emote) Type() string {
	return "emote"
}

func (e *Emote) GetName() string {
	return e.Alias
}

func (e *Emote) RandomImage() (image.Image, error) {
	img, err := image.Random(e.Images, e.LastImageURLServed)
	if err != nil {
		return image.Image{}, err
	}
	e.LastImageURLServed = img.Url
	return img, nil
}

func (e *Emote) Store() error {
	q := fmt.Sprintf(`INSERT INTO %s (alias) VALUES ($1)`, table)
	_, err := dbC.Exec(q, e.Alias)
	return err
}

func (e *Emote) AddImage(text string, url string) error {
	img, err := image.Validate(table, text, url)
	if err != nil {
		return err
	}
	all, err := image.Add(e.Images, img)
	if err != nil {
		return err
	}
	e.Images = all
	q := fmt.Sprintf(`UPDATE %s SET images = images || '%s'::jsonb WHERE alias = %s`, table, string(img), e.Alias)
	_, err = dbC.Exec(q)
	return err
}

func GetAll() ([]*Emote, error) {
	emotes := []*Emote{}
	err := dbC.Select(&emotes, `SELECT alias, images FROM emotes`)
	if err != nil {
		return nil, err
	}
	return emotes, nil
}
