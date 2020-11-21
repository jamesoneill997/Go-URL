package urlshort

import "net/http"

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r http.Request) {
		p := r.URL.Path

		path, ok := pathsToUrls[p]

		if ok {
			http.Redirect(w, r, path, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL.
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var unmYAML []map[string]string

	e := yamlV2.Unmarshal(yaml, &unmYAML)

	//if yaml is of invalid format
	if e != nil {
		return nil, err
	}

	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func buildMap(parsedYaml []map[string]string) map[string]string {

	pathURLMap := make(map[string]string)

	for _, entry := range parsedYaml {
		key := entry["path"]
		pathUrlMap[key] = entry["url"]
	}
	return pathURLMap
}
