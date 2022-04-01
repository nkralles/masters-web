package pgdriver

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

var whitespace = ' '

func ColumnEqualArg(col string, arg int) string {
	return fmt.Sprintf("%s = $%d", col, arg)
}

func ColumnEqual(col string, eq string) string {
	if len(eq) > 0 {
		var buffer bytes.Buffer
		buffer.WriteString(col)
		buffer.WriteString(" = ")
		buffer.WriteRune('\'')
		buffer.WriteString(eq)
		buffer.WriteRune('\'')
		return buffer.String()
	}
	return ""
}

func ColumnNotEqual(col string, eq string) string {
	if len(eq) > 0 {
		var buffer bytes.Buffer
		buffer.WriteString(col)
		buffer.WriteString(" != ")
		buffer.WriteRune('\'')
		buffer.WriteString(eq)
		buffer.WriteRune('\'')
		return buffer.String()
	}
	return ""
}

func ColumnIntEqual(col string, eq int) string {
	if eq > 0 {
		var buffer bytes.Buffer
		buffer.WriteString(col)
		buffer.WriteString(" = ")
		buffer.WriteString(strconv.Itoa(eq))
		return buffer.String()
	}
	return ""
}

func ColumnLike(col string, like string) string {
	if len(like) > 0 {
		var buffer bytes.Buffer
		buffer.WriteString(col)
		buffer.WriteString(" ILIKE ")
		buffer.WriteRune('\'')
		buffer.WriteRune('%')
		buffer.WriteString(like)
		buffer.WriteRune('%')
		buffer.WriteRune('\'')
		return buffer.String()
	}
	return ""
}

func ColumnEqualBool(col string, eq bool) string {
	var buffer bytes.Buffer
	buffer.WriteString(col)
	buffer.WriteString(" = ")
	buffer.WriteString(strconv.FormatBool(eq))
	return buffer.String()
}

func Limit(num int) string {
	if num > 0 {
		return fmt.Sprintf("LIMIT %d", num)
	}
	return ""
}

func Limit64(num int64) string {
	if num > 0 {
		return fmt.Sprintf("LIMIT %d", num)
	}
	return ""
}

func SelectColumnsMap(cols []string, m map[string]string) string {
	var newcols []string
	for _, col := range cols {
		if c, ok := m[col]; ok {
			newcols = append(newcols, c)
		} else {
			newcols = append(newcols, col)
		}
	}
	return SelectColumns(newcols)
}

func SelectColumns(cols []string) string {
	if len(cols) > 0 {
		var buffer bytes.Buffer
		for i, col := range cols {
			buffer.WriteString(col)
			if i < len(cols)-1 {
				buffer.WriteRune(',')
			}
		}
		return buffer.String()
	}
	return "*"
}

func WhereAnds(ands []string) string {
	if len(ands) > 0 {
		var buffer bytes.Buffer
		buffer.WriteString("WHERE ")
		buffer.WriteString(Ands(ands))
		return buffer.String()
	}
	return ""
}

func WhereOrs(ands []string) string {
	if len(ands) > 0 {
		var buffer bytes.Buffer
		buffer.WriteString("WHERE ")
		buffer.WriteString(Ors(ands))
		return buffer.String()
	}
	return ""
}

// Simple ands where all values are easily stringed
func Ands(ands []string) string {

	if len(ands) > 0 {
		var buffer bytes.Buffer
		buffer.WriteRune('(')
		buffer.WriteRune(whitespace)
		for i, and := range ands {
			if len(and) > 0 {
				buffer.WriteString(and)
				if i < len(ands)-1 {
					buffer.WriteString(" AND ")
				}
			}
		}
		buffer.WriteRune(whitespace)
		buffer.WriteRune(')')
		return buffer.String()
	}
	return ""
}

// Simple ands where all values are easily stringed
func Ors(ors []string) string {

	if len(ors) > 0 {
		var buffer bytes.Buffer
		buffer.WriteRune('(')
		buffer.WriteRune(whitespace)
		for i, or := range ors {
			if len(or) > 0 {
				buffer.WriteString(or)
				if i < len(ors)-1 {
					buffer.WriteString(" or ")
				}
			}
		}
		buffer.WriteRune(whitespace)
		buffer.WriteRune(')')
		return buffer.String()
	}
	return ""
}

type OrderByItem struct {
	Column    string
	Direction OrderDirection
}
type OrderDirection string

func getOrderDirection(asc bool) OrderDirection {
	if asc {
		return ASC
	}
	return DESC
}

var DESC OrderDirection = "desc"
var ASC OrderDirection = "asc"

func OrderByItems(i []OrderByItem) string {
	var orders []string
	if len(i) > 0 {
		for _, item := range i {
			orders = append(orders, fmt.Sprintf("%s %s", item.Column, item.Direction))
		}
		return "order by " + strings.Join(orders, ",")
	}
	return ""
}

func OrderByMap(o map[string]string) string {
	var orders []string
	for k, v := range o {
		orders = append(orders, fmt.Sprintf("%s %s", k, v))
	}
	return strings.Join(orders, ",")
}

func OrderByRaw(orders []string) string {
	if len(orders) > 0 {
		var buffer bytes.Buffer
		buffer.WriteString("ORDER BY ")
		for i, order := range orders {
			buffer.WriteString(order)
			if i < len(orders)-1 {
				buffer.WriteString(",")
			}
		}
		return buffer.String()
	}
	return ""
}

func OrderBy(orders []string, asc bool) string {
	if len(orders) > 0 {
		var buffer bytes.Buffer
		buffer.WriteString("ORDER BY ")
		for i, order := range orders {
			if len(order) == 0 {
				return ""
			}
			buffer.WriteString(order)
			if i < len(orders)-1 {
				buffer.WriteString(",")
			}
		}
		if asc {
			buffer.WriteString(" ASC")
		} else {
			buffer.WriteString(" DESC")
		}
		return buffer.String()
	}
	return ""
}

func AnyIntInArray(col string, vals []int, not bool) string {
	if len(vals) > 0 {
		var buffer bytes.Buffer
		if not {
			buffer.WriteString("not")
		}
		buffer.WriteRune('(')
		buffer.WriteString(col)
		buffer.WriteString(" = any(array[")
		for i, val := range vals {
			buffer.WriteString(fmt.Sprintf("%d", val))
			if i < len(vals)-1 {
				buffer.WriteString(",")
			}
		}
		buffer.WriteRune(']')
		buffer.WriteRune(')')
		buffer.WriteRune(')')
		return buffer.String()
	}
	return ""
}

func AnyInt64InArray(col string, vals []int64, not bool) string {
	if len(vals) > 0 {
		var buffer bytes.Buffer
		if not {
			buffer.WriteString("not")
		}
		buffer.WriteRune('(')
		buffer.WriteString(col)
		buffer.WriteString(" = any(array[")
		for i, val := range vals {
			buffer.WriteString(fmt.Sprintf("%d", val))
			if i < len(vals)-1 {
				buffer.WriteString(",")
			}
		}
		buffer.WriteRune(']')
		buffer.WriteRune(')')
		buffer.WriteRune(')')
		return buffer.String()
	}
	return ""
}

func AnyInArray(col string, vals []string, not bool) string {
	if len(vals) > 0 {
		var buffer bytes.Buffer
		if not {
			buffer.WriteString("not")
		}
		buffer.WriteRune('(')
		buffer.WriteString(col)
		buffer.WriteString(" = any(array[")
		for i, val := range vals {
			buffer.WriteRune('\'')
			buffer.WriteString(val)
			buffer.WriteRune('\'')
			if i < len(vals)-1 {
				buffer.WriteString(",")
			}
		}
		buffer.WriteRune(']')
		buffer.WriteRune(')')
		buffer.WriteRune(')')
		return buffer.String()
	}
	return ""
}

func ArrayContains(arrCol string, vals []string) string {
	if len(vals) > 0 {
		var buffer bytes.Buffer
		buffer.WriteString(arrCol)
		buffer.WriteString(" && ARRAY[")
		for i, val := range vals {
			buffer.WriteRune('\'')
			buffer.WriteString(val)
			buffer.WriteRune('\'')
			if i < len(vals)-1 {
				buffer.WriteString(",")
			}
		}
		buffer.WriteRune(']')
		return buffer.String()
	}
	return ""
}

func ValueInArrayCol(val string, col string) string {
	var buffer bytes.Buffer
	buffer.WriteRune('\'')
	buffer.WriteString(val)
	buffer.WriteRune('\'')
	buffer.WriteString(fmt.Sprintf(" = any (%s)", col))
	return buffer.String()
}

func Offset(offset int) string {
	if offset > 0 {
		return fmt.Sprintf("OFFSET %d", offset)
	}
	return ""
}

func Offset64(offset int64) string {
	if offset > 0 {
		return fmt.Sprintf("OFFSET %d", offset)
	}
	return ""
}

func Froms(f map[string]string) string {
	var froms []string
	for k, v := range f {
		froms = append(froms, fmt.Sprintf("%s as %s", v, k))
	}
	return strings.Join(froms, ",")
}

func StringArray(vals []string) string {
	var buffer bytes.Buffer
	buffer.WriteString("ARRAY[")

	for i, val := range vals {
		buffer.WriteRune('\'')
		buffer.WriteString(val)
		buffer.WriteRune('\'')
		if i < len(vals)-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteRune(']')
	return buffer.String()
}

func CreateTsVectorFromQuery(val, tsName string) string {
	strs := strings.Split(val, " ")
	var buffer bytes.Buffer
	for i, str := range strs {
		buffer.WriteString(str)
		if i < len(strs)-1 {
			buffer.WriteString(" | ")
		}
	}
	return fmt.Sprintf(",to_tsquery('%s') %s", buffer.String(), tsName)
}
