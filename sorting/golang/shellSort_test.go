package sorting

import "testing"

func Test_shellSort(t *testing.T) {
	type args struct {
		items []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "insertion sort test 1 positive 25 items",
			args:args{items:generateSlice(25)},
		},
		{
			name: "insertion sort test 1 positive 100 items" ,
			args:args{items:generateSlice(100)},
		},
		{
			name: "insertion sort test 1 positive 1000 items" ,
			args:args{items:generateSlice(1000)},
		},
		{
			name: "insertion sort test 1 positive 10000 items" ,
			args:args{items:generateSlice(10000)},
		},
		{
			name: "insertion sort test 1 positive 100000 items" ,
			args:args{items:generateSlice(100000)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shellSort(tt.args.items)
		})
	}
}

func BenchmarkShellSortTest(b *testing.B) {
	genSlice := generateSlice(b.N)


	b.StartTimer()
	shellSort(genSlice)
}
