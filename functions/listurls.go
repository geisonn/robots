package functions

import (
    "bufio"
    "os"
    "fmt"
    "net/http"
    "io"
    "strings"
    
)

var endpoints []string = []string{
    "/robots.txt",
}

//read urls from a file
func ReadLines(path string)([]string, error){
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        lines = append(lines, scanner.Text())
    }
    for _, line := range lines {
        for _, endpoint := range endpoints {
            fullUrl := line + endpoint
            if !strings.HasPrefix(fullUrl, "http://") && !strings.HasPrefix(fullUrl, "https://"){
                request := "http://" + fullUrl
                resp, err := http.Get(request)
                if err != nil {
                    if !hideFailsEnabled{
                        fmt.Printf("Error making request: Site inactive (%s)\n", fullUrl)
                    }
                    continue
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
                resp, err := http.Get(fullUrl)
                if err != nil {
                    if !hideFailsEnabled{
                        fmt.Printf("Error making request: Site inactive(%s)\n", fullUrl)
                    }
                    continue
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
    return lines, nil
}