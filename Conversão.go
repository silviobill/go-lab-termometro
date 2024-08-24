// declarar o pacote principal
package main

//importar a função fmt
import "fmt"

//declaraçõa da variável CONST do ponto de ebulição
const ebulicaoK = 373.0

//funcao principal
func main() {
	var tempK = ebulicaoK   //variável do valor da tempurado em F
	var tempC = tempK - 273 //conversão dos valores
	//apareça na tela do resultado
	fmt.Printf(`A temperatura de ebulição da água em Kelvin é %g°, enquanto em Celsius é %g°.`, tempK, tempC)
}
