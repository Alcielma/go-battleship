package design

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Row organiza widgets em uma linha, com espaçamento e altura opcional fixa
// <USAR A ALTURA DO PAI PARA ALINHAMENTO NO EIXO SECUNDARIO (Y)>
type Row struct {
	Pos        Point   //posição inicial
	Spacing    float32 //espaçamento horizontal entre elementos
	Children   []Widget
	MainAlign  Align // alinhamento dos elementos no eixo principal (center, start, end)
	CrossAlign Align //eixo cruzado
	size       Size  //para calculo de tamanho caso necessario
}

// NewRow cria uma linha e já calcula a posição de todos os widgets,
// alinhando verticalmente e no eixo secundario de acordo com o alinhamento dado
// construtor
func NewRow(
	pos Point,
	spacing float32,
	parentSize Size,
	mainAlign Align,
	crossAlign Align,
	children []Widget,
) *Row {

	r := &Row{
		Pos:        pos,
		Spacing:    spacing,
		Children:   children,
		MainAlign:  mainAlign,
		CrossAlign: crossAlign,
	}

	// 1) posiciona como Start/Start
	r.init()

	// 2) se ambos Start, não faz nada
	if mainAlign == Start && crossAlign == Start {
		return r
	}

	// 3) aplica alinhamentos relativos ao retângulo do pai iniciado em r.Pos
	if mainAlign != Start {
		r.alignMain(parentSize)
	}
	if crossAlign != Start {
		r.alignCross(parentSize)
	}

	return r
}

// posiciona filhos como Start/Start
func (r *Row) init() {
	cursorX := r.Pos.X

	for _, w := range r.Children {
		size := w.GetSize()

		w.SetPos(Point{
			X: cursorX,
			Y: r.Pos.Y,
		})

		cursorX += size.W + r.Spacing
	}
	r.calcSize()
}

// Update chama Update de todos os filhos
func (r *Row) Update() {
	for _, w := range r.Children {
		w.Update()
	}
}

// Draw chama Draw de todos os filhos
func (r *Row) Draw(screen *ebiten.Image) {
	for _, w := range r.Children {
		w.Draw(screen)
	}
}

// alinhamento no eixo principal (horizontal)
func (r *Row) alignMain(parentSize Size) {
	content := r.GetSize()

	var offsetX float32
	switch r.MainAlign {
	case Start:
		return
	case Center:
		offsetX = (parentSize.W - content.W) / 2
	case End:
		offsetX = parentSize.W - content.W
	}

	for _, w := range r.Children {
		p := w.GetPos()
		p.X += offsetX
		w.SetPos(p)
	}
}

// alinhamento no eixo cruzado (vertical)
func (r *Row) alignCross(parentSize Size) {
	for _, w := range r.Children {
		size := w.GetSize()
		p := w.GetPos()

		switch r.CrossAlign {
		case Start:
			continue
		case Center:
			p.Y = r.Pos.Y + (parentSize.H-size.H)/2
		case End:
			p.Y = r.Pos.Y + (parentSize.H - size.H)
		}

		w.SetPos(p)
	}
}

// tamanho ocupado pela row, calculado no init
func (r *Row) calcSize() {
	var totalW float32
	var maxH float32

	for i, w := range r.Children {
		s := w.GetSize()

		if i > 0 {
			totalW += r.Spacing
		}
		totalW += s.W

		if s.H > maxH {
			maxH = s.H
		}
	}

	r.size = Size{W: totalW, H: maxH}
}

func (r *Row) GetPos() Point {
	return r.Pos
}

func (r *Row) GetSize() Size {
	return r.size
}

func (r *Row) SetSize(size Size) {
}
func (r *Row) SetPos(point Point) {
}
