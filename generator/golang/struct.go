package golang

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type GoStruct struct {
	header string
}

func (pg GoStruct) generateStruct() string {
	err, goH := pg.generate("struct")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	ioutil.WriteFile(goH.DestinationPath, goH.Output, 0644)
	return "Generated " + goH.ClassName + " => " + goH.DestinationPath
}

func (pg GoStruct) GetStructTemplate() []string {
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

`, pg.header)
	return strings.Split(templ, "\n")
}
