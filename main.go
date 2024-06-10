package main

import (
	"fmt"
	"log"
	"time"

	"ia_driver/database"
	"ia_driver/models"
	"ia_driver/utils"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sjwhitworth/golearn/base"
)

func main() {
	// Inicializar banco de dados
	db, err := database.InitDB("drivers.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Treinar o modelo com dados existentes
	dataset, err := models.TrainModel(db)
	if err != nil {
		log.Fatal(err)
	}
	trainData, testData := base.InstancesTrainTestSplit(dataset, 0.70)
	model := models.TrainAndEvaluateModel(trainData, testData)

	// Configurar o consumidor Kafka
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "vehicle-data-consumer-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	// Configurar o produtor Kafka
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	// Subscrever ao tópico
	consumer.SubscribeTopics([]string{"vehicle-data"}, nil)

	// Loop para consumir mensagens
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
			continue
		}

		// Processar a mensagem recebida
		processMessage(msg.Value, db, producer, model)
	}
}

func processMessage(value []byte, db *database.DB, producer *kafka.Producer, model *models.Model) {
	// Aqui você processa a mensagem recebida
	fmt.Printf("Received message: %s\n", string(value))

	// Exemplo de dados recebidos
	data := []struct {
		Position [2]float64
		Time     time.Time
	}{
		// Dados de exemplo
		{Position: [2]float64{0, 0}, Time: time.Now()},
		{Position: [2]float64{1, 1}, Time: time.Now().Add(1 * time.Second)},
		{Position: [2]float64{2, 2}, Time: time.Now().Add(2 * time.Second)},
	}

	var driverID string
	var prevVelocity float64
	for i := 1; i < len(data)-1; i++ {
		acceleration := utils.CalculateAcceleration(data[i-1].Position, data[i].Position, data[i-1].Time, data[i].Time)
		if i > 1 {
			braking := utils.CalculateBraking(prevVelocity, acceleration, data[i-1].Time, data[i].Time)
			isSharpTurn := utils.CalculateSharpTurn(data[i-1].Position, data[i].Position, data[i+1].Position)

			fmt.Printf("Acceleration: %.2f, Braking: %.2f, Sharp Turn: %t\n", acceleration, braking, isSharpTurn)

			// Identificar motorista com base nos padrões detectados usando o modelo treinado
			driverID = models.IdentifyDriver(model, acceleration, braking, isSharpTurn)
			if driverID != "" {
				break
			}
		}
		prevVelocity = acceleration
	}

	// Enviar o ID do motorista via Kafka se encontrado
	if driverID != "" {
		topic := "identified-driver"
		msg := kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(driverID),
		}
		producer.Produce(&msg, nil)
	}
}
