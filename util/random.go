package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Devuelve un entero aleatorio entre min y max (inclusive)
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Genera una cadena aleatoria de longitud n utilizando el alfabeto proporcionado.
// La cadena generada puede contener cualquier combinación de letras minúsculas.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Genera un nombre de propietario aleatorio para una cuenta bancaria.
func RandomOwner() string {
	return RandomString(6)
}

// Genera un saldo aleatorio para una cuenta bancaria.
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Genera una moneda aleatoria para una cuenta bancaria.
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "ARS"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
