package rbmq

import (
	"errors"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	jsonType = "application/json"
)

type Client struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewClient(serverURL string) (*Client, error) {
	conn, err := amqp.Dial(serverURL)
	if err != nil {
		return nil, fmt.Errorf("new rabbitmq client - dial %s error: %w", serverURL, err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("new rabbitmq client - create channel error: %w", err)
	}

	c := &Client{
		conn:    conn,
		channel: ch,
	}

	return c, nil
}

func (c *Client) CreateQueue(queue string) error {
	if c.channel == nil {
		return errors.New("nil channel")
	}

	_, err := c.channel.QueueDeclare(
		queue, true, false, false, false, nil,
	)
	if err != nil {
		return fmt.Errorf("rbmq: declare queue %s error %w", queue, err)
	}

	return nil
}

func (c *Client) PublishJSONMessage(queue string, data []byte) error {
	if c.channel == nil {
		return errors.New("nil channel")
	}

	msg := amqp.Publishing{
		ContentType: jsonType,
		Body:        data,
	}
	if err := c.channel.Publish("", queue, false, false, msg); err != nil {
		return fmt.Errorf("publish JSON msg to %s error: %w", queue, err)
	}

	return nil
}

func (c *Client) GetConsumerChannel(queue string) (<-chan amqp.Delivery, error) {
	if c.channel == nil {
		return nil, errors.New("nil channel")
	}

	msg, err := c.channel.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		return nil, fmt.Errorf("consume queue %s error: %w", queue, err)
	}

	return msg, nil
}

func (c *Client) Close() error {
	if c.channel == nil {
		return errors.New("closing nil rbmq channel")
	}

	if err := c.channel.Close(); err != nil {
		return fmt.Errorf("closing rbmq channel error: %w", err)
	}

	c.channel = nil

	if c.conn == nil {
		return fmt.Errorf("closing nil rbmq connection")
	}

	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("closing rmbq connection error: %w", err)
	}

	return nil
}
