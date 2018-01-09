package test_peg

import (
	"fmt"
	"strconv"
	"testing"
	"github.com/yhirose/go-peg"
	"github.com/stretchr/testify/require"
	"time"
)

func TestCalculator(t *testing.T) {
	parser, err := peg.NewParser(`
    # Simple calculator
    EXPR         <- ATOM (BINOP ATOM)*
    ATOM         <- VALUE / '(' EXPR ')'
    BINOP        <- < [-+/*] >
    VALUE        <- NUMBER / X
    X            <- 'X'
    NUMBER       <- < [0-9]+ >
    %whitespace  <- [ \t]*
    ---
    # Expression parsing option
    %expr  = EXPR   # Rule to apply 'precedence climbing method' to
    %binop = L + -  # Precedence level 1
    %binop = L * /  # Precedence level 2
`)
	require.Nil(t, err)
	// Setup semantic actions
	g := parser.Grammar
	g["EXPR"].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		val := v.ToInt(0)
		if v.Len() > 1 {
			ope := v.ToStr(1)
			rhs := v.ToInt(2)
			switch ope {
			case "+": val += rhs
			case "-": val -= rhs
			case "*": val *= rhs
			case "/": val /= rhs
			}
		}
		return val, nil
	}
	g["BINOP"].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		return v.Token(), nil
	}
	g["NUMBER"].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		return strconv.Atoi(v.Token())
	}
	x := 10
	g["X"].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		return x, nil
	}

	// Parse
	input := "1 + 2 * 3 * (4 - 5 + 6) / 7 - 8 + X"
	val, err := parser.ParseAndGetValue(input, nil)
	require.Nil(t, err)
	require.Equal(t, 7, val)
}

const (
	searchGrammar = `
    EXPR        <- ATOM JOIN EXPR / INVERSE? ATOM
    ATOM        <- COMPARE / '(' EXPR ')'
    COMPARE     <- PARAM BINOP VALUE / PARAM 'in' MULTIVALUE / PARAM 'ip_range' '(' STRING COMMA STRING ')'
    MULTIVALUE  <- '(' VALUE (COMMA VALUE)* ')'
    VALUE       <- NUMBER / STRING
    STRING      <- SQUOTE ((ESCAPE ESCAPE) / (ESCAPE SQUOTE) / GENERAL / DQUOTE)* SQUOTE /
			       DQUOTE ((ESCAPE ESCAPE) / (ESCAPE DQUOTE) / GENERAL / SQUOTE)* DQUOTE
    INVERSE     <- 'not'
    JOIN        <- 'and' / 'or'
    BINOP       <- '=' / '<>' / '!=' / '<' / '>'
    SQUOTE      <- "'"
	DQUOTE      <- '"'
    ESCAPE      <- "\\"
    COMMA       <- ','
	GENERAL     <- &(!SQUOTE !DQUOTE !ESCAPE) .
    PARAM       <- < [A-Za-z_] [A-Za-z0-9_.]* >
    NUMBER      <- < [-+]?[0-9]+ >
    %whitespace <- [ \t]*
`
)

func TestParser(t *testing.T) {
	start := time.Now()
	parser, err := peg.NewParser(searchGrammar)
	end := time.Now()
	fmt.Println(end.Sub(start))
	require.Nil(t, err)
	fillActions(parser.Grammar, []string{"EXPR", "ATOM", "JOIN", "COMPARE", "INVERSE", "PARAM", "BINOP",
		"VALUE", "MULTIVALUE", "NUMBER", "STRING", "GENERAL", "COMMA", "SQUOTE", "DQUOTE"})
	require.Nil(t, parser.Parse(`p=1`, nil))
	require.Nil(t, parser.Parse(`p1=1 and p2=2`, nil))
	require.Nil(t, parser.Parse(`(p in ("1",'2',3))`, nil))
	require.Nil(t, parser.Parse(`not (p=1)`, nil))
	require.Nil(t, parser.Parse(`(p1=1) and (p2=2)`, nil))
	require.Nil(t, parser.Parse(`(p1=1) and p2=2`, nil))
	require.Nil(t, parser.Parse(`((p=1) and (p=2))or(v=3)`, nil))
	require.Nil(t, parser.Parse(`(p='\\\'1 234\\\'') and (p=2)`, nil))
	require.Nil(t, parser.Parse(`(p="\\\"1 234\\\"") and (p=2)`, nil))
}

func TestMultiValue(t *testing.T) {
	parser, err := peg.NewParser(`
    COMPARE     <- PARAM 'in'
    PARAM       <- < [A-Za-z_] [A-Za-z0-9_]* >
    MULTIOP     <- 'in' / 'ip_range'
    MULTIVALUE  <- '(' VALUE (COMMA VALUE)* ')'
    VALUE       <- NUMBER / STRING
    STRING      <- SQUOTE ((ESCAPE ESCAPE) / (ESCAPE SQUOTE) / GENERAL / DQUOTE)* SQUOTE /
			       DQUOTE ((ESCAPE ESCAPE) / (ESCAPE DQUOTE) / GENERAL / SQUOTE)* DQUOTE
    INVERSE     <- 'not'
    JOIN        <- 'and' / 'or'
    BINOP       <- '=' / '<>' / '!=' / '<' / '>'
    SQUOTE      <- "'"
	DQUOTE      <- '"'
    ESCAPE      <- "\\"
    COMMA       <- ','
	GENERAL     <- &(!SQUOTE !DQUOTE !ESCAPE) .
    NUMBER      <- [-+]?[0-9]+
    %whitespace <- [ \t]*
`)
	require.Nil(t, err)
	fillActions(parser.Grammar, []string{"VALUE", "MULTIVALUE", "NUMBER", "STRING", "GENERAL", "COMMA", "SQUOTE", "DQUOTE"})
	require.Nil(t, parser.Parse(`p in`, nil))
}

func TestSequence(t *testing.T) {
	parser, err := peg.NewParser(`
    EXPR        <- ATOM JOIN EXPR / ATOM
    ATOM        <- PARAM / '(' EXPR ')'
    PARAM       <- < [A-Za-z_] [A-Za-z0-9_]* >
    JOIN        <- 'and' / 'or'
    %whitespace <- [ \t]*
`)
	require.Nil(t, err)
	fillActions(parser.Grammar, []string{"EXPR", "ATOM", "JOIN", "PARAM"})
	require.Nil(t, parser.Parse(`p and p`, nil))
}

func TestSimpleParser(t *testing.T) {
	parser, err := peg.NewParser(`
	EXPR     <- SINGLE / '(' EXPR ')' / OR_EXPR / AND_EXPR
    OR_EXPR  <- SINGLE / '(' EXPR ')' 'or' OR_EXPR
    AND_EXPR <- SINGLE / '(' EXPR ')' 'and' AND_EXPR
    SINGLE   <- PARAM OP NUMBER
    PARAM    <- [A-Za-z_] [A-Za-z0-9_.]*
    OP       <- '=' / '<>' / '!=' / '<' / '>' / 'in' / 'ip_range'
	NUMBER   <- '1'
	%whitespace <-  [ \t]*
`)
	require.Nil(t, err)
	g := parser.Grammar
	eprs := []string{"EXPR", "OR_EXPR", "AND_EXPR", "SINGLE", "PARAM", "OP", "NUMBER"/*, "VALUE", "STRING"*/}
	for _, exp := range eprs {
		e := exp
		g[exp].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
			t.Log(e, ":", v.Token())
			return nil, nil
		}
	}
	err = parser.Parse(`p=1`, nil)
	require.Nil(t, err)
}

func fillActions(g map[string]*peg.Rule, acts []string) {
	for _, act := range acts {
		e := act
		g[act].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
			fmt.Println(e, ":", v.Token())
			return nil, nil
		}
	}
}