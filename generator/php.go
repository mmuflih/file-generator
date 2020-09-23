package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type phpGo struct {
	args []string
}

type PhpGo interface {
	Generate()
	getTemplate() []string
	generateClass() string
}

func NewPhpGo(args []string) PhpGo {
	return &phpGo{args}
}

func (pg phpGo) getTemplate() []string {
	templ := `<?php

/**
 * Created by Muhammad Muflih Kholidin
 * at NOW
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

namespace DummyNamespace;

class DummyClass
{
	public function __construct()
	{
	}
}
`
	return strings.Split(templ, "\n")
}

func (pg phpGo) Generate() {
	switch pg.args[1] {
	case "class":
		fmt.Println(pg.generateClass())
	}
}

func (pg phpGo) generateClass() string {
	path := pg.args[2]
	items := strings.Split(path, "/")
	className := items[len(items)-1]
	fileName := className + ".php"
	namespace := path[0 : len(path)-len(className)-1]
	destinationPath := namespace + "/" + fileName
	/** check existing php class */
	if _, err := os.Stat(destinationPath); err == nil {
		return "File already exist"
	}

	/** create folder */
	_, err := os.Stat(namespace)
	if os.IsNotExist(err) {
		createFolder(namespace)
	}

	destination, err := os.Create(destinationPath)
	if err != nil {
		fmt.Println(err, "error read destination")
		return err.Error()
	}

	defer destination.Close()

	lines := pg.getTemplate()
	var newLines []string
	for _, line := range lines {
		strData := strings.Replace(line, "DummyNamespace", strings.ReplaceAll(strings.Title(namespace), "/", "\\"), -1)
		strData = strings.Replace(strData, "DummyClass", className, -1)
		strData = strings.Replace(strData, "NOW", time.Now().Format("2006-01-02 15:04:05"), -1)
		newLines = append(newLines, strData)
	}
	output := strings.Join(newLines, "\n")
	ioutil.WriteFile(destinationPath, []byte(output), 0644)
	return "Generated " + className + " => " + destinationPath
}

func createFolder(namespace string) {
	fmt.Println(namespace)
	err := os.MkdirAll(namespace, 0755)
	if err != nil {
		fmt.Println(err, "creating namespace folder")
	}
}
