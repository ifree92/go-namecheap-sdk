package namecheap

import "fmt"

// GetDomains retrieves all the domains available on account.
func (c *Client) GetDomains() ([]Domain, error) {
	var domainsResponse DomainsResponse
	params := map[string]string{
		"Command": "namecheap.domains.getList",
	}
	req, err := c.NewRequest(params)
	if err != nil {
		return nil, err
	}
	resp, err := c.Http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = c.decode(resp.Body, &domainsResponse)
	if err != nil {
		return nil, err
	}
	if len(domainsResponse.Errors) > 0 {
		apiErr := domainsResponse.Errors[0]
		return nil, fmt.Errorf("%s (%s)", apiErr.Message, apiErr.Number)
	}
	return domainsResponse.CommandResponse.Domains, nil
}
