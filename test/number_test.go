package test

import (
	"fmt"
	"testing"

	"github.com/javadabadoo/bnsn_chllng_01/model"
)



type SingleNumberTestStructure struct {
  TextNumber []string
  ExpectedByteValue byte
}


var (
  OkNumbersList = []*SingleNumberTestStructure {
    {
      TextNumber: []string{
        " _ ",
        "| |",
        "|_|",
      },
      ExpectedByteValue: 0,
    },
    {
      TextNumber: []string{
        "   ",
        "  |",
        "  |",
      },
      ExpectedByteValue: 1,
    },
    {
      TextNumber: []string{
        " _ ",
        " _|",
        "|_ ",
      },
      ExpectedByteValue: 2,
    },
    {
      TextNumber: []string{
        " _ ",
        " _|",
        " _|",
      },
      ExpectedByteValue: 3,
    },
    {
      TextNumber: []string{
        "   ",
        "|_|",
        "  |",
      },
      ExpectedByteValue: 4,
    },
    {
      TextNumber: []string{
        " _ ",
        "|_ ",
        " _|",
      },
      ExpectedByteValue: 5,
    },
    {
      TextNumber: []string{
        " _ ",
        "|_ ",
        "|_|",
      },
      ExpectedByteValue: 6,
    },
    {
      TextNumber: []string{
        " _ ",
        "  |",
        "  |",
      },
      ExpectedByteValue: 7,
    },
    {
      TextNumber: []string{
        " _ ",
        "|_|",
        "|_|",
      },
      ExpectedByteValue: 8,
    },
    {
      TextNumber: []string{
        " _ ",
        "|_|",
        " _|",
      },
      ExpectedByteValue: 9,
    },
  }
  FailNumbersList = []*SingleNumberTestStructure {
    {
      TextNumber: []string{
        " _  ",
        "| |",
        "|_|",
      },
      ExpectedByteValue: 0,
    },
    {
      TextNumber: []string{
        "   ",
        "| |",
        "|_|",
      },
      ExpectedByteValue: 0,
    },
    {
      TextNumber: []string{
        " __",
        "| |",
        "|_|",
      },
      ExpectedByteValue: 0,
    },
  }
)



func TestOkNumbers(t *testing.T) {
  errors := testNumber(OkNumbersList)
  for _, err := range errors {
    t.Errorf("Valio queso: %v", err)
  }
}



func TestFailNumbers(t *testing.T) {
  errors := testNumber(FailNumbersList)
  if len(errors) != len(FailNumbersList) {
    t.Errorf(
      "Deberiamos tener el mismo numero de errores que de numeros procesados pero tenemos %d de %d",
      len(errors),
      len(FailDigitsLines))
  }
}



func testNumber(numbers []*SingleNumberTestStructure) (errors []error) {
  for _, numberTest := range numbers {
    number := &model.Number{
      TextNumber: numberTest.TextNumber,
    }
    if err := number.ValidateText(); err != nil {
      errors = append(errors, err)
      continue
    }
    if byteValue, err := number.GetByteValue(); err != nil {
      if byteValue != numberTest.ExpectedByteValue {
        errors = append(errors, fmt.Errorf("Expected: %d Got: %d For: \n'%s'", numberTest.ExpectedByteValue, byteValue, numberTest.TextNumber))
      }
    }
  }
  return errors
}
