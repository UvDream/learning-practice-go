package student

import "fmt"

// Student 解释
type Student struct {
	id   int64
	name string
}

// Studentmanagement 解释
type Studentmanagement struct {
	AllStudent map[int64]Student
}

// ShowStudent 查看
func (s Studentmanagement) ShowStudent() {
	fmt.Println("查看所有学生")
	for _, v := range s.AllStudent {
		fmt.Printf("学号:%d,姓名:%s\n", v.id, v.name)
	}

}

// AddStudent 增加
func (s Studentmanagement) AddStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&name)

	// 根据用户输入
	newStu := Student{
		id:   id,
		name: name,
	}
	// 放入map
	s.AllStudent[newStu.id] = newStu
}

// EditStudent 修改
func (s Studentmanagement) EditStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入需要修改的学生学号:")
	fmt.Scanln(&id)
	m, ok := s.AllStudent[id]
	if !ok {
		fmt.Println("查无此人!!!!!")
	} else {
		fmt.Printf("你要修改的学号:%d 姓名:%s\n", m.id, m.name)
		fmt.Print("请输入学生姓名:")
		fmt.Scanln(&name)
		// 修改1
		m.name = name
		s.AllStudent[id] = m
	}
	// fmt.Print("请输入学生名:")
	// fmt.Scanln(&name)
}

// DeleteStudent 删除
func (s Studentmanagement) DeleteStudent() {
	var (
		id int64
	)
	fmt.Print("请输入需要修改的学生学号:")
	fmt.Scanln(&id)
	_, ok := s.AllStudent[id]
	if !ok {
		fmt.Println("查无此人!!!!!")
	} else {
		// map删除键值对
		delete(s.AllStudent, id)
	}

}
