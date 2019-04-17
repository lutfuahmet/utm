package utm

import "net/url"

// full url = http://abc.com/?utm_source=source_test&utm_medium=medium_test&utm_campaign=campaign_test
// base url = http://abc.com

func Generate(base, source, medium, campaign string) (string, error) {

	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	query := u.Query()
	query.Set("utm_source", source)
	query.Set("utm_medium", medium)
	query.Set("utm_campaign", campaign)

	u.RawQuery = query.Encode()
	return u.String(), nil
}
