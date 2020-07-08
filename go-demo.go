/*
A plugin representing all events handled by the Kong Go plugin ecosystem.

The docs are relatively dispersed as the Go plugin ecosystem is relatively new.

This link provides the list of events handled by the Go plugin system:
https://docs.konghq.com/2.0.x/go/#handling-events

This link provides more detail about those events:
https://docs.konghq.com/2.0.x/plugin-development/custom-logic/#available-contexts

Function comments are stolen from the above link!
*/
package main

import (
	"github.com/Kong/go-pdk"
)

// Config contains the schema defined for the Kong Admin API to populate data
// https://docs.konghq.com/2.0.x/go/#configuration-structure
type Config struct {
	Path   string
	Reopen bool
}

// New ensures that Kong knows what your configuration is
// https://docs.konghq.com/2.0.x/go/#new-constructor
func New() interface{} {
	return &Config{}
}

// Certificate is executed during the SSL certificate serving phase of the SSL handshake.
// HTTP/HTTPS requests
func (conf Config) Certificate(kong *pdk.PDK) {
	kong.Log.Info("go-demo, Certificate event")
}

// Rewrite is executed for every request upon its reception from a client as a rewrite phase handler.
// NOTE in this phase neither the Service nor the Consumer have been identified, hence this handler will only be
// executed if the plugin was configured as a global plugin!
// HTTP/HTTPS requests
func (conf Config) Rewrite(kong *pdk.PDK) {
	kong.Log.Info("go-demo, Rewrite event")
}

// Access is executed for every request from a client and before it is being proxied to the upstream service.
// HTTP/HTTPS requests
func (conf Config) Access(kong *pdk.PDK) {
	kong.Log.Info("go-demo, Access event")
}

// Preread is executed once for every connection.
// Streaming requests
func (conf Config) Preread(kong *pdk.PDK) {
	kong.Log.Info("go-demo, Preread event")
}

// Log is executed when the last response byte has been sent to the client for HTTP/HTTPS,
// or executed once for each connection after it has been closed for streaming.
// HTTP/HTTPS and Streaming requests
func (conf Config) Log(kong *pdk.PDK) {
	kong.Log.Info("go-demo, Log event")
}
