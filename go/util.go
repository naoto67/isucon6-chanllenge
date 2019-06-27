package main

import (
	"fmt"
	"html"
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
		panic(err)
	}
}

func newHtmlify(w http.ResponseWriter, r *http.Request, content string) string {
	if content == "" {
		return ""
	}
	content = html.EscapeString(content)
	var rep_data []string
	data, ok := repCache.Get("reps")
	if ok {
		rep_data = data.([]string)
	}
	replacer := strings.NewReplacer(rep_data...)
	content = replacer.Replace(content)

	return strings.Replace(content, "\n", "<br />\n", -1)
}
