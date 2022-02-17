package main

import "fmt"

//interface == contrato
type transporter interface {
	calculateFreight(value float32) float32
}

type correios struct{}

func (c correios) calculateFreight(value float32) float32 {
	return value * 0.3
}

type fedex struct{}

func (f fedex) calculateFreight(value float32) float32 {
	return value * 0.5
}

type novaTransportadora struct{}

func (n novaTransportadora) calculateFreight(value float32) float32 {
	return value * 0.8
}

func main() {
	value := float32(53.1)

	freights := make(map[string]float32)

	freights["Correios"] = getFreight(correios{}, value)
	freights["Fedex"] = getFreight(fedex{}, value)
	freights["Nova transportadora"] = getFreight(novaTransportadora{}, value)

	fmt.Println(freights)

}

func getFreight(t transporter, value float32) float32 {
	return t.calculateFreight(value)
}
