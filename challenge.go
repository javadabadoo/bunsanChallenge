package main

import (
	"fmt"
	"os"

	"github.com/javadabadoo/bnsn_chllng_01/util"
)



func main() {
  inputFilePath := os.Args[1]
  outputFilePath := os.Args[2]
  fmt.Printf("Leyendo archivo de `%s` y guardando en '%s'\n", inputFilePath, outputFilePath)

  numbers, errors := util.GetNumbersFromFile(inputFilePath)
  for errorIndex, err := range errors {
    fmt.Printf("Error [%d]: %v", errorIndex, err)
  }

  output := util.CreateOutput(numbers)
  util.WriteToFile(outputFilePath, output)
}
