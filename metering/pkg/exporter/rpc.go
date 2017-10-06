package exporter

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"github.com/tangfeixiong/go-to-docker/metering/pb"
)

type client_gRPC struct {
	address string
	tls     bool
}

func NewClient_gRPC(address string) *client_gRPC {
	c := &client_gRPC{
		address: address,
	}
	return c
}

func NewClientTLS_gRPC(address string) {
	c := NewClient_gRPC(address)
	demoTLS()
	c.tls = true
}

func (c *client_gRPC) Transit(content *pb.MeteringReqResp) (*pb.MeteringReqResp, error) {
	if c.tls {
		return c.transitTLS(content)
	}
	address := c.address
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Printf("fail to dial: %v\n", err)
	}
	defer conn.Close()
	client := pb.NewCollectorServiceClient(conn)

	copts := []grpc.CallOption{}
	// copts := append(copts, grpc.EmptyCallOption{})
	resp, err := client.Transit(context.Background(), content, copts...)
	return resp, err
}

func (c *client_gRPC) transitTLS(content *pb.MeteringReqResp) (*pb.MeteringReqResp, error) {
	println("grpc with tls")
	var opts []grpc.DialOption
	creds := credentials.NewClientTLSFromCert(demoCertPool, demoAddr)
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(demoAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCollectorServiceClient(conn)

	copts := []grpc.CallOption{}
	// copts := append(copts, grpc.EmptyCallOption{})
	resp, err := client.Transit(context.Background(), content, copts...)
	println(resp, err)
	return resp, err
}

const (
	port = 12305
)

var (
	demoKeyPair  *tls.Certificate
	demoCertPool *x509.CertPool
	demoAddr     string
)

func demoTLS() {
	var err error
	pair, err := tls.X509KeyPair([]byte(Cert), []byte(Key))
	if err != nil {
		panic(err)
	}
	demoKeyPair = &pair
	demoCertPool = x509.NewCertPool()
	ok := demoCertPool.AppendCertsFromPEM([]byte(Cert))
	if !ok {
		panic("bad certs")
	}
	demoAddr = fmt.Sprintf("localhost:%d", port)
}

const (
	Key = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEA2KEH2CLMVmPLGwuExXDUk4TdQInkD0AAe22a1ixKxmRdSXkf
lDKu4cHL0N5ohJBonL/udGWkI6AI51lvJ54zl9SxPoZdWzkdca5H9Cw/7Es3SYyd
eVBkEo341D4hZqpFJM+vPReCQyTCHKGR6XnmN8E9V2RhTKpiagqfSpNmj08Rg3o3
Ap7IzaRRnqmGyKLgJEC2hEXj7nIvHtpm4SCjm2px4bRKQRxREuujOGfqBDh6uYQF
L46PI+QKWDN5J9+fROU9Y0HRFy7JnTkNyTuUxhjB5r2KReQ7cGGpUJzRIyyUDp/u
3+N58AoyQHf/bXpWE0oSfOhwNABVaRWWgxOP2wIDAQABAoIBAEyvtLcummd6rEvG
qBm894PBZRGTvsgnQARxqH5o74+Lt/pqhmNQDdidYiluklFbTa0vxJov4Qs7e+tq
HY9I0brN8HDR3/qLHYFA0Pf/MiHT/p5qyNRJQSPQXmXEGM7fN9rwKnjV+acLPCwm
hiDApl7WaYCmaEtbhbtER19/Dq9sB2vwF2hWvaS8TkpMe6otPVO+3h8MZfrCEQCY
sKpZRJEGaHGX+Oe14NlmO+WhpCR9YXftq5En//zZie5TXiSM596IDriUYf9pWr99
Xeo1JggyrGkQvmGfi5u5qpieX4QBjBKdSqOhiaUKaLQ9T6+mnijp3Cfl/HAkhS5d
RQcmJwECgYEA6+gAvAvPaWxeTY2hiyZfL7trbwJAjrNPxCim6eJJZbhByXnG7oVT
JZQo0pNaE/Exszv0BoaxURKNHM5mU6dEZRgwaiY4UUuKyK2OkSHEu8sOwpXcGI7k
udjEbRaeYBcNAmB25qgisubjbveDJiwK1hbG9T+pSn3E+VotRU2Vk3UCgYEA6xSu
wxSq79llO25O5BQ39DHKfTPeb6KBRdtCIM1L9FQH5CswVQcv1BwqvtvP724cQLkR
5fSQyG5G9qHXmh2dj9pB65h05wlO+F03DP1pQB25QyiogyrWwNmWGnrdsJRo1lAh
pbEWFP+/26n+VtFBcbDcClQsvPL0gOz7hiAuvA8CgYB9SYwKUuNnBAzZd1zSQCDR
guI95J5Qq16zuTtcf7endEJMNIa4asqL7LH5lBSE/tX8cNzbEnHdstKK9/tUdkNW
xZAA8Cd81XfxuGs9HQgVDHTcVya7TDihk0RPA3I9akCYgI7lVWqIRSOI7Z8TiNSA
ezxTR+orC7yvCXt9kQTdeQKBgF2z+dFCzLwcMJDW8FVThdYtfqQXZ8Ohx9ubgSlo
C62RTS/y0yohWjw3GgbHwYOTpWlbG7pImOl7o4etjS4ePe7YNcx+EaMB/9tZ9JaV
8D0hW/ZcH4dhLQbj9EQL05AOKBe9CxxrkPy/0K7zfLEIagiyUZNAaDDMuw8k50FY
VKibAoGAMLZlWDtCHA4J5GhLqRFzOzt2I650EOu/kNhGtJ/8YybgtMVaoN50PGfk
Dr7+TS/DxJzY7h0yNakDg6KZKT4U4qLh74VFaHCnADyQfQnJK+1ffhNhdeoSzp+L
zpDUVEXH6eEeRWmyxoWjWnsquube0gRKf2BQ+yYjk+CUwL/Aqk4=
-----END RSA PRIVATE KEY-----`
	Cert = `-----BEGIN CERTIFICATE-----
MIIEBjCCAu6gAwIBAgIJALzaDcEdLBD7MA0GCSqGSIb3DQEBBQUAMF8xCzAJBgNV
BAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQxGDAWBgNVBAMTD2xvY2FsaG9zdDoxMDAwMDAeFw0xNjAy
MTgwMzU5NDJaFw0yNjAyMTUwMzU5NDJaMF8xCzAJBgNVBAYTAkFVMRMwEQYDVQQI
EwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQx
GDAWBgNVBAMTD2xvY2FsaG9zdDoxMDAwMDCCASIwDQYJKoZIhvcNAQEBBQADggEP
ADCCAQoCggEBANihB9gizFZjyxsLhMVw1JOE3UCJ5A9AAHttmtYsSsZkXUl5H5Qy
ruHBy9DeaISQaJy/7nRlpCOgCOdZbyeeM5fUsT6GXVs5HXGuR/QsP+xLN0mMnXlQ
ZBKN+NQ+IWaqRSTPrz0XgkMkwhyhkel55jfBPVdkYUyqYmoKn0qTZo9PEYN6NwKe
yM2kUZ6phsii4CRAtoRF4+5yLx7aZuEgo5tqceG0SkEcURLrozhn6gQ4ermEBS+O
jyPkClgzeSffn0TlPWNB0RcuyZ05Dck7lMYYwea9ikXkO3BhqVCc0SMslA6f7t/j
efAKMkB3/216VhNKEnzocDQAVWkVloMTj9sCAwEAAaOBxDCBwTAdBgNVHQ4EFgQU
7JqKxmk2/4aClcix32bvTr0MUkQwgZEGA1UdIwSBiTCBhoAU7JqKxmk2/4aClcix
32bvTr0MUkShY6RhMF8xCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRl
MSEwHwYDVQQKExhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQxGDAWBgNVBAMTD2xv
Y2FsaG9zdDoxMDAwMIIJALzaDcEdLBD7MAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcN
AQEFBQADggEBAGo0MdEPAV6EH2mhIXBJb6qjg7X0kGqmh10UzmNc/r4N0lcfoPc3
q91N3tAk2zxASW16FPumd3eRtn5FdEWLTK2SAJkP24g6199pUbcEvzHas5/awRI3
PFwNJ+cqsYkXxsW09/cvRBFqMqrkavvoMfCwQhMJwGnql+BeN4mBS00JglHWSfDT
e8T2yhkPc0+FuAH4ZfmdZUb+yPAv+liT+lCw+vUEsN8mnam8lZKCzhROVfmgKEHM
Ze0aj9tzK3Su1tjAEzN4arrajCopkJA2aDI2i8EZ+2Zx1qbhNXwJd3E9MYs9WmLf
RX7r0aSW3Y9r+/SmjYJLXB36CwbcjLHmQN0=
-----END CERTIFICATE-----`
)
