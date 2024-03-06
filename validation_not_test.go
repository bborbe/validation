// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation_test

import (
	"context"
	stderrors "errors"

	"github.com/bborbe/errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/validation"
	"github.com/bborbe/validation/mocks"
)

var _ = Describe("Not", func() {
	var ctx context.Context
	var err error
	var hasValidation *mocks.ValidationHasValidation
	BeforeEach(func() {
		ctx = context.Background()
		hasValidation = &mocks.ValidationHasValidation{}
	})
	JustBeforeEach(func() {
		err = validation.Not(hasValidation).Validate(ctx)
	})
	Context("valid", func() {
		BeforeEach(func() {
			hasValidation.ValidateReturns(validation.Error)
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
	})
	Context("invalid", func() {
		BeforeEach(func() {
			hasValidation.ValidateReturns(nil)
		})
		It("returns error", func() {
			Expect(err).NotTo(BeNil())
			Expect(errors.Is(err, validation.Error)).To(BeTrue())
		})
	})
	Context("error", func() {
		BeforeEach(func() {
			hasValidation.ValidateReturns(stderrors.New("banana"))
		})
		It("returns error", func() {
			Expect(err).NotTo(BeNil())
			Expect(errors.Is(err, validation.Error)).To(BeFalse())
		})
	})

})
