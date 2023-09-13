package cyj

type Options interface {
	apply(*Client) *Client
}

type OptionsFunc func(*Client)

func (f OptionsFunc) apply(o *Client) *Client {
	f(o)
	return o
}
func Host(host string) Options {
	return OptionsFunc(func(o *Client) {
		o.host = host
	})
}
func ContentKey(contentKey string) Options {
	return OptionsFunc(func(o *Client) {
		o.contentKey = contentKey
	})
}
func ParamKey(paramKey string) Options {
	return OptionsFunc(func(o *Client) {
		o.paramKey = paramKey
	})
}

func Token(token string) Options {
	return OptionsFunc(func(o *Client) {
		o.token = token
	})
}
