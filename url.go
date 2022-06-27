package g

import "net/url"

func EncodeURL(s string) string {
	parsedURL, err := url.Parse(s)
	if err != nil {
		return err.Error()
	}

	parsedURL.RawQuery = parsedURL.Query().Encode()

	return parsedURL.String()
}
