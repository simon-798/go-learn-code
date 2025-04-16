package basics

import (
	"fmt"
	"strconv"
)

func mainTypeConvertMethod() {

	typeConvertMethod5()
}

func typeConvertMethod1() {
	var i int32 = 17
	var b byte = 5
	var f float32

	// 数字类型可以直接强转
	f = float32(i) / float32(b)
	ff := float32(i)
	fmt.Println("ff:", ff)
	fmt.Printf("f 的值为: %f\n", f)

	// 当int32类型强转成byte时，高位被直接舍弃
	var i2 int32 = 256
	var b2 = byte(i2)
	fmt.Printf("b2 的值为: %d\n", b2)
}
func typeConvertMethod2() {
	str := "hello, 123, 你好"
	var bytes = []byte(str)
	var runes = []rune(str)
	fmt.Printf("bytes 的值为: %v \n", bytes)
	fmt.Printf("runes 的值为: %v \n", runes)

	str2 := string(bytes)
	str3 := string(runes)
	fmt.Printf("str2 的值为: %v \n", str2)
	fmt.Printf("str3 的值为: %v \n", str3)
}

func typeConvertMethod3() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("字符串转换为int: %d \n", num)
	str1 := strconv.Itoa(num)
	fmt.Printf("int转换为字符串: %s \n", str1)

	ui64, err := strconv.ParseUint(str, 10, 32)
	fmt.Printf("字符串转换为uint64: %d \n", num)

	str2 := strconv.FormatUint(ui64, 2)
	fmt.Printf("uint64转换为字符串: %s \n", str2)

	str3 := fmt.Sprintf("%d", int32(ui64))
	fmt.Printf("str3:%s \n", str3)
}

func typeConvertMethod4() {
	var i interface{} = "3"
	a, ok := i.(string)
	if ok {
		fmt.Println("a:", a)
	} else {
		fmt.Println("conversion failed")
	}

	var j interface{} = "test"
	switch v := j.(type) {
	case int:
		fmt.Println("i is a int", v)
	case string:
		fmt.Println("i is a string", v)
	default:
		fmt.Println("i is unknown type", v)
	}
}

func typeConvertMethod5() {
	/*var aa = DigitSupplier{value: 1}
	fmt.Println(aa)*/
	var a Supplier = &DigitSupplier{value: 1}
	fmt.Println(a.Get())

	b, ok := (a).(*DigitSupplier)
	fmt.Println(b, ok)
}

type Supplier interface {
	Get() string
}

type DigitSupplier struct {
	value int
}

func (i *DigitSupplier) Get() string {
	return fmt.Sprintf("%d", i.value)
}

func typeConvertMethod6() {
	a := SameFieldA{
		name:  "a",
		value: 1,
	}

	b := SameFieldB(a)
	fmt.Printf("conver SameFieldA to SameFieldB, value is : %d \n", b.getValue())

	// 只能结构体类型实例之间相互转换，指针不可以相互转换
	var c interface{} = &a
	_, ok := c.(*SameFieldB)
	fmt.Printf("c is *SameFieldB: %v \n", ok)
}

type SameFieldA struct {
	name  string
	value int
}

type SameFieldB struct {
	name  string
	value int
}

func (s *SameFieldB) getValue() int {
	return s.value
}
