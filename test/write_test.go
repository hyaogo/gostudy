package test

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestWrite(test *testing.T) {
	var b bytes.Buffer
	b.Write([]byte("hello"))
	fmt.Fprintf(&b, "world")
	b.WriteTo(os.Stdout)
}
