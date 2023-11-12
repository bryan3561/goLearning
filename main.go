package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("default!")
	app := App{}

	app.Run()
}

type App struct {
	// ...
}

func (a *App) Run() error {
	http.Handle("/", http.FileServer(http.Dir("./web/static")))

	http.ListenAndServe(":8080", nil)

	return nil
}
