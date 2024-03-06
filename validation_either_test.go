// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation_test

import (
	"context"
	stderrors "errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/validation"
)

var _ = Describe("Either", func() {
	var ctx context.Context
	var err error
	var valid validation.HasValidation
	var inValid validation.HasValidation
	var either validation.Either
	BeforeEach(func() {
		ctx = context.Background()
		valid = validation.HasValidationFunc(func(ctx context.Context) error {
			return nil
		})
		inValid = validation.HasValidationFunc(func(ctx context.Context) error {
			return stderrors.New("banana")
		})
	})

	JustBeforeEach(func() {
		err = either.Validate(ctx)
	})
	Context("one valid", func() {
		BeforeEach(func() {
			either = validation.Either{
				inValid,
				valid,
				inValid,
			}
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
	})
	Context("multi valid", func() {
		BeforeEach(func() {
			either = validation.Either{
				valid,
				valid,
				inValid,
			}
		})
		It("returns error", func() {
			Expect(err).NotTo(BeNil())
		})
	})
	Context("all invalid", func() {
		BeforeEach(func() {
			either = validation.Either{
				inValid,
				inValid,
				inValid,
			}
		})
		It("returns error", func() {
			Expect(err).NotTo(BeNil())
		})
	})
	Context("empty", func() {
		BeforeEach(func() {
			either = validation.Either{}
		})
		It("returns error", func() {
			Expect(err).NotTo(BeNil())
		})
	})
	Context("one invalid", func() {
		BeforeEach(func() {
			either = validation.Either{
				inValid,
			}
		})
		It("returns error", func() {
			Expect(err).NotTo(BeNil())
		})
	})
	Context("one valid", func() {
		BeforeEach(func() {
			either = validation.Either{
				valid,
			}
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
	})
})
