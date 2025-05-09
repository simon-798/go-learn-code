package lesson3

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 导入MySQL驱动
	"github.com/jmoiron/sqlx"
	"strings"
)

/*
*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
func Job3Method() {
	if err := sqlxInitDB(); err != nil {
		fmt.Printf("init db fail, err:%v\n", err)
	}

	createTable()

	employeeList := []Employee{
		{Name: "李四", Department: "技术部", Salary: 288888.88},
		{Name: "王五", Department: "销售部", Salary: 18888.88},
		{Name: "赵六", Department: "技术部", Salary: 888888.88},
	}

	insertCount, err := batchInsert(employeeList)
	if err != nil {
		fmt.Printf("batch insert fail, err:%v\n", err)
	}

	fmt.Printf("insert count:%d\n", insertCount)

	//查询 employees 表中所有部门为 "技术部" 的员工信息
	query1 := `select * from employees where department = ?`
	var employeeList1 []Employee
	err = DB.Select(&employeeList1, query1, "技术部") //Select()：用于查询多条记录到切片
	if err != nil {
		fmt.Printf("query1 fail, err:%v\n", err)
	}
	fmt.Printf("查询 employees 表中所有部门为 \"技术部\" 的员工信息,employeeList1:%v\n", employeeList1)

	//查询 employees 表中工资最高的员工信息
	query2 := `select * from employees order by salary desc limit 1`
	var employee2 Employee
	err = DB.Get(&employee2, query2) //Get()：用于查询单条记录到结构体
	if err != nil {
		fmt.Printf("query2 fail, err:%v\n", err)
	}

	fmt.Printf("查询 employees 表中工资最高的员工信息,employee2:%v\n", employee2)

}

// Employee 结构体定义（与数据库表结构映射）
type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

var DB *sqlx.DB

func sqlxInitDB() (err error) {
	dsn := "root:20250423qwER@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB err:%v\n", err)
		return err
	}

	DB = db

	DB.SetMaxOpenConns(30)
	DB.SetMaxIdleConns(10)

	return nil

}

func createTable() {
	createTableSql := `CREATE TABLE IF NOT EXISTS employees (
    	id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    	name varchar(255) NOT NULL DEFAULT '',
    	department varchar(255) NOT NULL DEFAULT '',
    	salary Decimal(18,2) NOT NULL DEFAULT '0.00'
		)`
	_, err := DB.Exec(createTableSql)
	if err != nil {
		fmt.Printf("create table err:%v\n", err)
	}

	fmt.Println("create table success")
}

func insertData() {
	insertSql := "insert into employees (name,department,salary) values (?,?,?)"
	insertResult, err := DB.Exec(insertSql, "张三", "销售部", 10008.88)
	if err != nil {
		fmt.Printf("insert data err:%v\n", err)
	}

	lastId, err := insertResult.LastInsertId()
	if err != nil {
		fmt.Printf("insert data err:%v\n", err)
	}

	fmt.Println("insert success, lastId:", lastId)

}

func batchInsert(employees []Employee) (int64, error) {

	batchSql := `INSERT INTO employees (name,department,salary) VALUES `

	// 每条记录3个参数
	placeholders := make([]string, 0, len(employees))
	params := make([]interface{}, 0, len(employees)*3) // 每条记录3个参数

	for _, employee := range employees {
		placeholders = append(placeholders, "(?,?,?)") // 每个括号对应一条记录
		params = append(params, employee.Name, employee.Department, employee.Salary)
	}
	fmt.Println("placeholders:", placeholders)
	fmt.Println("params:", params)

	// 拼接完整SQL
	batchSql += strings.Join(placeholders, ",")

	result, err := DB.Exec(batchSql, params...)
	if err != nil {
		return 0, fmt.Errorf("执行插入失败: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("获取影响行数失败: %w", err)
	}

	return rowsAffected, nil
}
