package network

import (
	"net"
	"reflect"
	"testing"
)

func Test_udpAddr(t *testing.T) {
	addr := udpAddr(":2022")
	if addr == nil {
		// it is not necessary, if addr is not valid application is exit with fatal error
		t.Error("failed addr")
	}
}

func TestServer(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    *Connection
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Server(tt.args.addr)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Server() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Server() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestClient(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    *Connection
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Client(tt.args.addr)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Client() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Client() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestConnection_carrier(t *testing.T) {
	type fields struct {
		Conn *net.UDPConn
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		c := &Connection{
			Conn: tt.fields.Conn,
		}
		c.carrier()
	}
}

func Test_parseRequest(t *testing.T) {
	type args struct {
		bRequest []byte
	}
	tests := []struct {
		name    string
		args    args
		want    request
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := parseRequest(tt.args.bRequest)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. parseRequest() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. parseRequest() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_request_getRPC(t *testing.T) {
	type fields struct {
		Struct string
		Method string
		Packet []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
		want1  reflect.Value
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		r := &request{
			Struct: tt.fields.Struct,
			Method: tt.fields.Method,
			Packet: tt.fields.Packet,
		}
		got, got1 := r.getRPC()
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. request.getRPC() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. request.getRPC() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func Test_request_call(t *testing.T) {
	type fields struct {
		Struct string
		Method string
		Packet []byte
	}
	type args struct {
		rpcInterface interface{}
		method       reflect.Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		r := &request{
			Struct: tt.fields.Struct,
			Method: tt.fields.Method,
			Packet: tt.fields.Packet,
		}
		r.call(tt.args.rpcInterface, tt.args.method)
	}
}

func TestClose(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		Close()
	}
}

func TestConnection_Close(t *testing.T) {
	type fields struct {
		Conn *net.UDPConn
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		c := &Connection{
			Conn: tt.fields.Conn,
		}
		c.Close()
	}
}

func TestReply(t *testing.T) {
	type args struct {
		rpcname string
		data    []byte
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		Reply(tt.args.rpcname, tt.args.data)
	}
}

func TestConnection_Reply(t *testing.T) {
	type fields struct {
		Conn *net.UDPConn
	}
	type args struct {
		rpcname string
		data    []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		c := &Connection{
			Conn: tt.fields.Conn,
		}
		c.Reply(tt.args.rpcname, tt.args.data)
	}
}

func TestSetRPC(t *testing.T) {
	type args struct {
		userRPC []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		SetRPC(tt.args.userRPC...)
	}
}

func Test_newRPC(t *testing.T) {
	type args struct {
		c interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantRpc RPC
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if gotRpc := newRPC(tt.args.c); !reflect.DeepEqual(gotRpc, tt.wantRpc) {
			t.Errorf("%q. newRPC() = %v, want %v", tt.name, gotRpc, tt.wantRpc)
		}
	}
}

func Test_getRPCname(t *testing.T) {
	type args struct {
		c interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := getRPCname(tt.args.c); got != tt.want {
			t.Errorf("%q. getRPCname() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
