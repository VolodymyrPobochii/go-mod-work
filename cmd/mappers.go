package main

import (
	"fmt"
	"strconv"
)

type User struct {
	Name string
	Age  int
}

type Account struct {
	Type  string
	Owner string
}

type Mapper[S, T any] interface {
	Map(in S, out T)
	Map2(in S) T
}

type mapper1 struct {
}

func (*mapper1) Map(in *int, out *string) {
	*out = strconv.FormatInt(int64(*in), 10)
}

func (*mapper1) Map2(in *int) *string {
	str := strconv.FormatInt(int64(*in), 10)
	return &str
}

func NewIntToStringMapper() Mapper[*int, *string] {
	return &mapper1{}
}

type mapper2 struct{}

func (*mapper2) Map(in *User, out *Account) {
	if in.Age > 33 {
		out.Type = "adult"
	} else {
		out.Type = "junior"
	}
	out.Owner = in.Name
}

func (*mapper2) Map2(in *User) *Account {
	out := Account{}
	if in.Age > 33 {
		out.Type = "adult"
	} else {
		out.Type = "junior"
	}
	out.Owner = in.Name
	return &out
}

type UAMapper Mapper[*User, *Account]

func NewUserToAccountMapper() UAMapper {
	return &mapper2{}
}

const initVal = "n/a"

func main() {
	m := NewIntToStringMapper()
	//checkReflection(mapper)

	out := "initVal"
	fmt.Println("out = ", out)
	in := 13
	m.Map(&in, &out)
	fmt.Println("out = ", out)
	in = 33
	out2 := m.Map2(&in)
	fmt.Println("out2 = ", out2)

	uam := NewUserToAccountMapper()

	u := User{
		Name: "Killa Bill",
		Age:  37,
	}
	a := Account{}
	fmt.Println("before: a = ", a)
	uam.Map(&u, &a)
	fmt.Println("after: a = ", a)

	a2 := uam.Map2(&u)
	fmt.Println("*a2 = ", a2)
	fmt.Println("a2 = ", *a2)
}

//func checkReflection(mapper Mapper[*int, *string]) {
//mt := reflect.TypeOf(mapper)
//fmt.Println("mapper type:", mt)
//mte := reflect.TypeOf(mapper).Elem()
//fmt.Println("mapper type element:", mte)
//mti := mt.Implements(reflect.TypeOf((*Mapper[int, string])(nil)).Elem())
//fmt.Println("mapper type implements Mapper[int, string]:", mti)
//mv := reflect.ValueOf(mapper)
//fmt.Println("mapper value:", mv)
//mvt := mv.Type()
//fmt.Println("mapper value type:", mvt)
//mvi := mv.Interface()
//fmt.Println("mapper value interface:", mvi)
//
//var ei interface{} = int64(10)
//fmt.Println("empty interface type:", reflect.TypeOf(ei))
//fmt.Println("empty interface type element:", reflect.TypeOf(&ei).Elem())
//fmt.Println("empty interface value:", reflect.ValueOf(ei))
//fmt.Println("empty interface value interface:", reflect.ValueOf(ei).Interface())
//}
