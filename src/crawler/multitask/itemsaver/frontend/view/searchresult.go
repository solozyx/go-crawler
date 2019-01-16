package view

import(
	"html/template"
	"io"
	"crawler/multitask/itemsaver/frontend/model"
)

type SearchResultView struct{
	tpl *template.Template
}

func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		tpl:template.Must(template.ParseFiles(filename)),
	}
}

func (s SearchResultView)Render(w io.Writer,data model.SearchResult) error {
	return s.tpl.Execute(w,data)
}