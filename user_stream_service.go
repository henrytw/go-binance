package binance

import (
	"context"
)

// StartUserStreamService create listen key for user stream service
type StartUserStreamService struct {
	c *Client
}

// Do send request
func (s *StartUserStreamService) Do(ctx context.Context, opts ...RequestOption) (listenKey string, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/api/v1/userDataStream",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}
	j, err := newJSON(data)
	if err != nil {
		return "", err
	}
	listenKey = j.Get("listenKey").MustString()
	return listenKey, nil
}

// KeepaliveUserStreamService update listen key
type KeepaliveUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *KeepaliveUserStreamService) ListenKey(listenKey string) *KeepaliveUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   "PUT",
		endpoint: "/api/v1/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// CloseUserStreamService delete listen key
type CloseUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *CloseUserStreamService) ListenKey(listenKey string) *CloseUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *CloseUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/api/v1/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// StartMarginUserStreamService create listen key for user stream service
type StartMarginUserStreamService struct {
	c *Client
}

// Do send request
func (s *StartMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (listenKey string, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/userDataStream",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}
	j, err := newJSON(data)
	if err != nil {
		return "", err
	}
	listenKey = j.Get("listenKey").MustString()
	return listenKey, nil
}

// KeepaliveMarginUserStreamService update listen key
type KeepaliveMarginUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *KeepaliveMarginUserStreamService) ListenKey(listenKey string) *KeepaliveMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   "PUT",
		endpoint: "/sapi/v1/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// CloseUserStreamService delete listen key
type CloseMarginUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *CloseMarginUserStreamService) ListenKey(listenKey string) *CloseMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *CloseMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/sapi/v1/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}
