package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"strings"
)

func prepareHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("X-Forwarded-Host"); h != "" {
			baseUrl, _ = url.Parse("http://" + h)
		} else {
			baseUrl, _ = url.Parse("http://" + r.Host)
		}
		fn(w, r)
	}
}

func myHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Fprintf(os.Stderr, "%+v", err)
				debug.PrintStack()
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		prepareHandler(fn)(w, r)
	}
}

func pathURIEscape(s string) string {
	return (&url.URL{Path: s}).String()
}

func notFound(w http.ResponseWriter) {
	code := http.StatusNotFound
	http.Error(w, http.StatusText(code), code)
}

func badRequest(w http.ResponseWriter) {
	code := http.StatusBadRequest
	http.Error(w, http.StatusText(code), code)
}

func forbidden(w http.ResponseWriter) {
	code := http.StatusForbidden
	http.Error(w, http.StatusText(code), code)
}

func panicIf(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func newHtmlify(w http.ResponseWriter, r *http.Request, content string, rep []string) string {
	if content == "" {
		return ""
	}
	replacer := strings.NewReplacer(rep...)
	content = replacer.Replace(content)

	return strings.Replace(content, "\n", "<br />\n", -1)
}
