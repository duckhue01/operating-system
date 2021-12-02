package main

import (
	"fmt"
	// "time"
	// "os"
	"syscall"
	// "os"
)

// Write a program that calls fork(). Before calling fork(), have the
// main process access a variable (e.g., x) and set its value to some-
// thing (e.g., 100). What value is the variable in the child process?
// What happens to the variable when both the child and parent change
// the value of x?

// func _1() {
// 	foo := 4

// 	id, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
// 	if id == 0 {
// 		foo = 100
// 		fmt.Println("In child:", id, foo, syscall.Getpid())

// 	} else {
// 		foo = 10000
// 		fmt.Println("In parent:", id, foo, syscall.Getpid())

// 	}

// }

// Write a program that opens a file (with the open() system call)
// and then calls fork() to create a new process. Can both the child
// and parent access the file descriptor returned by open()? What
// happens when they are writing to the file concurrently, i.e., at the
// same time?
// func _2() {

// 	syscall.Close(syscall.Stdout)
// 	syscall.Open("/home/duckhue01/coding-data/os/text.txt", syscall.O_RDWR, 0)
// 	id, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
// 	if id == 0 {
// 		fmt.Println("In child:")

// 	} else {
// 		fmt.Println("In parent:")
// 	}
// }

// Write another program using fork(). The child process should
// print “hello”; the parent process should print “goodbye”. You should
// try to ensure that the child process always prints first; can you do
// this without calling wait() in the parent?

// func _3() {

// 	id, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)

// 	if id == 0 {
// 		fmt.Println("parent")
// 	} else {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("children")
// 	}

// }

// Write a program that calls fork() and then calls some form of
// exec() to run the program /bin/ls. See if you can try all of the
// variants of exec(), including (on Linux) execl(), execle(),
// execlp(), execv(), execvp(), and execvpe(). Why do
// you think there are so many variants of the same basic call?

// func _4() {
// }

// Now write a program that uses wait() to wait for the child process
// to finish in the parent. What does wait() return? What happens if
// you use wait() in the child?

func _5() {

	id, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)

	if id == 0 {

		fmt.Println("children")

	} else {
		i1, i2, i3 := syscall.Syscall(syscall.SYS_WAIT4, 0, 0, 0)
		fmt.Println("parent", id, i1, i2, i3)
	}

}

// Write a slight modification of the previous program, this time us-
// ing waitpid() instead of wait(). When would waitpid() be
// useful?

// Write a program that creates a child process, and then in the child
// closes standard output (STDOUT FILENO). What happens if the child
// calls printf() to print some output after closing the descriptor?

// Write a program that creates two children, and connects the stan-
// dard output of one to the standard input of the other, using the
// pipe() system call.

func main() {
	_5()
}
