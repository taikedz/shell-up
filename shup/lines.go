package shup
 
import (
    "bufio"
    "os"
)
 
func ReadLines(path string) ([]string, error) {
    // https://golangdocs.com/golang-read-file-line-by-line
    readFile, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    defer readFile.Close()
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)
    var lines []string

    for fileScanner.Scan() {
        lines = append(lines, fileScanner.Text())
    }

    return lines, nil
}
