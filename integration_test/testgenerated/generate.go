// Copyright (c) 2018 Palantir Technologies. All rights reserved.
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

// This directory contains all of the tests that run using code generated by conjure-go. These tests are typically
// integration tests that make use of the code generated by conjure-go to verify its behavior. The generator generates
// all of the generated files that are used by these tests.

//go:generate go run $GOFILE auth/auth-service.yml auth
//go:generate go run $GOFILE client/client-service.yml client
//go:generate go run $GOFILE errors/errors.yml errors
//go:generate go run $GOFILE objects/objects.yml objects
//go:generate go run $GOFILE post/post-service.yml post
//go:generate go run $GOFILE queryparam/query-service.yml queryparam

package main

import (
	"fmt"
	"os"

	"github.com/palantir/godel-conjure-plugin/ir-gen-cli-bundler/conjureircli"

	"github.com/palantir/conjure-go/conjure"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("task requires input YAML file and output directory as arguments")
		os.Exit(1)
	}
	if err := run(os.Args[1], os.Args[2]); err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}
}

func run(in, out string) error {
	irBytes, err := conjureircli.InputPathToIR(in)
	if err != nil {
		return err
	}
	conjureDef, err := conjure.FromIRBytes(irBytes)
	if err != nil {
		return err
	}
	return conjure.Generate(conjureDef, out)
}