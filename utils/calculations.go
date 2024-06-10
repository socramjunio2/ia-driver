package utils

import (
	"math"
	"time"
)

// CalculateAcceleration calcula a aceleração com base nas posições e tempos fornecidos
func CalculateAcceleration(prevPos, currPos [2]float64, prevTime, currTime time.Time) float64 {
	deltaX := currPos[0] - prevPos[0]
	deltaY := currPos[1] - prevPos[1]
	distance := math.Sqrt(deltaX*deltaX + deltaY*deltaY)
	deltaTime := currTime.Sub(prevTime).Seconds()
	velocity := distance / deltaTime

	return velocity / deltaTime // aceleração
}

// CalculateBraking calcula a desaceleração (frenagem) com base nas velocidades e tempos fornecidos
func CalculateBraking(prevVelocity, currVelocity float64, prevTime, currTime time.Time) float64 {
	deltaTime := currTime.Sub(prevTime).Seconds()
	deceleration := (currVelocity - prevVelocity) / deltaTime

	return deceleration // desaceleração
}

// CalculateSharpTurn determina se há uma curva acentuada com base nas posições fornecidas
func CalculateSharpTurn(prevPos, currPos, nextPos [2]float64) bool {
	vec1X := currPos[0] - prevPos[0]
	vec1Y := currPos[1] - prevPos[1]
	vec2X := nextPos[0] - currPos[0]
	vec2Y := nextPos[1] - currPos[1]

	dotProduct := vec1X*vec2X + vec1Y*vec2Y
	mag1 := math.Sqrt(vec1X*vec1X + vec1Y*vec1Y)
	mag2 := math.Sqrt(vec2X*vec2X + vec2Y*vec2Y)
	angle := math.Acos(dotProduct / (mag1 * mag2))

	return angle > math.Pi/4 // Exemplo para detectar curvas acentuadas (ângulo > 45 graus)
}
