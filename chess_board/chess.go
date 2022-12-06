package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var whites []interface{}
var blacks []interface{}

const (
	order   = string("KQRBNP")
	columns = string("abcdefgh")
)

type Piece struct {
	isWhite bool
	piece   string
	col     int
	row     int
}

type Black struct {
	Piece
}

type White struct {
	Piece
}

func result(sliceInterface []interface{}) string {
	result := make([]string, 0, 16)

	for i, _ := range sliceInterface {
		iItem := reflect.ValueOf(sliceInterface[i])
		result = append(result, pieceWithPosition(iItem.FieldByName("piece").String(),
			iItem.FieldByName("row").Int(),
			iItem.FieldByName("col").Int()))

	}
	return strings.Join(result, ",")
}

func createPiece(piecesLine string, row int) {
	for c, pChar := range piecesLine {
		if pChar == 124 || pChar == 46 || pChar == 58 {
			continue
		}
		p := Piece{unicode.IsUpper(pChar), strings.ToUpper(string(pChar)), c / 4, row}
		if p.isWhite {
			whites = append(whites, White{p})
		} else {
			blacks = append(blacks, Black{p})
		}
	}
}

func pieceWithPosition(piece string, row int64, col int64) string {
	builder := strings.Builder{}
	if piece != "P" {
		builder.WriteString(piece)
	}
	builder.WriteString(string(columns[col]))
	builder.WriteString(strconv.FormatInt(row, 10))
	return builder.String()
}

func sortSlice(sliceInterface []interface{}) {
	sort.Slice(sliceInterface, func(i, j int) bool {

		iItem := reflect.ValueOf(sliceInterface[i])
		jItem := reflect.ValueOf(sliceInterface[j])

		iPiece := iItem.FieldByName("piece").String()
		jPiece := jItem.FieldByName("piece").String()

		iIndex := strings.Index(order, strings.ToUpper(iPiece))
		jIndex := strings.Index(order, strings.ToUpper(jPiece))

		iRow := iItem.FieldByName("row").Int()
		jRow := jItem.FieldByName("row").Int()

		iCol := iItem.FieldByName("col").Int()
		jCol := jItem.FieldByName("col").Int()

		if iIndex != jIndex {
			return iIndex < jIndex
		} else if iRow != jRow {
			if iItem.FieldByName("isWhite").Bool() {
				return iRow < jRow
			} else {
				return iRow > jRow
			}
		} else if iCol != jCol {
			return iCol < jCol
		}
		return false
	})
}

func main() {
	//file, err := os.Open(os.Args[1])
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer func(file *os.File) {
	//	err := file.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(file)
	//
	//scanner := bufio.NewScanner(file)
	//counter := 0
	//for r := 8; r > 0; {
	//	for scanner.Scan() {
	//		if counter%2 != 0 {
	//			piecesLine := strings.ReplaceAll(scanner.Text(), "\r\n", "")[2:]
	//			createPiece(piecesLine, r)
	//			r--
	//		}
	//		counter++
	//	}
	//}

	for r := 8; r > 0; r-- {
		func() string {
			text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			return strings.ReplaceAll(text, "\r\n", "")
		}()
		piecesLine := func() string {
			text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			return strings.ReplaceAll(text, "\r\n", "")
		}()[2:]
		for c, pChar := range piecesLine {
			if pChar == 124 || pChar == 46 || pChar == 58 {
				continue
			}
			p := Piece{unicode.IsUpper(pChar), strings.ToUpper(string(pChar)), c / 4, r}
			if p.isWhite {
				whites = append(whites, White{p})
			} else {
				blacks = append(blacks, Black{p})
			}
		}
	}

	sortSlice(whites)
	sortSlice(blacks)
	fmt.Println("White: ", result(whites))
	fmt.Println("Black: ", result(blacks))
}
