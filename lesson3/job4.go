package lesson3

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

/**
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

func Job4Method() {

	if err := sqlxInitBookDB(); err != nil {
		fmt.Printf("init db fail, err:%v\n", err)
	}

	/*createBooksTable()
	bookList := []Book{
		{Title: "深入Go语言", Author: "张三", Price: 68.5},
		{Title: "系统编程实践", Author: "李四", Price: 55},
		{Title: "数据库原理", Author: "王五", Price: 45},
	}
	insertCount, err := batchBooksInsert(bookList)
	if err != nil {
		fmt.Printf("batch insert fail, err:%v\n", err)
	}

	fmt.Printf("insert count:%d\n", insertCount)*/

	query := "select * from books where price > ?"
	var result []Book
	err := DB_BOOK.Select(&result, query, 50)
	if err != nil {
		fmt.Printf("query fail, err:%v\n", err)
	}
	fmt.Printf("查询价格大于50元的书籍,result:%#v\n", result)

	//分页查询，limit表示每页多少条数据，offset表示从第几条数据开始
	pageSize := 10
	pageNumber := 1
	offset := (pageNumber - 1) * pageSize

	pageQuery := "select * from books limit ? offset ?"
	var pageResult []Book
	err = DB_BOOK.Select(&pageResult, pageQuery, pageSize, offset)
	if err != nil {
		fmt.Printf("page query fail, err:%v\n", err)
	}
	fmt.Printf("分页查询,pageResult:%#v\n", pageResult)
}

type Book struct {
	Id     uint64  `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

var DB_BOOK *sqlx.DB

func sqlxInitBookDB() (err error) {
	dsn := "root:20250423qwER@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB err:%v\n", err)
		return err
	}

	DB_BOOK = db

	DB_BOOK.SetMaxOpenConns(30)
	DB_BOOK.SetMaxIdleConns(10)

	return nil

}

func createBooksTable() {
	createTableSql := `CREATE TABLE IF NOT EXISTS books (
    	id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    	title varchar(255) NOT NULL DEFAULT '',
    	author varchar(64) NOT NULL DEFAULT '',
    	price Decimal(18,2) NOT NULL DEFAULT '0.00'
		)`

	_, err := DB_BOOK.Exec(createTableSql)
	if err != nil {
		fmt.Printf("create table err:%v\n", err)
	}

	fmt.Println("create table success")
}

func batchBooksInsert(books []Book) (int64, error) {

	batchSql := `INSERT INTO books (title,author,price) VALUES `

	// 每条记录3个参数
	placeholders := make([]string, 0, len(books))
	params := make([]interface{}, 0, len(books)*3) // 每条记录3个参数

	for _, book := range books {
		placeholders = append(placeholders, "(?,?,?)") // 每个括号对应一条记录
		params = append(params, book.Title, book.Author, book.Price)
	}
	fmt.Println("placeholders:", placeholders)
	fmt.Println("params:", params)

	// 拼接完整SQL
	batchSql += strings.Join(placeholders, ",")

	result, err := DB_BOOK.Exec(batchSql, params...)
	if err != nil {
		return 0, fmt.Errorf("执行插入失败: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("获取影响行数失败: %w", err)
	}

	return rowsAffected, nil
}
