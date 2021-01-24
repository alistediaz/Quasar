# Quasar
# Nivel 1
# Desde la función main se ejecutan las dos funciones requeridas:

func main() {
	x, y := funciones.GetLocation(500, 424.26, 707.106)
	fmt.Printf("%v %v\n", x, y)

	msg := funciones.GetMessage(
		[]string{"", "este", "es", "un", "mensaje"},
		[]string{"este", "", "un", "mensaje"},
		[]string{"", "", "un", "mensaje", ""})
	fmt.Printf("%v ", msg)
}
