// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package peer provides functions for extracting information from a
// CER or CEA, and associating with a Context.
//
// Example:
//
//	func handleXYZ(c diam.Conn, m *diam.Message) {
//		meta, ok := peer.FromContext(c.Context())
//		if ok {
//			log.Println(meta)
//		}
//	}
//
// See the Metadata type for details.
package peer
