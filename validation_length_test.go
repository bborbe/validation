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

var _ = DescribeTable("LengthEqual",
	func(list []any, length int, expectedError bool) {
		hasValidation := validation.LengthEqual(list, length)
		if expectedError {
			Expect(hasValidation.Validate(context.Background())).NotTo(BeNil())
		} else {
			Expect(hasValidation.Validate(context.Background())).To(BeNil())
		}
	},
	Entry("3 == 2", []any{"a", "b", "c"}, 2, true),
	Entry("3 == 3", []any{"a", "b", "c"}, 3, false),
	Entry("3 == 4", []any{"a", "b", "c"}, 4, true),
)

var _ = DescribeTable("LengthLt",
	func(list []any, length int, expectedError bool) {
		hasValidation := validation.LengthLt(list, length)
		if expectedError {
			Expect(hasValidation.Validate(context.Background())).NotTo(BeNil())
		} else {
			Expect(hasValidation.Validate(context.Background())).To(BeNil())
		}
	},
	Entry("3 < 2", []any{"a", "b", "c"}, 2, true),
	Entry("3 < 3", []any{"a", "b", "c"}, 3, true),
	Entry("3 < 4", []any{"a", "b", "c"}, 4, false),
	Entry("3 < 5", []any{"a", "b", "c"}, 5, false),
)

var _ = DescribeTable("LengthLe",
	func(list []any, length int, expectedError bool) {
		hasValidation := validation.LengthLe(list, length)
		if expectedError {
			Expect(hasValidation.Validate(context.Background())).NotTo(BeNil())
		} else {
			Expect(hasValidation.Validate(context.Background())).To(BeNil())
		}
	},
	Entry("3 <= 2", []any{"a", "b", "c"}, 2, true),
	Entry("3 <= 3", []any{"a", "b", "c"}, 3, false),
	Entry("3 <= 4", []any{"a", "b", "c"}, 4, false),
	Entry("3 <= 5", []any{"a", "b", "c"}, 5, false),
)

var _ = DescribeTable("LengthGt",
	func(list []any, length int, expectedError bool) {
		hasValidation := validation.LengthGt(list, length)
		if expectedError {
			Expect(hasValidation.Validate(context.Background())).NotTo(BeNil())
		} else {
			Expect(hasValidation.Validate(context.Background())).To(BeNil())
		}
	},
	Entry("3 > 2", []any{"a", "b", "c"}, 2, false),
	Entry("3 > 3", []any{"a", "b", "c"}, 3, true),
	Entry("3 > 4", []any{"a", "b", "c"}, 4, true),
	Entry("3 > 5", []any{"a", "b", "c"}, 5, true),
)

var _ = DescribeTable("LengthGe",
	func(list []any, length int, expectedError bool) {
		hasValidation := validation.LengthGe(list, length)
		if expectedError {
			Expect(hasValidation.Validate(context.Background())).NotTo(BeNil())
		} else {
			Expect(hasValidation.Validate(context.Background())).To(BeNil())
		}
	},
	Entry("3 >= 2", []any{"a", "b", "c"}, 2, false),
	Entry("3 >= 3", []any{"a", "b", "c"}, 3, false),
	Entry("3 >= 4", []any{"a", "b", "c"}, 4, true),
	Entry("3 >= 5", []any{"a", "b", "c"}, 5, true),
)
