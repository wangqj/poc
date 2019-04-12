package model

type CallOptions struct {
	ExactLabels bool
	Key         string
	Labels      Labels
}

type CallOption func(*CallOptions)

//WithExactLabels tell model service to return only one kv matches the labels
func WithExactLabels() CallOption {
	return func(o *CallOptions) {
		o.ExactLabels = true
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
