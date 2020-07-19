package test

import (
	"testing"

	"github.com/javadabadoo/bnsn_chllng_01/util"
)

const (
  OkAccountsFilepath = "./resources/ok_accounts.txt"
  FailAccountsFilepath = "./resources/fail_accounts.txt"
)

func TestOkReadFile(t *testing.T) {
  accountNumbers, errors := util.GetNumbersFromFile(OkAccountsFilepath)
  for _, err := range errors {
    t.Errorf("Error al procesar el archivo de cuentas. %v", err)
  }
  if len(accountNumbers) == 0 {
    t.Error("No se encontraron numeros de cuenta en el archivo")
  }
}

func TestFailReadFile(t *testing.T) {
  accountNumbers, errors := util.GetNumbersFromFile(FailAccountsFilepath)
  if len(accountNumbers) != 0 {
    t.Errorf("Expected 0 numbers, Got %d (and %d Errors)", len(accountNumbers), len(errors))
  }
}
