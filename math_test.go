package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(11, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é invalido: Resultado %d. Esperado: %d", total, 30)
	}
}
