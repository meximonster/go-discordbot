package pet

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/meximonster/go-discordbot/image"
)

var (
	dbC   *sqlx.DB
	table = "pets"
)

type Pet struct {
	Name               string
	Images             []image.Image
	LastImageURLServed string
}

func NewDB(db *sqlx.DB) {
	dbC = db
}

func (p *Pet) Type() string {
	return "pet"
}

func (p *Pet) GetName() string {
	return p.Name
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
	images, err := json.Marshal(p.Images)
	if err != nil {
		return err
	}
	q := fmt.Sprintf(`INSERT INTO %s (alias,images) VALUES ($1,$2)`, table)
	dbC.MustExec(q, p.Name, images)
	return nil
}

func (p *Pet) AddImage(text string, url string) error {
	img, err := image.ValidateImage(table, text, url)
	if err != nil {
		return err
	}
	q := fmt.Sprintf(`UPDATE %s SET images = images || '%s'::jsonb WHERE alias = %s`, table, string(img), p.Name)
	dbC.MustExec(q)
	return nil
}

func GetAll() []Pet {
	pets := []Pet{}
	dbC.Select(&pets, `SELECT * FROM pets`)
	return pets
}
