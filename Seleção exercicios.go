Seleção ex01.go

package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Digite o primeiro numero:")
	fmt.Scan(&n1)
	fmt.Print("Digite o segundo numero:")
	fmt.Scan(&n2)

	if n1 > n2 {
		fmt.Println("o maior numero é ", n1)
	} else {
		fmt.Println("o maior numero é ", n2)
	}

}



seleçao ex02.go

package main

import "fmt"

func main() {
	var n1, n2, n3 int
	fmt.Print("Indique um número: ")
	fmt.Scan(&n1)
	fmt.Print("Indique um número: ")
	fmt.Scan(&n2)
	fmt.Print("Indique um número: ")
	fmt.Scan(&n3)

	if n1 < n2 && n1 < n3 {
		fmt.Print("Seu menor número é: ", n1)
	}
	if n2 < n1 && n2 < n3 {
		fmt.Print("Seu menor número é: ", n2)
	}
	if n3 < n1 && n3 < n2 {
		fmt.Print("Seu menor número é: ", n3)
	}
}


seleçao ex03.go

package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Digite o numero:")
	fmt.Scan(&n1)

	if n1%2 == 0 {
		fmt.Println("o numero é ", n1, "par")
	} else {
		fmt.Println("o numero é ", n1, "impar")

	}
}


seleçao ex04.go

package main

import "fmt"

func main() {

	var peso, altura float64
	fmt.Print("peso:  ")
	fmt.Scan(&peso)
	fmt.Print("altura : ")
	fmt.Scan(&altura)

	IMC := peso / (altura * altura)
	fmt.Println("IMC ", IMC)
	if IMC > 18.5 {
		fmt.Print("esta abaixo do peso ideal ", IMC)

	} else if IMC >= 18.5 && IMC < 24 {
		fmt.Print("peso ideal ", IMC)

	} else if IMC >= 25 && IMC < 30 {
		fmt.Print("peso ideial ", IMC)

	}

}


seleçao ex05.go

main

import "fmt"

func main() {

	var n1 int
	fmt.Print("Digite o numero:")
	fmt.Scan(&n1)

	if n1%3 == 0 {
		fmt.Println("o numero é ", n1, "é divisor de 3")
	} else if n1%5 == 0 {
		fmt.Println("o numero é ", n1, "é divisor de 5")
	} else if n1%3 == 0 && n1%5 == 0 {
		fmt.Println("o numero é ", n1, "é divisor de 3 e 5")

	}
}



seleçao ex06.go

package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Qual é primeiro numero ? ")
	fmt.Scan(&n1)
	fmt.Print("Qual é segundo numero ? ")
	fmt.Scan(&n2)

	if n1 >= 0 && n2 >= 0 {

		multiplicaçao := n1 * n2
		fmt.Println("A multiplicaçao é ", multiplicaçao)

	} else if n1 < 0 || n2 < 0 {
		soma := n1 + n2
		fmt.Println("a soma é ", soma)
	}

}



seleçao ex07.go


package main

import "fmt"

func main() {
	var n1 float64
	fmt.Print("salario : ")
	fmt.Scan(&n1)

	if n1 < 1000 {

		salario := (n1 * (0.10)) + n1
		fmt.Println("novo salario: ", salario)

	} else if n1 >= 1000 {
		salario := (n1 * (0.05)) + n1
		fmt.Println("novo salario: ", salario)
	}

}


seleçao ex08.go


package main

import "fmt"

func main() {

	var n int
	fmt.Print("dia: ")
	fmt.Scan(&n)

	if n == 1 {

		fmt.Print("dom ")
	} else if n == 2 {

		fmt.Print("seg")
	} else if n == 3 {

		fmt.Print("ter")
	} else if n == 4 {

		fmt.Print("qua")
	} else if n == 5 {

		fmt.Print("qui")
	} else if n == 6 {

		fmt.Print("sex")
	} else if n == 7 {

		fmt.Print("sab")
	}
}




