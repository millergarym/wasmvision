package cv

import (
	"github.com/orsinium-labs/wypes"
	"gocv.io/x/gocv"
)

// Net is a wrapper around gocv.Net for DNN image processing.
type Net struct {
	ID       wypes.UInt32
	Name     string
	Filename string

	Net gocv.Net
}

// NewNet creates a new Net.
func NewNet(model string) *Net {
	return &Net{
		Name: model,
	}
}

// SetNet sets the gocv.Net for the Net.
func (n *Net) SetNet(net gocv.Net) {
	n.Net = net
}

// Close closes the Net.
func (n *Net) Close() {
	n.Net.Close()
}