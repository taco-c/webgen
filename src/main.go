package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		return
	}
	filename := os.Args[1]

	outputDir := "."
	if len(os.Args) > 2 {
		outputDir = os.Args[2]
	}
	// outputDir = fmt.Sprintf("%s%c%s", getFileDir(outputDir), os.PathSeparator, outputDir)
	outputDir, _ = filepath.Abs(outputDir)
	fmt.Println("Output:", outputDir)

	xmlData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	root := getFileDir(filename)
	fmt.Println("Source root:", root)

	var website Website
	if err := xml.Unmarshal(xmlData, &website); err != nil {
		panic(err)
	}

	styleFile := fmt.Sprintf("%s%c%s", root, os.PathSeparator, website.Stylesheet)
	styleData, err := ioutil.ReadFile(styleFile)
	if err != nil {
		panic(err)
	}
	globalStyle := string(styleData)

	templateFile := fmt.Sprintf("%s%c%s", root, os.PathSeparator, website.TemplateFile)
	for _, page := range website.Pages {
		page.Style = globalStyle + page.Style
		writePage(&page, outputDir, templateFile)
	}
}

func getFileDir(filename string) string {
	path, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}
	root := filepath.Dir(path)
	return root
}

func writePage(page *Page, root, templateFilename string) {
	templateBytes, err := ioutil.ReadFile(templateFilename)
	if err != nil {
		panic(err)
	}
	html := string(templateBytes)

	replace(&html, "title1", page.Title1)
	replace(&html, "title2", page.Title2)
	replace(&html, "title3", page.Title3)
	replace(&html, "style", page.Style)
	replace(&html, "content", page.Content.InnerXML)

	fullpath := fmt.Sprintf("%s%c%s%c%s.html", root, os.PathSeparator, page.Path, os.PathSeparator, page.Filename)
	fmt.Println(fullpath)

	// Clean me up, please
	f, err := os.Create(fullpath)
	if err != nil {
		os.MkdirAll(fmt.Sprintf("%s%c%s", root, os.PathSeparator, page.Path), os.ModeDir)
		f, err2 := os.Create(fullpath)
		if err2 != nil {
			panic(err)
		}
		defer f.Close()
	}
	defer f.Close()

	f.WriteString(html)
}

func replace(text *string, name, value string) {
	r, err := regexp.Compile(fmt.Sprintf(`{{\s*%s\s*}}`, name))
	if err != nil {
		panic(err)
	}
	*text = r.ReplaceAllString(*text, value)
}
