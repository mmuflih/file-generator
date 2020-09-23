package generator

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type fileGo struct {
	args []string
}

type FileGo interface {
	Generate()
	getTemplate() []string
	getStructTemplate() []string
	generateStruct() string
	generateGo() string
	generate(string) (error, *goHelper)
}

func NewGo(args []string) FileGo {
	return &fileGo{args}
}

func (pg fileGo) getTemplate() []string {
	templ := `package DummyPackage

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

`
	return strings.Split(templ, "\n")
}

func (pg fileGo) getStructTemplate() []string {
	templ := `package DummyPackage

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

 type DummyStruct struct {

 }
`
	return strings.Split(templ, "\n")
}

func (pg fileGo) Generate() {
	switch pg.args[1] {
	case "struct":
		fmt.Println(pg.generateStruct())
		break
	default:
		fmt.Println(pg.generateGo())
	}
}

func (pg fileGo) generate(tipe string) (error, *goHelper) {
	path := pg.args[1]
	items := strings.Split(path, "/")
	className := items[len(items)-1]
	fileName := strings.ToLower(className) + ".go"
	filePath := path[0 : len(path)-len(className)-1]
	destinationPath := filePath + "/" + fileName
	/** check existing php class */
	if _, err := os.Stat(destinationPath); err == nil {
		return errors.New("File already exist"), nil
	}

	/** create folder */
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		createFolder(filePath)
	}

	destination, err := os.Create(destinationPath)
	if err != nil {
		fmt.Println(err, "error read destination")
		return err, nil
	}

	defer destination.Close()

	lines := pg.getTemplate()
	if tipe == "struct" {
		lines = pg.getStructTemplate()
	}
	var newLines []string
	var packageName = items[len(items)-2]
	for _, line := range lines {
		strData := strings.Replace(line, "DummyPackage", strings.ReplaceAll(strings.ToLower(packageName), "/", "\\"), -1)
		if tipe == "struct" {
			strData = strings.Replace(strData, "DummyStruct", className, -1)
		}
		newLines = append(newLines, strData)
	}
	output := strings.Join(newLines, "\n")
	return nil, &goHelper{
		DestinationPath: destinationPath,
		ClassName:       className,
		Output:          []byte(output),
	}
}

func (pg fileGo) generateStruct() string {
	err, goH := pg.generate("struct")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	ioutil.WriteFile(goH.DestinationPath, goH.Output, 0644)
	return "Generated " + goH.ClassName + " => " + goH.DestinationPath
}

func (pg fileGo) generateGo() string {
	err, goH := pg.generate("")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	ioutil.WriteFile(goH.DestinationPath, goH.Output, 0644)
	return "Generated " + goH.ClassName + " => " + goH.DestinationPath
}

type goHelper struct {
	DestinationPath string
	ClassName       string
	Output          []byte
}
