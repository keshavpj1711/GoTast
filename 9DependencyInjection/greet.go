package main

import (
	// "bytes"
	"fmt"
	"io"
	"os"
)

// Instead of *bytes.Buffer we use io.Writer type
// Reason:
// *bytes.Buffer implements io.Writer and Greet does not specify what kind of Writer it wants 
// basically any type which implement io.Writer works
// So In the testing scenario,
// you use a bytes.Buffer because it allows you to capture the output from Greet and verify its content.
// In the real application,
// you use os.Stdout (standard output) as the writer for Greet.

func Greet(writer io.Writer, name string)  {

  // This is what it looks like inside a Printf()
  fmt.Fprintf(writer, "Hello %s", name)
  
  // One thing to note is that this line stdout even if our test fail 
  // fmt.Printf("Hello %s", name)
}

func main()  {
  Greet(os.Stdout, "Keshav")
}
