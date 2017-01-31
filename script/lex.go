//chim

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	shlex "github.com/flynn-archive/go-shlex"
	"github.com/xwb1989/sqlparser"
)

type QueryToken struct {
	Token   int
	Content string
}

type QueryObject struct {
	Query  string
	Error  string
	Tokens []QueryToken
}

type CmdObject struct {
	Command string
	Tokens  []string
}

func ParseSQL(query string, qrObj *QueryObject) {
	//select *,123,456,'7a8xxxx/*aahihi\'*/xxxx9'/*hihi*/,"zzzz", ` + "`aaaa`" + ` from information_schema.tables, hihi, (select * from master_db) where 1=2 or hihi>(select 100 from abcd) union select * from information_schema.columns where (select count(*) from abcd)>100
	_, err := sqlparser.Parse(query)
	if err != nil {
		//fmt.Println("[SQL Parser]", err)
	}
	qrObj.Error = fmt.Sprintf("%s", err)
	qrObj.Query = query
	qrObj.Tokens = make([]QueryToken, 0)
	tokens := sqlparser.NewStringTokenizer(query)
	for {
		token, content := tokens.Scan()
		if token == 0 {
			break
		}
		qrObj.Tokens = append(qrObj.Tokens, QueryToken{Token: token, Content: string(content)})
		//fmt.Println(token, string(content))
	}
}

func ParseCMD(cmd string, cmdObj *CmdObject) {
	//echo 'bbb' > comment/'aaaa'
	cmdObj.Tokens = make([]string, 0)
	cmdObj.Command = cmd

	ioBuf := new(bytes.Buffer)
	ioBuf.Write([]byte(cmd))

	tokens, err := shlex.NewLexer(ioBuf)
	if err != nil {
		fmt.Println("Bi loi cmd roi", err)
	} else {
		for {
			token, err := tokens.NextWord()
			if err != nil {
				cmdObj.Tokens = append(cmdObj.Tokens, fmt.Sprintf("%s", err))
				break
			} else {
				cmdObj.Tokens = append(cmdObj.Tokens, token)
			}
		}
	}
}

func main() {
	fmt.Println("Lexer testing...")
	fmt.Println("1. Sql")
	fmt.Println("2. Command")
	c := 1
	fmt.Print("> ")
	fmt.Scanf("%d", &c)
	if c != 1 { //Command
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
		for {

			fmt.Print("Enter command: ")
			text, _ := reader.ReadString('\n')
			var cmdObj CmdObject
			ParseCMD(text, &cmdObj)
			fmt.Println("> NumTokens:", len(cmdObj.Tokens))
			fmt.Println("> Tokens:", cmdObj.Tokens)
			fmt.Println()
		}
	} else { //Sql
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
		for {

			fmt.Print("Enter query: ")
			text, _ := reader.ReadString('\n')
			var qrObj QueryObject
			ParseSQL(text, &qrObj)
			fmt.Println("> Error:", qrObj.Error)
			fmt.Println("> Tokens:", qrObj.Tokens)
			fmt.Println()
		}
	}

}
