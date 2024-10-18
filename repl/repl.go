// REPL : Read Evaluate Print Loop
// REPL is a simple interactive programming environment that takes user input, evaluates it, and returns the result to the user.

package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/YuneshShrestha/Interpretor/token"

	"github.com/YuneshShrestha/Interpretor/lexer"
)

const PROMPT = ">> "

/* : read from the input source until encountering a newline, take
the just read line and pass it to an instance of our lexer and finally print all the tokens the lexer
gives us until we encounter EOF.
*/
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in) // scanner is a bufio.Scanner that reads from the input source

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan() // read from the input source until encountering a newline
		if !scanned {
			return
		}
		line := scanner.Text() // take the just read line
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
