package goctfd

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
)

type GetFilesParams struct {
	Type     *string `schema:"type,omitempty"`
	Location *string `schema:"location,omitempty"`
	Q        *string `schema:"q,omitempty"`
	Field    *string `schema:"field,omitempty"`
}

func (client *Client) GetFiles(params *GetFilesParams, opts ...option) ([]*File, error) {
	files := []*File{}
	if err := get(client, "/files", params, &files, opts...); err != nil {
		return nil, err
	}
	return files, nil
}

type PostFilesParams struct {
	File      *InputFile
	CSRF      string // XXX should not be part of the request
	Challenge int    // TODO May be additional i.e. pages don't need it
}

func (client *Client) PostFiles(params *PostFilesParams, opts ...option) ([]*File, error) {
	// Maps parameters to values
	b, ct, err := encodeMultipart(map[string]any{
		"file":      params.File,
		"nonce":     params.CSRF,
		"challenge": params.Challenge,
		"type":      "challenge",
	})
	if err != nil {
		return nil, err
	}

	files := []*File{}

	// Process request directly, as it does not use the REST flow
	req, _ := http.NewRequest(http.MethodPost, "/files", b)
	req.Header.Set("Content-Type", ct)
	if err := call(client, req, &files, opts...); err != nil {
		return nil, err
	}
	return files, nil
}

func (client *Client) GetFile(id string, opts ...option) (*File, error) {
	file := &File{}
	if err := get(client, "/files/"+id, nil, &file, opts...); err != nil {
		return nil, err
	}
	return file, nil
}

func (client *Client) DeleteFile(id string, opts ...option) error {
	return delete(client, "/files/"+id, nil, nil, opts...)
}

type InputFile struct {
	Name    string
	Content []byte
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
