package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	m := genStateModel()
	seq := m.generateSequence()
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	pre := seq[0]
	for sc.Scan() {
		ob := sc.Text()
		next := m.nextHiddenState(pre, ob)
		seq = append(seq, next)
		pre = next
	}
	printElectedSeq(seq)
}

func printElectedSeq(seq []hiddenState) {
	length := len(seq)
	var (
		lastRate float64
		last     string
	)
	for s, i := range seq[length-1] {
		if i.rate >= lastRate {
			lastRate = i.rate
			last = s
		}
	}
	// follow parents
	for i := length; i > 0; i-- {
		fmt.Printf("step: %d: state:%s\n", i-1, last)
		last = seq[i-1][last].parent
	}
}

type model struct {
	states   map[string]struct{}
	observes map[string]struct{}

	starts                map[string]float64
	transitionProbability map[string]map[string]float64
	emissionProbability   map[string]map[string]float64
}

var errInvalid = errors.New("invalid")

func newModel() *model {
	return &model{
		states:   make(map[string]struct{}),
		observes: make(map[string]struct{}),
		starts:   make(map[string]float64),

		transitionProbability: make(map[string]map[string]float64),
		emissionProbability:   make(map[string]map[string]float64),
	}
}

type hiddenState map[string]hiddenInfo

type hiddenInfo struct {
	rate   float64
	parent string
}

func (m *model) generateSequence() []hiddenState {
	seq := make([]hiddenState, 1)
	initState := make(hiddenState, len(m.starts))
	for s, r := range m.starts {
		initState[s] = hiddenInfo{rate: r}
	}
	seq[0] = initState
	return seq
}

func (m *model) registerState(s string, rate float64) error {
	if _, ok := m.states[s]; ok {
		return errInvalid
	}
	m.states[s] = struct{}{}
	m.starts[s] = rate
	return nil
}
func (m *model) registerObserves(o string) error {
	if _, ok := m.observes[o]; ok {
		return errInvalid
	}
	m.observes[o] = struct{}{}
	return nil
}
func (m *model) setTransitionProbability(p, n string, rate float64) error {
	if _, ok := m.states[p]; !ok {
		return errInvalid
	}
	if _, ok := m.states[n]; !ok {
		return errInvalid
	}
	if _, ok := m.transitionProbability[p]; !ok {
		m.transitionProbability[p] = map[string]float64{}
	}
	m.transitionProbability[p][n] = rate
	return nil
}

func (m *model) setEmissionProbability(s, o string, rate float64) error {
	if _, ok := m.states[s]; !ok {
		return errInvalid
	}
	if _, ok := m.observes[o]; !ok {
		return errInvalid
	}
	if _, ok := m.emissionProbability[s]; !ok {
		m.emissionProbability[s] = map[string]float64{}
	}
	m.emissionProbability[s][o] = rate
	return nil
}

func (m *model) nextHiddenState(preHidden hiddenState, o string) hiddenState {
	nextState := make(hiddenState, len(m.states))
	for s := range m.states {
		nextInfo := hiddenInfo{}
		for pre, info := range preHidden {
			probability := info.rate * m.coocurrenceProbability(pre, s, o)
			if probability >= nextInfo.rate {
				nextInfo.rate = probability
				nextInfo.parent = pre
			}
		}
		nextState[s] = nextInfo
	}
	// for canonicalize
	sum := 0.0
	for _, info := range nextState {
		sum += info.rate
	}
	for s := range nextState {
		info := nextState[s]
		info.rate = info.rate / sum
		nextState[s] = info
	}
	return nextState
}

func (m *model) coocurrenceProbability(p, n, o string) float64 {
	return m.transitionRate(p, n) * m.emissionRate(n, o)
}

func (m *model) transitionRate(p, n string) float64 {
	if _, ok := m.transitionProbability[p]; !ok {
		return 0.0
	}
	return m.transitionProbability[p][n]
}

func (m *model) emissionRate(s, o string) float64 {
	if _, ok := m.emissionProbability[s]; !ok {
		return 0.0
	}
	return m.emissionProbability[s][o]
}

func genStateModel() *model {
	m := newModel()
	for _, p := range []struct {
		state string
		rate  float64
	}{
		{"Rainy", 0.3},
		{"Sunny", 0.4},
		{"Cloudy", 0.3},
	} {
		m.registerState(p.state, p.rate)
	}
	for _, ob := range []string{"sleep", "game", "eat"} {
		m.registerObserves(ob)
	}
	for _, p := range []struct {
		pre, next string
		rate      float64
	}{
		{"Rainy", "Rainy", 0.4},
		{"Rainy", "Sunny", 0.3},
		{"Rainy", "Cloudy", 0.3},
		{"Sunny", "Rainy", 0.2},
		{"Sunny", "Sunny", 0.7},
		{"Sunny", "Cloudy", 0.1},
		{"Cloudy", "Rainy", 0.4},
		{"Cloudy", "Sunny", 0.1},
		{"Cloudy", "Cloudy", 0.5},
	} {
		m.setTransitionProbability(p.pre, p.next, p.rate)
	}
	for _, p := range []struct {
		state, obs string
		rate       float64
	}{
		{"Rainy", "sleep", 0.5},
		{"Rainy", "game", 0.4},
		{"Rainy", "eat", 0.1},
		{"Sunny", "sleep", 0.2},
		{"Sunny", "game", 0.7},
		{"Sunny", "eat", 0.1},
		{"Cloudy", "sleep", 0.2},
		{"Cloudy", "game", 0.2},
		{"Cloudy", "eat", 0.6},
	} {
		m.setEmissionProbability(p.state, p.obs, p.rate)
	}
	return m
}
