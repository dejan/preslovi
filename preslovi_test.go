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

func ExampleLatinicom_digraph() {
	fmt.Println(Latinicom("ЉИЉА Љ. Љиљановић"))
	// Output: LJILJA Lj. Ljiljanović
}

func ExampleLatinicom_digraph2() {
	fmt.Println(Latinicom("ЂОРЂЕ Ђ. Ђорђевић"))
	// Output: ĐORĐE Đ. Đorđević
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

func ExampleLatinicomAscii_digraph2() {
	fmt.Println(LatinicomAscii("ЂОРЂЕ Ђ. Ђорђевић"))
	// Output: DJORDJE Dj. Djordjevic
}

func ExampleLatinicomAscii_digraph3() {
	fmt.Println(LatinicomAscii("ĐORĐE Đ. Đorđević"))
	// Output: DJORDJE Dj. Djordjevic
}
