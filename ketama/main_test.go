package ketama

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRingBuild(t *testing.T) {
	var nodeList []Node
	for i := 0; i < 1000; i++ {
		nodeList = append(nodeList, Node{
			Ip:     fmt.Sprintf("127.0.0.%d", i),
			Port:   "8000",
			Weight: 100,
		})
	}
	type args struct {
		nodeList []Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{nodeList: nodeList},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RingBuild(tt.args.nodeList)
		})
	}
}

func TestNodeLocation(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want Node
	}{
		{
			name: "test1",
			args: args{key: "1"},
			want: Node{
				Ip:     "127.0.0.682",
				Port:   "8000",
				Weight: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NodeLocation(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
