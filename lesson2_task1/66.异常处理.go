package main

import (
	"errors"
	"fmt"
)

/*
《error》

			error 是一种接口类型，用于表示和处理错误情况。

			一、声明错误：
				你可以使用内置的 errors.New 函数来创建一个新的错误。这个函数接受一个字符串作为参数，并返回一个实现了 error 接口的值。
				err := errors.New("这是一个错误")

			二、返回错误
				函数可以通过返回多个值来报告错误。通常，这是第二个返回值。
				func divide(a, b float64) (float64, error) {
		    		if b == 0 {
		        		return 0, errors.New("除数不能为零")
		    		}
					return a / b, nil
				}
			三、处理错误：
				调用可能返回错误的函数时，应该检查错误值，并根据需要处理错误。
				func main() {
	    			result, err := divide(10, 0)
					if err != nil {
	        			fmt.Println("错误:", err)
	    			} else {
	        			fmt.Println("结果:", result)
	   			 	}
				}

			四、自定义错误
				除了使用 errors.New 创建简单的错误消息外，你还可以通过实现 error 接口来自定义错误类型。
*/
func main() {

	res, err := divide(2, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func divide(num1 int, num2 int) (result int, err error) {
	err = nil
	if num2 == 0 {
		err = errors.New("除数不能为0")
		return
		//  return 0, errors.New("除数不能为零")
	}
	result = num1 / num2
	return
}
