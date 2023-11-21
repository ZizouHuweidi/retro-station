package uow

import (
	data2 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/contracts/data"
)

type catalogContext struct {
	gameRepository data2.GameRepository
}

func (c *catalogContext) Games() data2.GameRepository {
	return c.gameRepository
}
