package main

import (
	"fmt"
	"os"
	"log"
	"io"
)

func main(){

	//Did we have min. 1 Filename as arg?
	if len(os.Args) == 1{
		fmt.Println("Mindestens eine Datei als Parameter erwartet")
		os.Exit(1)
	}


	//iterate over all files
	for _, fname := range os.Args[1:]{

		//open file
		f,err := os.Open(fname)
		if err !=nil {
			log.Printf("Fehler beim Öffnen der Datei %s\n%s\n",fname,err)
			f.Close()
			continue
		}

		//write file to stdout
		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			log.Printf("Fehler bei der Ausgabe von %s\n%s\n",fname, err)
		}
		f.Close()

	}

}
