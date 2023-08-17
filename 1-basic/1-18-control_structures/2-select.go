/*
	  	select 与 switch 非常相似, select语句只能用于信道的读写操作, 语法:
			select {
			case communication clause  :
				statement(s);
			case communication clause  :
				statement(s);
			// 你可以定义任意数量的 case
			default : // 可选
				statement(s)
			}

基本概念理解:

	a.通道: Go 协程之间交互的一种机制,
	b.数据传输：通道是用来在不同协程之间传递数据的管道。一个协程（goroutine）可以把数据发送到通道，而另一个协程可以从通道接收这些数据。这种机制
	  确保了数据的有序传递和同步。
	c.阻塞与非阻塞：当从通道接收数据时，如果通道为空，接收操作会阻塞，直到有数据可以接收。类似地，当向通道发送数据时，如果通道已满，发送操作会阻塞，
	  直到有空间可以发送。这有助于实现协程之间的同步。
	d.同步：通道可以用于协调多个协程的执行。一个协程可以等待另一个协程发送数据到通道，从而实现同步操作，确保在特定条件满足之前等待。
	e.安全性：由于通道的操作是原子的，通道可以用来实现多个协程之间的数据共享和同步，避免了数据竞争和并发问题。
	f.关闭通道：通道可以被显式地关闭，以通知接收方没有更多的数据会到达。接收方可以通过多重赋值操作判断通道是否关闭，从而避免在关闭通道后继续接收数据。

博客资料:

	· https://www.jianshu.com/p/de4bc02e7c72
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建无缓冲通道
	ch1 := make(chan int)
	ch2 := make(chan int)

	// selectCase1(ch1, ch2)
	// selectCase2(ch1, ch2)
	// selectCase3(ch1, ch2)
	selectCase4(ch1, ch2)
}

func selectCase1(ch1 chan int, ch2 chan int) {

	// 启动一个协程来发送数据到通道
	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	// 启动一个协程将值"3"发送到通道 ch2
	go func() {
		ch2 <- 3
	}()

	// 协程的启动需要一定的耗时, 当执行到select的时候数据还没有准备好就会走的default
	time.Sleep(time.Second * 3)

	// select随机从通道接收数据
	select {
	case i := <-ch1: // 不带判断条件,而是一个通信操作(i1 = <-ch1 >>>> 接受 c1 通道的值并赋给变量 i1)
		fmt.Printf("从ch1读取了数据%d\n", i)
	case j := <-ch2:
		fmt.Printf("从ch2读取了数据%d\n", j)
	default:
		// 当代码块有 default 语句的时候, 不会被阻塞, 反之则会被阻塞
		fmt.Println("我是默认操作")
	}
}

// select 的使用场景——竞争选举
func selectCase2(ch1 chan int, ch2 chan int) {
	select {
	case i := <-ch1:
		fmt.Printf("从ch1读取了数据%d", i)
	case j := <-ch2:
		fmt.Printf("从ch2读取了数据%d", j)
	}
}

// select 的使用场景——超时处理
func selectCase3(ch1 chan int, ch2 chan int) {
	select {
	case str := <-ch1:
		fmt.Println("receive str", str)
	case <-time.After(time.Second * 5): // 如果超过5秒没有响应就执行此操作
		fmt.Println("timeout!!")
	}
}

// select 的使用场景——判断buffered channel是否阻塞
// 这个例子很经典，比如我们有一个有限的资源（这里用buffer channel实现），我们每一秒向bufChan传送数据，
// 由于生产者的生产速度大于消费者的消费速度，故会触发default语句，这个就很像我们web端来显示并发过高的提示了，
// 小伙伴们可以尝试删除go func中的time.Sleep(5*time.Second)，看看是否还会触发default语句
func selectCase4(ch1 chan int, ch2 chan int) {
	// 创建有限资源的通道
	bufChan := make(chan int, 5)

	// 启动协程
	go func() {
		time.Sleep(time.Second)
		for {
			<-bufChan
			time.Sleep(5 * time.Second)
		}
	}()

	for {
		select {
		case bufChan <- 1:
			// 生产速度大于消费速度
			fmt.Println("add success")
			time.Sleep(time.Second)
		default:
			fmt.Println("资源已满，请稍后再试")
			time.Sleep(time.Second)
		}
	}
}

// select 的使用场景——阻塞main函数
func selectCase5(ch1 chan int, ch2 chan int) {
	bufChan := make(chan int)

	go func() {
		for {
			bufChan <- 1
			time.Sleep(time.Second)
		}
	}()
	// 这里要注意上面一定要有一直活动的goroutine,否则会报deadlock。大家还可以把select{}换成for{}试一下，打开系统管理器看下CPU的占用变化
	go func() {
		for {
			fmt.Println(<-bufChan)
		}
	}()

	// 有时候我们会让main函数阻塞不退出，如http服务，我们会使用空的select{}来阻塞main goroutine
	select {}
}
