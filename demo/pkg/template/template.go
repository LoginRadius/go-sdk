package template

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func Render(w http.ResponseWriter, r *http.Request, name string,
	v map[string]interface{}) error {

	if v == nil {
		v = map[string]interface{}{}
	}

	cwd, _ := os.Getwd()

	files := []string{
		filepath.Join(cwd, "../../ui/html/base.html"),
		filepath.Join(cwd, "../../ui/html/"+name+".html"),
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		return err
	}

	err = ts.ExecuteTemplate(w, "base", v)

	return err
}
