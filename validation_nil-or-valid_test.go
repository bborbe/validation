// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/validation"
)

var _ = Describe("NilOrValid", func() {
	var ctx context.Context
	var err error
	var value validation.HasValidation
	BeforeEach(func() {
		ctx = context.Background()
	})
	JustBeforeEach(func() {
		err = validation.NilOrValid(value).Validate(ctx)
	})
	Context("nil", func() {
		BeforeEach(func() {
			value = nil
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
	})
	Context("valid", func() {
		BeforeEach(func() {
			value = validation.HasValidationFunc(func(ctx context.Context) error {
				return nil
			})
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
	})
	Context("invalid", func() {
		BeforeEach(func() {
			value = validation.HasValidationFunc(func(ctx context.Context) error {
				return validation.Error
			})
		})
		It("returns error", func() {
			Expect(err).NotTo(BeNil())
		})
	})
})
