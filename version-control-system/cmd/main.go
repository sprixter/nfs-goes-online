package main

import (
	"log"
	"os"
	"path/filepath"
)

func gitinit() {
	//Create a .git directory in the working directory
	gitdir := filepath.Join(".", ".git")
	err := os.MkdirAll(gitdir, 0777)
	if err != nil {
		log.Fatalf("Error creating .gitx directory : %v", err)
	}

	//Create objects directory in the .git directory
	gitobjects := filepath.Join(gitdir, "objects")
	err = os.MkdirAll(gitobjects, 0777)
	if err != nil {
		log.Fatalf("Error creating objects directory : %v", err)
	}

	//Create a HEAD file
	githead := filepath.Join(gitdir, "HEAD.txt")
	githeadfile, err := os.OpenFile(githead, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer githeadfile.Close()

}

func gitcommit() {

}
