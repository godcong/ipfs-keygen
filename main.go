package main

import (
	"flag"
	"github.com/libp2p/go-libp2p-pnet"
	"io"
	"log"
	"os"
	"path/filepath"
)

var file = flag.String("file", "swarm.key", "set the key output path")

func main() {
	flag.Parse()

	_ = os.MkdirAll(filepath.Dir(*file), os.ModePerm)
	openFile, e := os.OpenFile(*file, os.O_CREATE|os.O_SYNC|os.O_RDWR, os.ModePerm)
	if e != nil {
		panic(e)
	}
	reader, e := pnet.GenerateV1PSK()
	if e != nil {
		panic(e)
	}
	i, e := io.Copy(openFile, reader)
	if e != nil {
		panic(e)
	}
	log.Println("success with bytes:", i)
}
