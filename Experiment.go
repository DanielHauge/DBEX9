package main

import "time"

var start time.Time
var t time.Time

func TestStringQuery(input []string,f func(job string)([]person))[]time.Duration{
	res := []time.Duration{}
	for _, in := range input{
		start = time.Now()
		_ = f(in)
		t = time.Now()
		res = append(res, t.Sub(start))
	}
	return res
}

func TestIntQuery(id []int, f func(i int)([]person))[]time.Duration{
	res := []time.Duration{}
	for _, in := range id{
		start = time.Now()
		_ = f(in)
		t = time.Now()
		res = append(res, t.Sub(start))
	}
	return res

}

func TestIntArrayQuery(id []int, f func(i []int)([]person))[]time.Duration{
	res := []time.Duration{}

	start = time.Now()
	_ = f(id)
	t = time.Now()
	res = append(res, t.Sub(start))

	return res

}
