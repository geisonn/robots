package functions

import (
	"fmt"
)

func Help(){
		fmt.Println("Usage:")
		fmt.Println("  -h, -help           Show help")
		fmt.Println("  -hf, -hidefails     Hide failed attempts")
		fmt.Println("  -l, -list <file>    List of URLs to process")
		fmt.Println("  -u, -url <URL>      Target URL to test")
	}