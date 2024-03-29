package imagedownloader

import (
    "fmt"
    "io"
    "net/http"
    "net/url"
    "os"
    "path/filepath"
    "strings"
)

var (
    fileName    string
    fullUrlFile string
)

func main() {

    fullUrlFile = "http://www.golangprograms.com/skin/frontend/base/default/logo.png"
    fmt.Println(testfunc())
    // Build fileName from fullPath
    buildFileName()

    // Create blank file
    file := createFile()

    // Put content on file
    putFile(file, httpClient())

}


func Test() string {
	

	return "dummy"

}

func putFile(file *os.File, client *http.Client) {
    resp, err := client.Get(fullUrlFile)
    path, _ := filepath.Abs(filepath.Dir(file.Name()))
    fmt.Println(path + "/" + file.Name())

    checkError(err)

    defer resp.Body.Close()

    size, err := io.Copy(file, resp.Body)

    defer file.Close()

    checkError(err)

    fmt.Println("Just Downloaded a file %s with size %d", fileName, size)
}

func buildFileName() {
    fileUrl, err := url.Parse(fullUrlFile)
    checkError(err)

    path := fileUrl.Path
   
    segments := strings.Split(path, "/")

    fileName = segments[len(segments)-1]
}

func httpClient() *http.Client {
    client := http.Client{
        CheckRedirect: func(r *http.Request, via []*http.Request) error {
            r.URL.Opaque = r.URL.Path
            return nil
        },
    }

    return &client
}

func createFile() *os.File {
    file, err := os.Create(fileName)
   
    checkError(err)
    return file
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
