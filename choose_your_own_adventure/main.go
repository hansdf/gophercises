package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func main() {

	arcs := loadTemplates("story.json")

	introArc := arcs["intro"]
	// mushroomArc := arcs["mushroom-path"]

	http.HandleFunc("/intro", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "This is the beginning of the magical Enchanted Grove adventure 🐞 ")
		fmt.Fprintf(w, introArc.Title)
		fmt.Fprintf(w, introArc.Story[0])
		// fmt.Fprintf(w, introArc)

	})

	fmt.Println(introArc.Title)
	fmt.Println(introArc.Story)
	fmt.Println(introArc.Options[2])

	// http.ListenAndServe(":8080", nil)

}

func loadTemplates(filename string) map[string]Arc {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var arcs map[string]Arc
	err = json.Unmarshal(file, &arcs)
	if err != nil {
		panic(err)
	}

	return arcs
}
