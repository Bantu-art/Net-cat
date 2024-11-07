package netcat

import (
	"net"
	"time"
)

// mock net.Conn
type MockTcpConn struct {
	InputBuffer  []byte // holds the name that will be returned as the eventual client name in Client struct
	OutputBuffer []byte // just need it for the write function
}

func (conn *MockTcpConn) Read(b []byte) (n int, err error) {
	copy(b, conn.InputBuffer)
	return len(conn.InputBuffer), nil
}

func (conn *MockTcpConn) Write(b []byte) (n int, err error) {
	conn.OutputBuffer = append(conn.OutputBuffer, b...)
	return len(b), nil
}

func (conn MockTcpConn)GetString() string {
	return string(conn.OutputBuffer)
}

// implement the rest of the functions to satisfy the net.Conn interface
func (conn MockTcpConn) Close() error                       { return nil }
func (conn MockTcpConn) LocalAddr() net.Addr                { return nil }
func (conn MockTcpConn) RemoteAddr() net.Addr               { return nil }
func (conn MockTcpConn) SetDeadline(t time.Time) error      { return nil }
func (conn MockTcpConn) SetReadDeadline(t time.Time) error  { return nil }
func (conn MockTcpConn) SetWriteDeadline(t time.Time) error { return nil }
