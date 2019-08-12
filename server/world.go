package server

import (
	"encoding/json"
	"fmt"
	"github.com/Anrop/Arma-Worlds-API/config"
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

type steamWorkshopResponse struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

type worldResponse struct {
	Name          string                 `json:"name"`
	Title         string                 `json:"title"`
	Size          sizeResponse           `json:"size"`
	SteamWorkshop *steamWorkshopResponse `json:"steamWorkshop"`
	Layers        []layerResponse        `json:"layers"`
}

func ConvertWorld(config config.Config, world database.World) worldResponse {
	topographicLayer := layerResponse{
		Title: "Topographic",
		Url:   fmt.Sprintf("%s/%s/{z}/%s_{x}_{y}.png", config.TopographicTilesBaseURL, world.Name, world.Name),
	}

	var steamWorkshop *steamWorkshopResponse
	if world.SteamWorkshopID != nil {
		steamWorkshopUrl := fmt.Sprintf("https://steamcommunity.com/sharedfiles/filedetails/%d", *world.SteamWorkshopID)
		steamWorkshop = &steamWorkshopResponse{*world.SteamWorkshopID, steamWorkshopUrl}
	}

	zoom := int(math.Ceil(math.Log2(float64(world.Size / 256))))

	return worldResponse{
		Name:          world.Name,
		Title:         world.Title,
		Size:          sizeResponse{world.Size, world.Size, zoom},
		SteamWorkshop: steamWorkshop,
		Layers:        []layerResponse{topographicLayer},
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
			worldsResponse = append(worldsResponse, ConvertWorld(*s.config, world))
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		json.NewEncoder(w).Encode(worldsResponse)
	}
}
