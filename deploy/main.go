package main

	
import (
	"os"
	"fmt"
	"version"
	
)
var (
	builtOn   string
	builtAt   string
	goVersion string
	gitAuthor string
	gitCommit string
  )
func main() {
    fmt.Println("hello world")
	//fmt.Fprintf(os.Stderr, "Git Commit: %s\n", gitCommit)
}