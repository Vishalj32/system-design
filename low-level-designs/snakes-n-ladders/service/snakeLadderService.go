package service

import (
	"container/list"
	"fmt"
	"math/rand"

	"github.com/Vishalj32/system-design/low-level-designs/snakes-n-ladders/models"
)

type SnakeLadderService struct {
	NoOfDice        int
	NumberOfPlayers int
	Board           models.SnakeLadderBoard
	Players         *list.List
}

func NewSnakeLadderService(boardSize, noOfDice int) SnakeLadderService {
	return SnakeLadderService{
		NoOfDice: noOfDice,
		Board:    models.NewSnakeLadderBoard(boardSize),
	}
}

func (s *SnakeLadderService) SetPlayers(players []models.Player) {
	playerList := list.New()
	for _, v := range players {
		playerList.PushBack(v)
	}
	s.Players = playerList
	s.NumberOfPlayers = playerList.Len()

	playerPositions := make(map[string]int, 0)
	for _, v := range players {
		playerPositions[v.Id] = 0
	}
	s.Board.PlayerPositions = playerPositions
}

func (s *SnakeLadderService) SetSnakes(snakes []models.Snakes) {
	s.Board.Snakes = snakes
}

func (s *SnakeLadderService) SetLadders(ladders []models.Ladders) {
	s.Board.Ladders = ladders
}

func (s *SnakeLadderService) getPositionAfterTurn(newPosition int) int {
	var previousPosition int

	if newPosition != previousPosition {
		previousPosition = newPosition
		for _, snake := range s.Board.Snakes {
			if snake.Start == newPosition {
				newPosition = snake.End
			}
		}

		for _, ladder := range s.Board.Ladders {
			if ladder.Start == newPosition {
				newPosition = ladder.End
			}
		}
	}
	return previousPosition
}

func (s *SnakeLadderService) roll() int {
	return rand.Intn(6*s.NoOfDice) + 1
}

func (s *SnakeLadderService) hasPlayerWon(player models.Player) bool {
	playerPosition := s.Board.PlayerPositions[player.Id]
	winningPosition := s.Board.Size

	return playerPosition == winningPosition
}

func (s *SnakeLadderService) isGameCompleted() bool {
	currentPlayers := s.Players.Len()
	return currentPlayers < s.NumberOfPlayers
}

func (s *SnakeLadderService) movePlayer(player models.Player, position int) {
	oldPosition := s.Board.PlayerPositions[player.Id]
	newPosition := oldPosition + position

	if newPosition > s.Board.Size {
		newPosition = oldPosition
	} else {
		newPosition = s.getPositionAfterTurn(newPosition)
	}

	s.Board.PlayerPositions[player.Id] = newPosition
	fmt.Println(player.Name, "rolled a", position, "and moved from", oldPosition, "to", newPosition)
}

func (s *SnakeLadderService) StartGame() {
	for !s.isGameCompleted() {
		diceRoll := s.roll()
		current := s.Players.Front()
		currentPlayer := current.Value.(models.Player)
		s.movePlayer(currentPlayer, diceRoll)
		if s.hasPlayerWon(currentPlayer) {
			fmt.Println("Player", currentPlayer.Name, "won the game")
			s.Players.Remove(current)
		} else {
			s.Players.Remove(current)
			s.Players.PushBack(currentPlayer)
		}
	}
}
