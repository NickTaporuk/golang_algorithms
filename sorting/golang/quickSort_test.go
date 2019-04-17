package sorting

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		items []int
	}
	tests := []struct {
		name string
		args args
		want []int
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
			if got := quickSort(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}


func BenchmarkQuickTest(b *testing.B) {
	genSlice := generateSlice(b.N)


	b.StartTimer()
	quickSort(genSlice)
}
