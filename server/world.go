package server

import (
	"encoding/json"
	"fmt"
	"github.com/Anrop/Arma-Worlds-API/config"
	"github.com/Anrop/Arma-Worlds-API/database"
	"math"
	"net/http"
	"strings"
)

type layerResponse struct {
	Title string `json:"title"`
	URL   string `json:"url"`
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
	LocationsURL  string                 `json:"locationsUrl,omitempty"`
}

func convertWorld(config config.Config, world database.World) worldResponse {
	var layers []layerResponse
	var locationsURL string

	if len(config.TopographicTilesURL) > 0 {
		topographicLayer := layerResponse{
			Title: "Topographic",
			URL:   strings.ReplaceAll(config.TopographicTilesURL, "{world}", world.Name),
		}
		layers = append(layers, topographicLayer)
	}

	if len(config.SatelliteTilesURL) > 0 {
		satelliteLayer := layerResponse{
			Title: "Satellite",
			URL:   strings.ReplaceAll(config.SatelliteTilesURL, "{world}", world.Name),
		}
		layers = append(layers, satelliteLayer)
	}

	if len(config.LocationsURL) > 0 {
		locationsURL = strings.ReplaceAll(config.LocationsURL, "{world}", world.Name)
	}

	var steamWorkshop *steamWorkshopResponse
	if world.SteamWorkshopID != nil {
		steamWorkshopURL := fmt.Sprintf("https://steamcommunity.com/sharedfiles/filedetails/%d", *world.SteamWorkshopID)
		steamWorkshop = &steamWorkshopResponse{*world.SteamWorkshopID, steamWorkshopURL}
	}

	zoom := int(math.Ceil(math.Log2(float64(world.Size / 256))))

	return worldResponse{
		Name:          world.Name,
		Title:         world.Title,
		Size:          sizeResponse{world.Size, world.Size, zoom},
		SteamWorkshop: steamWorkshop,
		Layers:        layers,
		LocationsURL:  locationsURL,
	}
}

func (s *Server) worlds() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		worlds, err := s.database.FetchWorlds(r.Context())

		if err != nil {
			s.error(w, r, err)
			return
		}

		var worldsResponse []worldResponse
		for _, world := range *worlds {
			worldsResponse = append(worldsResponse, convertWorld(*s.config, world))
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		json.NewEncoder(w).Encode(worldsResponse)
	}
}
