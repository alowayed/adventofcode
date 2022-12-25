package interval

import (
	"reflect"
	"testing"
)

func TestValid(t *testing.T) {
	type args struct {
		i Interval
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{{
		name: "single",
		args: args{
			i: Interval{Start: 1, End: 1},
		},
		wantErr: false,
	}, {
		name: "multi",
		args: args{
			i: Interval{Start: 1, End: 2},
		},
		wantErr: false,
	}, {
		name: "invalid",
		args: args{
			i: Interval{Start: 2, End: 1},
		},
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Valid(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSubInterval(t *testing.T) {
	type args struct {
		inner  Interval
		outter Interval
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{{
		name: "outter > inner",
		args: args{
			inner:  Interval{Start: 2, End: 3},
			outter: Interval{Start: 1, End: 4},
		},
		want: true,
	}, {
		name: "outter == inner",
		args: args{
			inner:  Interval{Start: 2, End: 3},
			outter: Interval{Start: 2, End: 3},
		},
		want: true,
	}, {
		name: "outter < inner",
		args: args{
			inner:  Interval{Start: 1, End: 4},
			outter: Interval{Start: 2, End: 3},
		},
		want: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SubInterval(tt.args.inner, tt.args.outter)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubInterval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SubInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []Interval
		wantErr bool
	}{{
		name: "single",
		args: args{
			s: "1-2",
		},
		want: []Interval{
			{Start: 1, End: 2},
		},
	}, {
		name: "double",
		args: args{
			s: "1-2,3-4",
		},
		want: []Interval{
			{Start: 1, End: 2},
			{Start: 3, End: 4},
		},
	}, {
		name: "multiple",
		args: args{
			s: "1-2,3-4,6-9",
		},
		want: []Interval{
			{Start: 1, End: 2},
			{Start: 3, End: 4},
			{Start: 6, End: 9},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOverlap(t *testing.T) {
	type args struct {
		a Interval
		b Interval
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{{
		name: "no overlap",
		args: args{
			a: Interval{Start: 1, End: 2},
			b: Interval{Start: 3, End: 4},
		},
		want: false,
	}, {
		name: "partial",
		args: args{
			a: Interval{Start: 1, End: 2},
			b: Interval{Start: 2, End: 3},
		},
		want: true,
	}, {
		name: "full",
		args: args{
			a: Interval{Start: 1, End: 4},
			b: Interval{Start: 2, End: 3},
		},
		want: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Overlap(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Overlap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Overlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
