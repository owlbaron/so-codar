package grouper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/marcobarao/parser/internal/automata/errors"
	"github.com/marcobarao/parser/internal/automata/tokens/grouper"
)

var _ = Describe("Grouper", func() {
	DescribeTable(
		"Providing an input",
		func(input rune, succeed bool, accepted bool, err error) {
			automata := grouper.NewGrouperAutomata()

			if succeed {
				Expect(automata.Input(input)).Should(Succeed())
			} else {
				Expect(automata.Input(input)).Should(MatchError(err))
			}

			Expect(automata.Accepted()).To(Equal(accepted))
		},
		Entry("When it is '(' should succeed and be accepted", '(', true, true, nil),
		Entry("When it is ')' should succeed and be accepted", ')', true, true, nil),
		Entry("When it is ';' should succeed and be accepted", ';', true, true, nil),
		Entry("When it is '.' should not succeed and not be accepted", '.', false, false, TransitionNotFoundError{State: 0, Input: '.'}),
		Entry("When it is 'a' should not succeed and not be accepted", 'a', false, false, TransitionNotFoundError{State: 0, Input: 'a'}),
	)
})
