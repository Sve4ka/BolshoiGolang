package storage

import (
	"strconv"
	"testing"
)

type testCase struct {
	name string
	key  string
	val  string
}

func TestSetGet(t *testing.T) {
	cases := []testCaseType{
		{"hello world", "hello", "world", KindString},
		{"int 2", "int", "2", KindDigit},
		{"int with string", "ints", "12s", KindString},
	}

	s, err := InitStorage()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s.Set(c.key, c.val)
			sValue := s.Get(c.key)

			if *sValue != c.val {
				t.Errorf("Get(%v) = %v, want %v", c.key, sValue, c.val)
			}

			kValue := s.GetKind(c.key)
			if kValue != c.t {
				t.Errorf("GetKind(%v) = %v, want %v", c.key, kValue, c.t)
			}
		})
	}
}

type testCaseType struct {
	name string
	key  string
	val  string
	t    string
}

func TestSetType(t *testing.T) {
	cases := []testCaseType{
		{"hello world", "hello", "world", KindString},
		{"int 2", "int", "2", KindDigit},
		{"int with string", "ints", "12s", KindString},
	}

	s, err := InitStorage()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s.Set(c.key, c.val)
			value := s.getTest(c.key)

			if value.s != c.val {
				t.Errorf("Get(%v) = %v, want %v", c.key, value.s, c.val)
			}

			if value.kind != c.t {
				t.Errorf("GetKind(%v[%v]) = %v, want %v", c.key, c.val, value.kind, c.t)
			}
		})
	}
}

type testCaseNil struct {
	name string
	key  string
	val  *string
	t    string
}

func TestGetNil(t *testing.T) {
	cases := []testCaseNil{
		{"nil", "nil", nil, KindUnknown},
	}
	s, err := InitStorage()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := s.Get(c.key)
			if c.val != out {
				t.Errorf("Get(%v) = %v, want %v", c.key, out, c.val)
			}
			outT := s.GetKind(c.key)
			if c.t != outT {
				t.Errorf("GetKind(%v[%v]) = %v, want %v", c.key, c.val, outT, c.t)
			}
		})
	}
}

type bench struct {
	name  string
	count int
}

var casesInit = []bench{
	{"5", 5},
	{"10", 10},
	{"100", 100},
	{"500", 500},
	{"1000", 10000},
}

func BenchmarkInitStorage(b *testing.B) {
	for _, tCase := range casesInit {
		b.Run(tCase.name, func(bb *testing.B) {
			for i := 0; i < tCase.count; i++ {
				_, err := InitStorage()
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

var casesSet = []bench{
	{"5", 5},
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkSet(b *testing.B) {
	for _, tCase := range casesSet {
		b.Run(tCase.name, func(bb *testing.B) {
			casesi := []testCaseType{}
			for i := 0; i < tCase.count; i++ {
				casesi = append(casesi, testCaseType{"int" + strconv.Itoa(i), "int" + strconv.Itoa(i), strconv.Itoa(i), KindDigit})
			}
			s, err := InitStorage()
			if err != nil {
				b.Fatal(err)
			}
			bb.ResetTimer()
			for i := 0; i < tCase.count; i++ {
				s.Set(casesi[i].key, casesi[i].val)
			}

		})
	}
}

func BenchmarkGet(b *testing.B) {
	for _, tCase := range casesSet {
		b.Run(tCase.name, func(bb *testing.B) {
			casesi := []testCaseType{}
			for i := 0; i < tCase.count; i++ {
				casesi = append(casesi, testCaseType{"int" + strconv.Itoa(i), "int" + strconv.Itoa(i), strconv.Itoa(i), KindDigit})
			}
			s, err := InitStorage()
			if err != nil {
				b.Fatal(err)
			}
			for i := 0; i < tCase.count; i++ {
				s.Set(casesi[i].key, casesi[i].val)
			}
			bb.ResetTimer()
			for i := 0; i < tCase.count; i++ {
				s.Get(casesi[i].key)
			}

		})
	}
}

func BenchmarkGetKind(b *testing.B) {
	for _, tCase := range casesSet {
		b.Run(tCase.name, func(bb *testing.B) {
			casesi := []testCaseType{}
			for i := 0; i < tCase.count; i++ {
				casesi = append(casesi, testCaseType{"int" + strconv.Itoa(i), "int" + strconv.Itoa(i), strconv.Itoa(i), KindDigit})
			}
			s, err := InitStorage()
			if err != nil {
				b.Fatal(err)
			}
			for i := 0; i < tCase.count; i++ {
				s.Set(casesi[i].key, casesi[i].val)
			}
			bb.ResetTimer()
			for i := 0; i < tCase.count; i++ {
				t := s.GetKind(casesi[i].key)
				if t != casesi[i].t {
					b.Fatalf("GetKind(%v) = %v, want %v", casesi[i].key, t, casesi[i].t)
				}
			}

		})
	}
}

func BenchmarkSetGet(b *testing.B) {
	for _, tCase := range casesSet {
		b.Run(tCase.name, func(bb *testing.B) {
			casesi := []testCaseType{}
			for i := 0; i < tCase.count; i++ {
				casesi = append(casesi, testCaseType{"int" + strconv.Itoa(i), "int" + strconv.Itoa(i), strconv.Itoa(i), KindDigit})
			}
			s, err := InitStorage()
			if err != nil {
				b.Fatal(err)
			}
			bb.ResetTimer()
			for i := 0; i < tCase.count; i++ {
				s.Set(casesi[i].key, casesi[i].val)
				s.Get(casesi[i].key)
			}

		})
	}
}
