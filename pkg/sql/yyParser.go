// Code generated by goyacc - DO NOT EDIT.

package sql

import __yyfmt__ "fmt"

import (
	"strconv"
)

type yySymType struct {
	yys     int
	str     string
	strList []string
	boolean bool

	stmt     Statment
	stmtList []Statment

	createStmt     *CreateStmt
	tableDef       *TableDef
	columnDef      *ColumnDef
	indexDef       *IndexDef
	tableOption    *TableOption
	columnDataType ColumnDataType

	selectStmt      *SelectStmt
	sqlFunc         *SqlFunction
	selectFieldList []*SelectField
	selectStmtFrom  *SelectStmtFrom
	whereExprList   []BoolExpr
	orderFieldList  []*OrderField
	selectStmtLimit *SelectStmtLimit
	compareOp       CompareOp

	insertStmt *InsertStmt
	valueList  [][]string
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault  = 57378
	yyEofCode  = 57344
	AND        = 57365
	ASC        = 57363
	BY         = 57360
	COMP_GE    = 57376
	COMP_LE    = 57375
	COMP_NE    = 57374
	CONSTRAINT = 57352
	COUNT      = 57371
	CREATE     = 57346
	DEFAULT    = 57355
	DESC       = 57364
	FROM       = 57357
	INDEX      = 57350
	INSERT     = 57367
	INTO       = 57368
	KEY        = 57351
	LIMIT      = 57361
	MAX        = 57372
	MIN        = 57373
	NOT        = 57353
	NULL       = 57354
	OFFSET     = 57362
	OR         = 57366
	ORDER      = 57359
	PRIMARY    = 57348
	SELECT     = 57356
	TABLE      = 57347
	UNIQUE     = 57349
	VALUE      = 57369
	VALUES     = 57370
	VARIABLE   = 57377
	WHERE      = 57358
	yyErrCode  = 57345

	yyMaxDepth = 200
	yyTabOfs   = -76
)

var (
	yyPrec = map[int]int{
		OR:  0,
		AND: 1,
		'+': 2,
		'-': 2,
		'*': 3,
		'/': 3,
	}

	yyXLAT = map[int]int{
		44:    0,  // ',' (48x)
		57377: 1,  // VARIABLE (44x)
		41:    2,  // ')' (40x)
		59:    3,  // ';' (38x)
		57389: 4,  // Expr (29x)
		57361: 5,  // LIMIT (26x)
		40:    6,  // '(' (16x)
		57344: 7,  // $end (11x)
		57346: 8,  // CREATE (11x)
		57367: 9,  // INSERT (11x)
		57356: 10, // SELECT (11x)
		57359: 11, // ORDER (10x)
		57365: 12, // AND (9x)
		57357: 13, // FROM (9x)
		57366: 14, // OR (9x)
		57355: 15, // DEFAULT (6x)
		57408: 16, // VaribleList (6x)
		57354: 17, // NULL (5x)
		60:    18, // '<' (4x)
		61:    19, // '=' (4x)
		62:    20, // '>' (4x)
		57376: 21, // COMP_GE (4x)
		57375: 22, // COMP_LE (4x)
		57374: 23, // COMP_NE (4x)
		57351: 24, // KEY (4x)
		57363: 25, // ASC (3x)
		57386: 26, // CompareOp (3x)
		57364: 27, // DESC (3x)
		57350: 28, // INDEX (3x)
		57391: 29, // IndexKey (3x)
		57353: 30, // NOT (3x)
		57358: 31, // WHERE (3x)
		57409: 32, // WhereExprList (3x)
		57379: 33, // AggregateFunction (2x)
		57381: 34, // Ascend (2x)
		57383: 35, // ColumnDef (2x)
		57371: 36, // COUNT (2x)
		57387: 37, // CreateStmt (2x)
		57390: 38, // IndexDef (2x)
		57392: 39, // InsertStmt (2x)
		57372: 40, // MAX (2x)
		57373: 41, // MIN (2x)
		57348: 42, // PRIMARY (2x)
		57396: 43, // PrimaryDef (2x)
		57398: 44, // SelectStmt (2x)
		57400: 45, // SelectStmtLimit (2x)
		57403: 46, // Stmt (2x)
		57349: 47, // UNIQUE (2x)
		57369: 48, // VALUE (2x)
		57407: 49, // ValueGroup (2x)
		57370: 50, // VALUES (2x)
		57380: 51, // AllowNull (1x)
		57360: 52, // BY (1x)
		57382: 53, // ColumnDataType (1x)
		57384: 54, // ColumnGroup (1x)
		57385: 55, // ColumnList (1x)
		57388: 56, // DefaultValue (1x)
		57393: 57, // InsertStmtInto (1x)
		57394: 58, // InsertStmtValue (1x)
		57368: 59, // INTO (1x)
		57362: 60, // OFFSET (1x)
		57395: 61, // OrderFieldList (1x)
		57397: 62, // SelectFieldList (1x)
		57399: 63, // SelectStmtFrom (1x)
		57401: 64, // SelectStmtOrder (1x)
		57402: 65, // SelectStmtWhere (1x)
		57410: 66, // start (1x)
		57404: 67, // StmtList (1x)
		57347: 68, // TABLE (1x)
		57405: 69, // TableDef (1x)
		57406: 70, // TableOption (1x)
		57378: 71, // $default (0x)
		42:    72, // '*' (0x)
		43:    73, // '+' (0x)
		45:    74, // '-' (0x)
		47:    75, // '/' (0x)
		57352: 76, // CONSTRAINT (0x)
		57345: 77, // error (0x)
	}

	yySymNames = []string{
		"','",
		"VARIABLE",
		"')'",
		"';'",
		"Expr",
		"LIMIT",
		"'('",
		"$end",
		"CREATE",
		"INSERT",
		"SELECT",
		"ORDER",
		"AND",
		"FROM",
		"OR",
		"DEFAULT",
		"VaribleList",
		"NULL",
		"'<'",
		"'='",
		"'>'",
		"COMP_GE",
		"COMP_LE",
		"COMP_NE",
		"KEY",
		"ASC",
		"CompareOp",
		"DESC",
		"INDEX",
		"IndexKey",
		"NOT",
		"WHERE",
		"WhereExprList",
		"AggregateFunction",
		"Ascend",
		"ColumnDef",
		"COUNT",
		"CreateStmt",
		"IndexDef",
		"InsertStmt",
		"MAX",
		"MIN",
		"PRIMARY",
		"PrimaryDef",
		"SelectStmt",
		"SelectStmtLimit",
		"Stmt",
		"UNIQUE",
		"VALUE",
		"ValueGroup",
		"VALUES",
		"AllowNull",
		"BY",
		"ColumnDataType",
		"ColumnGroup",
		"ColumnList",
		"DefaultValue",
		"InsertStmtInto",
		"InsertStmtValue",
		"INTO",
		"OFFSET",
		"OrderFieldList",
		"SelectFieldList",
		"SelectStmtFrom",
		"SelectStmtOrder",
		"SelectStmtWhere",
		"start",
		"StmtList",
		"TABLE",
		"TableDef",
		"TableOption",
		"$default",
		"'*'",
		"'+'",
		"'-'",
		"'/'",
		"CONSTRAINT",
		"error",
	}

	yyTokenLiteralStrings = map[int]string{
		57361: "LIMIT",
		57346: "CREATE",
		57367: "INSERT",
		57356: "SELECT",
		57359: "ORDER",
		57365: "AND",
		57357: "FROM",
		57366: "OR",
		57355: "DEFAULT",
		57354: "NULL",
		57376: ">=",
		57375: "<=",
		57374: "!=",
		57351: "KEY",
		57363: "ASC",
		57364: "DESC",
		57350: "INDEX",
		57353: "NOT",
		57358: "WHERE",
		57371: "COUNT",
		57372: "MAX",
		57373: "MIN",
		57348: "PRIMARY",
		57349: "UNIQUE",
		57369: "VALUE",
		57370: "VALUES",
		57360: "BY",
		57368: "INTO",
		57362: "OFFSET",
		57347: "TABLE",
		57352: "CONSTRAINT",
	}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {66, 1},
		2:  {67, 1},
		3:  {67, 2},
		4:  {46, 1},
		5:  {46, 1},
		6:  {46, 1},
		7:  {37, 8},
		8:  {69, 1},
		9:  {69, 1},
		10: {69, 1},
		11: {69, 3},
		12: {69, 3},
		13: {69, 3},
		14: {35, 4},
		15: {38, 5},
		16: {38, 6},
		17: {43, 5},
		18: {29, 1},
		19: {29, 1},
		20: {53, 1},
		21: {51, 0},
		22: {51, 1},
		23: {51, 2},
		24: {56, 0},
		25: {56, 1},
		26: {56, 2},
		27: {56, 2},
		28: {70, 0},
		29: {39, 5},
		30: {57, 1},
		31: {57, 2},
		32: {54, 3},
		33: {55, 0},
		34: {55, 1},
		35: {58, 2},
		36: {58, 2},
		37: {49, 3},
		38: {49, 5},
		39: {44, 4},
		40: {44, 7},
		41: {62, 1},
		42: {62, 3},
		43: {62, 1},
		44: {62, 3},
		45: {33, 4},
		46: {33, 4},
		47: {33, 4},
		48: {63, 2},
		49: {65, 0},
		50: {65, 2},
		51: {32, 3},
		52: {32, 5},
		53: {32, 5},
		54: {32, 5},
		55: {32, 5},
		56: {26, 1},
		57: {26, 1},
		58: {26, 1},
		59: {26, 1},
		60: {26, 1},
		61: {26, 1},
		62: {64, 0},
		63: {64, 3},
		64: {61, 2},
		65: {61, 4},
		66: {34, 0},
		67: {34, 1},
		68: {34, 1},
		69: {45, 0},
		70: {45, 2},
		71: {45, 4},
		72: {45, 4},
		73: {16, 1},
		74: {16, 3},
		75: {4, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [148][]uint16{
		// 0
		{8: 83, 84, 85, 37: 80, 39: 82, 44: 81, 46: 79, 66: 77, 78},
		{7: 76},
		{7: 75, 83, 84, 85, 37: 80, 39: 82, 44: 81, 46: 223},
		{7: 74, 74, 74, 74},
		{7: 72, 72, 72, 72},
		// 5
		{7: 71, 71, 71, 71},
		{7: 70, 70, 70, 70},
		{68: 180},
		{1: 92, 4: 156, 57: 155, 59: 157},
		{1: 92, 4: 87, 33: 88, 36: 89, 40: 90, 91, 62: 86},
		// 10
		{104, 3: 7, 5: 106, 13: 105, 45: 102, 63: 103},
		{35, 3: 35, 5: 35, 13: 35},
		{33, 3: 33, 5: 33, 13: 33},
		{6: 99},
		{6: 96},
		// 15
		{6: 93},
		{1, 1, 1, 1, 5: 1, 1, 11: 1, 1, 1, 1, 1, 17: 1, 1, 1, 1, 1, 1, 1, 25: 1, 27: 1, 30: 1, 1},
		{1: 94},
		{2: 95},
		{29, 3: 29, 5: 29, 13: 29},
		// 20
		{1: 97},
		{2: 98},
		{30, 3: 30, 5: 30, 13: 30},
		{1: 100},
		{2: 101},
		// 25
		{31, 3: 31, 5: 31, 13: 31},
		{3: 154},
		{3: 27, 5: 27, 11: 27, 31: 116, 65: 115},
		{1: 92, 4: 113, 33: 114, 36: 89, 40: 90, 91},
		{1: 92, 4: 112},
		// 30
		{1: 107},
		{108, 3: 6, 60: 109},
		{1: 111},
		{1: 110},
		{3: 4},
		// 35
		{3: 5},
		{3: 28, 5: 28, 11: 28, 31: 28},
		{34, 3: 34, 5: 34, 13: 34},
		{32, 3: 32, 5: 32, 13: 32},
		{3: 14, 5: 14, 11: 142, 64: 141},
		// 40
		{1: 92, 4: 118, 32: 117},
		{3: 26, 5: 26, 11: 26, 128, 14: 127},
		{18: 121, 120, 122, 124, 123, 125, 26: 119},
		{1: 92, 4: 126},
		{1: 20},
		// 45
		{1: 19},
		{1: 18},
		{1: 17},
		{1: 16},
		{1: 15},
		// 50
		{2: 25, 25, 5: 25, 11: 25, 25, 14: 25},
		{1: 92, 4: 135, 6: 136},
		{1: 92, 4: 129, 6: 130},
		{18: 121, 120, 122, 124, 123, 125, 26: 133},
		{1: 92, 4: 118, 32: 131},
		// 55
		{2: 132, 12: 128, 14: 127},
		{2: 21, 21, 5: 21, 11: 21, 21, 14: 21},
		{1: 92, 4: 134},
		{2: 23, 23, 5: 23, 11: 23, 23, 14: 23},
		{18: 121, 120, 122, 124, 123, 125, 26: 139},
		// 60
		{1: 92, 4: 118, 32: 137},
		{2: 138, 12: 128, 14: 127},
		{2: 22, 22, 5: 22, 11: 22, 22, 14: 22},
		{1: 92, 4: 140},
		{2: 24, 24, 5: 24, 11: 24, 24, 14: 24},
		// 65
		{3: 7, 5: 106, 45: 152},
		{52: 143},
		{1: 92, 4: 145, 61: 144},
		{149, 3: 13, 5: 13},
		{10, 3: 10, 5: 10, 25: 147, 27: 148, 34: 146},
		// 70
		{12, 3: 12, 5: 12},
		{9, 3: 9, 5: 9},
		{8, 3: 8, 5: 8},
		{1: 92, 4: 150},
		{10, 3: 10, 5: 10, 25: 147, 27: 148, 34: 151},
		// 75
		{11, 3: 11, 5: 11},
		{3: 153},
		{7: 36, 36, 36, 36},
		{7: 37, 37, 37, 37},
		{6: 160, 54: 159},
		// 80
		{6: 46},
		{1: 92, 4: 158},
		{6: 45},
		{48: 169, 50: 168, 58: 167},
		{1: 92, 43, 4: 163, 16: 162, 55: 161},
		// 85
		{2: 166},
		{164, 2: 42},
		{3, 2: 3},
		{1: 92, 4: 165},
		{2, 2: 2},
		// 90
		{48: 44, 50: 44},
		{3: 179},
		{6: 171, 49: 178},
		{6: 171, 49: 170},
		{174, 3: 40},
		// 95
		{1: 92, 4: 163, 16: 172},
		{164, 2: 173},
		{39, 3: 39},
		{6: 175},
		{1: 92, 4: 163, 16: 176},
		// 100
		{164, 2: 177},
		{38, 3: 38},
		{174, 3: 41},
		{7: 47, 47, 47, 47},
		{1: 92, 4: 181},
		// 105
		{6: 182},
		{1: 92, 4: 187, 24: 191, 28: 192, 188, 35: 184, 38: 185, 42: 190, 186, 47: 189, 69: 183},
		{217, 2: 216},
		{68, 2: 68},
		{67, 2: 67},
		// 110
		{66, 2: 66},
		{1: 92, 4: 207, 53: 206},
		{1: 92, 4: 202},
		{24: 191, 28: 192, 197},
		{24: 193},
		// 115
		{1: 58},
		{1: 57},
		{6: 194},
		{1: 92, 4: 163, 16: 195},
		{164, 2: 196},
		// 120
		{59, 2: 59},
		{1: 92, 4: 198},
		{6: 199},
		{1: 92, 4: 163, 16: 200},
		{164, 2: 201},
		// 125
		{60, 2: 60},
		{6: 203},
		{1: 92, 4: 163, 16: 204},
		{164, 2: 205},
		{61, 2: 61},
		// 130
		{55, 2: 55, 15: 55, 17: 209, 30: 210, 51: 208},
		{56, 2: 56, 15: 56, 17: 56, 30: 56},
		{52, 2: 52, 15: 213, 56: 212},
		{54, 2: 54, 15: 54},
		{17: 211},
		// 135
		{53, 2: 53, 15: 53},
		{62, 2: 62},
		{51, 92, 51, 4: 215, 17: 214},
		{50, 2: 50},
		{49, 2: 49},
		// 140
		{3: 48, 70: 221},
		{1: 92, 4: 187, 24: 191, 28: 192, 188, 35: 218, 38: 219, 42: 190, 220, 47: 189},
		{65, 2: 65},
		{64, 2: 64},
		{63, 2: 63},
		// 145
		{3: 222},
		{7: 69, 69, 69, 69},
		{7: 73, 73, 73, 73},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer, stmt *[]Statment) int {
	const yyError = 77

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 2:
		{
			*stmt = append(*stmt, yyS[yypt-0].stmt)
		}
	case 3:
		{
			*stmt = append(*stmt, yyS[yypt-0].stmt)
		}
	case 4:
		{
			yyVAL.stmt = Statment(yyS[yypt-0].createStmt)
		}
	case 5:
		{
			yyVAL.stmt = Statment(yyS[yypt-0].selectStmt)
		}
	case 6:
		{
			yyVAL.stmt = Statment(yyS[yypt-0].insertStmt)
		}
	case 7:
		{
			yyVAL.createStmt = &CreateStmt{Table: yyS[yypt-5].str, Def: yyS[yypt-3].tableDef, Option: yyS[yypt-1].tableOption}
		}
	case 8:
		{
			yyVAL.tableDef = &TableDef{
				Column: []*ColumnDef{yyS[yypt-0].columnDef},
				Index:  []*IndexDef{},
			}
		}
	case 9:
		{
			yyVAL.tableDef = &TableDef{
				Column: []*ColumnDef{},
				Index:  []*IndexDef{yyS[yypt-0].indexDef},
			}
		}
	case 10:
		{
			yyVAL.tableDef = &TableDef{
				Column:  []*ColumnDef{},
				Index:   []*IndexDef{},
				Primary: yyS[yypt-0].indexDef,
			}
		}
	case 11:
		{
			yyVAL.tableDef.Column = append(yyVAL.tableDef.Column, yyS[yypt-0].columnDef)
		}
	case 12:
		{
			yyVAL.tableDef.Index = append(yyVAL.tableDef.Index, yyS[yypt-0].indexDef)
		}
	case 13:
		{
			if yyVAL.tableDef.Primary == nil {
				yyVAL.tableDef.Primary = yyS[yypt-0].indexDef
			} else {
				__yyfmt__.Printf("重复定义主键 %v %v ", yyVAL.tableDef.Primary, yyS[yypt-0].indexDef)
				goto ret1
			}
		}
	case 14:
		{
			yyVAL.columnDef = &ColumnDef{FieldName: yyS[yypt-3].str, FieldType: yyS[yypt-2].columnDataType, AllowNull: yyS[yypt-1].boolean, Default: yyS[yypt-0].str}
		}
	case 15:
		{
			yyVAL.indexDef = &IndexDef{IndexName: yyS[yypt-3].str, IndexField: yyS[yypt-1].strList}
		}
	case 16:
		{
			yyVAL.indexDef = &IndexDef{IndexName: yyS[yypt-3].str, IndexField: yyS[yypt-1].strList, Unique: true}
		}
	case 17:
		{
			yyVAL.indexDef = &IndexDef{Primary: true, IndexField: yyS[yypt-1].strList}
		}
	case 20:
		{
			t, ok := typeMapping[yyS[yypt-0].str]

			if ok {
				yyVAL.columnDataType = t
			} else {
				__yyfmt__.Printf("不支持的数据类型 %s", yyS[yypt-0].str)
				goto ret1
			}
		}
	case 21:
		{
			yyVAL.boolean = true
		}
	case 22:
		{
			yyVAL.boolean = true
		}
	case 23:
		{
			yyVAL.boolean = false
		}
	case 24:
		{
			yyVAL.str = ""
		}
	case 25:
		{
			yyVAL.str = ""
		}
	case 26:
		{
			yyVAL.str = ""
		}
	case 27:
		{
			yyVAL.str = yyS[yypt-0].str
		}
	case 28:
		{
			yyVAL.tableOption = nil
		}
	case 29:
		{
			yyVAL.insertStmt = &InsertStmt{Into: yyS[yypt-3].str, ColumnList: yyS[yypt-2].strList, ValueList: yyS[yypt-1].valueList}
		}
	case 30:
		{
			yyVAL.str = yyS[yypt-0].str
		}
	case 31:
		{
			yyVAL.str = yyS[yypt-0].str
		}
	case 32:
		{
			yyVAL.strList = yyS[yypt-1].strList
		}
	case 33:
		{
			yyVAL.strList = nil
		}
	case 35:
		{
			yyVAL.valueList = yyS[yypt-0].valueList
		}
	case 36:
		{
			yyVAL.valueList = yyS[yypt-0].valueList
		}
	case 37:
		{
			yyVAL.valueList = [][]string{yyS[yypt-1].strList}
		}
	case 38:
		{
			yyVAL.valueList = append(yyS[yypt-4].valueList, yyS[yypt-1].strList)
		}
	case 39:
		{
			yyVAL.selectStmt = &SelectStmt{Filed: yyS[yypt-2].selectFieldList, Limit: yyS[yypt-1].selectStmtLimit}
		}
	case 40:
		{
			yyVAL.selectStmt = &SelectStmt{Filed: yyS[yypt-5].selectFieldList, From: yyS[yypt-4].selectStmtFrom, Where: yyS[yypt-3].whereExprList, Order: yyS[yypt-2].orderFieldList, Limit: yyS[yypt-1].selectStmtLimit}
		}
	case 41:
		{
			yyVAL.selectFieldList = []*SelectField{&SelectField{Field: yyS[yypt-0].str}}
		}
	case 42:
		{
			yyVAL.selectFieldList = append(yyS[yypt-2].selectFieldList, &SelectField{Field: yyS[yypt-0].str})
		}
	case 43:
		{
			yyVAL.selectFieldList = []*SelectField{&SelectField{Expr: yyS[yypt-0].sqlFunc}}
		}
	case 44:
		{
			yyVAL.selectFieldList = append(yyS[yypt-2].selectFieldList, &SelectField{Expr: yyS[yypt-0].sqlFunc})
		}
	case 45:
		{
			yyVAL.sqlFunc = &SqlFunction{Type: AGGREGATE_FUNCTION, Func: yyS[yypt-3].str, Args: []interface{}{yyS[yypt-1].str}}
		}
	case 46:
		{
			yyVAL.sqlFunc = &SqlFunction{Type: AGGREGATE_FUNCTION, Func: yyS[yypt-3].str, Args: []interface{}{yyS[yypt-1].str}}
		}
	case 47:
		{
			yyVAL.sqlFunc = &SqlFunction{Type: AGGREGATE_FUNCTION, Func: yyS[yypt-3].str, Args: []interface{}{yyS[yypt-1].str}}
		}
	case 48:
		{
			yyVAL.selectStmtFrom = &SelectStmtFrom{
				Table: yyS[yypt-0].str,
			}
		}
	case 49:
		{
			yyVAL.whereExprList = nil
		}
	case 50:
		{
			yyVAL.whereExprList = yyS[yypt-0].whereExprList
		}
	case 51:
		{
			yyVAL.whereExprList = []BoolExpr{&WhereField{Field: yyS[yypt-2].str, Value: yyS[yypt-0].str, Compare: yyS[yypt-1].compareOp}}
		}
	case 52:
		{
			yyS[yypt-1].compareOp.Negate()
			filed := &WhereField{Field: yyS[yypt-2].str, Value: yyS[yypt-0].str, Compare: yyS[yypt-1].compareOp}
			if len(yyVAL.whereExprList) == 1 {
				yyVAL.whereExprList[0].Negate()
				yyVAL.whereExprList = append(yyVAL.whereExprList, filed)
				yyVAL.whereExprList = []BoolExpr{&WhereExpr{Negation: true, Cnf: yyVAL.whereExprList}}
			} else {
				yyVAL.whereExprList = []BoolExpr{&WhereExpr{Negation: true, Cnf: []BoolExpr{&WhereExpr{Negation: true, Cnf: yyVAL.whereExprList}, filed}}}
			}
		}
	case 53:
		{
			yyVAL.whereExprList = append(yyVAL.whereExprList, &WhereField{Field: yyS[yypt-2].str, Value: yyS[yypt-0].str, Compare: yyS[yypt-1].compareOp})
		}
	case 54:
		{
			if len(yyVAL.whereExprList) == 1 {
				yyVAL.whereExprList[0].Negate()
				yyVAL.whereExprList = append(yyVAL.whereExprList, &WhereExpr{Negation: true, Cnf: yyS[yypt-1].whereExprList})
				yyVAL.whereExprList = []BoolExpr{&WhereExpr{Negation: true, Cnf: yyVAL.whereExprList}}
			} else {
				yyVAL.whereExprList = []BoolExpr{&WhereExpr{Negation: true, Cnf: []BoolExpr{&WhereExpr{Negation: true, Cnf: yyVAL.whereExprList}, &WhereExpr{Negation: true, Cnf: yyS[yypt-1].whereExprList}}}}
			}
		}
	case 55:
		{
			yyVAL.whereExprList = append(yyVAL.whereExprList, yyS[yypt-1].whereExprList...)
		}
	case 56:
		{
			yyVAL.compareOp = EQ
		}
	case 57:
		{
			yyVAL.compareOp = LT
		}
	case 58:
		{
			yyVAL.compareOp = GT
		}
	case 59:
		{
			yyVAL.compareOp = LE
		}
	case 60:
		{
			yyVAL.compareOp = GE
		}
	case 61:
		{
			yyVAL.compareOp = NE
		}
	case 62:
		{
			yyVAL.orderFieldList = nil
		}
	case 63:
		{
			yyVAL.orderFieldList = yyS[yypt-0].orderFieldList
		}
	case 64:
		{
			yyVAL.orderFieldList = []*OrderField{&OrderField{Field: yyS[yypt-1].str, Asc: yyS[yypt-0].boolean}}
		}
	case 65:
		{
			yyVAL.orderFieldList = append(yyS[yypt-3].orderFieldList, &OrderField{Field: yyS[yypt-1].str, Asc: yyS[yypt-0].boolean})
		}
	case 66:
		{
			yyVAL.boolean = true
		}
	case 67:
		{
			yyVAL.boolean = true
		}
	case 68:
		{
			yyVAL.boolean = false
		}
	case 69:
		{
			yyVAL.selectStmtLimit = nil
		}
	case 70:
		{
			size, err := strconv.Atoi(yyS[yypt-0].str)
			if err != nil {
				yylex.Error(err.Error())
				goto ret1
			}
			yyVAL.selectStmtLimit = &SelectStmtLimit{Size: size}
		}
	case 71:
		{

			offset, err := strconv.Atoi(yyS[yypt-2].str)
			if err != nil {
				yylex.Error(err.Error())
				goto ret1
			}
			size, err := strconv.Atoi(yyS[yypt-0].str)
			if err != nil {
				yylex.Error(err.Error())
				goto ret1
			}
			yyVAL.selectStmtLimit = &SelectStmtLimit{Offset: offset, Size: size}
		}
	case 72:
		{
			offset, err := strconv.Atoi(yyS[yypt-0].str)
			if err != nil {
				yylex.Error(err.Error())
				goto ret1
			}
			size, err := strconv.Atoi(yyS[yypt-2].str)
			if err != nil {
				yylex.Error(err.Error())
				goto ret1
			}
			yyVAL.selectStmtLimit = &SelectStmtLimit{Offset: offset, Size: size}
		}
	case 73:
		{
			yyVAL.strList = []string{yyS[yypt-0].str}
		}
	case 74:
		{
			yyVAL.strList = append(yyS[yypt-2].strList, yyS[yypt-0].str)
		}
	case 75:
		{
			str, err := TrimQuote(yyS[yypt-0].str)
			if err != nil {
				yylex.Error(err.Error())
				goto ret1
			}
			yyVAL.str = str
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
