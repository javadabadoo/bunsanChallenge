package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"

	"github.com/javadabadoo/bnsn_chllng_01/model"
)

// IsValidAccount Verifica el numero de cuenta mediante el calculo de su checksum
func IsValidAccount(digits []byte) (isValid bool) {
  checksum := 0
  for digitIndex := 0; digitIndex < len(digits); digitIndex++ {
    checksum += (int(digits[digitIndex]) + 1) * int(digits[digitIndex])
  }
  return checksum % 11 == 0
}



// GetNumbersFromFile Regresa los numeros de cuenta encapuslados en un slice de model.Numbers
// y una lista de los errores encontrados. Tomando como fuente el path de un archivo en Filesystem
func GetNumbersFromFile(accountsFilePath string) (accountNumbers []*model.Numbers, errors []error) {
  var fileContent []byte
  var err error
  if fileContent, err = ioutil.ReadFile(accountsFilePath); err == nil {
    accountNumbers, errors = GetNumbersFromString(string(fileContent))
  } else {
    errors = append(errors, err)
  }
  return accountNumbers, errors
}



// GetNumbersFromString Regresa los numeros de cuenta encapuslados en un slice de model.Numbers
// y una lista de los errores encontrados tomando como fuente una cadena de texto
func GetNumbersFromString(accountsString string) (accountNumbers []*model.Numbers, errors []error) {
  return parseString(accountsString)
}



// ParseString Convierte el Texto a numeros
func parseString(text string)(accountNumbers []*model.Numbers, errors []error) {
  accountNumbersArray := regexp.MustCompile(`[\r\n](?m)^$(\r|\n|\r\n)?`).Split(text, -1)

  // Iteración hasta len(accountNumbersArray)-1 para que ignore la ultima linea vacia
  for accountIndex := 0; accountIndex < len(accountNumbersArray) -1; accountIndex++ {
    accountNumber := &model.Numbers {
      TextNumbers: regexp.MustCompile("\n").Split(accountNumbersArray[accountIndex], -1),
    }
    if err := accountNumber.ValidateText(); err != nil {
      errors = append(errors, err)
      continue
    }
    accountNumbers = append(accountNumbers, accountNumber)
  }

  return accountNumbers, errors
}



func WriteToFile(outputFilepath, output string) (err error) {
  var f *os.File
  if f, err = os.Create(outputFilepath); err != nil {
    return err
  }
  if _, err = f.WriteString(output); err != nil {
    return err
  }
  return f.Close()
}



func CreateOutput(numbers []*model.Numbers) (output string) {
  for _, number := range numbers {
    byteValues := number.GetByteValues()
    stringNumber, status := BuildStringLine(byteValues)
    output += fmt.Sprintf("%s\t%s\n", stringNumber, status)
  }
  return output
}



func BuildStringLine(number []byte) (stringNumber, status string) {
  status = model.StatusOk
  for _, n := range number {
    if n == 99 {
      stringNumber += "?"
      status = model.StatusIll
    } else {
      stringNumber += strconv.Itoa(int(n))
    }
  }
  if status == model.StatusOk && !IsValidAccount(number) {
    status = model.StatusErr
  }
  return stringNumber, status
}
