package test

import (
	"testing"

	"github.com/javadabadoo/bnsn_chllng_01/util"
)



type ChecksumTestStructure struct {
  Number []byte
  ExpectedValue bool
}



var (
  ChecsumsValues = []*ChecksumTestStructure {
    {
      Number: []byte {1, 1, 1, 1, 1, 1, 1, 1, 1},
      ExpectedValue: false,
    },
    {
      Number: []byte {2, 2, 2, 2, 2, 2, 2, 2, 2},
      ExpectedValue: false,
    },
    {
      Number: []byte {3, 4, 5, 8, 8, 2, 8, 6, 5},
      ExpectedValue: false,
    },
    {
      Number: []byte {4, 5, 7, 5, 0, 8, 0, 0, 0},
      ExpectedValue: false,
    },
  }
)





func TestChecsums(t *testing.T) {
  for _, checksumCase := range ChecsumsValues {
    if util.IsValidAccount(checksumCase.Number) != checksumCase.ExpectedValue {
      t.Errorf("La validación del numero %v no cumple con el resultado esperado", checksumCase.Number)
    }
  }
}

