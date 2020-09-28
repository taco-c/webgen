package main

// Website holds information about everything.
type Website struct {
	TemplateFile string `xml:"template-file"`
	Stylesheet   string `xml:"stylesheet"`
	Pages        []Page `xml:"page"`
}

// A Page is a html-document-to-be.
type Page struct {
	Path     string  `xml:"path"`
	Filename string  `xml:"filename"`
	Style    string  `xml:"style"`
	Title1   string  `xml:"title1"`
	Title2   string  `xml:"title2"`
	Title3   string  `xml:"title3"`
	Content  Content `xml:"content"`
}

// Content is the main content of the page, in HTML.
type Content struct {
	InnerXML string `xml:",innerxml"`
}
