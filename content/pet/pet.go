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

func (p *Pet) Type() string {
	return "pet"
}

func (p *Pet) GetName() string {
	return p.Alias
}

func (p *Pet) RandomImage() (image.Image, error) {
	img, err := image.RandomImage(p.Images, p.LastImageURLServed)
	if err != nil {
		return image.Image{}, err
	}
	p.LastImageURLServed = img.Url
	return img, nil
}

func (p *Pet) Store() error {
	q := fmt.Sprintf(`INSERT INTO %s (alias) VALUES ($1)`, table)
	dbC.MustExec(q, p.Alias)
	return nil
}

func (p *Pet) AddImage(text string, url string) error {
	img, err := image.ValidateImage(table, text, url)
	if err != nil {
		return err
	}
	all, err := image.AddImage(p.Images, img)
	if err != nil {
		return err
	}
	p.Images = all
	q := fmt.Sprintf(`UPDATE %s SET images = images || '%s'::jsonb WHERE alias = %s`, table, string(img), p.Alias)
	dbC.MustExec(q)
	return nil
}

func GetAll() ([]*Pet, error) {
	pets := []*Pet{}
	err := dbC.Select(&pets, `SELECT alias, images FROM pets`)
	if err != nil {
		return nil, err
	}
	return pets, nil
}
