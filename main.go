package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Article .
type Article struct {
	ID          int
	Link        string
	Name        string
	Description string
	Cover       string
}

var globalDescription string = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ullamcorper tortor ut tempus iaculis. Quisque nec sem vitae ex rutrum iaculis vel vel enim. Quisque at tempus tortor. Aliquam imperdiet faucibus vulputate. Ut iaculis velit vitae lobortis tempus. Vivamus fermentum suscipit pulvinar. Proin quam purus, auctor vel lacinia in, iaculis vitae nisl. Cras ac est tellus. Vestibulum commodo orci sed arcu sagittis commodo. Vivamus ipsum mi, mattis non vulputate ut, lobortis eu lacus. Mauris eleifend sem urna, ac ultricies ante vestibulum eget.

Sed lobortis vel purus at sagittis. Donec ultrices odio turpis, eu commodo orci lacinia non. Aenean ultrices libero lorem, id lobortis augue aliquet id. Suspendisse et egestas dolor, quis iaculis justo. Ut placerat venenatis elit sit amet dapibus. Aenean eu ipsum suscipit, gravida felis in, tincidunt diam. Sed convallis consectetur turpis in varius. Phasellus tincidunt ultrices mi.

`

var articles = []Article{
	{1,
		"https://www.google.com",
		"Google needs more developers",
		globalDescription,
		"https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRe8i6HbyeJRMojFLmuFrHWXcvd_ni6WHXHKZ1Yj-cVnKlG7yXJ",
	},
	{2,
		"https://www.gazzetanews.com",
		"Financial issues makes the Firefox fired many developers",
		globalDescription,
		"https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRe8i6HbyeJRMojFLmuFrHWXcvd_ni6WHXHKZ1Yj-cVnKlG7yXJ",
	},
	{3,
		"https://www.oracle.com",
		"Oracle is dead",
		globalDescription,
		"https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRe8i6HbyeJRMojFLmuFrHWXcvd_ni6WHXHKZ1Yj-cVnKlG7yXJ",
	},
	{4,
		"https://www.hackernews.com",
		"Jeff Bezzos want to you leave the power point",
		globalDescription,
		"https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRe8i6HbyeJRMojFLmuFrHWXcvd_ni6WHXHKZ1Yj-cVnKlG7yXJ",
	},
	{5,
		"https://www.hackernews.com",
		"A new zero day was discovery",
		globalDescription,
		"https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRe8i6HbyeJRMojFLmuFrHWXcvd_ni6WHXHKZ1Yj-cVnKlG7yXJ",
	},
}

// Page .
type Page struct {
	Articles []Article
}

var page = &Page{
	Articles: articles,
}

func handler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./template/index.html")
	t.Execute(w, page)
}

func details(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("id")
	if query == "" {
		http.Error(w, "Query inválida", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(query)
	if err != nil {
		http.Error(w, "Tipo  inválido", http.StatusBadRequest)
		return
	}

	t, _ := template.ParseFiles("./template/detail.html")
	var search Article

	for _, article := range articles {
		if article.ID == id {
			search = article
			break
		}
	}

	t.Execute(w, search)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler)
	http.HandleFunc("/details", details)
	http.HandleFunc("/articles/new", newArticle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
