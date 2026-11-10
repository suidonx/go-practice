package main

import "fmt"

const (
	AppName = "MyApp"
	AppVersion = "1.0.0"
	MaxRetries = 3
)

const Pi = 3.14
const PiTyped float64 = 3.141592653589793
const Big = 1 << 100

func main() {
	fmt.Println("AppName:",AppName)
	fmt.Println("AppVersion:",AppVersion)
	fmt.Println("MaxRetries:",MaxRetries)

	// 型なし
	var f32 float32 = Pi
	var f64 float64 = Pi
	fmt.Printf("f32: %.6f, f64: %.6f\n", f32, f64)

	// 型付き
	var f32Typed float32 = float32(PiTyped)
	var f32Typed float32 = PiTyped
	var f64Typed float64 = PiTyped
	fmt.Printf("f32Typed: %.6f, f64Typed: %.6f\n", f32Typed, f64Typed)

	// 巨大な定数
	const Small = Big >> 99
	fmt.Println("Big >> 99 = %\n", Small)

	// intに代入できない
	var i int = Big
	fmt.Println("Big = %d\n", i)
}