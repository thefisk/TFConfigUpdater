package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

// Define a struct to store config in, must contain all keys in file
type Config struct {
	OrgName string            `hcl:"org_name"`
	MyMap   map[string]string `hcl:"my_map"`
	MyList  []int             `hcl:"mylist"`
}

func decode(f string) {
	var config Config

	err := hclsimple.DecodeFile(f, nil, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}

	log.Printf("Configuration is %#v", config)
}

func main() {

	dir, err := os.MkdirTemp("", "clone-home")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir)

	// git.PlainClone to download a repo
	_, err2 := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      "https://github.com/thefisk/azure-in-terraform",
		Progress: os.Stdout,
	})

	if err2 != nil {
		fmt.Println(err2)
	}

	file := filepath.Join(dir, "sample.hcl")
	fmt.Printf("file: %v\n", file)

	decode(file)

}
