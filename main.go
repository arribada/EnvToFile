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
	// TODO remove once balena supports dots in the env variables
	for _, pair := range os.Environ() {
		d := strings.SplitN(pair, "=", 2)
		os.Setenv(strings.ReplaceAll(strings.ReplaceAll(d[0], "_", "."), "..", "_"), d[1])
	}
	log.SetFlags(log.Ltime | log.Lshortfile)
	app := kingpin.New(filepath.Base(os.Args[0]), `
	A simple tool that reads all env variables and writes each one to a file.
	Useful in environments where the only way to add configuratios is through env variables and
	need to use docker images which accept configs only through files.
	`)
	app.HelpFlag.Short('h')
	dest := app.Flag("dir", "destination to which it should write the files").
		Short('d').Required().ExistingDir()
	kingpin.MustParse(app.Parse(os.Args[1:]))

	for _, env := range os.Environ() {
		d := strings.SplitN(env, "=", 2)
		destPath := filepath.Join(*dest, d[0])
		expanded := os.ExpandEnv(d[1])
		err := ioutil.WriteFile(destPath, []byte(expanded), 0644)
		if err != nil {
			log.Printf("could not write env var:%v to:%v \n", d[0], destPath)
		}
		log.Printf("wrote env var:%v to:%v \n", d[0], destPath)
		if os.Getenv("DEBUG") != "" {
			log.Printf("env content:%v\n", d[1])
		}
	}
	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
}
