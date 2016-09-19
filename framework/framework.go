package framework

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type AppData struct {
	AppName    string
	Author     string
	Date       string
	TplPath    string
	OutputPath string
}

func New(appName string, author string, tplPath string, outputPath string) (error, AppData) {
	t := time.Now()
	app := AppData{appName, author, fmt.Sprintf("%d-%.2d-%.2d", t.Year(), t.Month(), t.Day()), tplPath, outputPath}

	return nil, app
}

func (app *AppData) Generate() error {
	fmt.Println("Processing...")

	//walk file list
	fileList, err := walkFileList(app.TplPath)
	if err != nil {
		//log err
		return err
	}

	//create file
	err = app.createAppFiles(fileList)
	if err != nil {
		return err
	}

	fmt.Println("Success.")

	return nil
}

func walkFileList(path string) ([]string, error) {
	var fileList = make([]string, 0)
	i := 0
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		fileList = append(fileList, path)
		i++
		return nil
	})
	return fileList, err
}

func (app *AppData) createAppFiles(fileList []string) error {
	//remove if exist
	err := os.RemoveAll(app.OutputPath + "/" + app.AppName)
	if err != nil {
		return err
	}

	//create file
	for i := 0; i < len(fileList); i++ {
		app.createFile(fileList[i])
	}

	return nil
}

func (app *AppData) createFile(filePath string) error {
	// rename file
	replacedFileName := strings.Replace(filePath, app.TplPath, app.OutputPath+"/"+app.AppName, -1)
	tplFileName := strings.Replace(replacedFileName, ".tpl", "", -1)
	fileNameTpl, err := template.New("file").Parse(tplFileName)
	if err != nil {
		return err
	}

	// create buffer
	buffer := new(bytes.Buffer)
	err = fileNameTpl.Execute(buffer, app)
	if err != nil {
		return err
	}

	realFileName := buffer.String()

	dir, _ := filepath.Split(realFileName)

	// create dir
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	// create file
	fout, err := os.Create(realFileName)
	defer fout.Close()
	if err != nil {
		return err
	}

	// build template
	// content := make([]byte, 10240)
	fin, err := os.Open(filePath)
	defer fin.Close()
	if err != nil {
		return err
	}
	// _, _ = fin.Read(content)
	contentTpl, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}
	buffer.Reset()
	err = contentTpl.Execute(buffer, app)
	// _, _ = fout.Write(buffer)
	buffer.WriteTo(fout)

	return nil
}
