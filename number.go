package main
import ("fmt"
	"time"
)
/*type book struct{
	title string
	price int
}
*/
func main(){
//	var a int=3
//	a :=[]int{0,1,2,3}
//	var b string="hello"
/*	var a int=3
	var b *int
	b=&a
	fmt.Println("address: %x\n",&a)
	fmt.Println(*b)
*/
//	fmt.Println(book{"tong",12})
/*
	var book1 book
	book1.title="tong"
	book1.price=12
	fmt.Println("book's price is %x",book1.price)
	*/
	/*
	c:=make(chan int,10)
	for i:=0;i<10;i=i+1{
		c<-i
	}
	close(c)
	/*
	for i :=range c{
		fmt.Println(i)
	}
	*/
	/*
	var n1 int=11
	test()
	fmt.Println(n1)
	*/
//	var n1 int=10
//	test(&n1)
/*
	var n3 int
	n3=Sum(3,5)
	fmt.Println(n3)
	*/
	/*
	sli:=[]int{1,2,3,4}
	for index,value:=range sli{
		fmt.Printf("index=%d,value=%d",index,value)
}
*/
/*
var a animal
d:=dog{sound:"wang"}
d.wang()
a=d
fmt.Println(a)
*/
/*
di:=map[string]int{
	"a":1,
	"b":2,
}
for index,value :=range di{
fmt.Printf("%v\n%v",index,value)
}
*/
//var a bool=true
//fmt.Println(a)
/*
func say(s string){
	for i:=0;i<5;i=i+1{
		time.Sleep(100 * time.Millisecond)
		fmt.Printn(s)
}
}
*/
func say(){
	fmt.Println("hello")
}
say()
//fmt.Println("hello")
//go say("hello")
//say("world")
}

func test(n1 *int){
	*n1=*n1+10
	fmt.Println(n1)
}
func Sum(n1 int,n2 int)(int){
	return n1+n2
}
type animal interface{
	wang()
}
type dog struct{
	sound string
}
func (d dog)wang(){
	fmt.Printf("%s\n",d.sound)
}
