package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	//使用sql包是必须注入（至少）一个数据库驱动
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
)

// Student 定义学生结构体
type Student struct {
	StudentID string
	Name      string
	Age       int
	Gender    string
}

// 数据库连接信息
const (
	username = "root"
	password = "2003926jia.."
	hostname = "127.0.0.1:3306"
	dbname   = "student_management"
)

// 连接数据库
func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

// 添加学生信息
func addStudent(db *sql.DB) error {
	var studentID, name, gender string
	var age int
	fmt.Print("请输入学生学号: ")
	fmt.Scanln(&studentID)
	fmt.Print("请输入学生姓名: ")
	fmt.Scanln(&name)
	fmt.Print("请输入学生年龄: ")
	fmt.Scanln(&age)
	fmt.Print("请输入学生性别: ")
	fmt.Scanln(&gender)

	student := Student{
		StudentID: studentID,
		Name:      name,
		Age:       age,
		Gender:    gender,
	}
	query := "INSERT INTO students2 (student_id, name, age, gender) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, student.StudentID, student.Name, student.Age, student.Gender)
	return err
}

// 删除学生信息
func deleteStudent(db *sql.DB) error {
	var studentID string
	fmt.Print("请输入要删除的学生学号: ")
	fmt.Scanln(&studentID)
	query := "DELETE FROM students2 WHERE student_id = ?"
	result, err := db.Exec(query, studentID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("未找到学号为 %s 的学生信息", studentID)
	}
	return nil
}

// 修改学生信息
func updateStudent(db *sql.DB) error {
	var studentID, name, gender string
	var age int
	fmt.Print("请输入要修改的学生学号: ")
	fmt.Scanln(&studentID)
	fmt.Print("请输入新的学生姓名: ")
	fmt.Scanln(&name)
	fmt.Print("请输入新的学生年龄: ")
	fmt.Scanln(&age)
	fmt.Print("请输入新的学生性别: ")
	fmt.Scanln(&gender)

	student := Student{
		StudentID: studentID,
		Name:      name,
		Age:       age,
		Gender:    gender,
	}
	query := "UPDATE students2 SET name = ?, age = ?, gender = ? WHERE student_id = ?"
	result, err := db.Exec(query, student.Name, student.Age, student.Gender, studentID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("未找到学号为 %s 的学生信息，无法更新", student.StudentID)
	}
	return nil
}

// 查询单个学生信息
func getStudentByID(db *sql.DB) (*Student, error) {
	var studentID string
	fmt.Print("请输入要查询的学生学号: ")
	fmt.Scanln(&studentID)
	query := "SELECT student_id, name, age, gender FROM students2 WHERE student_id = ?"
	row := db.QueryRow(query, studentID)
	var student Student
	err := row.Scan(&student.StudentID, &student.Name, &student.Age, &student.Gender)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("未找到学号为 %s 的学生信息", studentID)
		}
		return nil, err
	}
	return &student, nil
}

// 查询所有学生信息
func getAllStudents(db *sql.DB) ([]Student, error) {
	query := "SELECT  student_id, name, age, gender FROM students2"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var students []Student
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.StudentID, &student.Name, &student.Age, &student.Gender)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return students, nil
}
func main() {
	//连接数据库
	db, err := connectDB()
	if err != nil {
		fmt.Println("数据库连接失败: ", err)
		return
	}
	defer db.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n请选择操作：")
		fmt.Println("1. 添加学生信息")
		fmt.Println("2. 删除学生信息")
		fmt.Println("3. 修改学生信息")
		fmt.Println("4. 查询单个学生信息")
		fmt.Println("5. 查询所有学生信息")
		fmt.Println("6. 退出")
		fmt.Print("请输入选项编号: ")
		if scanner.Scan() {
			choice := scanner.Text()
			switch strings.TrimSpace(choice) {
			case "1":
				err = addStudent(db)
				if err != nil {
					fmt.Println("学生信息添加失败: ", err)
				} else {
					fmt.Println("学生信息添加成功")
				}
			case "2":
				err = deleteStudent(db)
				if err != nil {
					fmt.Println("修改学生信息失败:", err)
				} else {
					fmt.Println("学生信息删除成功")
				}
			case "3":
				err = updateStudent(db)
				if err != nil {
					fmt.Println("修改学生信息失败:", err)
				} else {
					fmt.Println("学生信息修改成功")
				}
			case "4":
				student, err := getStudentByID(db)
				if err != nil {
					fmt.Println("查询学生信息失败:", err)
				} else {
					fmt.Printf("查询到的学生信息: 学号：%s, 姓名：%s, 年龄：%d, 性别：%s\n", student.StudentID, student.Name, student.Age, student.Gender)
				}
			case "5":
				students, err := getAllStudents(db)
				if err != nil {
					fmt.Println("查询学生信息失败:", err)
				} else {
					fmt.Println("所有学生信息如下：")
					for _, student := range students {
						fmt.Printf("查询到的学生信息: 学号：%s, 姓名：%s, 年龄：%d, 性别：%s\n", student.StudentID, student.Name, student.Age, student.Gender)
					}
				}
			case "6":
				fmt.Println("退出程序")
				return
			default:
				fmt.Println("无效的选项，请重新输入。")
			}
		}
	}
}
