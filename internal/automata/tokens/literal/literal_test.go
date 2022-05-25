package literal_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/marcobarao/parser/internal/automata/errors"
	"github.com/marcobarao/parser/internal/automata/tokens/literal"
)

var _ = Describe("Literal", func() {
	DescribeTable(
		"Providing an input",
		func(input string, succeed bool, accepted bool, expectedErr error) {
			var inputErr error = nil
			automata := literal.NewLiteralAutomata()

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
		Entry("When it is '123' should succeed and be accepted", "123", true, true, nil),
		Entry("When it is '000' should succeed and be accepted", "000", true, true, nil),
		Entry("When it is '0.372' should succeed and be accepted", "0.372", true, true, nil),
		Entry("When it is '.372' should succeed and be accepted", ".372", true, true, nil),
		Entry("When it is '.372e3' should succeed and be accepted", ".372e3", true, true, nil),
		Entry("When it is '.372e+3' should succeed and be accepted", ".372e+3", true, true, nil),
		Entry("When it is '0.37122e-51' should succeed and be accepted", "0.37122e-51", true, true, nil),
		Entry("When it is '0.37122e-.' should not succeed and not be accepted", "0.37122e-.", false, false, TransitionNotFoundError{State: 5, Input: '.'}),
		Entry("When it is 'abc' should not succeed and not be accepted", "abc", false, false, TransitionNotFoundError{State: 0, Input: 'a'}),
		Entry("When it is '37.2.' should not succeed and be accepted", "37.2.", false, true, TransitionNotFoundError{State: 3, Input: '.'}),
		Entry("When it is '25*7' should not succeed and be accepted", "25*7", false, true, TransitionNotFoundError{State: 1, Input: '*'}),
	)
})
