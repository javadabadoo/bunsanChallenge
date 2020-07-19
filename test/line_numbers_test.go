package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/javadabadoo/bnsn_chllng_01/model"
	"github.com/javadabadoo/bnsn_chllng_01/util"
)



type LineDigitsTestStructure struct {
  // TextNumbers Lineas de texto en las que se definen los digitos a procesar
  TextNumbers []string
  // ExpectedByteValues Valores esperados del procesamiento de los digitos
  ExpectedByteValues []byte
}

var (
  // OkDigitsLines Valores de prueba que deben ser procesados exitosamente
  OkDigitsLines = []*LineDigitsTestStructure {
     {
      TextNumbers: []string {
        "                           ",
        "  |  |  |  |  |  |  |  |  |",
        "  |  |  |  |  |  |  |  |  |",
      },
      ExpectedByteValues: []byte {1, 1, 1, 1, 1, 1, 1, 1, 1},
    },
    {
      TextNumbers: []string {
        " _  _  _  _  _  _  _  _  _ ",
        " _| _| _| _| _| _| _| _| _|",
        "|_ |_ |_ |_ |_ |_ |_ |_ |_ ",
      },
      ExpectedByteValues: []byte {2, 2, 2, 2, 2, 2, 2, 2, 2},
    },
    {
      TextNumbers: []string {
        " _  _  _  _  _  _  _  _  _ ",
        " _| _| _| _| _| _| _| _| _|",
        " _| _| _| _| _| _| _| _| _|",
      },
      ExpectedByteValues: []byte {3, 3, 3, 3, 3, 3, 3, 3, 3},
    },
    {
      TextNumbers: []string {
        "                           ",
        "|_||_||_||_||_||_||_||_||_|",
        "  |  |  |  |  |  |  |  |  |",
      },
      ExpectedByteValues: []byte {4, 4, 4, 4, 4, 4, 4, 4, 4},
    },
    {
      TextNumbers: []string {
        " _  _  _  _  _  _  _  _  _ ",
        "|_ |_ |_ |_ |_ |_ |_ |_ |_ ",
        " _| _| _| _| _| _| _| _| _|",
      },
      ExpectedByteValues: []byte {5, 5, 5, 5, 5, 5, 5, 5, 5},
    },
     {
      TextNumbers: []string {
        "    _  _     _  _  _  _  _ ",
        "  | _| _||_||_ |_   ||_||_|",
        "  ||_  _|  | _||_|  ||_| _|",
      },
      ExpectedByteValues: []byte {1, 2, 3, 4, 5, 6, 7, 8, 9},
    },
    {
      TextNumbers: []string {
        " _  _  _  _  _  _     _    ",
        "| | _||_||_ |_   ||_||_|  |",
        "|_| _||_||_| _|  |  | _|  |",
      },
      ExpectedByteValues: []byte {0, 3, 8, 6, 5, 7, 4, 9, 1},
    },
    {
      TextNumbers: []string {
        "    _  _     _  _  _  _  _ ",
        "  ||_|| ||_||_ |_| _| _||_ ",
        "  | _||_|  ||_||_||_  _| _|",
      },
      ExpectedByteValues: []byte {1, 9, 0, 4, 6, 8, 2, 3, 5},
    },
  }
  // FailDigitsLines Valores de prueba que deben ser procesados con error
  FailDigitsLines = []*LineDigitsTestStructure {
    {
      TextNumbers: []string {
        "       _  _     _  _  _  _  _ ",
        "| |  | _| _||_||  |_   ||_||_|",
        "|_|  ||_  _|  | _||_|  ||_| _|",
      },
      ExpectedByteValues: []byte {0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
    },
    {
      TextNumbers: []string {
        " _  _  _  _  _  _     _    ",
        " _||_||_   ||_   ||_||_|  ",
        " _||_||_|  | _|  |  | _|  |",
      },
      ExpectedByteValues: []byte {3, 8, 6, 7, 5, 7, 4, 9, 1},
    },
    {
      TextNumbers: []string {
        "|   _  _        _  _  _  _ ",
        "  ||_|| || ||_||_ |_| _| _|",
        "  | _||_||_|  ||_||_||_  _|",
      },
      ExpectedByteValues: []byte {1, 9, 0, 0, 4, 6, 8, 2, 3},
    },
  }
  OkChecksumValues = []*LineDigitsTestStructure {
    {
      TextNumbers: []string {
        "       _  _     _  _  _  _ ",
        "| |  | _| _||_||  |_   ||_|",
        "|_|  ||_  _|  | _||_|  ||_|",
      },
      ExpectedByteValues: []byte {0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
    },
    {
      TextNumbers: []string {
        " _  _  _  _  _  _     _    ",
        " _||_||_   ||_   ||_||_|   ",
        " _||_||_|  | _|  |  | _|  |",
      },
      ExpectedByteValues: []byte {3, 8, 6, 7, 5, 7, 4, 9, 1},
    },
    {
      TextNumbers: []string {
        "|   _  _        _  _  _  _ ",
        "  ||_|| || ||_||_ |_| _| _|",
        "  | _||_||_|  ||_||_||_  _|",
      },
      ExpectedByteValues: []byte {1, 9, 0, 0, 4, 6, 8, 2, 3},
    },
    {
     TextNumbers: []string {
       "                           ",
       "  |  |  |  |  |  |  |  |  |",
       "  |  |  |  |  |  |  |  |  |",
     },
    },
    {
     TextNumbers: []string {
       "    _  _     _  _  _  _  _ ",
       "  | _| _||_||_ |_   ||_||_|",
       "  ||_  _|  | _||_|  ||_| _|",
     },
     ExpectedByteValues: []byte {1, 2, 3, 4, 5, 6, 7, 8, 9},
   },
  }
)



func TestOkTextScan(t *testing.T) {
  errors := testLineNumbers(OkDigitsLines)
  for _, err := range errors {
    t.Errorf("Valio queso: %v", err)
  }
}



func TestFailTextScan(t *testing.T) {
  errors := testLineNumbers(FailDigitsLines)
  if len(errors) != len(FailDigitsLines) {
    t.Errorf(
      "Deberiamos tener el mismo numero de errores que de numeros procesados pero tenemos %d de %d",
      len(errors),
      len(FailDigitsLines))
  }
}



func TestCreateOutput(t *testing.T) {
  numbers := []*model.Numbers{}
  for _, number := range OkChecksumValues {
    numbers = append(
      numbers,
      &model.Numbers{
        TextNumbers: number.TextNumbers,
      },
    )
  }
  output := util.CreateOutput(numbers)
  if output == "" {
    t.Error("Expected: Output, Got: Empty")
  }
}



func testLineNumbers(lines []*LineDigitsTestStructure)(errors []error){
  for _, digits := range lines {
    numbers := model.Numbers {
      TextNumbers: digits.TextNumbers,
    }
    if err := numbers.ValidateText(); err != nil {
      errors = append(errors, err)
      continue
    }
    byteValues := numbers.GetByteValues()
    if !reflect.DeepEqual(digits.ExpectedByteValues, byteValues) {
      errors = append(errors, fmt.Errorf("Expected %v, Got: %v, For: %v", digits.ExpectedByteValues, byteValues, digits.TextNumbers))
    }
  }
  return errors
}
