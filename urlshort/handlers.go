package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//get url path
		path := r.URL.Path

		//check map for key
		if val, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, val, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var p2URLs []pathToURL

	err := yaml.Unmarshal(yml, &p2URLs)

	if err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string)

	for _, pu := range p2URLs {
		pathsToUrls[pu.Path] = pu.URL
	}

	return MapHandler(pathsToUrls, fallback), nil
}

type pathToURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
