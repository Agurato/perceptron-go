package main

import (
	"math"
	"testing"
)

func TestInitPerceptron(t *testing.T) {
	var p Perceptron

	p.InitPerceptron(10, []int{}, 10)
	if p.layerNb != 2 {
		t.Error(p.layerNb, "layers, expected 2")
	}

	p = Perceptron{}
	p.InitPerceptron(10, []int{10}, 10)

	for il, l := range p.layers {
		if l.position != p.layerNb-1 {
			for in, n := range l.neurons {
				for iw, w := range n.weights {
					if w <= 0 || w > 1/float64(l.size) {
						t.Error("layers[", il, "].neurons[", in, "].weights[", iw, "], got ", w, " | max = ", 1/float64(l.size))
					}
				}
			}
		}
	}
}

func TestAddLayer(t *testing.T) {
	var p Perceptron
	p.AddLayer(0)
	if p.layerNb != 0 {
		t.Error("p.layerNb =", p.layerNb, ", expected 0")
	}

	p.AddLayer(10)
	if p.layerNb != 1 {
		t.Error("p.layerNb =", p.layerNb, ", expected 1")
	}
	if p.layers[0].size != 10 {
		t.Error("p.layers[0].size =", p.layers[0].size, ", expected 10")
	}
	if p.layers[0].position != 0 {
		t.Error("p.layers[0].position =", p.layers[0].position, ", expected 0")
	}
	if len(p.layers[0].neurons) != 10 {
		t.Error("len(p.layers[0].neurons) =", len(p.layers[0].neurons), ", expected 10")
	}
}

func TestCalculateLayer(t *testing.T) {
	var p Perceptron

	p.AddLayer(2)
	p.AddLayer(2)

	p.layers[0].neurons[0].value = 0.2
	p.layers[0].neurons[0].weights = []float64{0.1, 0.5}
	p.layers[0].neurons[1].value = 0.7
	p.layers[0].neurons[1].weights = []float64{0.2, 0.6}

	p.CalculateLayer(1)

	if p.layers[1].neurons[0].value != 1/(1+math.Exp(-(0.2*0.1+0.7*0.2))) {
		t.Error("p.layers[1].neurons[0].value =", p.layers[1].neurons[0].value)
	}
	if p.layers[1].neurons[1].value != 1/(1+math.Exp(-(0.2*0.5+0.7*0.6))) {
		t.Error("p.layers[1].neurons[1].value =", p.layers[1].neurons[1].value)
	}
}

func TestComputeFromInput(t *testing.T) {
	var p Perceptron
	p.InitPerceptron(3, []int{3}, 3)

	p.layers[0].neurons[0].value = 0.1
	p.layers[0].neurons[1].value = 0.5
	p.layers[0].neurons[2].value = 0.9

	p.ComputeFromInput()

	if p.layers[2].neurons[0].value <= 0 || p.layers[2].neurons[0].value > 1 {
		t.Error("p.layers[2].neurons[0].value =", p.layers[2].neurons[0].value, ", expected in ]0, 1]")
	}
	if p.layers[2].neurons[1].value <= 0 || p.layers[2].neurons[1].value > 1 {
		t.Error("p.layers[2].neurons[1].value =", p.layers[2].neurons[1].value, ", expected in ]0, 1]")
	}
	if p.layers[2].neurons[2].value <= 0 || p.layers[2].neurons[2].value > 1 {
		t.Error("p.layers[2].neurons[2].value =", p.layers[2].neurons[2].value, ", expected in ]0, 1]")
	}
}