package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher
var dir = "/home/miguelarch/Downloads"
var directories map[string]string

func init() {
	log.SetOutput(os.Stdout)
	home, err := os.UserHomeDir()
	buffer, err := ioutil.ReadFile(home + "/.config/fileOrganizer/config.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(buffer), &directories)
	log.Println(directories)
	dir = home + directories["default"]
	for _, value := range directories {
		dir := home + value
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.Mkdir(dir, os.ModePerm)
			log.Println(dir + " was created")
		}
	}
	moveAll(dir, directories)
}

func move(filename string) {
	home, _ := os.UserHomeDir()
	log.Println(home)
	log.Println(directories["directory"])
	var dest string
	var source string
	var file, err = os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(file.Name())
	if file.IsDir() {
		dest = home + directories["directory"] + "/" + file.Name()
		source = filename
	} else {
		source = dir + "/" + file.Name()
		var extension = filepath.Ext(file.Name())[1:]
		log.Println("Extension: " + extension + " was found")
		if val, ok := directories[extension]; ok {
			dest = home + val + "/" + file.Name()
		} else {
			dest = home + directories["other"] + "/" + file.Name()
		}
	}
	os.Rename(source, dest)
	log.Println("File: " + source + " moved to " + dest)
}

func moveAll(dir string, directories map[string]string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		move(dir + "/" + f.Name())
	}
}

func main() {
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()
	if err := filepath.Walk(dir, watchDir); err != nil {
		fmt.Println("ERROR", err)
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Printf("EVENT! %#v\n", event)
				fmt.Printf(event.Op.String())
				if event.Op.String() == "CREATE" {
					move(event.Name)
				}
				//move(event.Name)
				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	<-done
}

func watchDir(path string, fi os.FileInfo, err error) error {
	if fi.Mode().IsDir() {
		return watcher.Add(path)
	}

	return nil
}
