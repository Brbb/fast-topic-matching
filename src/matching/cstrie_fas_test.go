package matching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCSTrieFASMatcher(t *testing.T) {
	assert := assert.New(t)
	var (
		m         = NewCSTrieMatcher()
		fascriber = Fascriber{Http, "fas123"}
	)

	sub0, err := m.Subscribe("forex.*", fascriber)
	assert.NoError(err)
	sub1, err := m.Subscribe("*.usd", fascriber)
	assert.NoError(err)
	sub2, err := m.Subscribe("forex.eur", fascriber)
	assert.NoError(err)

	assertEqual(assert, []Subscriber{fascriber}, m.Lookup("forex.eur"))
	assertEqual(assert, []Subscriber{}, m.Lookup("trade.jpy"))

	m.Unsubscribe(sub0)
	m.Unsubscribe(sub1)
	m.Unsubscribe(sub2)

	assertEqual(assert, []Subscriber{}, m.Lookup("forex.eur"))
	assertEqual(assert, []Subscriber{}, m.Lookup("forex"))
	assertEqual(assert, []Subscriber{}, m.Lookup("trade.jpy"))
	assertEqual(assert, []Subscriber{}, m.Lookup("forex.jpy"))
	assertEqual(assert, []Subscriber{}, m.Lookup("trade"))
}
