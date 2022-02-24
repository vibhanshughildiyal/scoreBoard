package amqp

import (
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
	"reflect"
	scoreErr "scoreBoard/internal/errors"
	"time"
)

// Amqp Interface
type Amqp interface {
	createQueue() (string, *scoreErr.Errors)
	PublishMessage(payload []byte) *scoreErr.Errors
	ConsumeMessage(result chan []byte)
}

// Config - an object for storing amqp config details
type Config struct {
	Url         string
	ChannelName string
}

//amqpImpl - an amqp Object
type amqpImpl struct {
	conf      Config
	Channel   *amqp.Channel
	QueueName string
}

var configObject Config

func Init(conf Config) {
	configObject = conf
}

// New - provides a new amqp connection
func New(conf Config) Amqp {
	amqpObject := &amqpImpl{}
	if reflect.ValueOf(conf).IsZero() {
		amqpObject.conf = configObject
	} else {
		amqpObject.conf = conf
	}
	amqpObject.createQueue()
	return amqpObject
}

// CreateQueue - provides the connection to arango bloom database
func (am *amqpImpl) createQueue() (string, *scoreErr.Errors) {
	conn, err := amqp.Dial(am.conf.Url)
	if err != nil {
		return "", scoreErr.Error(scoreErr.Code("1.0"), err)
	}
	//defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		return "", scoreErr.Error(scoreErr.Code("1.0"), err)
	}
	//defer channel.Close()
	am.Channel = channel
	queue, err := channel.QueueDeclare(am.conf.ChannelName, false, false, false, false, nil)
	if err != nil {
		return "", scoreErr.Error(scoreErr.Code("1.0"), err)
	}
	am.QueueName = queue.Name
	return queue.Name, nil
}

// PublishMessage - publish message to queue
func (am *amqpImpl) PublishMessage(payload []byte) *scoreErr.Errors {
	if err := am.Channel.Publish("", am.QueueName, false, false, amqp.Publishing{
		ContentType:     "application/json", //"text/plain",
		ContentEncoding: "",
		DeliveryMode:    0,
		Priority:        0,
		CorrelationId:   "",
		ReplyTo:         "",
		Expiration:      "",
		MessageId:       "",
		Timestamp:       time.Now(),
		Type:            "",
		UserId:          "",
		AppId:           "",
		Body:            payload,
	}); err != nil {
		return scoreErr.Error(scoreErr.Code("1.0"), err)
	}
	return nil
}

func (am *amqpImpl) ConsumeMessage(result chan []byte) {
	msgs, err := am.Channel.Consume(am.QueueName, "", true, false, false, false, nil)
	if err != nil {
		scoreErr.Error(scoreErr.Code("1.0"), err)
	}
	forever := make(chan bool)
	for msg := range msgs {
		result <- msg.Body
	}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
