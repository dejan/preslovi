package preslovi

import "fmt"

func ExampleLatinicom() {
	fmt.Println(Latinicom("Шабан Шаулић"))
	// Output: Šaban Šaulić
}

func ExampleLatinicom_second() {
	fmt.Println(Latinicom("ЊЕГОШ"))
	// Output: NJEGOŠ
}

func ExampleLatinicom_third() {
	fmt.Println(Latinicom("електродистрибуција"))
	// Output: elektrodistribucija
}

func ExampleLatinicom_fourth() {
	fmt.Println(Latinicom("Ђурђевак"))
	// Output: Đurđevak
}

func ExampleLatinicom_fifth() {
	fmt.Println(Latinicom("Latinica"))
	// Output: Latinica
}

func ExampleLatinicomAscii() {
	fmt.Println(LatinicomAscii("Шабан Шаулић"))
	// Output: Saban Saulic
}
