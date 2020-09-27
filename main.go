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
	templateString := string(templateBytes)

	replace(&templateString, "title1", page.Title1)
	replace(&templateString, "title2", page.Title2)
	replace(&templateString, "title3", page.Title3)
	replace(&templateString, "style", page.Style)
	replace(&templateString, "content", page.Content.InnerXML)

	filepath := fmt.Sprintf("%s%c%s.html", page.Path, os.PathSeparator, page.Filename)
	// fmt.Println(getFileDir(fmt.Sprintf("%s%c%s", page.Path, os.PathSeparator, page.Filename)))
	fmt.Println("Target:", root, "\\", filepath)
	fullpath := fmt.Sprintf("%s%c%s", root, os.PathSeparator, filepath)

	// Clean me up, please
	f, err := os.Create(fullpath)
	if err != nil {
		os.MkdirAll(fmt.Sprintf("%s\\%s", root, page.Path), os.ModeDir)
		f, err2 := os.Create(fullpath)
		if err2 != nil {
			panic(err)
		}
		defer f.Close()
	}
	defer f.Close()

	f.WriteString(templateString)
}

func replace(text *string, name, value string) {
	r, err := regexp.Compile(fmt.Sprintf(`{{\s*%s\s*}}`, name))
	if err != nil {
		panic(err)
	}
	*text = r.ReplaceAllString(*text, value)
}
