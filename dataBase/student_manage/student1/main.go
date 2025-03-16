package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Student 定义学生结构体
type Student struct {
	ID        int
	StudentId string
	Name      string
	Gender    string
}

// 链接数据库
const (
	username = "root"
	password = "2003926jia.."
	hostname = "127.0.0.1:3306"
	dbname   = "student_management"
)

func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//测试数据库连接
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

// 添加
func addStudent(db *sql.DB, student Student) error {
	query := "INSERT INTO students1 (student_id, name, gender) VALUES (?, ?, ?)"
	_, err := db.Exec(query, student.StudentId, student.Name, student.Gender)
	return err
}

// 删除
func deleteStudent(db *sql.DB, studentID string) error {
	query := "DELETE FROM students1 WHERE student_id = ?"
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

// 修改
func updateStudent(db *sql.DB, student Student) error {
	query := "UPDATE students1 SET name = ?, gender = ? WHERE student_id = ?"
	result, err := db.Exec(query, student.Name, student.Gender, student.StudentId)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("未找到学号为 %s 的学生信息，无法更新", student.StudentId)
	}
	return nil
}

// 查询单个学生信息
func getStudentByID(db *sql.DB, studentID string) (*Student, error) {
	query := "SELECT id, student_id, name, gender FROM students1 WHERE student_id = ?"
	row := db.QueryRow(query, studentID)
	var student Student
	err := row.Scan(&student.ID, &student.StudentId, &student.Name, &student.Gender)
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
	query := "SELECT id, student_id, name, gender FROM students1"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.ID, &student.StudentId, &student.Name, &student.Gender)
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
	// 连接数据库
	db, err := connectDB()
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	defer db.Close()

	// 添加学生信息
	newStudent1 := Student{
		StudentId: "2024001",
		Name:      "李四",
		Gender:    "女",
	}
	err = addStudent(db, newStudent1)
	newStudent2 := Student{
		StudentId: "2024002",
		Name:      "王五",
		Gender:    "男",
	}
	err = addStudent(db, newStudent2)
	newStudent3 := Student{
		StudentId: "2024003",
		Name:      "张三",
		Gender:    "女",
	}
	err = addStudent(db, newStudent3)
	newStudent4 := Student{
		StudentId: "2024004",
		Name:      "王一",
		Gender:    "女",
	}
	err = addStudent(db, newStudent4)
	if err != nil {
		fmt.Println("添加学生信息失败:", err)
	} else {
		fmt.Println("学生信息添加成功")
	}

	// 查询单个学生信息
	//student, err := getStudentByID(db, "2024002")
	//if err != nil {
	//	fmt.Println("查询学生信息失败:", err)
	//} else {
	//	fmt.Printf("查询到的学生信息：ID: %d, 学号: %s, 姓名: %s, 性别: %s\n", student.ID, student.StudentId, student.Name, student.Gender)
	//}

	// 修改学生信息
	//moStudent := Student{
	//	StudentId: "2024002",
	//	Name:      "李思思",
	//	Gender:    "女",
	//}
	//err = updateStudent(db, moStudent)
	//if err != nil {
	//	fmt.Println("修改学生信息失败:", err)
	//} else {
	//	fmt.Println("学生信息修改成功")
	//}

	// 查询所有学生信息
	students, err := getAllStudents(db)
	if err != nil {
		fmt.Println("查询所有学生信息失败:", err)
	} else {
		fmt.Println("所有学生信息如下：")
		for _, student := range students {
			fmt.Printf("ID: %d, 学号: %s, 姓名: %s, 性别: %s\n", student.ID, student.StudentId, student.Name, student.Gender)
		}
	}

	// 删除学生信息
	//err = deleteStudent(db, "2024002")
	//if err != nil {
	//	fmt.Println("删除学生信息失败:", err)
	//} else {
	//	fmt.Println("学生信息删除成功")
	//}
}
