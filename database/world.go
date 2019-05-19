package database

import (
	"context"
	"fmt"
	"os"
)

type World struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Title string `json:"title" db:"title"`
	Size  int    `json:"size" db:"size"`
}

func (db *Database) FetchWorlds(ctx context.Context) (*[]World, error) {
	var query = `
		SELECT
			worlds.id,
			worlds.name,
			worlds.title,
			worlds.size
		FROM
			worlds
		ORDER BY
			worlds.title
	`
	rows, err := db.client.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var worlds []World

	for rows.Next() {
		var world World
		err = rows.Scan(&world.ID, &world.Name, &world.Title, &world.Size)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading user from database: %q\n", err)
		} else {
			worlds = append(worlds, world)
		}
	}

	return &worlds, nil
}
