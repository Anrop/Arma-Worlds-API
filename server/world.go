package server

import (
	"encoding/json"
	"fmt"
	"github.com/Anrop/Arma-Worlds-API/database"
	"math"
	"net/http"
)

type layerResponse struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type sizeResponse struct {
	Height int `json:"height"`
	Width  int `json:"width"`
	Zoom   int `json:"zoom"`
}

type worldResponse struct {
	Name   string          `json:"name"`
	Title  string          `json:"title"`
	Size   sizeResponse    `json:"size"`
	Layers []layerResponse `json:"layers"`
}

func ConvertWorld(world database.World) worldResponse {
	topographicLayer := layerResponse{
		Title: "Topographic",
		Url:   fmt.Sprintf("https://maptiles.anrop.se/%s/{z}/%s_{x}_{y}.png", world.Name, world.Name),
	}

	zoom := int(math.Ceil(math.Log2(float64(world.Size / 256))))

	return worldResponse{
		Name:   world.Name,
		Title:  world.Title,
		Size:   sizeResponse{world.Size, world.Size, zoom},
		Layers: []layerResponse{topographicLayer},
	}
}

func (s *Server) Worlds() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		worlds, err := s.database.FetchWorlds(r.Context())

		if err != nil {
			ServerError(w, r, err)
			return
		}

		var worldsResponse []worldResponse
		for _, world := range *worlds {
			worldsResponse = append(worldsResponse, ConvertWorld(world))
		}

		json.NewEncoder(w).Encode(worldsResponse)
	}
}
