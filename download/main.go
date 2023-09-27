package main
 
import (
    "fmt"
    "io"
    "log"
    "net/http"
    "net/url"
    "os"
    "strings"
)
 
var (
    fileName    string
    fullURLFile string
)
 
func main() {
 
    fullURLFile = "https://assets.publishing.service.gov.uk/government/uploads/system/uploads/attachment_data/file/1187237/2023-09-26_-_Worker_and_Temporary_Worker.csv"
 
    // Build fileName from fullPath
    fileURL, err := url.Parse(fullURLFile)
    if err != nil {
        log.Fatal(err)
    }
    path := fileURL.Path
    segments := strings.Split(path, "/")
    fileName = segments[len(segments)-1]
 
    // Create blank file
    file, err := os.Create(fileName)
    if err != nil {
        log.Fatal(err)
    }
    client := http.Client{
        CheckRedirect: func(r *http.Request, via []*http.Request) error {
            r.URL.Opaque = r.URL.Path
            return nil
        },
    }
    // Put content on file
    resp, err := client.Get(fullURLFile)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
 
    size, err := io.Copy(file, resp.Body)
 
    defer file.Close()
 
    fmt.Printf("Downloaded a file %s with size %d", fileName, size)
 
}
