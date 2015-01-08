// Código de ejmplo para el capítulo 1.2 de "Construye Aplicaciones Web con Golang"
// Propósito: Muestra como crear un paquete simple llamado `mymath`
// Este paquete debe ser importado desde otro archivo Go para poder ejecutarse.
package mymath

func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}
