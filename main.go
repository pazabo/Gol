package main

import (
  "bufio"
  "fmt"
  "os"
  "regexp"
)

func main() {
  args := os.Args[1:]

  fmt.Printf("Running with following arguments: %v\n", args)

  if len(args) == 0 {
    fmt.Println("No arguments specified!")
    os.Exit(1)
  }

  scanner := bufio.NewScanner(os.Stdin)
  writer := bufio.NewWriter(os.Stdout)
  defer writer.Flush()

  switch args[0] {
  case "--grep":
    regex, _ := regexp.Compile(args[1])
    for scanner.Scan() {
      line := scanner.Bytes()
      if regex.Match(line) {
        writer.Write(line)
        writer.WriteByte('\n')
      }
    }
  case "--sed":
    regex, _ := regexp.Compile(args[1])
    text := []byte(args[2])
    for scanner.Scan() {
      line := scanner.Bytes()
      writer.Write(regex.ReplaceAll(line, text))
      writer.WriteByte('\n')
    }
  default:
    fmt.Printf("Unknown option %v\n", args[0])
    os.Exit(1)
  }
}
