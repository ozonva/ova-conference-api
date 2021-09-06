package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"ova-conference-api/internal/configs"
)

type Producer interface {
	Send(message ConferenceChangedMessage) error
	Close() error
}

type producer struct {
	syncProducer sarama.SyncProducer
	topic        string
}

func NewProducer(configuration configs.KafkaConfiguration) (Producer, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(configuration.Brokers, saramaConfig)
	if err != nil {
		return nil, err
	}

	return &producer{
		topic:        configuration.Topic,
		syncProducer: syncProducer,
	}, nil
}

func (p *producer) Send(message ConferenceChangedMessage) error {
	jsonMes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, _, err = p.syncProducer.SendMessage(
		&sarama.ProducerMessage{
			Topic:     p.topic,
			Partition: -1,
			Key:       sarama.StringEncoder(p.topic),
			Value:     sarama.StringEncoder(jsonMes),
		})
	return err
}

func (p *producer) Close() error {
	return p.syncProducer.Close()
}
