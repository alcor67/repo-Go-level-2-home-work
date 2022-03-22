package main

import (
	"fmt"
	"os"

	"github.com/alcor67/repo-Go-level-2-home-work/configuration"
)

func main() {

	conf, errConf := configuration.Load()

	if errConf != nil {
		fmt.Printf("произошла ошибка: %s \n", errConf)
		os.Exit(1)
	}

	if conf.MyDubFile != "" {
		//fmt.Println("===================================")
		fmt.Printf("список дубликатов файлов:\n")
		fmt.Fprintf(os.Stdout, conf.MyDubFile)
	}
}
