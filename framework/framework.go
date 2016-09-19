package framework

import (
    "fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type AppData struct {
	Appname    string
	Author     string
	Date       string
	TplPath    string
	OutputPath string
}

func New(appname string, author string, tplPath string, outputPath string) (error, AppData) {
	t := time.Now()
	app := AppData{appname, author, fmt.Sprintf("%d-%.2d-%.2d", t.Year(), t.Month(), t.Day()), tplPath, outputPath}

	return nil, app
}

func (app *AppData) Generate() error {
	fmt.Println("Processing...")

	//walk file list
	file_list, err := walkFileList(app.TplPath)
	if err != nil {
		//log err
		return err
	}

	//create file
	err = app.createAppFiles(file_list)
	if err != nil {
		return err
	}

	fmt.Println("Success.")

	return nil
}

func walkFileList(path string) ([]string, error) {
	var file_list = make([]string, 0)
	i := 0
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		file_list = append(file_list, path)
		i++
		return nil
	})
	return file_list, err
}

func (app *AppData) createAppFiles(file_list []string) error {
	//remove if exist
	err := os.RemoveAll(app.Appname)
	if err != nil {
		return err
	}

	//create file
	for i := 0; i < len(file_list); i++ {
		app.createFile(file_list[i])
	}

	return nil
}

func (app *AppData) createFile(file_path string) error {

	//rename file
	temp := strings.Replace(file_path, app.TplPath, app.Appname, -1)
	temp = strings.Replace(temp, ".tpl", "", -1)
	tmpl, err := template.New("fileName").Parse(temp)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, app)
    fmt.Println()

	dir, _ := filepath.Split(temp)

	//create dir
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	//create file
	fout, err := os.Create(temp)
	defer fout.Close()
	if err != nil {
		return err
	}

	//build template
	content := make([]byte, 10240)
	fin, err := os.Open(file_path)
	defer fin.Close()
	if err != nil {
		return err
	}
	_, _ = fin.Read(content)
	_, _ = fout.Write(content)

	return nil
}
