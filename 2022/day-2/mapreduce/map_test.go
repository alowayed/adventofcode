package mapreduce

import (
	"errors"
	"testing"
)

var (
	errAny = errors.New("any error")
)

func TestRPCMapper_Map(t *testing.T) {

	type args struct {
		input []string
	}

	tests := []struct {
		name    string
		m       *RPCMapper
		args    args
		want    int
		wantErr error
	}{{
		name: "1 round",
		m:    &RPCMapper{},
		args: args{
			input: []string{
				"A Y",
			},
		},
		want:    4,
		wantErr: nil,
	}, {
		name: "2 rounds",
		m:    &RPCMapper{},
		args: args{
			input: []string{
				"A Y",
				"B X",
			},
		},
		want:    5,
		wantErr: nil,
	}, {
		name: "3 rounds",
		m:    &RPCMapper{},
		args: args{
			input: []string{
				"A Y",
				"B X",
				"C Z",
			},
		},
		want:    12,
		wantErr: nil,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &RPCMapper{}
			got, err := m.Map(tt.args.input)
			if errors.Is(tt.wantErr, errAny) && err != nil {
				err = errAny
			}
			if err != tt.wantErr {
				t.Errorf("RPCMapper.Map() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RPCMapper.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
