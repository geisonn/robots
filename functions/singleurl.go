package functions

import (
    "fmt"
    "net/http"
    "io"
    "strings"
)

var endpointsSingle []string = []string{
    "/robots.txt",
}

//Single URL request
func UrlFlag(url string){

    HttpPrefix(url)

}

func HttpPrefix(url string) {
    for _, endpoint := range endpointsSingle {
        fullUrl := url + endpoint
        if !strings.HasPrefix(fullUrl, "http://") && !strings.HasPrefix(fullUrl, "https://"){
            request := "http://" + fullUrl
            resp, err := http.Get(request)
            if err != nil {
                fmt.Println("Error making request:", err)
            }

            defer resp.Body.Close()

            body, err := io.ReadAll(resp.Body)
            if err != nil {
                fmt.Println("Error to read the body:", err)
            }
            division := strings.Split(string(body), "\n")
            for _, divisions := range division{
                trimmed := strings.TrimSpace(divisions)
                if strings.HasPrefix(trimmed, "Disallow:"){
                    dir := strings.TrimSpace(strings.TrimPrefix(trimmed, "Disallow:"))
                    fmt.Println(dir)
                }
                if strings.HasPrefix(trimmed, "Allow:"){
                    dir := strings.TrimSpace(strings.TrimPrefix(trimmed, "Allow:"))
                    fmt.Println(dir)
                }
            }
        } else {
            request := fullUrl
            resp, err := http.Get(request)
            if err != nil {
                fmt.Println("Error making request:", err)
            }

            defer resp.Body.Close()

            body, err := io.ReadAll(resp.Body)
            if err != nil {
                fmt.Println("Error to read the body:", err)
            }
            division := strings.Split(string(body), "\n")
            for _, divisions := range division{
                trimmed := strings.TrimSpace(divisions)
                if strings.HasPrefix(trimmed, "Disallow:"){
                    dir := strings.TrimSpace(strings.TrimPrefix(trimmed, "Disallow:"))
                    fmt.Println(dir)
                }
                if strings.HasPrefix(trimmed, "Allow:"){
                    dir := strings.TrimSpace(strings.TrimPrefix(trimmed, "Allow:"))
                    fmt.Println(dir)
                }
            }

        }
    }
}