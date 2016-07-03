package core

import (
    "github.com/revel/revel/testing"
    "matrix/core"
)

type TemplateTest struct {
    testing.TestSuite
}

func (t *TemplateTest) TestFormatSQL() {
    sql := core.FormatText("name1", "SQL: {{.hello}}", map[string]interface{}{
        "hello" : "Hello World!",
    })

    t.AssertEqual("SQL: Hello World!", sql)
}
