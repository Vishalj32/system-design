package main

import (
	"github.com/Vishalj32/system-design/low-level-designs/snakes-n-ladders/models"
	"github.com/Vishalj32/system-design/low-level-designs/snakes-n-ladders/service"
)

func main() {

	snakeLadderService := service.NewSnakeLadderService(100, 1)
	var players []models.Player
	players = append(players, models.NewPlayer("John"), models.NewPlayer("Jane"))

	var snakes []models.Snakes
	snakes = append(snakes, models.CreateSnakes(99, 25), models.CreateSnakes(88, 54), models.CreateSnakes(29, 10))

	var ladders []models.Ladders
	ladders = append(ladders, models.CreateLadder(14, 28), models.CreateLadder(55, 97), models.CreateLadder(42, 78), models.CreateLadder(52, 92))

	snakeLadderService.SetPlayers(players)
	snakeLadderService.SetLadders(ladders)
	snakeLadderService.SetSnakes(snakes)

	snakeLadderService.StartGame()
}
