package operator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/marcobarao/parser/internal/automata/errors"
	"github.com/marcobarao/parser/internal/automata/tokens/operator"
)

var _ = Describe("Operator", func() {
	DescribeTable(
		"Providing an input",
		func(input string, succeed bool, accepted bool, expectedErr error) {
			var inputErr error = nil
			automata := operator.NewOperatorAutomata()

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
		Entry("When it is '+' should succeed and be accepted", "+", true, true, nil),
		Entry("When it is '/' should succeed and be accepted", "/", true, true, nil),
		Entry("When it is '**' should succeed and be accepted", "**", true, true, nil),
		Entry("When it is '!=' should succeed and be accepted", "!=", true, true, nil),
		Entry("When it is '>' should succeed and be accepted", ">", true, true, nil),
		Entry("When it is '==' should succeed and be accepted", "==", true, true, nil),
		Entry("When it is '.' should not succeed and not be accepted", ".", false, false, TransitionNotFoundError{State: 0, Input: '.'}),
		Entry("When it is 'abc' should not succeed and not be accepted", "abc", false, false, TransitionNotFoundError{State: 0, Input: 'a'}),
		Entry("When it is '=!' should not succeed and be accepted", "=!", false, true, TransitionNotFoundError{State: 3, Input: '!'}),
		Entry("When it is '*+' should not succeed and be accepted", "*+", false, true, TransitionNotFoundError{State: 2, Input: '+'}),
	)
})
