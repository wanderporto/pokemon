package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	timeNow := time.Now()

	pokemons := getNames()

	for _, name := range pokemons {
		go requestApi(`https://digimon-api.vercel.app/api/digimon/name/` + name)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Printf("Duration: ", time.Since(timeNow))

}

type Pokemon struct {
	Name  string `json:"name"`
	Img   string `json: "imagem"`
	Level string `json: "level"`
}

func getNames() []string {
	content, err := ioutil.ReadFile("pokemon.txt")

	if err != nil {
		panic(err.Error())
	}

	return strings.Split(string(content), "\n")
}

func requestApi(url string) {
	defer wg.Done()
	var pokemons []Pokemon

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	errJson := json.NewDecoder(res.Body).Decode(&pokemons)

	if errJson != nil {
		panic(errJson.Error())
	}

	fmt.Println(pokemons)

}
