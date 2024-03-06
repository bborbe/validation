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

type HasValidationType float64

func (p HasValidationType) Validate(ctx context.Context) error {
	return nil
}

var _ = Describe("NotNil", func() {
	var err error
	var value validation.HasValidation
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
		value = validation.All{}
	})
	JustBeforeEach(func() {
		err = validation.NotNilAndValid(value).Validate(ctx)
	})
	Context("valid", func() {
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
	})
	Context("nil", func() {
		BeforeEach(func() {
			value = nil
		})
		It("returns error", func() {
			Expect(err).NotTo(BeNil())
		})
	})
	Context("validation.HasValidation", func() {
		BeforeEach(func() {
			var v *HasValidationType
			value = v
		})
		It("returns error", func() {
			Expect(err).NotTo(BeNil())
		})
	})
})
