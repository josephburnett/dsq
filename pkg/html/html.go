package html

import (
	"html/template"
	"io"

	"github.com/josephburnett/dsq-golang/pkg/types"
)

func Render(w io.Writer, b *types.Board) error {
	t, err := template.New("page").Parse(tmp)
	if err != nil {
		return err
	}
	err = t.Execute(w, b)
	return err
}

const tmp = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Dou Shou Qi</title>
		<script>
			document.click = function (square) {
				// Select
				if (!document.selected) {
					document.selected = square;
					console.log("selected " + square);
					return;
				}
				// Unselect
				if (document.selected == square) {
					delete(document.selected);
					console.log("unselected " + square);
					return;
				}
				// Move
				console.log("moving from " + document.selected + " to " + square);
				// TODO: submit move
				delete(document.selected)
			}
		</script>
	</head>
	<body>
		<table style="border-style:solid;"><tbody>
			{{range $y, $row := . }}
			<tr>
				{{range $x, $cell := $row }}
				<td style="border-width:1px;border-style:solid;width:20px;height:20px;"
				    onclick="document.click('{{$x}}' + '{{$y}}');">
					{{ $cell }}
				</td>
				{{end}}
			</tr>
			{{end}}
		</tbody></table>
	</body>
</html>
`
