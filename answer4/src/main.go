package main

import (
	"strconv"

	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/:operation/:num1:int/:num2:int", &mainController{})
	beego.Run()
}

type mainController struct {
	beego.Controller
}

func (c *mainController) Get() {
	operation := c.Ctx.Input.Param(":operation")
	num1, _ := strconv.Atoi(c.Ctx.Input.Param(":num1"))
	num2, _ := strconv.Atoi(c.Ctx.Input.Param(":num2"))

	c.Data["operation"] = operation
	c.Data["num1"] = num1
	c.Data["num2"] = num2
	c.TplName = "result.html"

	switch operation {
	case "sum":
		c.Data["result"] = palindrome(num1, num2)

	default:
		c.TplName = "invalid-route.html"
	}

}

func palindrome(n1, n2 int) int {
	for i = n1; i <= n2; i++ {
		var reverse int = 0
		temp := i
		for {
			remainder = temp % 10
			reverse = reverse*10 + remainder
			temp = temp / 10
			if temp == 0 {
				break
			}
		}
		if reverse == i {
			sum++
		}
	}
	return sum
}
