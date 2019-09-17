package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	app := kingpin.New(filepath.Base(os.Args[0]), `
	A simple tool that reads all env variables and writes each one to a file.
	Useful in environments where the only way to add configuratios is through env variables and
	need to use docker images which accept configs only through files.
	`)
	app.HelpFlag.Short('h')
	dest := app.Flag("dir", "destination to which it should write the files").
		Short('d').ExistingDir()
	kingpin.MustParse(app.Parse(os.Args[1:]))

	for _, env := range os.Environ() {
		d := strings.Split(env, "=")
		destPath := filepath.Join(*dest, d[0])
		err := ioutil.WriteFile(destPath, []byte(d[1]), 0644)
		if err != nil {
			log.Printf("could not write env var:%v to:%v \n", d[0], destPath)
		}
		log.Printf("wrote env var:%v to:%v \n", d[0], destPath)
	}
	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
}
