package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processStack(e []string) float64 {
	//fmt.Printf("calculos ----> %s", e)
	result := 0.0
	fmt.Printf("valor slice a evaluar--> %s /n", e)
	for _, v := range e {
		//c := strings.Split(v, " ")
		fmt.Printf("datos--> %s", v)
		switch e[1] {
		case "*":
			num1, err := strconv.ParseFloat(e[0], 64)
			if err != nil {
				fmt.Println(err)
			}
			num2, err := strconv.ParseFloat(e[2], 64)
			if err != nil {
				fmt.Println(err)
			}
			result = num1 * num2
		case "/":
			num1, err := strconv.ParseFloat(e[0], 64)
			if err != nil {
				fmt.Println(err)
			}
			num2, err := strconv.ParseFloat(e[2], 64)
			if err != nil {
				fmt.Println(err)
			}
			if num2 == 0 {
				fmt.Println("error: you tried to divide by zero.")
				return 0.0
			}
			result = num1 / num2
		case "+":
			num1, err := strconv.ParseFloat(e[0], 64)
			if err != nil {
				fmt.Println(err)
			}
			num2, err := strconv.ParseFloat(e[2], 64)
			if err != nil {
				fmt.Println(err)
			}
			result = num1 + num2
		case "-":
			num1, err := strconv.ParseFloat(e[0], 64)
			if err != nil {
				fmt.Println(err)
			}
			num2, err := strconv.ParseFloat(e[2], 64)
			if err != nil {
				fmt.Println(err)
			}
			result = num1 - num2
		}
	}
	fmt.Printf("resultado--> %f \n", result)
	return result
}

func BuscaParentesis(e []string) ([]string, int, int) {
	//fmt.Print(e)
	hayPArentesis := false
	ParentesisIzq := "("
	ParentesisDer := ")"
	var i, j, ii, jj int

	for i, valor := range e {
		//fmt.Printf(" valor-->%s ParentesisDer-->%s valor i-->%d \n", valor, ParentesisDer, i)
		if valor == ParentesisDer {
			hayPArentesis = true
			ii = i
			fmt.Printf("el valor encontrado  %s   %d", valor, i) // El elemento fue encontrado en el slice
			break
		}

	}

	if hayPArentesis {
		for j := len(e) - 1; j >= 0; j-- {
			if e[j] == ParentesisIzq {
				println(j) // El elemento fue encontrado en el slice
				jj = j + 1
			}
		}
	}
	fmt.Printf("izqN -->%d derechaN -->%d", j, i)
	fmt.Printf("izq -->%d derecha -->%d", jj, ii)
	return e[jj:ii], jj + 1, ii

}
func ajustarSlice(e []string, inicio int, fin int) []string {
	// Calcular la longitud de la sección a quitar
	longitudSeccion := fin - inicio

	// Mover los elementos hacia la izquierda para cubrir la sección
	fmt.Printf("inicio %d Fin %d \n", inicio, fin)
	fmt.Printf("ajustarinicio %s \n", e[inicio-1:])
	fmt.Printf("ajustarFin %s \n", e[fin:])
	copy(e[inicio-1:], e[fin:])

	// Ajustar la longitud del slice
	//fmt.Printf(" \n valor antes return %s largo-->%d longitudSeccion-->%d", e, len(e), longitudSeccion)
	e = e[:len(e)-longitudSeccion]
	fmt.Printf("ajustarSlice e %s", e)
	return e
}

func Prioridad(e []string) (int, float64) {
	fmt.Printf("Prioridad a evaluar--> %s", e)
	miSlice := make([]string, 0)
	miSlice2 := make([]string, 0)
	res := 0.0
	i := 0
	for i, valor := range e {
		if valor == "/" {
			miSlice = append(miSlice, e[i-1], e[i], e[i+1])
			res = processStack(miSlice)
			println(res)
			if i == 1 {
				miSlice2 = append([]string{strconv.FormatFloat(res, 'f', 2, 64)}, e[i+2:]...)
				//fmt.Printf(" nuevo calculo---> %s,%d ,%s\n", miSlice2, len(miSlice2), e[i+2:])
			} else {
				miSlice2 = append(e[:i-1], []string{strconv.FormatFloat(res, 'f', 2, 64)}...)
				miSlice2 = append(miSlice2, e[i+2:]...)
			}
			if len(miSlice2) > 1 {
				res = processStack(miSlice2)
			}
		}
	}
	for i, valor := range e {
		if valor == "*" {
			miSlice = append(miSlice, e[i-1], e[i], e[i+1])
			res = processStack(miSlice)
			println(res)
			//fmt.Printf(" indices---> %d ,%d ,%d \n", i, i-1, i+1)
			//fmt.Printf(" valor a pegar y total---> %s ,%d,%d ,%s\n", miSlice[i+1:], len(e), len(miSlice), miSlice)
			if i == 1 {
				miSlice2 = append([]string{strconv.FormatFloat(res, 'f', 2, 64)}, e[i+2:]...)
				//fmt.Printf(" nuevo calculo---> %s,%d ,%s\n", miSlice2, len(miSlice2), e[i+2:])
			} else {
				miSlice2 = append(e[:i-1], []string{strconv.FormatFloat(res, 'f', 2, 64)}...)
				miSlice2 = append(miSlice2, e[i+2:]...)
			}
			if len(miSlice2) > 1 {
				res = processStack(miSlice2)
			}
		}
	}
	for i, valor := range e {
		if valor == "+" {
			miSlice = append(miSlice, e[i-1], e[i], e[i+1])
			res = processStack(miSlice)
			println(res)
			if i == 1 {
				miSlice2 = append([]string{strconv.FormatFloat(res, 'f', 2, 64)}, e[i+2:]...)
				//fmt.Printf(" nuevo calculo---> %s,%d ,%s\n", miSlice2, len(miSlice2), e[i+2:])
			} else {
				miSlice2 = append(e[:i-1], []string{strconv.FormatFloat(res, 'f', 2, 64)}...)
				miSlice2 = append(miSlice2, e[i+2:]...)
			}
			if len(miSlice2) > 1 {
				res = processStack(miSlice2)
			}
		}
	}
	for i, valor := range e {
		if valor == "-" {
			miSlice = append(miSlice, e[i-1], e[i], e[i+1])
			res = processStack(miSlice)
			println(res)
			if i == 1 {
				miSlice2 = append([]string{strconv.FormatFloat(res, 'f', 2, 64)}, e[i+2:]...)
				//fmt.Printf(" nuevo calculo---> %s,%d ,%s\n", miSlice2, len(miSlice2), e[i+2:])
			} else {
				miSlice2 = append(e[:i-1], []string{strconv.FormatFloat(res, 'f', 2, 64)}...)
				miSlice2 = append(miSlice2, e[i+2:]...)
			}
			if len(miSlice2) > 1 {
				res = processStack(miSlice2)
			}
		}
	}
	fmt.Printf(" res final---> %s,%d,resultado %f \n", miSlice2, len(miSlice2), res)
	return i, res
}
func EvaluaOperacion(e []string) {
	//fmt.Print("BuscaParentesis")
	//formulaOp := make([]string, 0)
	var inicio, fin int

	formulaOp, inicio, fin := BuscaParentesis(e)
	fmt.Printf("formula OP--> %s \n", formulaOp)
	indice, resultado := Prioridad(formulaOp)

	fmt.Printf("valores %d  %f\n", indice, resultado)
	//println(e)
	e = ajustarSlice(e, inicio, fin)
	fmt.Printf("slice ajustado %s ,inicio %d , fin %d \n", e, inicio, fin)
}

func main() {
	flag := false
	expressions := make([]string, 0)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("gocalc>")
		for scanner.Scan() {
			formula := strings.Split(scanner.Text(), " ")
			//for _, v := range formula {
			expressions = append(expressions, formula[:]...)
			//}

			//fmt.Print("EvaluaOperacion")
			EvaluaOperacion(expressions)
			if flag {
				res := processStack(expressions)
				fmt.Printf("final  %f \n", res)
			}

			//fmt.Println(expressions)
			fmt.Print("gocalc>")
		}
	}
}
