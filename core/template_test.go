package core

import (
	"fmt"
	"testing"
)

func Test_FormatSQL(t *testing.T) {
	sql := FormatText("name1", "SQL: {{.hello}}", map[string]interface{}{
		"hello": "Hello World!",
	})

	fmt.Println(sql)

	if sql == "SQL: Hello World!" {
		t.Log("pass!")
	} else {
		t.Error("error!")
	}
}
