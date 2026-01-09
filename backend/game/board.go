package game

type Board struct {
	positions [10][10]Position;
}

func PlaceShip(b *Board, ship Ship, x int, y int) bool {
	if (!checkPosition(b, ship, x, y)) {
		return false;
	}

	if (isHorizontal(&ship)) {
		for i := y; i < y + ship.Size; i++ {
			placeShip(&b.positions[x][i], &ship);
		}
	} else {
		for i := x; i < x + ship.Size; i++ {
			placeShip(&b.positions[i][y], &ship);
		}
	}

	return true;

}

func checkPosition(b *Board, ship Ship, x int, y int) bool {
	if (isHorizontal(&ship)) { //se o navio estiver na horizontal:
		if (y + ship.Size > 10) { // verifica se o navio ultrapassa os limites do tabuleiro
			return false;
		}

		for i := y; i < y + ship.Size; i++ { //se a posição não está bloqueada
			if (!isValidPosition(b.positions[x][i])) {
				return false;
			}
		}
	} else { // se o navio estiver na vertical:
		if (ship.Size + x > 10) {
			return false;
		}

		for i := x; i < x + ship.Size; i++ {
			if (!isValidPosition(b.positions[i][y])) {
				return false;
			}
		}
	}
	// se todas as verificações passarem, a posição é válida
	return true; 
}

func PrintBoard(b *Board) {
	for i :=0; i<10; i++ { // itera pelas linhas
		for j:=0; j<10; j++ { // itera pelas colunas
			if (isAttacked(b.positions[i][j])) { // se a posição foi atacada
				if (getShipReference(b.positions[i][j]) != nil) {
					print("x "); // posição atacada com navio
					continue;
				}

				print("o "); // posição atacada sem navio
				continue;
			} else if (getShipReference(b.positions[i][j]) != nil) {
				print("B "); // marca como bloqueada.
				continue;
			}

			//posição valida e não atacada.
			print("- ");
		}
		print("\n"); // nova linha apos cada linha do tabuleiro

	}
}