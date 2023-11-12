package main

type Response struct {
	CurrentPage int     `json:"current_page"`
	Data        []Breed `json:"data"`
}

type Breed struct {
	Breed   string `json:"breed"`
	Country string `json:"country"`
	Origin  string `json:"origin"`
	Coat    string `json:"coat"`
	Pattern string `json:"pattern"`
}
