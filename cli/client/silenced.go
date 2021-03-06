package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path"

	"github.com/sensu/sensu-go/types"
)

// CreateSilenced creates a new silenced  entry.
func (client *RestClient) CreateSilenced(silenced *types.Silenced) error {
	b, err := json.Marshal(silenced)
	if err != nil {
		return err
	}

	res, err := client.R().SetBody(b).Post("/silenced")
	if err != nil {
		return err
	}

	if res.StatusCode() >= 400 {
		return UnmarshalError(res)
	}

	return nil
}

// DeleteSilenced deletes a silenced entry.
func (client *RestClient) DeleteSilenced(id string) error {
	res, err := client.R().Delete(fmt.Sprintf("/silenced/%s", url.PathEscape(id)))
	if err != nil {
		return err
	}
	if res.StatusCode() >= 400 {
		return UnmarshalError(res)
	}
	return nil
}

// ListSilenceds fetches all silenced entries from configured Sensu instance
func (client *RestClient) ListSilenceds(namespace, sub, check string) ([]types.Silenced, error) {
	if sub != "" && check != "" {
		name, err := types.SilencedName(sub, check)
		if err != nil {
			return nil, err
		}
		silenced, err := client.FetchSilenced(name)
		if err != nil {
			return nil, err
		}
		return []types.Silenced{*silenced}, nil
	}
	endpoint := "/silenced"
	if sub != "" {
		endpoint = path.Join(endpoint, "subscriptions", url.PathEscape(sub))
	} else if check != "" {
		endpoint = path.Join(endpoint, "checks", url.PathEscape(check))
	}
	resp, err := client.R().SetQueryParam("namespace", namespace).Get(endpoint)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, UnmarshalError(resp)
	}

	var result []types.Silenced
	err = json.Unmarshal(resp.Body(), &result)
	return result, err
}

// FetchSilenced fetches a silenced entry from configured Sensu instance
func (client *RestClient) FetchSilenced(name string) (*types.Silenced, error) {
	resp, err := client.R().Get(path.Join("/silenced", url.PathEscape(name)))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, UnmarshalError(resp)
	}
	var result types.Silenced
	return &result, json.Unmarshal(resp.Body(), &result)
}

// UpdateSilenced updates a silenced entry from configured Sensu instance
func (client *RestClient) UpdateSilenced(s *types.Silenced) error {
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}
	resp, err := client.R().SetBody(b).Put(path.Join("/silenced", url.PathEscape(s.Name)))
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 400 {
		return UnmarshalError(resp)
	}
	return nil
}
