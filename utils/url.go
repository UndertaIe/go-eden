package utils

import "net/url"

func UrlJoin(baseUrl string, joinUrl string) (string, error) {
	base, err := url.Parse(baseUrl)
	if err != nil {
		return "", nil
	}
	join, err := url.Parse(joinUrl)
	if err != nil {
		return "", nil
	}
	return base.ResolveReference(join).String(), nil
}
