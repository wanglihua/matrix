package controllers

import (
	"database/sql"
	"fmt"
	"strings"

	"bytes"

	"github.com/revel/revel"
)

type CodeGenerate struct {
	*revel.Controller
}

func (c CodeGenerate) Index() revel.Result {
	//session := c.DbSession

	xx := strings.ToUpper("ddd")

	fmt.Println(xx)

	return c.RenderTemplate("code_generate/code_generate_index.html")
}

func (c CodeGenerate) Code() revel.Result {
	var result_type string
	c.Params.Bind(&result_type, "result_type")

	var sql_query string
	c.Params.Bind(&sql_query, "sql_query")

	db_driver, _ := revel.Config.String("db.driver")
	db_url, _ := revel.Config.String("db.url")

	db, _ := sql.Open(db_driver, db_url)
	rows, err := db.Query(sql_query)
	if err != nil {
		return c.RenderJson(map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}
	columns, _ := rows.Columns()

	/*
		type ManagerCount struct {
			ManagerCount int `xorm:"'manager_count'"`
		}
		var manager_count ManagerCount
		sql := `SELECT count(DISTINCT m.id) AS manager_count FROM hd_itsm_engineer e INNER JOIN hd_itsm_engineer_manager m ON e.id = m.engineer_id WHERE e.id = ?;`
		_, err := db_session.Sql(sql, engineer_id).Get(&manager_count)
		core.HandleError(err)
	*/

	under_score_to_title := func(under_score_string string) string {
		var buffer bytes.Buffer
		is_after_under_score := false
		for i, ch := range under_score_string {
			if ch == []rune("_")[0] {
				is_after_under_score = true
			} else if i == 0 || is_after_under_score {
				buffer.WriteString(strings.ToUpper(string(ch)))
				is_after_under_score = false
			} else {
				buffer.WriteRune(ch)
			}
		}
		return buffer.String()
	}

	result_struct_type := under_score_to_title(result_type)

	var code_buffer bytes.Buffer
	code_buffer.WriteString(fmt.Sprintf("type %s struct {\n", result_struct_type))

	for _, column := range columns {
		code_buffer.WriteString(fmt.Sprintf("\t%s Core.NullString `xorm:\"'%s'\"`\n", under_score_to_title(column), column))
	}
	code_buffer.WriteString("}\n\n")

	code_buffer.WriteString(fmt.Sprintf("var %s %s\n", result_type, result_struct_type))
	code_buffer.WriteString(fmt.Sprintf("sql := `%s`\n", sql_query))
	code_buffer.WriteString(fmt.Sprintf("_, err := db_session.Sql(sql).Get(&%s)\n", result_type))
	code_buffer.WriteString("core.HandleError(err)")

	return c.RenderJson(map[string]interface{}{
		"success": true,
		"message": "生成成功！",
		"code":    code_buffer.String(),
	})
}
