package queue

import (
	"fmt"
	"log"
	"reflect"
)

type QueueType int

const (
	RabbitMQ QueueType = iota
)

type Queue struct {
	qc QueueConnection
}

type QueueConnection interface {
	Publish([]byte) error
	Consume(chan<- QueueDto) error
}

func New(qt QueueType, cfg any) (q *Queue, err error) {
	rt := reflect.TypeOf(cfg)
	switch qt {
	case RabbitMQ:
		if rt.Name() != "RabbitMQConfig" {
			return nil, fmt.Errorf("config needs to be of type RabbitMQConfig")

		}
		conn, err := newRabbitConn(cfg.(RabbitMQConfig))
		if err != nil {
			return nil, err
		}
		q.qc = conn

	default:
		log.Fatal("type not expected")
	}
	return
}

func (q *Queue) Publish(msg []byte) error {
	return q.qc.Publish(msg)
}

func (q *Queue) Consume(chandto chan<- QueueDto) error {
	return q.qc.Consume(chandto)
}
