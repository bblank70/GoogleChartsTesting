package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"

	"github.com/bblank70/GoogleChartsTesting/pkg/handlers"
)

var tpl *template.Template

type President struct {
	Name       string  `json:"President"`
	DeathCause string  `json:"Cause of Death"`
	Age        float64 `json:"Age"`
	AvgRank    int     `json:"AvgRank"`
	Height     float64 `json:"Height"`
	Weight     float64 `json:"Weight"`
}

type Collection struct {
	Presidents []President `json:"Data"`
}

var Dataslice []President

// type RecommendataionResult []Business

// init instantiates the templates, they must be .tmpl extenstions
// func init() {
// 	tpl = template.Must(template.ParseGlob("templates/*.html"))
// }

func main() {

	jsonFile, err := os.Open("./json/Presidents.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened JSON FILE")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// byteValue, _ := ioutil.ReadAll(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)
	var c Collection
	json.Unmarshal(byteValue, &c)

	for i := 0; i < len(c.Presidents); i++ {
		presRecord := President{
			Name:       c.Presidents[i].Name,
			DeathCause: c.Presidents[i].DeathCause,
			Age:        c.Presidents[i].Age,
			AvgRank:    c.Presidents[i].AvgRank,
			Height:     c.Presidents[i].Height,
			Weight:     c.Presidents[i].Weight,
		}

		Dataslice = append(Dataslice, presRecord)

	}
	fmt.Println(Dataslice)

	// these are our paths
	http.HandleFunc("/", handlers.Index)
	// http.HandleFunc("/verify", verifyer)
	// http.HandleFunc("/response", responder)
	// //this starts the server
	http.ListenAndServe(":6060", nil)

}

// ////////////////////////////////////////////

// index handles the route to the home page
