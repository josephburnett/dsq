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
    post({move: document.selected + square})
    delete(document.selected)
}

        </script>
    </head>
    <body style="font-family:monospace;">
        <div>+--+--+--+--+--+--+--+</div>
        {{range $y, $row := . }}<div>{{range $x, $square := $row }}|<span onclick="document.click('{{$x}}'+'{{$y}}');">{{ $square }}</span>{{end}}|</div>
        <div>+--+--+--+--+--+--+--+</div>
        {{end}}
    </body>
</html>
`
