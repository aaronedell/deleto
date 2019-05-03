package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	pb "gopkg.in/cheggaaa/pb.v1"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var (
		dir = flag.String("dir", ".", "source of files to delete")
		num = flag.Int("num", 123, "number of files to delete")
	)
	flag.Parse()
	files, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Fatal(err)
	}
	total := len(files)
	bar := pb.StartNew(*num)
	totaldeleted := 0
	for i := 0; i < *num; i++ {
		rando := randO(1, total)
		selectedfile := filepath.Join(*dir, files[rando].Name())
		if _, err := os.Stat(selectedfile); err == nil {
			try := os.Remove(selectedfile)
			if try != nil {
				log.Fatal(try)
			}
			bar.Increment()
			totaldeleted++
		} else if os.IsNotExist(err) {
			i--
		} else {
			log.Fatal(err)
		}
	}
	bar.Finish()
	log.Printf("Total files deleted: %d", totaldeleted)
}

func randO(min, max int) int {
	randomnumber := rand.Intn(max-min) + min
	return randomnumber
}
