package viterbit

// Viterbit Viterbit
func Viterbit(observations []string, states []string,
	startProbality map[string]float64, transformProbality,
	emissionProbability map[string]map[string]float64) []string {
	path := make(map[string][]string, len(states))
	for _, state := range states {
		path[state] = make([]string, len(observations))
	}

	curProbalities := make(map[string]float64)

	for _, state := range states {
		curProbalities[state] = startProbality[state] * emissionProbability[state][observations[0]]
	}

	for i := 1; i < len(observations); i++ {
		lastProbalities := curProbalities
		curProbalities = make(map[string]float64)
		for _, curState := range states {
			var maxProbality float64
			var state string
			for _, lastState := range states {
				probality := lastProbalities[lastState] * transformProbality[lastState][curState] * emissionProbability[curState][observations[i]]
				if probality > maxProbality {
					maxProbality = probality
					state = lastState
				}
			}
			curProbalities[curState] = maxProbality
			path[curState] = append(path[curState], state)
		}
	}
	var maxPath []string
	var maxProbality float64
	for _, state := range states {
		path[state] = append(path[state], state)
		if curProbalities[state] > maxProbality {
			maxPath = path[state]
			maxProbality = curProbalities[state]
		}
	}
	return maxPath
}
