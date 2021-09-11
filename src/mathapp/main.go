package main
import(
    "mymath"
    "fmt"
    "time"
)
type Person struct{
    Name string 
    Age int
}
func main(){
    fmt.Println("<\tMODULE_TEST\t")
    fmt.Println("Hello,world. Sqrt(2)=%v\n",mymath.Sqrt(2))
    fmt.Println("\tARRAY_TEST\t")
    var arr_1 [5] int
    fmt.Println(arr_1)
    var arr_2=[5]int{1,2,3,4,5}
    fmt.Println(arr_2)
    arr_3:=[5] int{1,2,3,4,5}
    fmt.Println(arr_3)
    arr_4:=[...] int{1,2,3,4,6,7,}
    fmt.Println(arr_4)
    arr_5:=[5] int {0:3,1:5,4:6}
    fmt.Println(arr_5)
    fmt.Println("\t2D ARRAY\t")
    var arr_6 =[3][5] int {{1,2,3,4,5},{9,8,7,6,5},{3,4,5,6,7}}
    fmt.Println(arr_6)
    arr_7 := [3][5] int {{1,2,3,4,5},{9,8,7,6,5},{3,4,5,6,7}}
    fmt.Println(arr_7)
    arr_8 := [...][5] int {{1,2,3,4,5},{9,8,7,6,5},{0:3,1:5,4:6}}
    fmt.Println(arr_8)

    fmt.Println("\t\tSLICE\t\t")
    var sli_1 [] int
    fmt.Println("len=%d cap=%d slice=%v\n",len(sli_1),cap(sli_1),sli_1)
    var sli_2=[] int {1,2,3,4,5}
    fmt.Println("len=%d cap=%d slice=%v\n",len(sli_2),cap(sli_2),sli_2)

    fmt.Println("\t\tSTRUCT\t\t")
    var p1 Person
    p1.Name="Tom"
    p1.Age=30
    fmt.Println("p1=",p1)
    var p2 = Person{Name:"Burke",Age:31}
    fmt.Println("p2=",p2)
    p4:=struct{
        Name string 
        Age int
    }{Name:"匿名",Age:33}
    fmt.Println("p3=",p4)
    fmt.Println("\t\tMAP\t\t")
    var pTest1 map[int]string
    pTest1 = make(map[int]string)
    pTest1[1]="Tom"
    fmt.Println("p1:",pTest1)
    var pTest2 map[int]string=map[int]string{}
    pTest2[1]="Tom"
    fmt.Println("p2:",pTest2)
    var p3 map[int]string =make(map[int]string)
    p3[1]="Tom"
    fmt.Println("p3:",p3)
    fmt.Println("\t\tLOOP\t\t")
    person := [3] string{"Tom","Aaron","John"}
    fmt.Printf("len=%d cap=%d array=%v\n",len(person),cap(person),person)
    fmt.Println("")
    for i:=range person{
        fmt.Printf("person[%d]:%s\n",i,person[i])

    }
    fmt.Println("")
    for i:=0;i<len(person);i++{
        fmt.Printf("person[%d]:%s\n",i,person[i])
    }
    fmt.Println("")
    for _,name:=range person{
        fmt.Println("name:",name)
    }
    fmt.Println("\t\tFUNC\t\t")
    Hello()
    fmt.Println("\t\tCHAN\t\t")
    fmt.Println("----------------HELLO START------------------")
    go Hello()
    time.Sleep(1*time.Second)
    fmt.Println("----------------HELLO END--------------------")
    

}
func Hello(){
    fmt.Println("Hello GO!");
}
