package test

import (
	"fmt"
	"testing"
)

func Test数组指针(t *testing.T) {
	var arr = [5]int{1,2,3,4,5};
	var p = &arr;
	fmt.Println(p)//go1.16.5显示 &[1 2 3 4 5]

	//此时p称为是数组arr的指针
	//此时如果通过p访问数组元素是 (*p)[2]，结果是3

}
func Test指针数组(t *testing.T) {
	var arr1 [5]int = [5]int{1,2,3,4,5};
	var arr2 [5]int = [5]int{6,7,8,9,0};
	var p1 *[5]int = &arr1;
	var p2 *[5]int = &arr2;
	var pArr [2]*[5]int = [2]*[5]int{p1,p2};
	fmt.Println(pArr);//[0xc00006a030, 0xc00006a060]
	//此时pArr称为是指针数组
	//此时如果通过pArr访问数组元素是 (*pArr[0])[2],结果是3
}
func Test切片名(t *testing.T) {
	var slice []int = []int{1,2,3,4,5}
	fmt.Printf("%p\n", slice)//0xc000012346;
	fmt.Println(slice[0])
}
func Test二重指针(t *testing.T) {
	var slice []int = []int{1,2,3,4,5}
	fmt.Printf("%p    %T\n", slice,slice)//0xc000012346;
	fmt.Println(slice[0])
	p := &slice;
	fmt.Printf("%p   %T\n", p,p);//0xc000012446;
}
