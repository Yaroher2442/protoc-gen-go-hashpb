// Copyright 2021-2022 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(Generate)
}
