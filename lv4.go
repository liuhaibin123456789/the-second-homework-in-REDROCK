package main

import "fmt"

type sortMethod interface {
	sort() []int32
}
type MyType []int32
func (p MyType)sort() []int32 {
	for i := 0; i < len(p)-1; i++ {
		for j:=0;j<len(p)-1-i;j++ {
			if p[j+1]<p[j] {
				p[j],p[j+1]=p[j+1],p[j]
			}
		}
	}
	return p
}
func main() {
	myType := MyType{
		2, 1, 4, 5, 6, 7, 8,
	}

	fmt.Println(sortMethod.sort(myType))
}