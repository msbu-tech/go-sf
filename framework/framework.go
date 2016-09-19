package framework

import (
    "fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type StTplData struct {
	Appname string
	Author  string
	Date    string
}

func NewApp(app_name string, tpl_name string) error {
	//walk file list
	file_list, err := walkFileList(tpl_name)
	if err != nil {
		//log err
		return err
	}

	//create file
	err = createAppFiles(file_list, app_name, tpl_name)
	if err != nil {
		return err
	}

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

func createAppFiles(file_list []string, app_name string, tpl_name string) error {
	//remove if exist
	err := os.RemoveAll(app_name)
	if err != nil {
		return err
	}

	//create file
	for i := 0; i < len(file_list); i++ {
		createFile(file_list[i], app_name, tpl_name)
	}

	return nil
}

func createFile(file_path string, app_name string, tpl_name string) error {

	app_tpl_data := StTplData{"demo", "cuishichao", "2016-09-17"}

	//rename file
	temp := strings.Replace(file_path, tpl_name, app_name, -1)
	temp = strings.Replace(temp, ".tpl", "", -1)
	tmpl, err := template.New("fileName").Parse(temp)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, app_tpl_data)
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
