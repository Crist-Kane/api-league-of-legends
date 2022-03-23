package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Data struct {
	Type    string `json:"type"`
	Format  string `json:"format"`
	Version string `json:"version"`
	Data    map[string]struct {
		Version string `json:"version"`
		Id      string `json:"id"`
		Key     string `json:"key"`
		Name    string `json:"name"`
		Title   string `json:"title"`
		Blurb   string `json:"blurb"`
		Info    struct {
			Attack     int `json:"attack"`
			Defense    int `json:"defense"`
			Magic      int `json:"magic"`
			Difficulty int `json:"difficulty"`
		} `json:"info"`
		Image struct {
			Full   string `json:"full"`
			Sprite string `json:"sprite"`
			Group  string `json:"group"`
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
		} `json:"image"`
		Tags    []string `json:"tags"`
		Partype string   `json:"partype"`
		Stats   struct {
			Hp                   float64 `json:"hp"`
			Hpperlevel           float64 `json:"hpperlevel"`
			Mp                   float64 `json:"mp"`
			Mpperlevel           float64 `json:"mpperlevel"`
			Movespeed            float64 `json:"movespeed"`
			Armor                float64 `json:"armor"`
			Armorperlevel        float64 `json:"armorperlevel"`
			Spellblock           float64 `json:"spellblock"`
			Spellblockperlevel   float64 `json:"spellblockperlevel"`
			Attackrange          float64 `json:"attackrange"`
			Hpregen              float64 `json:"hpregen"`
			Hpregenperlevel      float64 `json:"hpregenperlevel"`
			Mpregen              float64 `json:"mpregen"`
			Mpregenperlevel      float64 `json:"mpregenperlevel"`
			Crit                 float64 `json:"crit"`
			Critperlevel         float64 `json:"critperlevel"`
			Attackdamage         float64 `json:"attackdamage"`
			Attackdamageperlevel float64 `json:"attackdamageperlevel"`
			Attackspeedperlevel  float64 `json:"attackspeedperlevel"`
			Attackspeed          float64 `json:"attackspeed"`
		} `json:"stats"`
	} `json:"data"`
}

const port = "8776"

func main() {
	url := "http://ddragon.leagueoflegends.com/cdn/12.5.1/data/fr_FR/champion.json"

	imageServer := http.FileServer(http.Dir("images"))
	http.Handle("/images/", http.StripPrefix("/images/", imageServer))

	httpClient := http.Client{
		Timeout: time.Second * 8,
	}

	tmpl, err := template.ParseFiles("static/html/index.html")
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	response := Data{}
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/spawn", func(w http.ResponseWriter, r *http.Request) {
		tmpl = template.Must(template.ParseFiles("static/html/index.html"))
		tmpl.Execute(w, response)
	})
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		tmpl = template.Must(template.ParseFiles("static/html/menu.html"))
		button2 := r.FormValue("deuxieme")
		if button2 == "Click pour entrer" {
			http.Redirect(w, r, "http://localhost:"+port+"/spawn", 301)
		}
		tmpl.Execute(w, response)
	})
	http.HandleFunc("/lore/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.ReplaceAll(r.URL.Path, "/lore/", "")
		tmpl = template.Must(template.ParseFiles("static/html/lore.html"))
		fmt.Println(response.Data[id])
		tmpl.Execute(w, response.Data[id])
	})
	http.HandleFunc("/", Handle404)

	http.ListenAndServe(":"+port, nil)
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/html/error.html")
	button := r.FormValue("Ici")
	if button == "Allez vers le Site" {
		http.Redirect(w, r, "http://localhost:"+port+"/home", 301)
	}
	tmpl.Execute(w, r)
}
