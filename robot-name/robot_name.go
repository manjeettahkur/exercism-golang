package robotname

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Robot struct {
	Nombre string
}

// Name returns the name
func (r *Robot) Name() string {
	if r.Nombre == "" {
		r.SetName()
	}
	fmt.Println(r.Nombre)
	return r.Nombre
}

// SetName sets the name (initially or after Reset) for the Robot
func (r *Robot) SetName() {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var buffer bytes.Buffer
	for i := 0; i < 5; i++ {
		if i < 2 {
			buffer.WriteString(string(charset[seededRand.Intn(len(charset))]))
		} else {
			buffer.WriteString(strconv.Itoa(rand.Intn(9)))
		}
	}
	r.Nombre = buffer.String()
}

// Reset calls SetName in order to assign a new name
func (r *Robot) Reset() {
	r.SetName()
}
