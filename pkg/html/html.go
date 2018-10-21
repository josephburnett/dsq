package html

import (
	"fmt"
	"html/template"
	"io"

	"github.com/josephburnett/dsq/pkg/types"
)

type HtmlBoard struct {
	*types.Board
}

type data struct {
	Board         *types.Board
	RenderedBoard template.HTML
	Message       []template.HTML
}

func (h *HtmlBoard) Render() template.HTML {
	board := "<table><tbody>"
	for y := 0; y < 9; y++ {
		board += "<tr>"
		for x := 0; x < 7; x++ {
			board += "<td>"
			board += h.RenderSquare(x, y)
			board += "</td>"
		}
		board += "</tr>"
	}
	board += "</tbody></table>"
	return template.HTML(board)
}

var startingPositions *types.Board

func init() {
	startingPositions = types.NewBoard()
}

func (h *HtmlBoard) RenderSquare(x, y int) string {
	squareWidth := float64(boardWidth)/7.0 + 4.0
	squareHeight := float64(boardHeight)/9.0 + 4.0
	p := h.Get(types.Point{x, y})
	ix := x
	iy := y
	switch p {
	case types.AMouse:
		ix = 0
		iy = 2
	case types.ACat:
		ix = 5
		iy = 1
	case types.AWolf:
		ix = 4
		iy = 2
	case types.ADog:
		ix = 1
		iy = 1
	case types.AHyena:
		ix = 2
		iy = 2
	case types.ATiger:
		ix = 6
		iy = 0
	case types.ALion:
		ix = 0
		iy = 0
	case types.AElephant:
		ix = 6
		iy = 2
	case types.BMouse:
		ix = 6
		iy = 6
	case types.BCat:
		ix = 1
		iy = 7
	case types.BWolf:
		ix = 2
		iy = 6
	case types.BDog:
		ix = 5
		iy = 7
	case types.BHyena:
		ix = 4
		iy = 6
	case types.BTiger:
		ix = 0
		iy = 8
	case types.BLion:
		ix = 6
		iy = 8
	case types.BElephant:
		ix = 0
		iy = 6
	case types.Empty:
		if startingPositions.Get(types.Point{x, y}) != types.Empty {
			ix = 3
			iy = 4
		}

	}
	offsetX := squareWidth * float64(ix)
	offsetY := squareHeight * float64(iy)
	square := fmt.Sprintf(
		`<span onclick="document.click('%v%v');"><div class="square" style="width:%vpx;height:%vpx;background-position:-%vpx -%vpx"><div></span>`,
		x, y, squareWidth, squareHeight, offsetX, offsetY)
	return square
}

func Render(w io.Writer, b *types.Board, msg []string) error {
	htmlMsg := make([]template.HTML, len(msg))
	for i, msg := range msg {
		htmlMsg[i] = template.HTML(msg)
	}
	d := data{
		Board:         b,
		RenderedBoard: (&HtmlBoard{b}).Render(),
		Message:       htmlMsg,
	}
	t, err := template.New("page").Parse(tmp)
	if err != nil {
		return err
	}
	err = t.Execute(w, d)
	return err
}

const tmp = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>Dou Shou Qi</title>
        <style>
            .square {
                background: url(images/board.png);
                display: inline-block;
            }
        </style>
        <script>

function log(msg) {
    var div = document.createElement("div");
    div.innerText = msg;
    document.body.appendChild(div);
}

// https://stackoverflow.com/questions/133925/javascript-post-request-like-a-form-submit
function post(params) {
    var form = document.createElement("form");
    form.setAttribute("method", "post");
    form.setAttribute("action", "/");
    for(var key in params) {
        if(params.hasOwnProperty(key)) {
            var hiddenField = document.createElement("input");
            hiddenField.setAttribute("type", "hidden");
            hiddenField.setAttribute("name", key);
            hiddenField.setAttribute("value", params[key]);
            form.appendChild(hiddenField);
        }
    }
    document.body.appendChild(form);
    form.submit();
}

document.board = {{ .Board.Marshal }}

document.click = function (square) {
    // Select
    if (!document.selected) {
        document.selected = square;
        log("Selected square " + square);
        return;
    }
    // Unselect
    if (document.selected == square) {
        delete(document.selected);
        log("Unselected square " + square);
        return;
    }
    // Move
    log("Moving from " + document.selected + " to " + square + "...");
    post({move: document.selected + square, board: document.board})
    delete(document.selected)
}

        </script>
    </head>
    <body style="font-family:monospace;">
        {{ .RenderedBoard }}
        <div style="margin-top:20px;"> </div>
        {{range .Message }}
        <div>{{ . }}</div>
        {{end}}
    </body>
</html>
`
