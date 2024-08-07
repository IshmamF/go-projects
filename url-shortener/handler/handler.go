package urlshort

import (
	"net/http"
	"gopkg.in/yaml.v3"
)

func redirectToSite(w http.ResponseWriter, r *http.Request, pathsToUrls map[string]string) bool{
	path := r.URL.Path
	if dest, ok := pathsToUrls[path]; ok {
		http.Redirect(w, r, dest, http.StatusSeeOther)
		return true
	}
	return false
}
// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if dontFallback := redirectToSite(w, r, pathsToUrls); dontFallback {
			return
		} else {
			fallback.ServeHTTP(w,r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	m := []T{}
	err := yaml.Unmarshal([]byte(yml), &m)
	paths := map[string]string{} 
	for _, path := range m {
		paths[path.Path] = path.URL
	}
	return MapHandler(paths, fallback), err
}

type T struct {
	Path string `yaml:"path"`
	URL string `yaml:"url"`
}
