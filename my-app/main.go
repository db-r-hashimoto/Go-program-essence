package main

import (
	"fmt"
	"reflect"
	"sort"
	"sync"

	"github.com/mattn/go-runewidth"
)

func test1() {
	var n int
	n = 1
	fmt.Println(n)

	// var a = int64
	x := 1
	y := 1.2

	// aはintなのでエラーになる
	// go build -buildvcs=false
	// # my-app
	// ./test1.go:10:10: int (type) is not an expression
	// a = 1 + ((float64(x) + 2) * float64(y))
	a := 1 + (float64(x) + 2) * float64(y)
	fmt.Println(a)
}

func test2() {
	const x = 1
	fmt.Println(x)

	y := 1

	f := 1.2 + (x + 2) * float64(y)

	fmt.Println(f)

	const n = 1
	a := 1 + n
	b := 1.2 + n
	fmt.Println(a)
	fmt.Println(b)
}

func test3() {
	// iotaに関して
	// https://qiita.com/curepine/items/2ae2f6504f0d28016411
	// 型なしの整数の連番、登場するたびに新しい宣言が実行される

	const (
		Apple = iota
		Orange
		Banana
	)

	fmt.Println(Apple) // 0
	fmt.Println(Orange) // 1
	fmt.Println(Banana) // 2

	const (
		Grape = iota + iota
		Cherry
		Peach
	)

	fmt.Println(Grape) // 0
	fmt.Println(Cherry) // 2
	fmt.Println(Peach) // 4

	const (
		a = iota + iota
		b
		c = iota + 3
	)

	fmt.Println(a) // 0
	fmt.Println(b) // 2
	fmt.Println(c) // 5

	const (
		zero  = iota != 1 // true
		one               // false
		two               // true
		three             // true
	)

	fmt.Println(zero)
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
	
	type Fruit int
	type Animal int

	const (
		f Fruit = iota
		r
		u
	)

	const (
		A Animal = iota
		N
		I
	)
	fmt.Println(f)
	fmt.Println(r)
	fmt.Println(u)
	fmt.Println(A)
	fmt.Println(N)
	fmt.Println(I)
}

// 関数の例
// Goではエラーが発生しうる関数の呼び出し時はまずerrをチェックして、即座に呼び出し元に返却することが推奨されている
// func FindUser(name String) (*User, error) {
// 	user, err = findUserFromList(users, name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

func test4() {
	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Print("two or ")
		fallthrough  // 明示的につ後にcaseに加工したい場合に使う、この場合は「two or three or four」と表示される
	case 3, 4:
		fmt.Println("three or four")
	default:
		fmt.Println("other")
	}
}

func test5() {
	// 配列とスライス
	// 配列は固定長、スライスは可変長と考えたらいい
	// どちらも0から始まる

	// 長さ4の配列の宣言
	var a [4]int
	a[0] = 1

	// スライスの宣言
	// 宣言時はbにはnilが入っていて、このまま要素にアクセスするとpanicになる
	var b []int
	fmt.Println(b)

	// 変数宣言もかねたスライスの作成
	c := make([]int, 3)
	// c := []int{1, 2, 3} という書き方もある
	c[0] = 1
	c[1] = 2
	c[2] = 3
	fmt.Println(c)

	// 固定長2次元の配列
	arr1 := [2][3]int {
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(arr1)

	// 2次元のスライス
	slc1 := [][]int {
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(slc1)
	slc1 = append(slc1, []int{7, 8, 9})
	fmt.Println(slc1)
	fmt.Printf("slc1 の長さは %d\n", len(slc1))

	// スライスには長さlenとは別でcap（容量）があり、makeで作成する際にlenとcapを指定できる
	// ただし、appendを繰り返すとメモリの再確保を毎回行うためにパフォーマンスが悪くなる
	// 事前に長さがわかるならlenとcapは最初に指定するのがいい
	slc2 := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		slc2 = append(slc2, i)
	}
	fmt.Println(slc2)

	// スライスは以下のように添え字で範囲指定して部分参照も可能
	fmt.Println(slc2[2:5]) // 2〜4番目の値が表示される


	// スライスから要素を削除する方法
	// 
	// 1. 新しくスライスを用意して詰め直す
	slc3 := make([]int, 0, len(slc2)/2)
	for i := 0; i < len(slc2); i++ {
		// 奇数を削除
		if i % 2 == 0 {
			slc3 = append(slc3, slc2[i])
		}
	}
	// slc2 = slc3  // この行でslc2は偶数のみ残る
	fmt.Println(slc3)

	// 2. appendを使う
	// l := 50
	// slc2 = append(slc2[:l], slc2[l + 1]...)

	// 3. 部分参照とcopyを使う（50だけ消える）
	l := 50
	slc2 = slc2[:l + copy(slc2[l:], slc2[l + 1:])]
	fmt.Println(slc2)
}

func test6() {
	name := "Taro"
	s := "Hello, "
	s += name

	fmt.Println(s)
	fmt.Printf("%c\n", s[0])

	// Helloをhelloに変える
	// GoのStringはイミュータブルなので、バイト列に一度変換して変更する
	b := []byte(s)
	b[0] = 'h'
	s = string(b)

	fmt.Println(s) // hello, Taroになる

	// runeを使ったStringの変更
	s2 := "こんにちわ世界"
	fmt.Println(s2)
	rs := []rune(s2)
	rs[4] = 'は'
	s2 = string(rs)
	fmt.Println(s2)
}

func test7() {
	// string型のキー、int型の値を持つmap
	// 宣言するだけではmにはnilが格納される
	// var m map[string]int

	// mapはmakeを使って作成する
	m := make(map[string]int)
	m["John"] = 21
	m["Bob"] = 18
	m["Mark"] = 33

	fmt.Println(m)

	// スライス同様capを指定できる
	m2 := make(map[string]int, 10)
	m2["John"] = 2
	m2["Bob"] = 19
	m2["Mark"] = 300
	fmt.Println(m2)

	// リテラルを使って初期値を再入することも可能
	m3 := map[string]int{
		"John": 20,
		"Bob": 10,
		"Mark": 30,
	}
	fmt.Println(m3)

	// mapの削除はdeleteを使う
	delete(m, "Bob")
	fmt.Println(m) // Bobが削除される

	// mapのキーと値を列挙する方法
	for k, v := range m2 {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

	// mapは順序を保持しないため、ソートしたい場合は先にキーを取り出してソートし、ソート後にfor-rangeを使う
	keys := []string{}
	for k := range m2 {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("key: %v, value: %v\n", k, m2[k])
	}

	// 存在しないキーを指定すると値の型のゼロ値が返る
	fmt.Println(m["zoo"]) // これは0が変える

	// キーの存在確認
	v, ok := m["John"]
	if ok {
		fmt.Println(v)
	}
}

func test8() {
	// typeで肩に名前をつけられる
	type MyString string

	var m MyString

	m = "foo"
	fmt.Println(m)
	s := string(m)
	fmt.Println(s)
}

type User struct {
	Name string
	Age int
	gender string  // 小文字からスタートすると他のパッケージから参照できない
}

func test9() {
	// 構造体

	var user User
	user.Name = "Bob"
	user.Age = 18
	user.gender = "male"

	fmt.Println(user)

	user2 := User{
		Name: "John",
		Age: 21,
	}
	fmt.Println(user2)

	user3 := User{
		Name: "Mark",
		Age: 33,
	}

	// 関数の引数としてstructを渡すと都度コピーされるため、コピーのオーバーヘッドをなくすためにポインタを使う手もある
	showName(&user3)
}

func showName(user *User) {
	fmt.Println(user.Name)
}


type Value1 int

func (v Value1) Add(n Value1) Value1 {
	return v + n
}

func method1() {
	v := Value1(1)
	v = v.Add(2)
	fmt.Println(v)
}

func test11() {
	v := 1
	p := &v   // &でポインタを得る
	*p = 2   // *で実態の参照ができる（実態の値を変更しているのでvの値は2になる）
	fmt.Println(v)

	// structの動的作成はnewを使う
	// この場合は、userは実はポインタ
	user := new(User)
	user.Name = "Bob"
	user.Age = 18
	fmt.Println(user)

	user2 := Bob()
	fmt.Println(user2)
}

func Bob() *User {
	user := User{
		Name: "Bob",
		Age: 18,
	}
	return &user
}

func test12() {
	var v interface{}
	v = 1
	n := v.(int)

	v = "こんにちは世界"
	s := v.(string)

	fmt.Println(n)
	fmt.Println(s)

	v = 300
	s, ok := v.(string)
	if !ok {
		fmt.Println("vはstringではない")
	} else {
		fmt.Printf("vはstring、v = %v \n", s)
	}

	PrintDetail(v)


	type V int
	var v1 V = 123
	PrintDetail(v1)
}

func PrintDetail(v interface{}) {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("int/int32/int64 型: ", v)
	case reflect.String:
		fmt.Println("string型: ", v)
	default:
		fmt.Println("知らない型")
	}
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age: age,
	}
}

type Speaker interface{
	Speak() error
}

type Dog struct {}
func (d *Dog) Speak() error {
	fmt.Println("BowWow")
	return nil
}

type Cat struct {}
func (c *Cat) Speak() error {
	fmt.Println("Meow")
	return nil
}

func DoSpeak(s Speaker) error {
	return s.Speak()
}

func test13() {
	dog := Dog{}
	DoSpeak(&dog)

	cat := Cat{}
	DoSpeak(&cat)
}

// deferで無名関数を呼び出すと、最後に無名関数が実行されるのでその時のnがキャプチャされるので2が出力される
func DoSomething1() {
	var n = 1
	defer func() {
		fmt.Println(n)
	}()
	n = 2
}

// deferで呼び出すと、最後にfmt.Printlnが実行されるが、呼び出し時のnが参照されるため1が出力される
func DoSomething2() {
	var n = 1
	defer fmt.Println(n)
	n = 2
}

func test14() {
	DoSomething1()
	DoSomething2()
}

func sendMessge(msg string) {
	println(msg)
}

func test15() {
	message := "Hello"
	sendMessge(message)
}

func server(ch chan string) {
	defer close(ch)
	ch <- "one"
	ch <- "two"
	ch <- "three"
}

func test16() {
	// var s string
	ch := make(chan string)
	go server(ch)

	// s = <- ch
	// fmt.Println(s)
	// s = <- ch
	// fmt.Println(s)
	// s = <- ch
	// fmt.Println(s)

	for st := range ch {
		fmt.Println(st)
	}
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
	test9()
	method1()
	test11()
	test12()
	test13()
	test14()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		test15()
	}()
	wg.Wait()
	test16()

	fmt.Println(runewidth.StringWidth("こんにちは"))
}