package main

import (
    "flag"
    "github.com/geisonn/robots/functions"
)

func main(){
    //help
    help := flag.Bool("h", false, "Help")
    flag.BoolVar(help, "help", false, "Help")
    //url
    url := flag.String("u", "", "Target URL")
    flag.StringVar(url, "url", "", "Target URL")
    //list
    list := flag.String("l", "", "URLs list")
    flag.StringVar(list, "list", "", "URLs list")
    //Hide fails
    hideFails := flag.Bool("hf", false, "Hide fails")
    flag.BoolVar(hideFails, "hidefails", false, "Hide fails")

    flag.Parse()

    if *help {
		functions.Help()
	}

    if *hideFails {
        functions.Hidefails()
    }

    if *url != "" {
        functions.UrlFlag(*url)
    }

    if *list != "" {
        functions.ReadLines(*list)
    }
}
