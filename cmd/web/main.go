package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/bblank70/GoogleChartsTesting/pkg/handlers"
)

//var tpl *template.Template

// TODO: WE NEED A TYPE FOR EVERY CHART/OBJECT
type Barchart struct {
	Star5 int `json:"5-Star"`
	Star4 int `json:"4-Star"`
	Star3 int `json:"3-Star"`
	Star2 int `json:"2-Star"`
	Star1 int `json:"1-Star"`
}

// /these will eventually get deprecated
type Collection struct {
	Ratings []Barchart       `json:"barchart"`
	History []BusinessRating `json:"recents"`
	//	Indicators []KPI
}

var Dataslice []Barchart

// geoslice and Geoslice allow for storage of the coordinates to map with leaflet

type Geomap map[string]float64

type BusinessRating struct {
	BusinessName string  `json:"name"`
	Stars        float64 `json:"stars"`
	YourRating   int64   `json:"YourRating"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Photopath    string  `json:"hyperlink"`
}

// THEN WE NEED TO COLLECT THEM INTO SOME CONTAINER (probably a struct)
// / Chart Data is the data object we will pass into the template
type DashboardData struct {
	Barchart Barchart
	Geo      []Geomap
	Recents  []BusinessRating
	//	KPI      Indicator
}

func main() {

	jsonFile, err := os.Open("./json/dashboard.json")

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
	var Dashboard DashboardData

	json.Unmarshal(byteValue, &c)
	fmt.Println("This is c:", c)

	for i := 0; i < len(c.Ratings); i++ {
		ChartRecord := Barchart{
			Star5: c.Ratings[i].Star5,
			Star4: c.Ratings[i].Star4,
			Star3: c.Ratings[i].Star3,
			Star2: c.Ratings[i].Star2,
			Star1: c.Ratings[i].Star1,
		}

		fmt.Println("ChartRecord is:", ChartRecord)
		Dashboard.Barchart = ChartRecord
	}

	var G = make(Geomap)
	var Geoslice []Geomap

	var RatingSlice []BusinessRating

	for i := 0; i < len(c.History); i++ {
		lat := c.History[i].Latitude
		long := c.History[i].Longitude
		G["lat"] = lat
		G["long"] = long
		RecentReview := BusinessRating{
			BusinessName: c.History[i].BusinessName,
			Stars:        c.History[i].Stars,
			YourRating:   c.History[i].YourRating,
			Photopath:    c.History[i].Photopath,
		}
		RatingSlice = append(RatingSlice, RecentReview)
		Geoslice = append(Geoslice, G)
	}
	Dashboard.Recents = RatingSlice
	Dashboard.Geo = Geoslice

	fmt.Println("The data exported to the template will be:", Dashboard)

	// these are our paths
	http.HandleFunc("/", handlers.Index)
	// http.HandleFunc("/verify", verifyer)
	// http.HandleFunc("/response", responder)
	// //this starts the server
	http.ListenAndServe(":6060", nil)

}

// ////////////////////////////////////////////

// index handles the route to the home page
