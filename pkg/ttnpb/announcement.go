// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ttnpb

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net"
	"strings"
)

// GetTLSConfig returns the TLS config for the announcement.
func (a *Announcement) GetTLSConfig() (*tls.Config, error) {
	if a.Certificate == "" {
		return nil, nil
	}
	if a.NetAddress == "" {
		return nil, errors.New("No address known for this component")
	}
	netAddress := strings.Split(a.NetAddress, ",")[0]
	netHost, _, _ := net.SplitHostPort(netAddress)
	rootCAs := x509.NewCertPool()
	ok := rootCAs.AppendCertsFromPEM([]byte(a.Certificate))
	if !ok {
		return nil, errors.New("Could not read component certificate")
	}
	return &tls.Config{ServerName: netHost, RootCAs: rootCAs}, nil
}
