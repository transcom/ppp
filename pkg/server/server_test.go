package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/testingsuite"
)

type serverSuite struct {
	logger Logger
	testingsuite.BaseTestSuite
	httpHandler http.Handler
}

func TestServerSuite(t *testing.T) {
	var httpHandler http.Handler
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Panic(err)
	}
	ss := &serverSuite{
		logger:      logger,
		httpHandler: httpHandler,
	}
	suite.Run(t, ss)
}

func (suite *serverSuite) readFile(filename string) []byte {
	testDataDir := "testdata"
	filePath := strings.Join([]string{testDataDir, filename}, "/")

	contents, err := ioutil.ReadFile(filepath.Clean(filePath))
	if err != nil {
		suite.T().Fatalf("failed to read file %s: %s", filename, err)
	}
	return contents

}

func (suite *serverSuite) TestParseSingleTLSCert() {

	keyPair, err := tls.X509KeyPair(
		suite.readFile("localhost.pem"),
		suite.readFile("localhost.key"))

	suite.NoError(err)

	httpsServer, err := CreateNamedServer(&CreateNamedServerInput{
		Host:         "127.0.0.1",
		Port:         8443,
		ClientAuth:   tls.NoClientCert,
		HTTPHandler:  suite.httpHandler,
		Logger:       suite.logger,
		Certificates: []tls.Certificate{keyPair},
	})
	suite.NoError(err)
	suite.Equal(len(httpsServer.TLSConfig.Certificates), 1)
	suite.Contains(httpsServer.TLSConfig.NameToCertificate, "localhost")
}

func (suite *serverSuite) TestParseBadTLSCert() {

	_, err := tls.X509KeyPair(
		suite.readFile("localhost-bad.pem"),
		suite.readFile("localhost.key"))

	suite.NotNil(err)
}

func (suite *serverSuite) TestParseMultipleTLSCerts() {

	keyPairLocalhost, err := tls.X509KeyPair(
		suite.readFile("localhost.pem"),
		suite.readFile("localhost.key"))

	suite.NoError(err)

	keyPairOffice, err := tls.X509KeyPair(
		suite.readFile("officelocal.pem"),
		suite.readFile("officelocal.key"))

	suite.NoError(err)

	httpsServer, err := CreateNamedServer(&CreateNamedServerInput{
		Host:        "127.0.0.1",
		Port:        8443,
		ClientAuth:  tls.NoClientCert,
		HTTPHandler: suite.httpHandler,
		Logger:      suite.logger,
		Certificates: []tls.Certificate{
			keyPairLocalhost,
			keyPairOffice,
		},
	})
	suite.NoError(err)
	suite.Equal(len(httpsServer.TLSConfig.Certificates), 2)
	suite.Contains(httpsServer.TLSConfig.NameToCertificate, "localhost")
	suite.Contains(httpsServer.TLSConfig.NameToCertificate, "officelocal")
}

func (suite *serverSuite) TestTLSConfigWithClientAuth() {

	keyPair, err := tls.X509KeyPair(
		suite.readFile("localhost.pem"),
		suite.readFile("localhost.key"))

	suite.NoError(err)

	caFile := suite.readFile("ca.pem")
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caFile)

	_, err = CreateNamedServer(&CreateNamedServerInput{
		Host:         "127.0.0.1",
		Port:         8443,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
		HTTPHandler:  suite.httpHandler,
		Logger:       suite.logger,
		Certificates: []tls.Certificate{keyPair},
	})
	suite.NoError(err)
}

func (suite *serverSuite) TestTLSConfigWithMissingCA() {

	keyPair, err := tls.X509KeyPair(
		suite.readFile("localhost.pem"),
		suite.readFile("localhost.key"))

	suite.NoError(err)

	_, err = CreateNamedServer(&CreateNamedServerInput{
		Host:         "127.0.0.1",
		Port:         8443,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		HTTPHandler:  suite.httpHandler,
		Logger:       suite.logger,
		Certificates: []tls.Certificate{keyPair},
	})
	suite.Equal(ErrMissingCACert, err)
}

func (suite *serverSuite) TestTLSConfigWithMisconfiguredCA() {

	keyPair, err := tls.X509KeyPair(
		suite.readFile("localhost.pem"),
		suite.readFile("localhost.key"))

	suite.NoError(err)

	caFile := suite.readFile("localhost-bad.pem")
	caCertPool := x509.NewCertPool()
	certOk := caCertPool.AppendCertsFromPEM(caFile)
	suite.False(certOk)

	_, err = CreateNamedServer(&CreateNamedServerInput{
		Host:         "127.0.0.1",
		Port:         8443,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
		HTTPHandler:  suite.httpHandler,
		Logger:       suite.logger,
		Certificates: []tls.Certificate{keyPair},
	})
	suite.Equal(ErrMissingCACert, err)
}

func (suite *serverSuite) TestHTTPServerConfig() {
	httpsServer, err := CreateNamedServer(&CreateNamedServerInput{
		Host:        "127.0.0.1",
		Port:        8080,
		HTTPHandler: suite.httpHandler,
		Logger:      suite.logger,
	})
	suite.NoError(err)
	suite.Equal(httpsServer.Addr, "127.0.0.1:8080")
	suite.Equal(suite.httpHandler, httpsServer.Handler)
}

func (suite *serverSuite) testTLSConfigWithRequest(tlsVersion uint16) {

	keyPair, err := tls.X509KeyPair(
		suite.readFile("localhost.pem"),
		suite.readFile("localhost.key"))
	suite.NoError(err)
	certificates := []tls.Certificate{keyPair}

	caFile := suite.readFile("ca.pem")
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caFile)

	// A handler that we can test with
	htmlBody := "<html><body>Hello, client</body></html>"
	httpHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the TLS connection from inside the request
		connState := r.TLS
		suite.Equal(tlsVersion, connState.Version)
		suite.True(connState.HandshakeComplete)
		suite.False(connState.DidResume)
		suite.Equal("", connState.NegotiatedProtocol)
		suite.True(connState.NegotiatedProtocolIsMutual)
		suite.Equal("localhost", connState.ServerName)
		suite.Equal("Snake Oil", connState.PeerCertificates[0].Subject.Organization[0])
		suite.Equal("Snake Oil", connState.VerifiedChains[0][0].Subject.Organization[0])

		// Now write out a message
		fmt.Fprintln(w, htmlBody)
	})

	host := "localhost"
	port := 7443
	srv, err := CreateNamedServer(&CreateNamedServerInput{
		Host:         host,
		Port:         port,
		HTTPHandler:  httpHandler,
		Logger:       suite.logger,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: certificates,
		ClientCAs:    caCertPool,
	})
	defer srv.Close()
	suite.NoError(err)

	// Start the Server
	go srv.ListenAndServeTLS()

	srv.WaitUntilReady()

	// Send a request
	clientTLSConfig := tls.Config{
		RootCAs:      caCertPool,
		Certificates: certificates,
		MinVersion:   tlsVersion,
		MaxVersion:   tlsVersion,
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &clientTLSConfig,
		},
	}
	res, err := client.Get(fmt.Sprintf("https://%s:%d", host, port))
	suite.NoError(err)

	// Read the response
	if res != nil {
		body, bodyErr := ioutil.ReadAll(res.Body)
		res.Body.Close()
		suite.NoError(bodyErr)
		suite.Equal(htmlBody+"\n", string(body))
	}

	// Check the TLS connection directly
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", host, port), &clientTLSConfig)
	defer conn.Close()
	suite.NoError(err)
}

func (suite *serverSuite) TestTLSConfigWithRequest() {
	var versions = map[string]uint16{"1.2": tls.VersionTLS12, "1.3": tls.VersionTLS13}
	for name, version := range versions {
		suite.T().Run(fmt.Sprintf("TLS version %s", name), func(t *testing.T) {
			suite.testTLSConfigWithRequest(version)
		})
	}
}

func (suite *serverSuite) TestTLSConfigWithRequestNoClientAuth() {

	keyPair, err := tls.X509KeyPair(
		suite.readFile("localhost.pem"),
		suite.readFile("localhost.key"))
	suite.NoError(err)
	certificates := []tls.Certificate{keyPair}

	caFile := suite.readFile("ca.pem")
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caFile)

	// A handler that we can test with
	httpHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("This handler should have never fired")
	})

	host := "localhost"
	port := 7443
	srv, err := CreateNamedServer(&CreateNamedServerInput{
		Host:         host,
		Port:         port,
		HTTPHandler:  httpHandler,
		Logger:       suite.logger,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: certificates,
		ClientCAs:    caCertPool,
	})
	defer srv.Close()
	suite.NoError(err)

	// Start the Server
	go srv.ListenAndServeTLS()

	srv.WaitUntilReady()

	// Send a request without TLS client side cert configuration, should return error
	client := &http.Client{
		Transport: &http.Transport{},
	}
	res, err := client.Get(fmt.Sprintf("https://%s:%d", host, port))
	suite.Nil(res)
	suite.Error(err)

	// Check the TLS connection directly
	// This should fail and conn should be nil
	config := tls.Config{}
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", host, port), &config)
	suite.Nil(conn)
	suite.Error(err)
}

func (suite *serverSuite) TestTLSConfigWithInvalidAuth() {

	keyPair, err := tls.X509KeyPair(
		suite.readFile("localhost.pem"),
		suite.readFile("localhost.key"))
	suite.NoError(err)
	certificates := []tls.Certificate{keyPair}

	caFile := suite.readFile("ca.pem")
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caFile)

	// A handler that we can test with
	httpHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("This handler should have never fired")
	})

	host := "localhost"
	port := 7443
	srv, err := CreateNamedServer(&CreateNamedServerInput{
		Host:         host,
		Port:         port,
		HTTPHandler:  httpHandler,
		Logger:       suite.logger,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: certificates,
		ClientCAs:    caCertPool,
	})
	defer srv.Close()
	suite.NoError(err)

	// Start the Server
	go srv.ListenAndServeTLS()

	srv.WaitUntilReady()

	// Send a request with self signed cert and CA, but doesn't match server used CA
	invalidKeyPair, err := tls.X509KeyPair(
		suite.readFile("localhost-invalid.pem"),
		suite.readFile("localhost-invalid.key"))
	suite.NoError(err)
	invalidCertificates := []tls.Certificate{invalidKeyPair}

	invalidCaFile := suite.readFile("invalid-ca.pem")
	invalidCaCertPool := x509.NewCertPool()
	invalidCaCertPool.AppendCertsFromPEM(invalidCaFile)

	config := tls.Config{
		RootCAs:      invalidCaCertPool,
		Certificates: invalidCertificates,
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &config,
		},
	}
	res, err := client.Get(fmt.Sprintf("https://%s:%d", host, port))
	suite.Nil(res)
	suite.Error(err)

	// Check the TLS connection directly
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", host, port), &config)
	suite.Nil(conn)
	suite.Error(err)
}
