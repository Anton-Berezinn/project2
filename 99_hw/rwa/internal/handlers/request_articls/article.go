package request_articls

import (
	"encoding/json"
	"fmt"
	"io"
	"rwa/internal/model"
)

func ReadBody(r io.ReadCloser) (model.TestArticle, error) {
	u := &model.Article{}
	resp, err := io.ReadAll(r)
	if err != nil {
		return u.TestArticle, fmt.Errorf("error in read %w", err)
	}
	err = json.Unmarshal(resp, u)
	if err != nil {
		return u.TestArticle, fmt.Errorf("error in unmarshal %w", err)
	}
	return u.TestArticle, nil
}
