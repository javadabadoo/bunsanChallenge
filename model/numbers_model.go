package model

import (
	"fmt"
	"strings"
)



const (
  StatusUnredeable = "UNREADEABLE"
  StatusOk = "Ok"
  StatusErr = "ERR"
  StatusIll = "ILL"
)



// Numbers Estructura de los digitos de los numeros de cuenta
type Numbers struct {
  TextNumbers []string
  Status string
}



// GetNumberValues Retorna el slice que representa cada uno de los digitos encontrados en TextNumbers
func (n *Numbers) GetNumberValues() (numbers []*Number) {
  for digitIndex := 0; digitIndex < len(n.TextNumbers[0]) / 3; digitIndex++ {
    numbers = append(
      numbers,
      n.extractTextNumberAt(n.TextNumbers, digitIndex),
    )
  }
  return numbers
}



// GetByteValues Regresa el slice que representa los digitos procesados
func (n *Numbers) GetByteValues() (byteValues []byte) {
  numbers := n.GetNumberValues()
  for _, number := range numbers {
    if value, e := number.GetByteValue(); e == nil {
      byteValues = append(byteValues, value)
    } else {
      byteValues = append(byteValues, 99)
    }
  }
  return byteValues
}



// ValidateText Verifica las condiciones para que se puedan procesar los digitos
func (n *Numbers) ValidateText() (err error) {
  if len(n.TextNumbers) != 3 {
    err = fmt.Errorf("No manches, no voy a poder interpretar los numeros, detecté %d lineas\n%s", len(n.TextNumbers), n.ToString())
  } else if len(n.TextNumbers[0]) != len(n.TextNumbers[1]) || len(n.TextNumbers[0]) != len(n.TextNumbers[2]) {
    err = fmt.Errorf(
      "Checale la longitud de las lineas, no checa. L1: %d chars,  L2: %d chars,  L3: %d chars\n%s",
      len(n.TextNumbers[0]),
      len(n.TextNumbers[1]),
      len(n.TextNumbers[2]),
      n.ToString())
  } else if len(n.TextNumbers[0]) != 27 {
    err = fmt.Errorf("La longitud esperada de las lineas no es valida, deberia ser 27\n%s", n.ToString())
  }
  if err != nil {
    n.Status = StatusUnredeable
  }
  return err
}



// ToString Texto imprimible
func (n *Numbers) ToString() (string) {
  return strings.Join(n.TextNumbers, "\n")
}



func (n *Numbers) extractTextNumberAt(text []string, at int) (number *Number) {
  substringFrom := at * 3
  return &Number {
    TextNumber: []string {
      text[0][substringFrom:substringFrom+3],
      text[1][substringFrom:substringFrom+3],
      text[2][substringFrom:substringFrom+3],
    },
  }
}
