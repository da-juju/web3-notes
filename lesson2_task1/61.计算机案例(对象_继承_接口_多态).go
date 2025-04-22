package main

import "fmt"

/*
《计算机的案例》

	使用对象，继承，接口以及多态的概念，来实现计算机的加减运算
*/

// BaseObj 父类（基类）
type BaseObj struct {
	num1 int
	num2 int
}

// Add 加法类 继承BaseObj
type Add struct {
	BaseObj // 匿名字段继承
}

// Sub 减法类 继承BaseObj
type Sub struct {
	BaseObj // 匿名字段继承
}

// Operator 接口
type Operator interface {
	// 计算的模板方法
	calculate() int
}

// Add类实现了接口的calculate()方法
func (obj *Add) calculate() int {
	return obj.num1 + obj.num2
}

// Sub类实现了接口的calculate()方法
func (obj *Sub) calculate() int {
	return obj.num1 - obj.num2
}

// 多态方法：根据入参不同，分别实现加减计算运算逻辑
// 形参：必须是接口类型
func calculateWho(op Operator) int {
	return op.calculate()
}

func getCalculateResult(num1 int, num2 int, op string) int {
	switch op {
	case "+":
		var addObj = Add{BaseObj{num1, num2}}
		return calculateWho(&addObj)
	case "-":
		var subObj = Sub{BaseObj{num1, num2}}
		return calculateWho(&subObj)
	default:
		return 0
	}
}

func main() {
	umn := getCalculateResult(1, 2, "-")
	fmt.Println("umn=", umn)
}
