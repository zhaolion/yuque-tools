package pretty

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type Dummy struct {
	A1 string `json:"a1"`
	A2 int    `json:"a2"`
}

func TestPrintStruct(t *testing.T) {
	got := Struct(&Dummy{A1: "A1", A2: 2})
	require.NotEmpty(t, got)
	fmt.Println(got)
}
