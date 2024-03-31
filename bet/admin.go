package bet

import (
	"github.com/meximonster/go-discordbot/configuration"
)

var admins = []*Admin{}

type Admin struct {
	Name    string
	Id      string
	Channel string
	Table   string
}

func InitAdmins(cfg []configuration.AdminConfig) {
	for _, c := range cfg {
		a := new(Admin)
		a.Name = c.Name
		a.Id = c.Id
		a.Channel = c.Channel
		a.Table = c.Table
		admins = append(admins, a)
	}
}

func GetAdmins() []*Admin {
	return admins
}

func GetTableFromChannel(channel string) string {
	for _, admin := range admins {
		if admin.Channel == channel {
			return admin.Table
		}
	}
	return ""
}

func IsAdmin(id string) bool {
	for _, admin := range admins {
		if admin.Id == id {
			return true
		}
	}
	return false
}

func IsBetCandidate(id string, channel string) bool {
	for _, admin := range admins {
		if admin.Id == id && admin.Channel == channel {
			return true
		}
	}
	return false
}

func IsBetChannel(channel string) bool {
	for _, admin := range admins {
		if admin.Channel == channel {
			return true
		}
	}
	return false
}
