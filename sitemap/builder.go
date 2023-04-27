package sitemap

import (
	"gophercises-link/linkparser"
	"net/http"
	"net/url"
)

func BuildFromUrlStr(siteUrlStr string) ([]*url.URL, error) {
	siteUrl, err := url.Parse(siteUrlStr)
	if err != nil {
		return nil, err
	}

	ret := []*url.URL{}
	queue := []*url.URL{siteUrl}

	// for BFS tracking use string type for efficiency
	visited := make(map[string]bool)
	for len(queue) != 0 {
		currPageUrl := queue[0]
		queue = queue[1:]

		currPageUrlStr := currPageUrl.String()
		if _, ok := visited[currPageUrlStr]; ok {
			continue
		}

		ret = append(ret, currPageUrl)
		visited[currPageUrlStr] = true
		resp, err := http.Get(currPageUrlStr)
		if err != nil {
			return nil, err
		}

		links, err := linkparser.ParseFromReader(resp.Body)
		if err != nil {
			return nil, err
		}

		queue, err = processLinksInBFS(queue, links, currPageUrl)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func processLinksInBFS(queue []*url.URL, links []linkparser.Link, currPageUrl *url.URL) ([]*url.URL, error) {
	for _, link := range links {
		parsedUrl, err := url.Parse(link.Href)
		if err != nil {
			return nil, err
		}

		if !parsedUrl.IsAbs() {
			queue = append(queue, currPageUrl.ResolveReference(parsedUrl))
			continue
		}

		if parsedUrl.Hostname() == currPageUrl.Hostname() {
			queue = append(queue, parsedUrl)
			continue
		}
	}

	return queue, nil
}
