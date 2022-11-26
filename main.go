package main

import "time"

type Note struct {
	Title       string
	Description string
	CreateOn    time.Time
}

var noteStore = make(map[string]Note)
var id int = 0

func main() {

}
