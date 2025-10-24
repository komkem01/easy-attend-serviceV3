package dto

type Kafka struct {
	CaPath   string
	CertPath string
	KeyPath  string
	Brokers  string
	Section  string
	Topics   []string
}
