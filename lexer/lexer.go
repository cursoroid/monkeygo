package lexer

type Lexer struct{
	input string
	position int
	readposition int
	ch byte
}

