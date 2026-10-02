package routes

import (
	"log"
	"net/http"
	"net/url"
)

const PatreonImageProxyRoute = "/patreon/"

func PatreonImageProxy(w http.ResponseWriter, r *http.Request) {
	patreonUrl := r.URL.Path[len(PatreonImageProxyRoute):]

	parsedURL, err := url.Parse(patreonUrl)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Write([]byte(parsedURL.String()))
}
