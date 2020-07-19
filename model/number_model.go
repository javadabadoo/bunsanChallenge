package model

import (
	"fmt"
	"reflect"
	"strings"
)



var (
  // SPC Caracter vacio
  SPC = " "
  // PIP Caracter pipe
  PIP = "|"
  // ___ Caracter guion bajo
  ___ = "_"

  // AllSupportedNumbers Numeros soportados para mapeo
  AllSupportedNumbers = map[byte][]string {
    0: {
      SPC + ___ + SPC,
      PIP + SPC + PIP,
      PIP + ___ + PIP,
    },
    1: {
      SPC + SPC + SPC,
      SPC + SPC + PIP,
      SPC + SPC + PIP,
    },
    2: {
      SPC + ___ + SPC,
      SPC + ___ + PIP,
      PIP + ___ + SPC,
    },
    3: {
      SPC + ___ + SPC,
      SPC + ___ + PIP,
      SPC + ___ + PIP,
    },
    4: {
      SPC + SPC + SPC,
      PIP + ___ + PIP,
      SPC + SPC + PIP,
    },
    5: {
      SPC + ___ + SPC,
      PIP + ___ + SPC,
      SPC + ___ + PIP,
    },
    6: {
      SPC + ___ + SPC,
      PIP + ___ + SPC,
      PIP + ___ + PIP,
    },
    7: {
      SPC + ___ + SPC,
      SPC + SPC + PIP,
      SPC + SPC + PIP,
    },
    8: {
      SPC + ___ + SPC,
      PIP + ___ + PIP,
      PIP + ___ + PIP,
    },
    9: {
      SPC + ___ + SPC,
      PIP + ___ + PIP,
      SPC + ___ + PIP,
    },
  }
)



// Number Estructura para el procesamiento de lineas de texto como numeros
type Number struct {
  TextNumber []string
  Status string
}



// GetByteValue Convierte el valor de TextNumber a byte
func (n *Number) GetByteValue() (byteValue byte, err error) {
  byteValue = 99
  for numInt, numText := range AllSupportedNumbers {
    if reflect.DeepEqual(numText, n.TextNumber) {
      return numInt, nil
    }
  }
  return byteValue, fmt.Errorf("Tsssss, no encontramos coincidencia de digito para el texto:\n%s", n.ToString())
}



// ToString Texto imprimible
func (n *Number) ToString() (string) {
  return strings.Join(n.TextNumber, "\n")
}


// ValidateText Verifica algunas relgas que deben de cumplirse para el procesamiento
func (n *Number) ValidateText() (err error) {
  if len(n.TextNumber) != 3 {
    err = fmt.Errorf("No manches, no voy a poder interpretar los numeros, detecté %d lineas", len(n.TextNumber))
  } else if len(n.TextNumber[0]) != len(n.TextNumber[1]) || len(n.TextNumber[0]) != len(n.TextNumber[2]) {
    err = fmt.Errorf(
      "Checale la longitud de las lineas, no checa. L1: %d chars,  L2: %d chars,  L3: %d chars",
      len(n.TextNumber[0]),
      len(n.TextNumber[1]),
      len(n.TextNumber[2]),
    )
  } else if len(n.TextNumber[0]) % 3 != 0 {
    err = fmt.Errorf("La longitud esperada de las lineas no es valida, deberia ser multiplo de 3")
  }
  return err
}
