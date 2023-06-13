package main

import (
  "fmt"
  "os"
  "path/filepath"
  "strings"
)

// mv lower underscore
// Rename all files or dirs, converting spaces to underscores, and uppercase to lowercase.

func main() {
  // Create k:v collection, with just the ks. The vs will implicitly be turned into empty structs.
  files := make(map[string]struct{})

  // for file in os.Args ...
  for _, file := range os.Args[1:] {

    abspath, _ := filepath.Abs(file)

    if _, exists := files[abspath]; exists {
      continue
    }

    // An empty struct, spelled struct{}{}, is like the empty set () in Haskell.
    files[abspath] = struct{}{}

    base := normalize(filepath.Base(file))
    dir := filepath.Dir(file)

    fileinfo, err := os.Stat(file)
    if err != nil {
      // file -> string -> (num_bytes, error)
      fmt.Fprintf(os.Stderr, "Warning: Cannot stat file '%s': %s\n", file, err.Error())
      continue
    }

    if fileinfo.Mode().IsDir() || fileinfo.Mode().IsRegular() {
      err := os.Rename(file, filepath.Join(dir, base))
      if err != nil {
        fmt.Fprintf(os.Stderr, "Warning: File not renamed '%s': %s\n", file, err.Error())
      }
    }

  }
}

func normalize(name string) string {
  return strings.ToLower(strings.Replace(name, " ", "_", -1))
}
