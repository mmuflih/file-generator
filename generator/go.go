package generator

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const header string = `
/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/
 `

type fileGo struct {
	args []string
}

type FileGo interface {
	Generate()
	getTemplate() []string
	getStructTemplate() []string
	generateStruct() string
	generateService() string
	generateGo() string
	generate(string) (error, *goHelper)
}

func NewGo(args []string) FileGo {
	return &fileGo{args}
}

func (pg fileGo) getTemplate() []string {
	templ := fmt.Sprintf(`package DummyPackage

%s

`, header)
	return strings.Split(templ, "\n")
}

func (pg fileGo) getStructTemplate() []string {
	templ := fmt.Sprintf(`package DummyPackage

%s

type DummyStruct struct {

}

func NewDummyStruct() *DummyStruct {
	return &DummyStruct{

	}
}

func (m DummyStruct) GenerateResponse() interface{} {
	return struct {

	}{
		
	}
}

func (m DummyStruct) GeneratePublicResponse() interface{} {
	return struct {

	}{
		
	}
}

`, header)
	return strings.Split(templ, "\n")
}

func (pg fileGo) getServiceTemplate() []string {
	templ := fmt.Sprintf(`package DummyPackage

%s

type DummyRepositoryRepository interface {

}

type dummyStructRepo struct {

}

func NewDummyRepositoryRepo() DummyRepositoryRepository {
	return &dummyStructRepo{}
}

`, header)
	return strings.Split(templ, "\n")
}

func (pg fileGo) generateHandlerTmp() []string {
	templ := fmt.Sprintf(`package DummyPackage

%s

type DummyHandlerHandler interface {
	Handle() (interface{}, error)
}

type dummyHandlerH struct {

}

func NewDummyHandlerHandler() DummyHandlerHandler {
	return &dummyHandlerH{}
}

func (h dummyHandlerH) Handle() (interface{}, error) {

}

`, header)
	return strings.Split(templ, "\n")
}

func (pg fileGo) generateReaderTmp() []string {
	templ := fmt.Sprintf(`package DummyPackage

%s

type DummyReaderReader interface {
	Read() (interface{}, error)
}

type dummyReaderR struct {

}

func NewDummyReaderReader() DummyReaderReader {
	return &dummyReaderR{}
}

func (h dummyReaderR) Read() (interface{}, error) {
	
}

`, header)
	return strings.Split(templ, "\n")
}

func (pg fileGo) Generate() {
	switch pg.args[1] {
	case "struct":
		fmt.Println(pg.generateStruct())
		break
	case "service":
		fmt.Println(pg.generateService())
		break
	case "handler":
		fmt.Println(pg.generateHandler())
		break
	case "reader":
		fmt.Println(pg.generateReader())
		break
	default:
		fmt.Println(pg.generateGo())
	}
}

func (pg fileGo) generate(tipe string) (error, *goHelper) {
	path := pg.args[1]
	lines := pg.getTemplate()
	if tipe == "struct" {
		path = pg.args[2]
		lines = pg.getStructTemplate()
	} else if tipe == "service" {
		path = pg.args[2]
		lines = pg.getServiceTemplate()
	} else if tipe == "handler" {
		path = pg.args[2]
		lines = pg.generateHandlerTmp()
	} else if tipe == "reader" {
		path = pg.args[2]
		lines = pg.generateReaderTmp()
	}
	items := strings.Split(path, "/")
	className := items[len(items)-1]
	fileName := strings.ToLower(className) + ".go"
	filePath := ""
	pathLength := len(path)
	classNameLength := len(className)
	fmt.Println(pathLength, classNameLength)
	if pathLength > classNameLength {
		filePath = path[0 : pathLength-classNameLength-1]
		filePath += "/"
	}
	if tipe == "service" {
		fileName = strings.ToLower(className) + "_service.go"
	} else if tipe == "handler" {
		fileName = strings.ToLower(className) + "_handler.go"
	} else if tipe == "reader" {
		fileName = strings.ToLower(className) + "_reader.go"
	}
	destinationPath := filePath + fileName
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

	var newLines []string
	packageName := "main"
	fmt.Println(len(items))
	if len(items) > 1 {
		packageName = items[len(items)-2]
	}
	for _, line := range lines {
		strData := strings.Replace(line, "DummyPackage", strings.ReplaceAll(strings.ToLower(packageName), "/", "\\"), -1)
		if tipe == "struct" {
			strData = strings.Replace(strData, "DummyStruct", className, -1)
		} else if tipe == "service" {
			strData = strings.Replace(strData, "dummyStruct", makeFirstLowerCase(className), -1)
			strData = strings.Replace(strData, "DummyRepository", className, -1)
		} else if tipe == "handler" {
			strData = strings.Replace(strData, "dummyHandler", makeFirstLowerCase(className), -1)
			strData = strings.Replace(strData, "DummyHandler", className, -1)
		} else if tipe == "reader" {
			strData = strings.Replace(strData, "dummyReader", makeFirstLowerCase(className), -1)
			strData = strings.Replace(strData, "DummyReader", className, -1)
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

func (pg fileGo) generateService() string {
	err, goH := pg.generate("service")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	ioutil.WriteFile(goH.DestinationPath, goH.Output, 0644)
	return "Generated " + goH.ClassName + " => " + goH.DestinationPath
}

func (pg fileGo) generateHandler() string {
	err, goH := pg.generate("handler")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	ioutil.WriteFile(goH.DestinationPath, goH.Output, 0644)
	return "Generated " + goH.ClassName + " => " + goH.DestinationPath
}

func (pg fileGo) generateReader() string {
	err, goH := pg.generate("reader")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	ioutil.WriteFile(goH.DestinationPath, goH.Output, 0644)
	return "Generated " + goH.ClassName + " => " + goH.DestinationPath
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

func makeFirstLowerCase(s string) string {

	if len(s) < 2 {
		return strings.ToLower(s)
	}

	bts := []byte(s)

	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}
