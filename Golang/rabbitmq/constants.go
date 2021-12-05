package rabbitmq

const (
	urlTemplate       = "amqp://%s:%s@%s/"
	contentType       = "text/plain"
	exchangeType      = "fanout"
	prefix            = "RabbitMq_"
	Url               = prefix + "Url"
	Username          = prefix + "Username"
	Password          = prefix + "Password"
	PublishExchange   = prefix + "PublishExchange"
	SubscribeExchange = prefix + "SubscribeExchange"
	PublishQueue      = prefix + "PublishQueue"
	SubscribeQueue    = prefix + "SubscribeQueue"
)
