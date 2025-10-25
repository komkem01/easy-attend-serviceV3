package config

import "github.com/easy-attend-serviceV3/internal/kafka/dto"

// Def.
var kafka = dto.Kafka{
	CaPath:   ``,
	CertPath: ``,
	KeyPath:  ``,
	Brokers:  `localhost:9092`,
	Topics: []string{
		TopicFileStatusUpdate,
	},
}

// TopicFileStatusUpdate is the Kafka topic for file status updates.
const (
	TopicFileStatusUpdate = "storage.file.status.update"
)
