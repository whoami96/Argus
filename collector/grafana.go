package collector

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type GrafanaResponse struct {
	Data struct {
		Groups []struct {
			Rules []struct {
				Name        string `json:"name"`
				State       string `json:"state"`
				Annotations struct {
					Summary string `json:"summary"`
				} `json:"annotations"`
				Alerts []struct {
					State string `json:"state"`
				} `json:"alerts"`
			} `json:"rules"`
		} `json:"groups"`
	} `json:"data"`
}

func FetchGrafanaAlerts(url string, token string) ([]Alert, error) {
	APiUrl := fmt.Sprintf("%s/api/prometheus/grafana/api/v1/rules", strings.TrimRight(url, "/"))
	req, err := http.NewRequest("GET", APiUrl, nil)
	if err != nil {
		return nil, err
	}
	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var ApiResp GrafanaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ApiResp); err != nil {
		return nil, err
	}

	var alerts []Alert
	for _, group := range ApiResp.Data.Groups {
		for _, rule := range group.Rules {
			if rule.State == "firing" || rule.State == "pending" {
				count := len(rule.Alerts)
				if count == 0 {
					count = 1
				}

				alerts = append(alerts, Alert{
					Name:    rule.Name,
					State:   rule.State,
					Source:  "grafana",
					Count:   count,
					Summary: rule.Annotations.Summary,
				})
			}
		}
	}

	return alerts, nil
}
