package pet

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/meximonster/go-discordbot/image"
)

var (
	dbC   *sqlx.DB
	table = "pets"
)

type Pet struct {
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

func (p *Pet) Type() string {
	return "pet"
}

func (p *Pet) GetName() string {
	return p.Alias
}

func (p *Pet) RandomImage() (image.Image, error) {
	img, err := image.Random(p.Images, p.LastImageURLServed)
	if err != nil {
		return image.Image{}, err
	}
	p.LastImageURLServed = img.Url
	return img, nil
}

func (p *Pet) Store() error {
	q := fmt.Sprintf(`INSERT INTO %s (alias) VALUES ($1)`, table)
	_, err := dbC.Exec(q, p.Alias)
	return err
}

func (p *Pet) AddImage(text string, url string) error {
	var concat string
	img, err := image.Validate(table, text, url)
	if err != nil {
		return err
	}
	all, err := image.Add(p.Images, img)
	if err != nil {
		return err
	}
	if len(p.Images) != 0 {
		concat = "images ||"
	}
	p.Images = all
	q := fmt.Sprintf(`UPDATE %s SET images = %s '%s'::jsonb WHERE alias = '%s'`, table, concat, string(img), p.Alias)
	_, err = dbC.Exec(q)
	return err
}

func GetAll() ([]*Pet, error) {
	pets := []*Pet{}
	err := dbC.Select(&pets, `SELECT alias, images FROM pets`)
	if err != nil {
		return nil, err
	}
	return pets, nil
}
