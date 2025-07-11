// cmd/blockchainnftvaultmanagerdevpro/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftvaultmanagerdevpro/internal/blockchainnftvaultmanagerdevpro"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftvaultmanagerdevpro.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
