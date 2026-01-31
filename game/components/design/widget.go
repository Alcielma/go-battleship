package design

import "github.com/hajimehoshi/ebiten/v2"

// Widget serve para facilitar posicionamento e tratamento de componentes
type Widget interface {

	// GetPos Retorna a posição e tamanho do widget
	GetPos() Point
	//SetPos para atualizar quando necessário
	SetPos(Point)
	GetSize() Size

	//Update Atualiza o widget//seus filhos no padrão ebiten
	Update()

	SetSize(Size)

	// Draw Desenha o widget/seus filhos
	Draw(screen *ebiten.Image)
}
