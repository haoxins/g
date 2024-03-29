package g

import "net/url"

func EncodeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	parsedURL.RawQuery = parsedURL.Query().Encode()

	return parsedURL.String(), nil
}

func GetURLHost(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return parsedURL.Host, nil
}

func GetURLPath(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return parsedURL.Path, nil
}
