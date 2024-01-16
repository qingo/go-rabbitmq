package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

// ExchangeOptions are used to configure an exchange.
// If the Passive flag is set the client will only check if the exchange exists on the server
// and that the settings match, no creation attempt will be made.
type ExchangeOptions struct {
	Name       string
	Kind       string // possible values: empty string for default exchange or direct, topic, fanout
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Passive    bool // if false, a missing exchange will be created on the server
	Args       Table
	Declare    bool
}

func getDefaultExchangeOptions() ExchangeOptions {
	return ExchangeOptions{
		Name:       "",
		Kind:       amqp.ExchangeDirect,
		Durable:    false,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Passive:    false,
		Args:       Table{},
		Declare:    false,
	}
}

func WithExchangeOptionsName(name string) func(*ExchangeOptions) {
	return func(options *ExchangeOptions) {
		options.Name = name
	}
}

func WithExchangeOptionsDurable(options *ExchangeOptions) {
	options.Durable = true
}

func WithExchangeOptionsKind(kind string) func(*ExchangeOptions) {
	return func(options *ExchangeOptions) {
		options.Kind = kind
	}
}

func WithExchangeOptionsAutoDelete(options *ExchangeOptions) {
	options.AutoDelete = true
}

// WithExchangeOptionsInternal ensures the exchange is an internal exchange
func WithExchangeOptionsInternal(options *ExchangeOptions) {
	options.Internal = true
}

// WithExchangeOptionsNoWait ensures the exchange is a no-wait exchange
func WithExchangeOptionsNoWait(options *ExchangeOptions) {
	options.NoWait = true
}

// WithExchangeOptionsDeclare stops this library from declaring the exchanges existance
func WithExchangeOptionsDeclare(options *ExchangeOptions) {
	options.Declare = true
}
