package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// json.Marshal() => Converte um Map ou um Struct para JSON.

// json.Unmarshal() => É o processo reverso. Transforma o JSON em uma Struct ou Map.

type cachorro struct {
	Nome  string `json:"nome"`
	Raça  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	//Struct
	cachorroEmJson := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	//Converter para Map

	cachorro2EmJson := `{"nome":"Toby", "raca":"Poodle"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJson), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)

}
