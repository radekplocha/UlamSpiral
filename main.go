package main

import (
	"fmt"
	"math"
	"strconv"
)

// PRZYPISANIE WARTOSCI KIERUNKOM UMOZLIWIAJACE LATWIEJSZE ICH POZNIEJSZE SPRAWDZANIE
const (
	RIGHT = iota 			// 0
	LEFT				// 1						
	UP				// 2
	DOWN				// 3
)

// FUNKCJA GENERUJĄCA SPIRALE N - WIELKOSC, PX PRZESUNIECIE PRAWO/LEWO, PY - PRZESUNIECIE GORA/DOL 
func generateUlumPrimes(n, px, py int) {
	
	// STWORZENIE PLANSZY JAKO SLICE
	board := make([][]string, n)
	for i := 0; i < n; i++ { 
		board[i] = make([]string, n) 
	}

	// PRZYPISANIE PIERWSZEMU KIEKUNKOWI RUCHU W PRAWO
	direction := RIGHT

	// OKRESLENIE SRODKA PLANSZY
	y := n / 2
	var x int
	// GDY PODAMY PARZYSTA WIELKOSC N PRZESUWAMY X JEDNA JEDNOSTKE W LEWO "WYSRODKOWANIE"
	if (n % 2 == 0) { 
		x = y - 1 
	} else {
		x = y 
	} 
		
	// PETLA ZAPISUJACA KOLEJNE WARTOSCI OD 1 DO N*N DO NASZEJ PLANSZY
	for j := 1; j <= n * n ; j++ {
		// SPRAWDZAMY CZY LICZBA JEST PIERWSZA, JEZELI NIE ZASTEPUJEMY JA MYSLNIKIEM
		if isPrime(j){
			board[y][x] = strconv.Itoa(j)
		} else {
			board[y][x] = "-"
		}
			
		// SPRAWDZAMY AKTUALNY KIERUNEK I GDY ZAJDZIE POTRZEBA ZMIENIAMY GO

		switch direction {
		case RIGHT : 
			if (board[y - 1][x] == "" && j > 1){ // JEZELI 1 MIEJSCE U GORY JEST WOLNE I J > 1 (CZYLI WSZYSTKIE POZA 1 RAZEM) 
				direction = UP 		     // TO ZMIENIAMY KIERUNEK NA GORA
			}
		case LEFT : 
			if (x == 0 || board[y + 1][x] == "") { // JEZELI 1 MIEJSCE W DOL JEST WOLNE ALBO X = 0 ( NA KRAWEDZI ) TO ZMIENIAMY KIERUNEK NA DOL 
				direction = DOWN 
			}
		case UP : 
			if (board[y][x - 1] == "") { // JEZELI 1 MIEJSCE W LEWO JEST WOLNE TO ZMIENIAMY KIERUNEK NA LEWO
				direction = LEFT 
			}
		case DOWN : 
			if (board[y][x + 1] == "") { // JEZELI 1 MIEJSCE W PRAWO JEST WOLNE TO ZMIENIAMY KIERUNEK NA PRAWO
				direction = RIGHT
		 	}
		}
		
		// AKTUALIZUJEMY X LUB Y W ZALEZNOSCI OD AKTUALNEGO KIERUNKU
		switch direction {
			case RIGHT : x += 1
			case LEFT : x -= 1
			case UP : y -= 1
			case DOWN : y += 1
		}
	}


	// ZNALEZIENIE DRUGI RAZ SRODKA PLANSZY, PONIEWAZ WARTOSCI X I Y ULEGLY ZMIANIE
	y = n / 2

	if (n % 2 == 0) { 
		x = y - 1 
	} else {
		x = y 
	}

	// PETLA WYSWIETLAJACA PLANSZE 20X10 Z PODANYM PRZESUNIECIEM PX, PY WZGLEDEM SRODKA
	for i := y - 5 + py ; i < y  + 5 + py ; i ++{
		for j := x - 10 - px  ; j < x  + 10 - px ; j ++{
			fmt.Printf("%5v", board[i][j])
		}
		fmt.Println()
	}


	// WYPISANIE CALEJ PLANSZY
	
	// for _, row := range board { 
	// 	fmt.Printf("%3v\n", row)
	// }


	fmt.Println()

}

// FUNKCJA GENERUJACA SPIRALE LICZB O PODANEJ WIELKOSCI N x N 
// ANALOGICZNIE DZIALA CO POPRZEDNIA Z WYJATKIEM SPRAWDZANIA CZY LICZBA JEST PIERWSZA, W TYM PRZYPADKU WYPISYWANIE SĄ 
// WSZYTKIE OD 1 DO N X N
func generateUlum(n, px, py int) {
	board := make([][]string, n)
	for i := 0; i < n; i++ { 
		board[i] = make([]string, n) 
	}
	direction := RIGHT
	y := n / 2
	var x int
	if (n % 2 == 0) { 
		x = y - 1 
	} else {
		x = y 
	}
 
	for j := 1; j <= n * n ; j++ {
		board[y][x] = strconv.Itoa(j)
		
		switch direction {
			case RIGHT : 
				if (x <= n - 1 && board[y - 1][x] == "" && j > 1){ 
					direction = UP 
				}
			case LEFT : 
				if (x == 0 || board[y + 1][x] == "") { 
					direction = DOWN 
				}
			case UP : 
				if (board[y][x - 1] == "") {
					direction = LEFT 
				}
			case DOWN : 
				if (board[y][x + 1] == "") { 
					direction = RIGHT
				}
		}
 
		switch direction {
			case RIGHT : x += 1
			case LEFT : x -= 1
			case UP : y -= 1
			case DOWN : y += 1
		}
	}

	y = n / 2

	if (n % 2 == 0) { 
		x = y - 1 
	} else {
		x = y 
	}

	for i := y - 5 + py ; i < y  + 5 + py ; i ++{
		for j := x - 10 - px  ; j < x  + 10 - px ; j ++{
			fmt.Printf("%5v", board[i][j])
		}
		fmt.Println()
	}

	// for _, row := range board { 
	// 	fmt.Printf("%3v\n", row)
	// }
	fmt.Println()
}
 


// FUNKCJA SPRAWDZAJĄCA CZY LICZBA JEST PIERWSZA
func isPrime(a int) bool {
	// 2 TO JEDYNA PARZYSTA LICZBA PIERSZA FUNKCJA ZWRACA TRUE
	if (a == 2) {
		return true 
	}
	// GDY PROGRAM DOSTANIE LICZBE 1 LUB MNIEJSZĄ LUB PARZYSTĄ (NIE LICZAC 2) ODRAZU ZWRACA FALSE
	if (a <= 1 || a % 2 == 0) { 
		return false 
	}
	// ZMIENNA KTORA OKRESLA MAKSYMALNA LICZBE DO JAKIEJ BEDA SPRAWDZANE DZIELNIKI PODANEJ LICZBY 
	max := int(math.Sqrt(float64(a)))

	// PETLA KTORA SPRAWDZA CZY PODANA LICZBA MA JAKIS DZIELNIK OD 3 DO PIERWIASTKA TEJ LICZBY ZAOKRAGLONEGO W DOL
	for n := 3; n <= max; n += 2 { 
		// JEZELI ZNAJDZIE TAKI DZIELNIK ZWRACA FALSE
		if (a % n == 0) {
			return false 
		} 
	}
	// FUNKCJA ZWRACA TRUE GDY ZADNA WARTOSC Z PETLI NIE SPELNIKA WARUNKU
	return true
}
 
func main() {
	generateUlum(30, 1, -3) // GENERUJE SPIRALE 30x30 I WYŚWIETJA JĄ 20x10 Z PRZESUNIĘCIEM 1 W PRAWO i 3 W DOŁ WZGLEGEM ŚRODKA
	generateUlumPrimes(30, 0, 0) // GENERUJE SPIRALE 30x30 I WYSWIETLA JĄ 20x10 BEZ PRZESUNIĘCIA ( WYSWIETLONE SA TYLKO LICZBY PIERWSZE )
}
