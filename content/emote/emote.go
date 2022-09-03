package emote

import (
	"encoding/json"
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

func (e *Emote) Type() string {
	return "emote"
}

func (e *Emote) GetName() string {
	return e.Alias
}

func (e *Emote) RandomImage() (image.Image, error) {
	img, err := image.RandomImage(e.Images, e.LastImageURLServed)
	if err != nil {
		return image.Image{}, err
	}
	e.LastImageURLServed = img.Url
	return img, nil
}

func (e *Emote) Store() error {
	images, err := json.Marshal(e.Images)
	if err != nil {
		return err
	}
	q := fmt.Sprintf(`INSERT INTO %s (Name,images) VALUES ($1,$2)`, table)
	dbC.MustExec(q, e.Alias, images)
	return nil
}

func (e *Emote) AddImage(text string, url string) error {
	img, err := image.ValidateImage(table, text, url)
	if err != nil {
		return err
	}
	q := fmt.Sprintf(`UPDATE %s SET images = images || '%s'::jsonb WHERE Name = %s`, table, string(img), e.Alias)
	dbC.MustExec(q)
	return nil
}

func GetAll() ([]Emote, error) {
	emotes := []Emote{}
	err := dbC.Select(&emotes, `SELECT alias, images FROM emotes`)
	if err != nil {
		return nil, err
	}
	return emotes, nil
}
