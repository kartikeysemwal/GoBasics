package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func add2(x int, y int) (int, int) {
	return x + y, x - y
}

func add4(x int, y int) (out1, out2 int) {
	out1 = x + y
	out2 = x - y
	return
}

func main() {
	// var num1 int = 1
	// num2 := 2

	// fmt.Println("Hello World")

	// fmt.Println(num1 + num2)

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// fmt.Println(add(2, 3))

	// sqrt := math.Sqrt(12)

	// fmt.Println(sqrt)
	// fmt.Printf("The output is %.2g \n", sqrt)

	// num := 2

	// if num < 5 {
	// 	fmt.Println("Hi")
	// } else {
	// 	fmt.Println("Hello")
	// }

	fmt.Println("Hello World")

	var num int = 20

	fmt.Println(num)
	fmt.Printf("%v %T\n", num, num)

	var intNum int = 42
	fmt.Printf("%v %T\n", intNum, intNum)

	var stringNum string
	stringNum = strconv.Itoa(intNum)
	fmt.Printf("%v %T\n", stringNum, stringNum)

	a := 3
	b := float32(10)
	fmt.Println(int(b) / a)

	var complexNum64 complex64 = 1 + 2i
	fmt.Printf("%v %T \n", real(complexNum64), real(complexNum64))
	fmt.Printf("%v %T \n", imag(complexNum64), imag(complexNum64))

	var complexNum128 complex128 = 1 + 2i
	fmt.Printf("%v %T \n", real(complexNum128), real(complexNum128))
	fmt.Printf("%v %T \n", imag(complexNum128), imag(complexNum128))

	var str string = "this is a string"

	fmt.Printf("%v %T \n", string(str[2]), str[2])

	var bytesArr = []byte(str)

	fmt.Printf("%v %T \n", bytesArr, bytesArr)

	var runeChar = 'r'
	fmt.Printf("%v %T \n", string(runeChar), runeChar)

	arr := [...]int{1, 2, 3}
	arrSlice := []int{1, 2, 3}
	fmt.Printf("%v %T \n", arr, arr)
	fmt.Printf("%v %T \n", arrSlice, arrSlice)

	arrMake := make([]int, 3, 100)
	fmt.Println(arrMake)
	fmt.Println("Length ", len(arrMake))
	fmt.Println("Capacity ", cap((arrMake)))

	arrMake = append(arrMake, 2)
	fmt.Println(arrMake)

	mapDataStructure := map[string]int{
		"a": 1,
		"b": 2, "c": 3, "d": 4}

	fmt.Println(mapDataStructure)

	pop, ok := mapDataStructure["f"]

	fmt.Println(pop, ok)

	if pop, ok := mapDataStructure["a"]; ok {
		fmt.Println(pop)
	}

	type doctorStruct struct {
		name string
		age  int
	}

	doctor := doctorStruct{
		name: "Hello",
		age:  10,
	}

	fmt.Println(doctor)

	doctorAnonymous := struct{ name string }{name: "test"}
	fmt.Println(doctorAnonymous)

	enumerateSlice := []int{1, 2, 3}

	for k, v := range enumerateSlice {
		fmt.Println(k, v)
	}

	sum(1, 2, 3, 4)

	var funcVar func(str string) = func(str string) {
		fmt.Println("Hello I am inside anonymous function. We have [" + str + "] as argument")
	}

	funcVar("test")

	greeterObj := greeterStruct{
		geeting: "hello",
		name:    "Ram",
	}

	greeterObj.greet()

	greet()

	var writer Writer = ConsoleWriter{}
	writer.Write([]byte{'a', 'b', 'c', 'd'})

	Write([]byte{'b', 'y', 't', 'e'})

	var intCounter IntCounter = IntCounter(1)
	var incrementor Incrementer = &intCounter
	incrementor.Increment()
	fmt.Println(intCounter)
	incrementor.Increment()
	fmt.Println(intCounter)

	var testInterface interface{} = 0

	switch testInterface.(type) {
	case int:
		fmt.Println("This is of type int")
	case float64:
		fmt.Println("This is of type float64")
	default:
		fmt.Println("This is default type")
	}

	var msgStr1 = "Hello"
	go func(msgStr2 string) {
		fmt.Println(msgStr2)
	}(msgStr1)

	msgStr1 = "Hello modified"

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	GoRoutine()

	WorkingWithChannels()

	printLogs()
}

func sum(values ...int) {
	for _, value := range values {
		fmt.Println(value)
	}
}

type greeterStruct struct {
	geeting string
	name    string
}

func (g greeterStruct) greet() {
	fmt.Println(g.geeting, g.name)
}

func greet() {
	fmt.Println("This is simple greet")
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, error := fmt.Println("Call in receiver write", string(data))
	return n, error
}

func Write(data []byte) (int, error) {
	n, error := fmt.Println("Call in simple write", string(data))
	return n, error
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (intCounter *IntCounter) Increment() int {
	*intCounter++
	return int(*intCounter)
}

type funcStruc struct {
}

// goroutines

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func GoRoutine() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	// m.RLock()
	fmt.Printf("Hello %v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	// m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}

// channels

var wgChannel = sync.WaitGroup{}

func WorkingWithChannels() {
	ch := make(chan int)
	// ch := make(chan int, 50) channel with 50 buffer
	wgChannel.Add(2)

	// go func(ch <-chan int) receive only channel
	go func() {
		i := <-ch
		fmt.Printf("Data received from channel %v\n", i)
		wgChannel.Done()
	}()

	// go func(ch chan<- int) send only channel
	go func() {
		ch <- 42
		wgChannel.Done()
	}()

	wgChannel.Wait()
}

// log channels

var logCh = make(chan string, 50)
var doneCh = make(chan struct{})

func printLogs() {
	go addLogs()

	logCh <- "Test log 1"
	logCh <- "Test log 2"

	time.Sleep(100 * time.Millisecond)

	// we can either write defer block to close the channel, or do it using an empty struct like doneCh
	doneCh <- struct{}{}
}

// func addLogs() {
// 	for entry := range logCh {
// 		fmt.Println(entry)
// 	}
// }

// logger with select
func addLogs() {
	for {
		select {
		case entry := <-logCh:
			fmt.Println(entry)
		case <-doneCh:
			break
			// default:
			// 	fmt.Println("Default block in addLogs")
		}
	}
}

// maps and slices when assigned from some other dataset point to the same dataset
// array and struct when assigned from some other dataset, they both have their own copy. To use the same use & in the begining

// golang don't support inheritance, rather it works on composition
// it uses embedded, and both of the struct are still independents

// defer -> panic -> recover

// stack (local memory), heap memory (shared memory)

// method: function that execute in contect of a type greet

// concrete types and behavioural types, concrete types such as struct, behavioural types such as interfaces

// if we want to create methods for struct we can do by doing receiver type, while in interfaces we can directly do it

// while other languages uses OS threads, Golang don't do it. Using goroutine it handles internally how it want to handle the threadpool

// if want to find race in the application run, go run <goFile>

// main is also a goroutine

// can close a channel if we are sure no data we will send later. Also after that there is no way to detect channel is closed or not from sender's side. Though from receiver side we can know by either applying forRange on the channel or ifOk
