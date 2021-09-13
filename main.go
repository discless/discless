package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/discless/discless/api"
	"log"
	"math/big"
	"net/http"
	"time"
)

var(
	cert 	[]byte
	key		[]byte
)

func genTLS() {
	priv, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Discless"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 180),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	certB, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		log.Fatalf("Failed to create certificate: %s", err)
	}
	out := &bytes.Buffer{}
	pem.Encode(out, &pem.Block{Type: "CERTIFICATE", Bytes: certB})
	cert = out.Bytes()
	kout := &bytes.Buffer{}
	b, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		panic(err)
	}
	pem.Encode(kout, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: b})
	key = kout.Bytes()
}

func main()  {
	genTLS()

	functionHandler := http.HandlerFunc(api.Apply)
	botHandler := http.HandlerFunc(api.NewBot)

	mux := http.NewServeMux()
	mux.Handle("/function", functionHandler)
	mux.Handle("/bot", botHandler)

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},

		Certificates: []tls.Certificate{},
	}

	kp, err := tls.X509KeyPair(cert,key)

	if err != nil {
		panic(err)
	}

	cfg.Certificates = append(cfg.Certificates, kp)

	srv := &http.Server{
		Addr:         "localhost:8443",
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Fatal(srv.ListenAndServeTLS("", ""))
}