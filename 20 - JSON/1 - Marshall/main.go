package main

import (
	"bytes"
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
	c := cachorro{"Rex", "Dálmata", 3}

	cachororEmJson, erro := json.Marshal(c)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachororEmJson)

	// bytes.NewBuffer(cachororEmJson) => Dá a saída em JSON, visivel e sem ser em Bytes.
	fmt.Println(bytes.NewBuffer(cachororEmJson))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachororEmJson2, erro2 := json.Marshal(c2)
	if erro2 != nil {
		log.Fatal(erro2)
	}

	fmt.Println(cachororEmJson2)

	// bytes.NewBuffer(cachororEmJson) => Dá a saída em JSON, visivel e sem ser em Bytes.
	fmt.Println(bytes.NewBuffer(cachororEmJson2))
}
