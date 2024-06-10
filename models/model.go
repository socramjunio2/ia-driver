package models

import (
	"ia_driver/database"
	"log"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

// Model estrutura para armazenar o classificador treinado
type Model struct {
	Classifier *knn.KNNClassifier
}

// TrainAndEvaluateModel treina e avalia um modelo KNN e retorna o modelo treinado
func TrainAndEvaluateModel(trainData, testData *base.InstancesView) *Model {
	if trainData == nil || testData == nil {
		log.Println("Train and test data not provided, skipping model training and evaluation")
		return nil
	}

	knn := knn.NewKnnClassifier("euclidean", "linear", 2)
	knn.Fit(trainData)

	predictions, err := knn.Predict(testData)
	if err != nil {
		log.Fatal(err)
	}

	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(evaluation.GetSummary(confusionMat))

	return &Model{Classifier: knn}
}

// IdentifyDriver usa o modelo treinado para identificar um motorista com base nos padrões detectados e retorna um ID único
func IdentifyDriver(model *Model, acceleration, braking float64, isSharpTurn bool) string {
	// Converter os dados de entrada em uma instância para previsão
	attrs := base.NewAttributes()
	attrs.Add(base.NewFloatAttribute("acceleration"))
	attrs.Add(base.NewFloatAttribute("braking"))
	attrs.Add(base.NewCategoricalAttribute("sharp_turn", "true", "false"))

	instance := base.NewDenseInstances()
	instance.AddAttributes(attrs...)

	sharpTurnStr := "false"
	if isSharpTurn {
		sharpTurnStr = "true"
	}

	instance.Add(base.NewDenseInstance().SetAll([]float64{
		acceleration,
		braking,
		float64(base.CatToFloat(attrs.Get("sharp_turn"), sharpTurnStr)),
	}))

	// Fazer a previsão usando o modelo treinado
	predictions, err := model.Classifier.Predict(instance)
	if err != nil {
		log.Fatal(err)
	}

	// Extrair o ID do motorista previsto
	driverID := base.GetClass(predictions, 0)
	return driverID
}

// TrainModel treina o modelo com os dados do banco de dados
func TrainModel(db *database.DB) (*base.Instances, error) {
	// Carregar dados do banco de dados
	rows, err := db.Query("SELECT acceleration, braking, sharp_turn, id FROM drivers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Criar dataset GoLearn
	attrs := base.NewAttributes()
	attrs.Add(base.NewFloatAttribute("acceleration"))
	attrs.Add(base.NewFloatAttribute("braking"))
	attrs.Add(base.NewCategoricalAttribute("sharp_turn", "true", "false"))
	attrs.Add(base.NewCategoricalAttribute("id"))

	dataset := base.NewDenseInstances()
	dataset.AddAttributes(attrs...)

	for rows.Next() {
		var acceleration float64
		var braking float64
		var sharpTurn bool
		var id int
		if err := rows.Scan(&acceleration, &braking, &sharpTurn, &id); err != nil {
			return nil, err
		}
		sharpTurnStr := "false"
		if sharpTurn {
			sharpTurnStr = "true"
		}

		dataset.Add(base.NewDenseInstance().SetAll([]float64{
			acceleration,
			braking,
			float64(base.CatToFloat(attrs.Get("sharp_turn"), sharpTurnStr)),
			float64(id),
		}))
	}

	return dataset, nil
}
