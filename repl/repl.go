package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkeyLang/lexer"
	"monkeyLang/token"
)

//REPL 是指Read-Eval-Print Loop（读取（求值求打印循环
// REPL 读取输入，将其发送到解释器进行求值，然后打印解释器的输出，最后重新开始，重复“读取读求值求打印”这个循环
const PROMPT = ">> "

// 输入的源代码中读取，直到读完一行代码，将读取的代码行传递给词法分析器实例，然后输出词法分析器生成的词法单元，直到遇到EOF
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
