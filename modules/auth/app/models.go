package app


type User struct {
    Id int64
    LoginName string `xorm: "varchar(255) notnull unique"`
    NickName string `xorm: "varchar(255) notnull"`
}
