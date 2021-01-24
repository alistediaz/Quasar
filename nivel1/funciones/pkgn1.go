package funciones

import "math"

//GetLocation ...
func GetLocation(distances ...float32) (x, y float32) {
	var xr float32
	var yr float32

	xr = 0.0
	yr = 0.0

	if cap(distances) == 3 {
		powD1 := float32(math.Pow(float64(distances[0]), 2))
		powD2 := float32(math.Pow(float64(distances[1]), 2))
		powD3 := float32(math.Pow(float64(distances[2]), 2))
		xr = (780000 - 2*powD1 + 3*powD2 - powD3) / -1600
		yr = (-1200*xr - 270000 + powD1 - powD2) / 200
	}
	return xr, yr
}

//GetMessage ...
func GetMessage(messages ...[]string) (msg string) {
	var resp string = ""
	var idx1 int = -1
	var idx2 int = -1
	var idx3 int = -1

	if cap(messages) == 3 {
		for i := 0; i < cap(messages[0]); i++ {
			for j := 0; j < cap(messages[1]); j++ {
				if messages[0][i] == messages[1][j] && messages[0][i] != "" {
					idx1 = i
					idx2 = j
					break
				}
			}

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

		for i := 0; i < cap(messages[1]); i++ {
			for j := 0; j < cap(messages[2]); j++ {
				if messages[1][i] == messages[2][j] && messages[1][i] != "" {
					idx2 = i
					idx3 = j
					break
				}
			}
		}

		if idx1 == -1 && idx2 == -1 && idx3 == -1 {
			for i := 0; i < 3; i++ {
				for j := 0; j < cap(messages[i]); j++ {
					if messages[i][j] != "" {
						resp += messages[i][j] + " "
					}
				}
			}
		} else {
			for i, j, k := idx1, idx2, idx3; true; i, j, k = i-1, j-1, k-1 {
				if i < 0 && j < 0 && k < 0 {
					break
				} else {
					if mens(messages[0], i) != "" {
						resp = mens(messages[0], i) + " " + resp
					} else if mens(messages[1], j) != "" {
						resp = messages[1][j] + " " + resp
					} else if mens(messages[2], k) != "" {
						resp = messages[2][k] + " " + resp
					}
				}
			}
			for i, j, k := idx1+1, idx2+1, idx3+1; true; i, j, k = i+1, j+1, k+1 {
				if i > (cap(messages[0])-1) && j > (cap(messages[1])-1) && k > (cap(messages[2])-1) {
					break
				} else {
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
