package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
)

const (
	apiUrl      = "https://catfact.ninja/breeds"
	limit       = 98
	outFileName = "out.json"
)

func main() {
	url := fmt.Sprintf("%s?limit=%d", apiUrl, limit)

	// получаем *http.Response, не забываем закрыть
	resp, err := http.Get(url)
	if err != nil {
		log.Println(errGetRequest, err)
		return

	}

	// проверяем статус, так как можем получить статус не 200
	if resp.StatusCode != http.StatusOK {
		log.Println(errResponseStatus)
		return
	}
	defer resp.Body.Close()

	// читаем тело ответа в срез байтов
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// парсим срез байтов и записываем в нашу структуру что бы исполнять манипуляции (группировка, сортировка)
	var res Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println(err)
		return
	}

	// распределям породы по странам, ключ мапы это страны, породы будут храниться в виде среза (группировка)
	originMap := make(map[string][]Breed)
	for _, breed := range res.Data {
		originMap[breed.Country] = append(originMap[breed.Country], breed)
	}

	// сортируем по названию породы (сортировка)
	for _, breed := range originMap {
		sort.Sort(breedNameSorter(breed))
	}

	// создание файла куда будем записывать
	file, err := os.Create(outFileName)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	// marshalIndent для того что бы в файл записался красивый json
	jsonRes, err := json.MarshalIndent(originMap, "", "    ")
	n, err := file.Write(jsonRes)
	if err != nil {
		log.Println(errWrite, err)
		return
	}

	log.Printf("Sucessfully wrote %d bytes\n", n)

}
