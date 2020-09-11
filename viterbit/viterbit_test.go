package viterbit

import (
	"fmt"
	"testing"
)

func TestViterbit(t *testing.T) {
	observations := []string{"normal", "cold", "dizzy"}
	states := []string{"Healthy", "Fever"}
	startProbability := map[string]float64{"Healthy": 0.6, "Fever": 0.4}

	transitionProbability := map[string]map[string]float64{
		"Healthy": {
			"Healthy": 0.7,
			"Fever":   0.3,
		},
		"Fever": {
			"Healthy": 0.4,
			"Fever":   0.6,
		},
	}

	emissionProbability := map[string]map[string]float64{
		"Healthy": {
			"normal": 0.5,
			"cold":   0.4,
			"dizzy":  0.1,
		},
		"Fever": {
			"normal": 0.1,
			"cold":   0.3,
			"dizzy":  0.6,
		},
	}
	fmt.Println(Viterbit(observations, states, startProbability, transitionProbability, emissionProbability))
}
