package mapreduce

import (
	"reflect"
	"sort"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		mapFunc func(int) (int, error)
		inCh    chan int
	}
	tests := []struct {
		name    string
		args    args
		in      []int
		want    []int
		wantErr error
	}{{
		name: "Square",
		args: args{
			mapFunc: func(n int) (int, error) { return n * n, nil },
			inCh:    make(chan int),
		},
		in:      []int{1, 2, 3, 4, 5, 6},
		want:    []int{1, 4, 9, 16, 25, 36},
		wantErr: nil,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			outCh, errCh := Map(tt.args.mapFunc, tt.args.inCh)

			for _, i := range tt.in {
				tt.args.inCh <- i
			}
			close(tt.args.inCh)

			out := []int{}
			var err error
			for {
				o := 0
				ok := false
				select {
				case o, ok = <-outCh:
					if !ok {
						outCh = nil
					} else {
						out = append(out, o)
					}
				case err, ok = <-errCh:
					if !ok {
						errCh = nil
					}
				}
				if outCh == nil && errCh == nil {
					break
				}
			}

			if err != tt.wantErr {
				t.Fatalf("Map() got error: %v, want error: %v", err, tt.wantErr)
			}

			sort.Ints(out)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("Map() = %v, want: %v", out, tt.want)
			}
		})
	}
}
