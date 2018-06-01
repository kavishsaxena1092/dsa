package main

import ("fmt")
func BubbleSort(arr[] int)[]int {
	for i:=1; i< len(arr); i++ {
		for j:=0; j < len(arr)-i; j++ {
			if (arr[j] > arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}


func main() {
	fmt.Print("Enter Array Length: ")
	var array_length int
	_, e := fmt.Scanf("%d", &array_length)
	if e != nil {
		fmt.Println(e)		
	}
	slice := make([]int, 0)
	var num int
	
	for i:=0; i<array_length; i++ {
	fmt.Print("Enter a Number : ")
	_,e := fmt.Scanf("%d", &num)
	slice = append(slice,num)
	if e != nil {
		fmt.Println(e)		
	}
	}
	sorted_array := BubbleSort(slice);
	fmt.Println("===Sorted Array===")
	fmt.Println(sorted_array)
}