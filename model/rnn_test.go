package model

import (
	"testing"
)

func TestCreateNetwork(t *testing.T) {
	nn := CreateNetwork(0.5, 10, 12, 20, 12)

	for i := 0; i < 4; i++ {
		if nn.Layers[i].Columns() != nn.Weights[i].Rows() {
			t.Errorf("Error with the creation of the RNN, the size of layers/weights is wrong.")
		}
	}
}