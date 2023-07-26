package models

import (
	"net"
	"reflect"
	"testing"
)

func TestParseAddr(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name     string
		args     args
		wantAddr Addr
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "tcp4",
			args: args{
				address: "127.0.0.1:8080",
			},
			wantAddr: Addr{
				IP:   net.ParseIP("127.0.0.1"),
				Port: 8080,
			},
			wantErr: false,
		},
		{
			name: "tcp6",
			args: args{
				address: "[::1]:8080",
			},
			wantAddr: Addr{
				IP:   net.ParseIP("::1"),
				Port: 8080,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAddr, err := ParseAddr(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAddr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAddr, tt.wantAddr) {
				t.Errorf("ParseAddr() = %v, want %v", gotAddr, tt.wantAddr)
			}
		})
	}
}

func TestParseAddr2(t *testing.T) {
	Addr, err := ParseAddr("127.0.0.1:8080")
	if err != nil {
		t.Fatalf("ParseAddr() error = %v", err)
	}
	t.Log(Addr, Addr.Network())
}
