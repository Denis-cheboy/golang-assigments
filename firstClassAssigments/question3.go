package main 
import (
"fmt"
"math"
)

func main(){
	// Assigment One=> Question 3
	var b byte=math.MaxUint8 //Byte is an alias for Uint8 so its maximum is 255
	var smallI int32=math.MaxInt32
	var bigI uint64=math.MaxUint64
    

	fmt.Printf("Before adding 1:\n")
	fmt.Printf("i: %v, %T\n", b, b)
	fmt.Printf("smallI: %v, %T\n", smallI, smallI)
	fmt.Printf("bigI: %v, %T\n", bigI, bigI)
 
	b+=1
	smallI+=1
	bigI+=1

    fmt.Printf("\nAfter adding 1:\n")
	fmt.Printf("i: %v, %T\n", b, b)
	fmt.Printf("smallI: %v, %T\n", smallI, smallI)
	fmt.Printf("bigI: %v, %T\n", bigI, bigI)
	
}

