// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package flex

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	fwtypes "github.com/hashicorp/terraform-provider-aws/internal/framework/types"
)

func TestExpand(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	testString := "test"
	testStringResult := "a"

	testARN := "arn:aws:securityhub:us-west-2:1234567890:control/cis-aws-foundations-benchmark/v/1.2.0/1.1" //lintignore:AWSAT003,AWSAT005

	testTimeStr := "2013-09-25T09:34:01Z"
	testTimeTime := errs.Must(time.Parse(time.RFC3339, testTimeStr))

	testCases := []struct {
		Context    context.Context //nolint:containedctx // testing context use
		TestName   string
		Source     any
		Target     any
		WantErr    bool
		WantTarget any
	}{
		{
			TestName: "nil Source and Target",
			WantErr:  true,
		},
		{
			TestName: "non-pointer Target",
			Source:   TestFlex00{},
			Target:   0,
			WantErr:  true,
		},
		{
			TestName: "non-struct Source",
			Source:   testString,
			Target:   &TestFlex00{},
			WantErr:  true,
		},
		{
			TestName: "non-struct Target",
			Source:   TestFlex00{},
			Target:   &testString,
			WantErr:  true,
		},
		{
			TestName:   "types.String to string",
			Source:     types.StringValue("a"),
			Target:     &testString,
			WantTarget: &testStringResult,
		},
		{
			TestName:   "empty struct Source and Target",
			Source:     TestFlex00{},
			Target:     &TestFlex00{},
			WantTarget: &TestFlex00{},
		},
		{
			TestName:   "empty struct pointer Source and Target",
			Source:     &TestFlex00{},
			Target:     &TestFlex00{},
			WantTarget: &TestFlex00{},
		},
		{
			TestName:   "single string struct pointer Source and empty Target",
			Source:     &TestFlexTF01{Field1: types.StringValue("a")},
			Target:     &TestFlex00{},
			WantTarget: &TestFlex00{},
		},
		{
			TestName: "does not implement attr.Value Source",
			Source:   &TestFlexAWS01{Field1: "a"},
			Target:   &TestFlexAWS01{},
			WantErr:  true,
		},
		{
			TestName:   "single string Source and single string Target",
			Source:     &TestFlexTF01{Field1: types.StringValue("a")},
			Target:     &TestFlexAWS01{},
			WantTarget: &TestFlexAWS01{Field1: "a"},
		},
		{
			TestName:   "single string Source and single *string Target",
			Source:     &TestFlexTF01{Field1: types.StringValue("a")},
			Target:     &TestFlexAWS02{},
			WantTarget: &TestFlexAWS02{Field1: aws.String("a")},
		},
		{
			TestName:   "single string Source and single int64 Target",
			Source:     &TestFlexTF01{Field1: types.StringValue("a")},
			Target:     &TestFlexAWS03{},
			WantTarget: &TestFlexAWS03{},
		},
		{
			TestName: "primtive types Source and primtive types Target",
			Source: &TestFlexTF03{
				Field1:  types.StringValue("field1"),
				Field2:  types.StringValue("field2"),
				Field3:  types.Int64Value(3),
				Field4:  types.Int64Value(-4),
				Field5:  types.Int64Value(5),
				Field6:  types.Int64Value(-6),
				Field7:  types.Float64Value(7.7),
				Field8:  types.Float64Value(-8.8),
				Field9:  types.Float64Value(9.99),
				Field10: types.Float64Value(-10.101),
				Field11: types.BoolValue(true),
				Field12: types.BoolValue(false),
			},
			Target: &TestFlexAWS04{},
			WantTarget: &TestFlexAWS04{
				Field1:  "field1",
				Field2:  aws.String("field2"),
				Field3:  3,
				Field4:  aws.Int32(-4),
				Field5:  5,
				Field6:  aws.Int64(-6),
				Field7:  7.7,
				Field8:  aws.Float32(-8.8),
				Field9:  9.99,
				Field10: aws.Float64(-10.101),
				Field11: true,
				Field12: aws.Bool(false),
			},
		},
		{
			TestName: "List/Set/Map of primitive types Source and slice/map of primtive types Target",
			Source: &TestFlexTF04{
				Field1: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("a"),
					types.StringValue("b"),
				}),
				Field2: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("a"),
					types.StringValue("b"),
				}),
				Field3: types.SetValueMust(types.StringType, []attr.Value{
					types.StringValue("a"),
					types.StringValue("b"),
				}),
				Field4: types.SetValueMust(types.StringType, []attr.Value{
					types.StringValue("a"),
					types.StringValue("b"),
				}),
				Field5: types.MapValueMust(types.StringType, map[string]attr.Value{
					"A": types.StringValue("a"),
					"B": types.StringValue("b"),
				}),
				Field6: types.MapValueMust(types.StringType, map[string]attr.Value{
					"A": types.StringValue("a"),
					"B": types.StringValue("b"),
				}),
			},
			Target: &TestFlexAWS05{},
			WantTarget: &TestFlexAWS05{
				Field1: []string{"a", "b"},
				Field2: aws.StringSlice([]string{"a", "b"}),
				Field3: []string{"a", "b"},
				Field4: aws.StringSlice([]string{"a", "b"}),
				Field5: map[string]string{"A": "a", "B": "b"},
				Field6: aws.StringMap(map[string]string{"A": "a", "B": "b"}),
			},
		},
		{
			TestName: "plural field names",
			Source: &TestFlexTF09{
				City: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("paris"),
					types.StringValue("london"),
				}),
				Coach: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("guardiola"),
					types.StringValue("mourinho"),
				}),
				Tomato: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("brandywine"),
					types.StringValue("roma"),
				}),
				Vertex: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("ab"),
					types.StringValue("bc"),
				}),
				Criterion: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("votes"),
					types.StringValue("editors"),
				}),
				Datum: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("d1282f78-fa99-5d9d-bd51-e6f0173eb74a"),
					types.StringValue("0f10cb10-2076-5254-bd21-d3f62fe66303"),
				}),
				Hive: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("Cegieme"),
					types.StringValue("Fahumvid"),
				}),
			},
			Target: &TestFlexAWS11{},
			WantTarget: &TestFlexAWS11{
				Cities: []*string{
					aws.String("paris"),
					aws.String("london"),
				},
				Coaches: []*string{
					aws.String("guardiola"),
					aws.String("mourinho"),
				},
				Tomatoes: []*string{
					aws.String("brandywine"),
					aws.String("roma"),
				},
				Vertices: []*string{
					aws.String("ab"),
					aws.String("bc"),
				},
				Criteria: []*string{
					aws.String("votes"),
					aws.String("editors"),
				},
				Data: []*string{
					aws.String("d1282f78-fa99-5d9d-bd51-e6f0173eb74a"),
					aws.String("0f10cb10-2076-5254-bd21-d3f62fe66303"),
				},
				Hives: []*string{
					aws.String("Cegieme"),
					aws.String("Fahumvid"),
				},
			},
		},
		{
			TestName: "capitalization field names",
			Source: &TestFlexTF10{
				FieldURL: types.StringValue("h"),
			},
			Target: &TestFlexAWS12{},
			WantTarget: &TestFlexAWS12{
				FieldUrl: aws.String("h"),
			},
		},
		{
			Context:  context.WithValue(ctx, ResourcePrefix, "Intent"),
			TestName: "resource name prefix",
			Source: &TestFlexTF16{
				Name: types.StringValue("Ovodoghen"),
			},
			Target: &TestFlexAWS18{},
			WantTarget: &TestFlexAWS18{
				IntentName: aws.String("Ovodoghen"),
			},
		},
		{
			TestName:   "single ARN Source and single string Target",
			Source:     &TestFlexTF17{Field1: fwtypes.ARNValue(testARN)},
			Target:     &TestFlexAWS01{},
			WantTarget: &TestFlexAWS01{Field1: testARN},
		},
		{
			TestName:   "single ARN Source and single *string Target",
			Source:     &TestFlexTF17{Field1: fwtypes.ARNValue(testARN)},
			Target:     &TestFlexAWS02{},
			WantTarget: &TestFlexAWS02{Field1: aws.String(testARN)},
		},
		{
			TestName: "timestamp pointer",
			Source: &TestFlexTimeTF01{
				CreationDateTime: fwtypes.TimestampValue(testTimeStr),
			},
			Target: &TestFlexTimeAWS01{},
			WantTarget: &TestFlexTimeAWS01{
				CreationDateTime: &testTimeTime,
			},
		},
		{
			TestName: "timestamp",
			Source: &TestFlexTimeTF01{
				CreationDateTime: fwtypes.TimestampValue(testTimeStr),
			},
			Target: &TestFlexTimeAWS02{},
			WantTarget: &TestFlexTimeAWS02{
				CreationDateTime: testTimeTime,
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.TestName, func(t *testing.T) {
			t.Parallel()

			testCtx := ctx //nolint:contextcheck // simplify use of testing context
			if testCase.Context != nil {
				testCtx = testCase.Context
			}

			err := Expand(testCtx, testCase.Source, testCase.Target)
			gotErr := err != nil

			if gotErr != testCase.WantErr {
				t.Errorf("gotErr = %v, wantErr = %v", gotErr, testCase.WantErr)
			}

			if gotErr {
				if !testCase.WantErr {
					t.Errorf("err = %q", err)
				}
			} else if diff := cmp.Diff(testCase.Target, testCase.WantTarget); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestExpandGeneric(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	testCases := []struct {
		Context    context.Context //nolint:containedctx // testing context use
		TestName   string
		Source     any
		Target     any
		WantErr    bool
		WantTarget any
	}{
		{
			TestName:   "single list Source and *struct Target",
			Source:     &TestFlexTF05{Field1: fwtypes.NewListNestedObjectValueOfPtr(ctx, &TestFlexTF01{Field1: types.StringValue("a")})},
			Target:     &TestFlexAWS06{},
			WantTarget: &TestFlexAWS06{Field1: &TestFlexAWS01{Field1: "a"}},
		},
		{
			TestName:   "single set Source and *struct Target",
			Source:     &TestFlexTF06{Field1: fwtypes.NewSetNestedObjectValueOfPtr(ctx, &TestFlexTF01{Field1: types.StringValue("a")})},
			Target:     &TestFlexAWS06{},
			WantTarget: &TestFlexAWS06{Field1: &TestFlexAWS01{Field1: "a"}},
		},
		{
			TestName:   "empty list Source and empty []struct Target",
			Source:     &TestFlexTF05{Field1: fwtypes.NewListNestedObjectValueOfValueSlice(ctx, []TestFlexTF01{})},
			Target:     &TestFlexAWS08{},
			WantTarget: &TestFlexAWS08{Field1: []TestFlexAWS01{}},
		},
		{
			TestName: "non-empty list Source and non-empty []struct Target",
			Source: &TestFlexTF05{Field1: fwtypes.NewListNestedObjectValueOfValueSlice(ctx, []TestFlexTF01{
				{Field1: types.StringValue("a")},
				{Field1: types.StringValue("b")},
			})},
			Target: &TestFlexAWS08{},
			WantTarget: &TestFlexAWS08{Field1: []TestFlexAWS01{
				{Field1: "a"},
				{Field1: "b"},
			}},
		},
		{
			TestName:   "empty list Source and empty []*struct Target",
			Source:     &TestFlexTF05{Field1: fwtypes.NewListNestedObjectValueOfSlice(ctx, []*TestFlexTF01{})},
			Target:     &TestFlexAWS07{},
			WantTarget: &TestFlexAWS07{Field1: []*TestFlexAWS01{}},
		},
		{
			TestName: "non-empty list Source and non-empty []*struct Target",
			Source: &TestFlexTF05{Field1: fwtypes.NewListNestedObjectValueOfSlice(ctx, []*TestFlexTF01{
				{Field1: types.StringValue("a")},
				{Field1: types.StringValue("b")},
			})},
			Target: &TestFlexAWS07{},
			WantTarget: &TestFlexAWS07{Field1: []*TestFlexAWS01{
				{Field1: "a"},
				{Field1: "b"},
			}},
		},
		{
			TestName:   "empty list Source and empty []struct Target",
			Source:     &TestFlexTF05{Field1: fwtypes.NewListNestedObjectValueOfValueSlice(ctx, []TestFlexTF01{})},
			Target:     &TestFlexAWS08{},
			WantTarget: &TestFlexAWS08{Field1: []TestFlexAWS01{}},
		},
		{
			TestName: "non-empty list Source and non-empty []struct Target",
			Source: &TestFlexTF05{Field1: fwtypes.NewListNestedObjectValueOfValueSlice(ctx, []TestFlexTF01{
				{Field1: types.StringValue("a")},
				{Field1: types.StringValue("b")},
			})},
			Target: &TestFlexAWS08{},
			WantTarget: &TestFlexAWS08{Field1: []TestFlexAWS01{
				{Field1: "a"},
				{Field1: "b"},
			}},
		},
		{
			TestName:   "empty set Source and empty []*struct Target",
			Source:     &TestFlexTF06{Field1: fwtypes.NewSetNestedObjectValueOfSlice(ctx, []*TestFlexTF01{})},
			Target:     &TestFlexAWS07{},
			WantTarget: &TestFlexAWS07{Field1: []*TestFlexAWS01{}},
		},
		{
			TestName: "non-empty set Source and non-empty []*struct Target",
			Source: &TestFlexTF06{Field1: fwtypes.NewSetNestedObjectValueOfSlice(ctx, []*TestFlexTF01{
				{Field1: types.StringValue("a")},
				{Field1: types.StringValue("b")},
			})},
			Target: &TestFlexAWS07{},
			WantTarget: &TestFlexAWS07{Field1: []*TestFlexAWS01{
				{Field1: "a"},
				{Field1: "b"},
			}},
		},
		{
			TestName: "non-empty set Source and non-empty []struct Target",
			Source: &TestFlexTF06{Field1: fwtypes.NewSetNestedObjectValueOfValueSlice(ctx, []TestFlexTF01{
				{Field1: types.StringValue("a")},
				{Field1: types.StringValue("b")},
			})},
			Target: &TestFlexAWS08{},
			WantTarget: &TestFlexAWS08{Field1: []TestFlexAWS01{
				{Field1: "a"},
				{Field1: "b"},
			}},
		},
		{
			TestName: "complex Source and complex Target",
			Source: &TestFlexTF07{
				Field1: types.StringValue("m"),
				Field2: fwtypes.NewListNestedObjectValueOfPtr(ctx, &TestFlexTF05{
					Field1: fwtypes.NewListNestedObjectValueOfPtr(ctx, &TestFlexTF01{
						Field1: types.StringValue("n"),
					}),
				}),
				Field3: types.MapValueMust(types.StringType, map[string]attr.Value{
					"X": types.StringValue("x"),
					"Y": types.StringValue("y"),
				}),
				Field4: fwtypes.NewSetNestedObjectValueOfValueSlice(ctx, []TestFlexTF02{
					{Field1: types.Int64Value(100)},
					{Field1: types.Int64Value(2000)},
					{Field1: types.Int64Value(30000)},
				}),
			},
			Target: &TestFlexAWS09{},
			WantTarget: &TestFlexAWS09{
				Field1: "m",
				Field2: &TestFlexAWS06{Field1: &TestFlexAWS01{Field1: "n"}},
				Field3: aws.StringMap(map[string]string{"X": "x", "Y": "y"}),
				Field4: []TestFlexAWS03{{Field1: 100}, {Field1: 2000}, {Field1: 30000}},
			},
		},
		{
			TestName: "map string",
			Source: &TestFlexTF11{
				FieldInner: fwtypes.NewMapValueOf(ctx, map[string]basetypes.StringValue{
					"x": types.StringValue("y"),
				}),
			},
			Target: &TestFlexAWS13{},
			WantTarget: &TestFlexAWS13{
				FieldInner: map[string]string{
					"x": "y",
				},
			},
		},
		{
			TestName: "object map",
			Source: &TestFlexTF12{
				FieldInner: fwtypes.NewObjectMapValueMapOf[TestFlexTF01](ctx, map[string]TestFlexTF01{
					"x": {
						Field1: types.StringValue("a"),
					}},
				),
			},
			Target: &TestFlexAWS14{},
			WantTarget: &TestFlexAWS14{
				FieldInner: map[string]TestFlexAWS01{
					"x": {
						Field1: "a",
					},
				},
			},
		},
		{
			TestName: "object map ptr target",
			Source: &TestFlexTF12{
				FieldInner: fwtypes.NewObjectMapValueMapOf[TestFlexTF01](ctx,
					map[string]TestFlexTF01{
						"x": {
							Field1: types.StringValue("a"),
						},
					},
				),
			},
			Target: &TestFlexAWS15{},
			WantTarget: &TestFlexAWS15{
				FieldInner: map[string]*TestFlexAWS01{
					"x": {
						Field1: "a",
					},
				},
			},
		},
		{
			TestName: "object map ptr source and target",
			Source: &TestFlexTF13{
				FieldInner: fwtypes.NewObjectMapValuePtrMapOf[TestFlexTF01](ctx,
					map[string]*TestFlexTF01{
						"x": {
							Field1: types.StringValue("a"),
						},
					},
				),
			},
			Target: &TestFlexAWS15{},
			WantTarget: &TestFlexAWS15{
				FieldInner: map[string]*TestFlexAWS01{
					"x": {
						Field1: "a",
					},
				},
			},
		},
		{
			TestName: "nested string map",
			Source: &TestFlexTF14{
				FieldOuter: fwtypes.NewListNestedObjectValueOfPtr(ctx, &TestFlexTF11{
					FieldInner: fwtypes.NewMapValueOf(ctx, map[string]basetypes.StringValue{
						"x": types.StringValue("y"),
					}),
				}),
			},
			Target: &TestFlexAWS16{},
			WantTarget: &TestFlexAWS16{
				FieldOuter: TestFlexAWS13{
					FieldInner: map[string]string{
						"x": "y",
					},
				},
			},
		},
		{
			TestName: "nested object map",
			Source: &TestFlexTF15{
				FieldOuter: fwtypes.NewListNestedObjectValueOfPtr(ctx, &TestFlexTF12{
					FieldInner: fwtypes.NewObjectMapValueMapOf[TestFlexTF01](ctx,
						map[string]TestFlexTF01{
							"x": {
								Field1: types.StringValue("a"),
							},
						},
					),
				}),
			},
			Target: &TestFlexAWS17{},
			WantTarget: &TestFlexAWS17{
				FieldOuter: TestFlexAWS14{
					FieldInner: map[string]TestFlexAWS01{
						"x": {
							Field1: "a",
						},
					},
				},
			},
		},
		{
			TestName: "complex nesting",
			Source: &TestFlexComplexNestTF01{
				DialogAction: fwtypes.NewListNestedObjectValueOfPtr(ctx, &TestFlexComplexNestTF02{
					Type:                fwtypes.StringEnumValue(TestEnumList),
					SlotToElicit:        types.StringValue("x"),
					SuppressNextMessage: types.BoolValue(true),
				}),
				Intent: fwtypes.NewListNestedObjectValueOfPtr(ctx, &TestFlexComplexNestTF03{
					Name: types.StringValue("x"),
					Slots: fwtypes.NewObjectMapValueMapOf[TestFlexComplexNestTF04](ctx, map[string]TestFlexComplexNestTF04{
						"x": {
							Shape: fwtypes.StringEnumValue(TestEnumList),
							Value: fwtypes.NewListNestedObjectValueOfPtr(ctx, &TestFlexComplexNestTF05{
								InterpretedValue: types.StringValue("y"),
							}),
						},
					}),
				}),
				SessionAttributes: fwtypes.NewMapValueOf(ctx, map[string]basetypes.StringValue{
					"x": basetypes.NewStringValue("y"),
				}),
			},
			Target: &TestFlexComplexNestAWS01{},
			WantTarget: &TestFlexComplexNestAWS01{
				DialogAction: &TestFlexComplexNestAWS02{
					Type:                TestEnumList,
					SlotToElicit:        aws.String("x"),
					SuppressNextMessage: aws.Bool(true),
				},
				Intent: &TestFlexComplexNestAWS03{
					Name: aws.String("x"),
					Slots: map[string]TestFlexComplexNestAWS04{
						"x": {
							Shape: TestEnumList,
							Value: &TestFlexComplexNestAWS05{
								InterpretedValue: aws.String("y"),
							},
						},
					},
				},
				SessionAttributes: map[string]string{
					"x": "y",
				},
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.TestName, func(t *testing.T) {
			t.Parallel()

			testCtx := ctx //nolint:contextcheck // simplify use of testing context
			if testCase.Context != nil {
				testCtx = testCase.Context
			}

			err := Expand(testCtx, testCase.Source, testCase.Target)
			gotErr := err != nil

			if gotErr != testCase.WantErr {
				t.Errorf("gotErr = %v, wantErr = %v", gotErr, testCase.WantErr)
			}

			if gotErr {
				if !testCase.WantErr {
					t.Errorf("err = %q", err)
				}
			} else if diff := cmp.Diff(testCase.Target, testCase.WantTarget); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}
