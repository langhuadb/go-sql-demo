package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/gin-gonic/gin"
)
type UserInfoModel struct {
	Id 			int
	UserId		string
	Name		string
	Gender		string
	Birthday	string
	Telphone	string
	Email		string
	Addr		string
}
func main() {
	router:=gin.Default()
	router.GET("/student", func(context *gin.Context) {
		db,err:=sql.Open("mysql","root:123@tcp(127.0.0.1:3306)/gocrud")
		if err != nil {
			panic(err)
		}
		log.Println("数据库连接成功。。。")
		user:=UserInfoModel{}
		row:=db.QueryRow("select * from user_info;")
		e:=row.Scan(&user.Id,&user.UserId,&user.Name,&user.Gender,&user.Birthday,&user.Telphone,&user.Email,&user.Addr)
		if e != nil {
			panic(e)
		}
		log.Println(user)
		db.Close()
	})

	router.Run(":8888")
}
