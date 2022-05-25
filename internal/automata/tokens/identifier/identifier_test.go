package identifier_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/marcobarao/parser/internal/automata/errors"
	"github.com/marcobarao/parser/internal/automata/tokens/identifier"
)

var _ = Describe("Identifier", func() {
	DescribeTable(
		"Providing an input",
		func(input string, succeed bool, accepted bool, expectedErr error) {
			var inputErr error = nil
			automata := identifier.NewIdentifierAutomata()

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
		Entry("When it is 'abc' should succeed and be accepted", "abc", true, true, nil),
		Entry("When it is 'abc12a' should succeed and be accepted", "abc12a", true, true, nil),
		Entry("When it is 'variavel' should succeed and be accepted", "variavel", true, true, nil),
		Entry("When it is '123abc' should not succeed and not be accepted", "123abc", false, false, TransitionNotFoundError{State: 0, Input: '1'}),
		Entry("When it is 'abc.' should not succeed and be accepted", "abc.", false, true, TransitionNotFoundError{State: 1, Input: '.'}),
		Entry("When it is 'abc*7' should not succeed and be accepted", "abc*7", false, true, TransitionNotFoundError{State: 1, Input: '*'}),
	)
})
