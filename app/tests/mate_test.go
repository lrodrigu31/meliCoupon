package testunitario

import "testing"

/*func TestSuma(t *testing.T) {
	total := Suma(5, 5)

	if total != 11 {
		t.Errorf("Suma incorrecta")
	}

}*/

func TestSuma(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)
		if total != item.c {
			t.Errorf("Suma incorrecta")
		}
	}
}
