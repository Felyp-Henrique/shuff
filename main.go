package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	move  *bool    = flag.Bool("m", false, "Do renamed(move) files after shuffle")
	files []string = nil
)

func main() {
	flag.Parse()

	if len(flag.Args()) <= 0 {
		flag.Usage()
		os.Exit(0)
	}

	files = flag.Args()

	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(files), func(i int, j int) {
		files[i], files[j] = files[j], files[i]
	})

	for index, file := range files {
		basename := filepath.Base(file)
		newfile := strings.Replace(file, basename, fmt.Sprintf("%d_%s", index, basename), -1)

		if *move {
			os.Rename(file, newfile)
		}

		fmt.Println(newfile)
	}
}
