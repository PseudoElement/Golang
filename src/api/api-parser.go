package api_main

import "net/http"

func MapQueryParams(req *http.Request, queryParams ...string) map[string]string {
	mapppedParams := make(map[string]string)

	for _, param := range queryParams {
		mapppedParams[param] = req.URL.Query().Get(param);
	}

	return mapppedParams;
}