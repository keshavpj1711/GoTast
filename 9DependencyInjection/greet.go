package main

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string)  {

  // This is what it looks like inside a Printf()
  fmt.Fprintf(writer, "Hello %s", name)
  
  // One thing to note is that this line stdout even if our test fail 
  // fmt.Printf("Hello %s", name)
}
