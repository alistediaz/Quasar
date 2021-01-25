package funciones

import "math"

//GetLocation ...
func GetLocation(distances ...float32) (x, y float32) {
	var xr float32
	var yr float32

	xr = 0.0
	yr = 0.0

	//valida que ingresen las tres distancias
	if cap(distances) == 3 {
		//Fue necesario calcular una ecuación que permitiera calcular x a partir de las tres distancias
		//cada satelite genera un círculo que conecta con el punto de la posición de la nave, calculé la intersección de los tres círculos
		//y construí las ecuaciones:
		// x = (780000 - 2*powD1 + 3*powD2 - powD3) / -1600
		// y = (-1200*xr - 270000 + powD1 - powD2) / 200
		powD1 := float32(math.Pow(float64(distances[0]), 2))
		powD2 := float32(math.Pow(float64(distances[1]), 2))
		powD3 := float32(math.Pow(float64(distances[2]), 2))
		xr = (780000 - 2*powD1 + 3*powD2 - powD3) / -1600
		yr = (-1200*xr - 270000 + powD1 - powD2) / 200
	}
	return xr, yr
}

//GetMessage
//Esta función busca el punto de intersección entre los 3 mensajes y desde ahí recorre hacia la izquierda y luego a la derecha:
// Ej:
//                  intersección
// mensaje1: <------------|-->
// mensaje2:       <------|----->
// mensaje3:    <---------|>
func GetMessage(messages ...[]string) (msg string) {
	var resp string = ""
	var idx1 int = -1
	var idx2 int = -1
	var idx3 int = -1

	//valida los tres mensajes disponibles
	if cap(messages) == 3 {

		//Este loop intersecta el primer mensaje con el segundo en la palabra que coincidan. Almacena ambos índices.
		for i := 0; i < cap(messages[0]); i++ {
			for j := 0; j < cap(messages[1]); j++ {
				if messages[0][i] == messages[1][j] && messages[0][i] != "" {
					idx1 = i
					idx2 = j
					break
				}
			}

			//Si no hay intersección del primer mensaje con el segundo, intenta intersectar el primer mensaje con el tercero.
			if idx1 == -1 {
				for j := 0; j < cap(messages[2]); j++ {
					if messages[0][i] == messages[2][j] && messages[0][i] != "" {
						idx1 = i
						idx3 = j
						break
					}
				}
			}
		}

		//intersecta el segundo mensaje con el tercero.
		for i := 0; i < cap(messages[1]); i++ {
			for j := 0; j < cap(messages[2]); j++ {
				if messages[1][i] == messages[2][j] && messages[1][i] != "" {
					idx2 = i
					idx3 = j
					break
				}
			}
		}

		//Si no hay intersección, el mensaje de auxilio viene totalmente distribuido en los tres parámetros. Los concateno uno tras otro.
		if idx1 == -1 && idx2 == -1 && idx3 == -1 {
			for i := 0; i < 3; i++ {
				for j := 0; j < cap(messages[i]); j++ {
					if messages[i][j] != "" {
						resp += messages[i][j] + " "
					}
				}
			}
		} else {
			//Este loop alínea los tres mensajes en sus puntos de intersección y los recorre en reversa
			for i, j, k := idx1, idx2, idx3; true; i, j, k = i-1, j-1, k-1 {
				//si se llegó al inicio de los tres mensajes, se temina el ciclo.
				if i < 0 && j < 0 && k < 0 {
					break
				} else {
					//concatena las palabras que no sean vacías, si es vacía, consulta en el siguiente mensaje.
					if mens(messages[0], i) != "" {
						resp = mens(messages[0], i) + " " + resp
					} else if mens(messages[1], j) != "" {
						resp = messages[1][j] + " " + resp
					} else if mens(messages[2], k) != "" {
						resp = messages[2][k] + " " + resp
					}
				}
			}
			//Este loop alínea los tres mensajes en sus puntos de intersección y los recorre hacia adelante.
			for i, j, k := idx1+1, idx2+1, idx3+1; true; i, j, k = i+1, j+1, k+1 {
				//si se llegó al final de los tres mensajes, se temina el ciclo.
				if i > (cap(messages[0])-1) && j > (cap(messages[1])-1) && k > (cap(messages[2])-1) {
					break
				} else {
					//concatena las palabras que no sean vacías, si es vacía, consulta en el siguiente mensaje.
					if mens(messages[0], i) != "" {
						resp += mens(messages[0], i) + " "
					} else if mens(messages[1], j) != "" {
						resp += messages[1][j] + " "
					} else if mens(messages[2], k) != "" {
						resp += messages[2][k] + " "
					}
				}
			}
		}
	}

	return resp
}

//Esta función permite obtener las palabras sin caer en el error de "Fuera de rango"
func mens(pal []string, i int) (resp string) {
	if i < 0 {
		resp = ""
	} else if i > cap(pal)-1 {
		resp = ""
	} else {
		resp = pal[i]
	}

	return resp
}
