package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type Pokemon struct {
	Name    string
	Url 	string
}

type AllPokemon struct {
	Results []Pokemon
}

type Pokemons struct {
    ID int
    Name string
}

func (a *App) GetPokemonList() []Pokemons {

    response, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=151")
    if err != nil {
        log.Fatal(err)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    var data AllPokemon
    json.Unmarshal(responseData, &data)

	var pokemon_array []Pokemons

    pokemonid := 0
    for _, value := range data.Results {
        pokemonid += 1
        pokemon_array = append(pokemon_array, Pokemons{pokemonid,value.Name})
    }

    return pokemon_array
}

func (a *App) GetImageUrlByPokemon(pokemon_id int) []string {

    url := fmt.Sprintf("%s%d%s", "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/dream-world/", pokemon_id, ".svg")
 
    var url_array []string

    url_array = append(url_array, url)
    return url_array
}