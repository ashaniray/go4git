package go4git

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func trims(s string) string {
	return strings.Trim(s, " ")
}

func SplitTwoAndTrim(s string, sep string) (string, string) {
	xs := strings.SplitN(s, sep, 2)
	if len(xs) == 2 {
		return trims(xs[0]), trims(xs[1])
	} else {
		return "", trims(xs[0])
	}
}

type Person struct {
	Name  string
	Email string
	Time  time.Time
}

func (p *Person) String() string {
	return fmt.Sprintf("%s <%s> %s", p.Name, p.Email, p.Time)
}

func parseTime(s string) (*time.Time, error) {
	ts, _ := SplitTwoAndTrim(s, " ")
	i, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return nil, err
	}
	tm := time.Unix(i, 0)

	return &tm, nil
}

func parsePerson(s string) *Person {
	name, rest := SplitTwoAndTrim(s, "<")
	email, rest := SplitTwoAndTrim(rest, ">")

	tm, _ := parseTime(rest)
	return &Person{Name: name, Email: email, Time: *tm}
}
