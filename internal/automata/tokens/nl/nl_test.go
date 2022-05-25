package nl_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/marcobarao/parser/internal/automata/errors"
	"github.com/marcobarao/parser/internal/automata/tokens/nl"
)

var _ = Describe("NL", func() {
	DescribeTable(
		"Providing an input",
		func(input string, succeed bool, accepted bool, expectedErr error) {
			var inputErr error = nil
			automata := nl.NewNLAutomata()

			for _, symbol := range input {
				inputErr = automata.Input(symbol)

				if inputErr != nil {
					break
				}
			}

			if succeed {
				Expect(inputErr).Should(Succeed())
			} else {
				Expect(inputErr).Should(MatchError(expectedErr))
			}

			Expect(automata.Accepted()).To(Equal(accepted))
		},
		Entry("When it is '\n' should succeed and be accepted", "\n", true, true, nil),
		Entry("When it is '\n,' should not succeed and be accepted", "\n,", false, true, TransitionNotFoundError{State: 1, Input: ','}),
		Entry("When it is '.' should not succeed and not be accepted", ".", false, false, TransitionNotFoundError{State: 0, Input: '.'}),
		Entry("When it is 'a' should not succeed and not be accepted", "a", false, false, TransitionNotFoundError{State: 0, Input: 'a'}),
	)
})
