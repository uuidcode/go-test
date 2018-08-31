package main

import (
	"html/template"
	"os"
)

type Project struct {
	Name     string
	Location string
	Title    string
}

func main() {
	template := template.New("template")
	template, _ = template.Parse("Hello, {{.Title}} {{.Name}} {{.Location}} {{.Name}}")
	project := Project{
		"test", "Seoul", "테스트",
	}

	template.Execute(os.Stdout, project)
}
