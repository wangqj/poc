package model

type CallOptions struct {
	ExactOne bool
	Key      string
	Labels   Labels
}

type CallOption func(*CallOptions)

//WithExactOne tell model service to return only one kv matches the labels
func WithExactOne() CallOption {
	return func(o *CallOptions) {
		o.ExactOne = true
	}
}

//WithKey find by key
func WithKey(key string) CallOption {
	return func(o *CallOptions) {
		o.Key = key
	}
}

//WithLabels find kv by labels
func WithLabels(labels Labels) CallOption {
	return func(o *CallOptions) {
		o.Labels = labels
	}
}
