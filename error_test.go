package err

import (
	"fmt"
        "testing"
)

func TestError (t *testing.T) {
        e1 := Create ("A")
        e2 := CreateModel2 ("1", "B", e1)
        fmt.Println (e1)
        fmt.Println (e2)
        fmt.Println (e1.GetId (), e1.Error (), e1.Unwrap ())
        fmt.Println (e2.GetId (), e2.Error (), e2.Unwrap ())
}
