package article

import (
	"database/sql"
	// 运行时请在本地装好该mysql驱动包
	// go get github.com/go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/pkg/errors"
)
// 定义一批变量以便在各个方法之间使用
var (
	db *sql.DB
	stmt *sql.Stmt
	rows *sql.Rows
)
// 开启连接
func openConn() (err error){
	// 此处请使用对应的账户密码
	db , err = sql.Open("mysql","root:xxx@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("数据库链接错误,", err)
		return
	}
	return nil
}

func closeConn() ( err error) {
	if db != nil {
		db.Close()
	}
	return 
}

func GetDetail(id string) (*Article,error) {
	sql := "select id,title,content,Created,Updated from t_article where id= ?"
	rows , err := query(sql,id)
	// fmt.Println(errors.Unwrap(err))
	// fmt.Printf("%+v\r\n",err)
	// fmt.Println("the error ",err)
	// if err !=nil {
	// 	fmt.Println(err)
	// 	return nil
	// }
	// if rows.Next(){
	// 	article := new(Article)
	// 	rows.Scan(&article.Id,&article.Title,&article.Content,&article.Created,&article.Updated)
	// 	closeConn();
	// 	return article
	// }
	return rows,err
}

func query ( querySql string,id string) (*Article ,error) {
	err := openConn()
	article := new(Article)
	err = db.QueryRow(querySql,id).Scan(&article.Id,&article.Title,&article.Content,&article.Created,&article.Updated)
	if err == sql.ErrNoRows {
		// fmt.Println("数据库链接错误:",err)
		return nil,errors.Wrapf(err,"数据库访问失败")
	}
	return article,err
	// if err!= nil {
	// 	fmt.Println("数据库链接错误")
	// 	return nil,err
	// }

	// stmt , err = db.Prepare (sql)
	// if err != nil {
	// 	fmt.Println("预处理失败")
	// 	return nil,err
	// }
	// rows , err := stmt.Query(args ...)
	// if err != nil {
	// 	fmt.Println("执行数据库操作失败")
	// 	return nil,err
	// }
	// return rows ,nil
}
