package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ExportRawParams struct {
	// Type could be "csv" (which returns the table dump in csv) or
	// anything (returns a zip).
	Type string `json:"type"`

	// Args contains all the arguments to export specific content.
	Args ExportRawArgsParams `json:"args"`
}

type ExportRawArgsParams struct {
	// Table, if specified, exports the given table in the specified type.
	Table *string `json:"table,omitempty"`
}

func (client *Client) ExportRaw(params *ExportRawParams, opts ...Option) ([]byte, error) {
	if params == nil {
		params = &ExportRawParams{}
	}
	i, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	// Build request and execute it manunally
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/exports/raw", bytes.NewBuffer(i))
	req.Header.Set("Content-Type", "application/json")
	req, client.sub.Transport = applyOpts(req, opts...)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("CTFd responded with unexpected status code: got %d", res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
