package Players

import (
	"github.com/user/RestApi/Models"
)

var Players []Models.Player

func AddPlayer(player Models.Player) {
	Players = append(Players, player)
}

func GetPlayer(id int) Models.Player {

	for _, player := range Players {

		if player.Id == id {
			return player
		}
	}
	return Models.Player{}
}

func DeletePlayer(id int) []Models.Player {

	for index, player := range Players {

		if player.Id == id {
			Players = append(Players[:index], Players[index+1:]...)
			break
		}
	}
	return Players
}
