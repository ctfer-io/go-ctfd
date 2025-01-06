package api

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

type GetFilesParams struct {
	Type     *string `schema:"type,omitempty"`
	Location *string `schema:"location,omitempty"`
	Q        *string `schema:"q,omitempty"`
	Field    *string `schema:"field,omitempty"`
}

func (client *Client) GetFiles(params *GetFilesParams, opts ...Option) ([]*File, error) {
	files := []*File{}
	if err := client.Get("/files", params, &files, opts...); err != nil {
		return nil, err
	}
	return files, nil
}

type PostFilesParams struct {
	Files     []*InputFile // XXX backend code shows it could be a list, but not the doc
	Challenge *int
	Location  *string
}

func (client *Client) PostFiles(params *PostFilesParams, opts ...Option) ([]*File, error) {
	// Maps parameters to values
	mp := map[string]any{
		"file":  params.Files,
		"nonce": client.nonce,
		"type":  "standard",
	}
	if params.Challenge != nil {
		mp["challenge"] = *params.Challenge
		mp["type"] = "challenge"
	}
	if params.Location != nil {
		mp["location"] = *params.Location
	}
	b, ct, err := encodeMultipart(mp)
	if err != nil {
		return nil, err
	}

	files := []*File{}

	// Process request directly, as it does not use the REST flow
	req, _ := http.NewRequest(http.MethodPost, "/files", b)
	req.Header.Set("Content-Type", ct)
	if err := client.Call(req, &files, opts...); err != nil {
		return nil, err
	}
	return files, nil
}

func (client *Client) GetFile(id string, opts ...Option) (*File, error) {
	file := &File{}
	if err := client.Get("/files/"+id, nil, &file, opts...); err != nil {
		return nil, err
	}
	return file, nil
}

func (client *Client) DeleteFile(id string, opts ...Option) error {
	return client.Delete("/files/"+id, nil, nil, opts...)
}

type InputFile struct {
	Name    string
	Content []byte
}

// GetFileContent is a helper leveraging the CTFd API that
// downloads a file's content given its location.
func (client *Client) GetFileContent(file *File, opts ...Option) ([]byte, error) {
	if file == nil {
		return nil, errors.New("can't get file from a nil value")
	}

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/files/%s", file.Location), nil)
	req = applyOpts(req, opts...)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}

// Returns a reader ready to be part of an HTTP request body.
func encodeMultipart(values map[string]any) (io.Reader, string, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	for k, v := range values {
		if v == nil {
			continue
		}

		var fw io.Writer
		var err error
		var content []byte
		// Avoid closer goleak
		if x, ok := v.(io.Closer); ok {
			defer x.Close()
		}
		// Add file content or field
		if file, ok := v.(*InputFile); ok {
			if file == nil {
				if fw, err = w.CreateFormFile(k, ""); err != nil {
					return nil, "", err
				}
				content = []byte{}
			} else {
				if fw, err = w.CreateFormFile(k, file.Name); err != nil {
					return nil, "", err
				}
				content = file.Content
			}
		} else if files, ok := v.([]*InputFile); ok {
			if files == nil {
				if fw, err = w.CreateFormFile(k, ""); err != nil {
					return nil, "", err
				}
				content = []byte{}
			} else {
				for _, file := range files {
					if fw, err = w.CreateFormFile(k, file.Name); err != nil {
						return nil, "", err
					}
					content = file.Content
				}
			}
		} else {
			if fw, err = w.CreateFormField(k); err != nil {
				return nil, "", err
			}
			content = getContent(v)
		}
		// Write value
		if _, err := fw.Write(content); err != nil {
			return nil, "", err
		}
	}
	w.Close()
	return &b, w.FormDataContentType(), nil
}

func getContent(v any) []byte {
	switch v := v.(type) {
	case *string:
		if v == nil {
			return []byte{}
		}
		return getContent(*v)
	case string:
		return []byte(v)
	case []byte:
		return v

	case *int:
		if v == nil {
			return []byte{}
		}
		return getContent(*v)
	case int:
		return []byte(strconv.Itoa(v))

	default:
		panic(fmt.Errorf("unhandled kind %T", v))
	}
}
