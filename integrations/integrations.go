package integrations

import (
	"github.com/meximonster/go-discordbot/integrations/pubg"
	"github.com/meximonster/go-discordbot/integrations/wow"
)

func Initialize(bnet_client_id string, bnet_client_secret string, pubg_api_key string, pubg_current_season string) {
	wow.LoadAuthVars(bnet_client_id, bnet_client_secret)
	pubg.InitAuth(pubg_api_key, pubg_current_season)
}
