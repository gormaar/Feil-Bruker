package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"path"
	"html/template"


)


func TestHandler(t *testing.T)	{
	t.Run("", func(t *testing.T){
		r , err := http.NewRequest("", "localholst:8080/", nil)
			if err != nil	{
				t.Error("Error when setting up server: ", r)
			}
			testhandler := httptest.NewRecorder()
			handler := http.HandlerFunc(Handler)

			handler.ServeHTTP(testhandler, r)
			if status := testhandler.Code; status != http.StatusOK	{
				t.Error("Error in Handler")
			}
	})
}

func TestPath(t *testing.T)	{
	t.Run("", func(t *testing.T) {
		path := path.Join("template", "index.html")
		if path != "template/index.html" {
			t.Error("Error in path: ", path)
		}
		tmpl, err := template.ParseFiles(path)
		if err != nil {
			t.Error("Error in parsing files", tmpl)
		}
	})
}


