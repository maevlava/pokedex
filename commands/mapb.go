package commands

import (
	"fmt"
	"github.com/maevlava/pokedex/model"
)

type PokeMapBackwardCommand struct {
	Pm *PokeMapCommand
}

func (n *PokeMapBackwardCommand) Name() string {
	return "mapb"
}

func (n *PokeMapBackwardCommand) Description() string {
	return "Return previous maps in pokemon"
}

func (n *PokeMapBackwardCommand) Execute(user *model.User, args ...string) error {
	loadedCount := len(n.Pm.PokeMaps)
	if loadedCount < 20 {
		return fmt.Errorf("map is not ready yet")
	}
	totalPage := (loadedCount + 19) / 20

	lastPrintedPage := (n.Pm.Page + totalPage - 1) % totalPage

	var newPage int
	if lastPrintedPage == 0 {
		newPage = 0
	} else {
		newPage = lastPrintedPage - 1
	}

	n.Pm.Page = newPage + 1

	start := newPage * 20
	end := start + 20
	if end > loadedCount {
		end = loadedCount
	}
	currentMaps := make([]model.Location, end-start)
	copy(currentMaps, n.Pm.PokeMaps[start:end])

	fmt.Printf("--- Page %d/%d ---\n", newPage+1, totalPage)
	for _, m := range currentMaps {
		fmt.Printf("Location: %s (ID: %d)\n", m.Name, m.ID)
	}

	return nil
}
