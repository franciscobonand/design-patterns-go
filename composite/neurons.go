package main

type NeuronInterface interface {
    Iter() []*Neuron
}

// Neuron is a scalar object that represents one neuron
type Neuron struct {
    In, Out []*Neuron
}

func (n *Neuron) ConnectTo(other *Neuron) {
    n.Out = append(n.Out, other)
    other.In = append(other.In, n)
}

func (n *Neuron) Iter() []*Neuron {
    return []*Neuron{n}
}

// NeuronLayer is a composite object that represents a set of neurons
type NeuronLayer struct {
    Neurons []Neuron
}

func (nl *NeuronLayer) Iter() []*Neuron {
    result := make([]*Neuron, 0)
    for i := range nl.Neurons {
        result = append(result, &nl.Neurons[i])
    }
    return result
}

func NewNeuronLayer(count int) *NeuronLayer {
    return &NeuronLayer{make([]Neuron, count)}
}


// Connect method works for both Neuron and NeuronLayer
func Connect(left, right NeuronInterface) {
    for _, l := range left.Iter() {
        for _, r := range right.Iter() {
            l.ConnectTo(r)
        }
    }
}

func main() {
    n1, n2 := &Neuron{}, &Neuron{}
    l1, l2 := NewNeuronLayer(3), NewNeuronLayer(4)

    Connect(n1, n2)
    Connect(n1, l1)
    Connect(l2, n1)
    Connect(l1, l2)
}
