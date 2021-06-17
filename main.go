package main

import (
	"flag"
	"fmt"
	"go_level_2/duplicate"
	"log"
	"os"
)

var (
	path   *string
	remove *bool
)

func main() {

	remove = flag.Bool("r", false, "remove duplicate")
	path = flag.String("p", "C:\\", "directory path")

	flag.Parse()

	Run(*path, *remove)
}

func Run(path string, remove bool) {
	duplicates, err := duplicate.GetDuplicateFile(path)

	if err != nil {
		log.Println(err)
	}

	if len(duplicates) == 0 {
		fmt.Println("no duplicates found")
		return
	}

	fmt.Println("Duplicates:")
	for i, item := range duplicates {
		fmt.Printf("%d.  %s", i+1, item)
		fmt.Println()
	}

	if remove {
		fmt.Print("remove duplicates? ", "confirm command: (y/n)  ")
		response := ""
		fmt.Fscan(os.Stdin, &response)

		if response == "y" || response == "Y" {
			err := duplicate.RemoveDuplicate(duplicates)

			if err != nil {
				log.Println(err)
				return
			}

			fmt.Println("Done!")
		}
	}
}
