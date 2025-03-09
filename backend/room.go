package main

type User struct {
	ID       int    `json:"ID"`
	Username string `json:"Username"`
}

type Room struct {
	ID          int    `json:"ID"`
	Label       string `json:"Label"`
	Description string `json:"Description"`
	Users       []User `json:"Users"`
}
