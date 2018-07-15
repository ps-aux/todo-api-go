package main

import "time"

type Todo struct {
	Id        string    `json:"id"`
	Completed bool      `json:"completed"`
	Name      string    `json:"name"`
	Due       time.Time `json:"due"`
}
