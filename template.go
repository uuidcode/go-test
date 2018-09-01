package main

import (
	"html/template"
	"os"
)

type Project struct {
	Name           string
	Location       string
	Title          string
	ThisIsNewField string
}

func main() {
	template := template.New("template")

	content :=
		`	Hello, 
	{{.Title}} {{.Name}} 
	{{.Location}} {{.Name}}`

	template, _ = template.Parse(content)

	project := Project{
		"test", "Seoul", "테스트", "한글",
	}

	template.Execute(os.Stdout, project)
}
