package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"fmt"
	"github.com/andygrunwald/go-trending"
	"log"
	"net/http"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/repos/:lang", GetReposLang),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	fmt.Println("Start API on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type Repo struct {
	Name string
	Language string
	Stars int8
}

type Repos struct {
	items []Repo
}


func GetReposLang(w rest.ResponseWriter, r *rest.Request) {
	code := r.PathParam("lang")
	trend := trending.NewTrending()

	// Show projects of today
	projects, err := trend.GetProjects(trending.TimeWeek, code)

	if projects == nil || err != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(projects)
}



