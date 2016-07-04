package main

import (
    //"bufio"  //缓存IO
    "fmt"
    //"io/ioutil" //io 工具包
    //"io"
    "os"
    textTemplate "text/template"
    "matrix/core"
    "matrix/smith/template"
    "matrix/smith/models"
    "matrix/smith"
    "bytes"
    "strings"
    "matrix/smith/fieldtype"
    "strconv"
)

var outputBaseDir = "z:\\codegen_output"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) (bool) {
    var exist = true;
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        exist = false;
    }
    return exist;
}

/**
    from: http://www.isharey.com/?p=143


  递归创建目录
  os.MkdirAll(path string, perm FileMode) error

  path  目录名及子目录
  perm  目录权限位
  error 如果成功返回nil，如果目录已经存在默认什么都不做

package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.MkdirAll("C:\\\\temp\\\\log", 0777)
    if err != nil {
        fmt.Printf("%s", err)
    } else {
        fmt.Print("Create Directory OK!")
    }
}
    //os.RemoveAll(path)

 */

func main() {

    entityList := []smith.Entity{
        models.GroupEntity,
    }

    result := RenderCodeTemplate("models", template.ModelsTemplate, map[string]interface{}{
        "tablePrefix": "hd_auth_",
        "entityList": entityList,
    })

    fmt.Println(result)

    result = RenderCodeTemplate("controller", template.ControllerTemplate, map[string]interface{}{
        "entity": models.GroupEntity,
    })

    fmt.Println(result)


    /*
    var wireteString = "测试n"
    var filename = "./output1.txt";
    var f    *os.File
    var err1 error;
    //***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************
    if checkFileIsExist(filename) {
        //如果文件存在
        f, err1 = os.OpenFile(filename, os.O_APPEND, 0666)  //打开文件
        fmt.Println("文件存在");
    } else {
        f, err1 = os.Create(filename)  //创建文件
        fmt.Println("文件不存在");
    }
    check(err1)
    n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
    check(err1)
    fmt.Printf("写入 %d 个字节n", n);

    //*****************************  第二种方式: 使用 ioutil.WriteFile 写入文件 ***********************************************
    var d1 = []byte(wireteString);
    err2 := ioutil.WriteFile("./output2.txt", d1, 0666)  //写入文件(字节数组)
    check(err2)

    //*****************************  第三种方式:  使用 File(Write,WriteString) 写入文件 ***********************************************
    f, err3 := os.Create("./output3.txt")  //创建文件
    check(err3)
    defer f.Close()
    n2, err3 := f.Write(d1)  //写入文件(字节数组)
    check(err3)
    fmt.Printf("写入 %d 个字节n", n2)
    n3, err3 := f.WriteString("writesn") //写入文件(字节数组)
    fmt.Printf("写入 %d 个字节n", n3)
    f.Sync()




    //***************************** 第四种方式:  使用 bufio.NewWriter 写入文件 ***********************************************
    w := bufio.NewWriter(f)  //创建新的 Writer 对象
    n4, err3 := w.WriteString("bufferedn")
    fmt.Printf("写入 %d 个字节n", n4)
    w.Flush()
    f.Close()
    */
}

func RenderCodeTemplate(tplName string, tplContent string, args  map[string]interface{}) string {
    template, err := textTemplate.New(tplName).Funcs(textTemplate.FuncMap{
        "RenderField": RenderField,
    }).Parse(tplContent)

    if err != nil {
        panic(err)
    }

    var buffer bytes.Buffer
    template.Execute(&buffer, args)
    return buffer.String()
}

func RenderField(field smith.Field) string {
    //field name
    fieldName := "    "
    nameCharLength := 15 //15 是字段名称在代码编辑器中占的字符数
    fieldName += field.Name
    fieldName += strings.Repeat(" ", nameCharLength - len(field.Name))


    //field type
    fieldType := ""
    //field column
    fieldColumnType := ""
    typeLength := 18 //18 是字段类型在代码编辑器中占的字符数
    if field.FieldType == fieldtype.Int {
        fieldType = "int"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType = "int"
    } else if field.FieldType == fieldtype.BigInt {
        fieldType = "int64"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType = "bigint"
    } else if field.FieldType == fieldtype.Decimal {
        fieldType = "float64"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType = "decimal(" + strconv.Itoa(field.Precision) + "," + strconv.Itoa(field.Scale) + ")"

    } else if field.FieldType == fieldtype.NVarchar {
        //GroupName  string           `xorm:"nvarchar(255) notnull unique index 'group_name'" json:"group_name"`
        fieldType = "string"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType += "nvarchar(" + strconv.Itoa(field.Length) + ")"

    } else if field.FieldType == fieldtype.DateTime {
        fieldType = "core.Time"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType += "datetime"

    } else if field.FieldType == fieldtype.Boolean {
        fieldType = "bool"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType += "bool"

    } else if field.FieldType == fieldtype.Create {
        fieldType = "core.Time"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType += "created"

    } else if field.FieldType == fieldtype.Update {
        fieldType = "core.Time"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType += "updated"

    } else if field.FieldType == fieldtype.Version {
        fieldType = "int"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType += "version"
    }

    fieldNull := ""
    if field.Null {
        fieldNull += "null"
    } else {
        fieldNull += "notnull"
    }

    fieldPrimaryKey := ""
    if field.Primarykey {
        fieldPrimaryKey = "pk autoincr "
    }

    fieldUnique := ""
    if field.Unique {
        fieldUnique = "unique "
    }

    fieldIndex := ""
    if field.Index {
        fieldIndex = "index "
    }

    fieldColumn := field.Column

    fieldCode := core.FormatText("field", `{{.fieldName}}{{.fieldType}}{{.tagchar}}xorm:"{{.fieldColumnType}} {{.fieldNull}} {{.fieldPrimaryKey}}{{.fieldUnique}}{{.fieldIndex}}'{{.fieldColumn}}'" json:"{{.fieldColumn}}"{{.tagchar}}`, map[string]interface{}{
        "tagchar": "`",
        "fieldName": fieldName,
        "fieldType": fieldType,
        "fieldColumnType": fieldColumnType,
        "fieldNull": fieldNull,
        "fieldPrimaryKey": fieldPrimaryKey,
        "fieldUnique": fieldUnique,
        "fieldIndex": fieldIndex,
        "fieldColumn": fieldColumn,
    })

    return fieldCode
}
