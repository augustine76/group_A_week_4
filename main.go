package main

import "fmt"

func main() {
	fmt.Println("Welcome to Group I's Week 4 Project!")
}

// Rohit Sthapit
// BubbleSort function sorts a slice of integers using the Bubble Sort algorithm
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

//Abhijith Augustine
//IsPrime function checks if a number is prime or not

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//Gul Mohammad Saifee
func calculateArea(length float64, width float64) float64 {
	return length * width
}

//Jagdish Sharma
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

//Anandu Prasad

func oddOrEven(num int) {
	if num%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}
}

// Jaskirat Singh
// Fibonacci Sequence
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
