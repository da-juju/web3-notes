package main

import "fmt"

type Stu struct {
	name    string
	sex     string
	age     int
	chinese float64
	math    float64
	english float64
}

// Hi /** 打招呼
func (s Stu) Hi() {
	fmt.Printf("大家好！我叫%s，今年%d岁了，是%s同学。\n", s.name, s.age, s.sex)
}

// SumAndAvg /* 计算总分和平均分
func (s Stu) SumAndAvg() {
	fmt.Println("总分：", s.chinese+s.math+s.english)
	fmt.Printf("平均分：%.2f", (s.chinese+s.math+s.english)/3)
}

func main() {
	stu1 := Stu{"小黄", "女", 18, 99, 66, 88}
	stu1.Hi()
	stu1.SumAndAvg()
}
