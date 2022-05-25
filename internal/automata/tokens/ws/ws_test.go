package ws_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/marcobarao/parser/internal/automata/errors"
	"github.com/marcobarao/parser/internal/automata/tokens/ws"
)

var _ = Describe("WS", func() {
	DescribeTable(
		"Providing an input",
		func(input string, succeed bool, accepted bool, expectedErr error) {
			var inputErr error = nil
			automata := ws.NewWSAutomata()

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
		Entry("When it is ' ' should succeed and be accepted", " ", true, true, nil),
		Entry("When it is '\t' should succeed and be accepted", "\t", true, true, nil),
		Entry("When it is '\t \t' should succeed and be accepted", "\t \t", true, true, nil),
		Entry("When it is '\t -' should not succeed and be accepted", "\t -", false, true, TransitionNotFoundError{State: 1, Input: '-'}),
		Entry("When it is '.' should not succeed and not be accepted", ".", false, false, TransitionNotFoundError{State: 0, Input: '.'}),
		Entry("When it is 'a' should not succeed and not be accepted", "a", false, false, TransitionNotFoundError{State: 0, Input: 'a'}),
	)
})
