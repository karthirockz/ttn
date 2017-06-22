// Copyright © 2017 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package types

// LocationMetadata contains GPS coordinates
type LocationMetadata struct {
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
	Altitude  int32   `json:"altitude,omitempty"`
	Source    string  `json:"location_source,omitempty"` // can be: gps, config, registry, estimate or unknown (unknown may be left out)
}
