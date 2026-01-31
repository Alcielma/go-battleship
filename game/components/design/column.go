package design

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// TODO: Função de centralizar componentes internamente -> tipos de centralizar (center, end, start)
// TODO: Text: fontSize, font, caso necessário, tipos de center
// TODO: Container -> color, child, size, pos, center (com base no seu proprio size)

// Column organiza widgets em uma coluna, com espaçamento e largura opcional fixa
// <USAR A LARGURA DO PAI PARA ALINHAMENTO NO EIXO SECUNDARIO (X)>
type Column struct {
	Pos Point //posição inicial

	Spacing float32 //espaçamento vertical entre elementos

	cursorY float32 //posição vertical do próximo elemento

	MainAlign Align // alinhamento dos elementos no eixo principal (center, start, end)

	CrossAlign Align //eixo cruzado

	Children []Widget //elementos da column

	size Size //caso necessario
}

//TODO: depois será necessário modularizar melhor NewColumn e NewRow.

// NewColumn cria uma coluna e já calcula a posição de todos os widgets
// alinhando no eixo principal (vertical) e no eixo cruzado (horizontal)
func NewColumn(
	pos Point,
	spacing float32,
	parentSize Size,
	mainAlign Align,
	crossAlign Align,
	children []Widget,
) *Column {

	c := &Column{
		Pos:        pos,
		Spacing:    spacing,
		Children:   children,
		MainAlign:  mainAlign,
		CrossAlign: crossAlign,
	}

	// posicionamento inicial (Start/Start)
	c.init()

	// se ambos forem Start, não faz nada
	if mainAlign == Start && crossAlign == Start {
		return c
	}

	if mainAlign != Start {
		c.alignMain(parentSize)
	}
	if crossAlign != Start {
		c.alignCross(parentSize)
	}

	return c
}

// Update chama Update de todos os filhos
func (c *Column) Update() {
	for _, w := range c.Children {
		w.Update()
	}
}

// Draw chama Draw de todos os filhos
func (c *Column) Draw(screen *ebiten.Image) {
	for _, w := range c.Children {
		w.Draw(screen)
	}
}

func (c *Column) alignMain(parentSize Size) {
	content := c.GetSize()

	var offsetY float32
	switch c.MainAlign {
	case Start:
		return
	case Center:
		offsetY = (parentSize.H - content.H) / 2
	case End:
		offsetY = parentSize.H - content.H
	}

	for _, w := range c.Children {
		p := w.GetPos()
		p.Y += offsetY
		w.SetPos(p)
	}
}

func (c *Column) alignCross(parentSize Size) {
	for _, w := range c.Children {
		size := w.GetSize()
		p := w.GetPos()

		switch c.CrossAlign {
		case Start:
			continue
		case Center:
			p.X = c.Pos.X + (parentSize.W-size.W)/2
		case End:
			p.X = c.Pos.X + (parentSize.W - size.W)
		}

		w.SetPos(p)
	}
}

func (c *Column) GetPos() Point {
	return c.Pos
}

func (c *Column) SetPos(point Point) {
	c.Pos = point
}

// calcula tamanho (apenas no init)
func (c *Column) calcSize() {
	var totalH float32
	var maxW float32

	for i, w := range c.Children {
		s := w.GetSize()

		totalH += s.H
		if i < len(c.Children)-1 {
			totalH += c.Spacing
		}

		if s.W > maxW {
			maxW = s.W
		}
	}

	//calcula size uma unica vez aqui
	c.size = Size{W: maxW, H: totalH}
}

// GetSize retorna dimensões da column
func (c *Column) GetSize() Size {
	return c.size
}

// unnecessary?
func (c *Column) SetSize(size Size) {
	//criar logica de setsize para todos os elementos terem o size do eixo cruzado??
}

// init serve para um primeiro posicionamento dos elementos (start x start)
func (c *Column) init() {
	cursorY := c.Pos.Y

	for i, w := range c.Children {
		size := w.GetSize()

		w.SetPos(Point{
			X: c.Pos.X, // cross como Start
			Y: cursorY, // main sequencial
		})

		cursorY += size.H
		if i < len(c.Children)-1 {
			cursorY += c.Spacing
		}
	}
	c.calcSize()
}
