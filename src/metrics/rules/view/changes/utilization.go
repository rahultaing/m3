// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package changes

import "github.com/m3db/m3/src/metrics/rules/view"

// UtilizationRuleChange is a utilization rule diff.
type UtilizationRuleChange struct {
	Op       Op                    `json:"op"`
	RuleID   *string               `json:"ruleID,omitempty"`
	RuleData *view.UtilizationRule `json:"ruleData,omitempty"`
}

type utilizationRuleChangesByOpAscNameAscIDAsc []UtilizationRuleChange

func (a utilizationRuleChangesByOpAscNameAscIDAsc) Len() int      { return len(a) }
func (a utilizationRuleChangesByOpAscNameAscIDAsc) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a utilizationRuleChangesByOpAscNameAscIDAsc) Less(i, j int) bool {
	if a[i].Op < a[j].Op {
		return true
	}
	if a[i].Op > a[j].Op {
		return false
	}
	// For adds and changes.
	if a[i].RuleData != nil && a[j].RuleData != nil {
		return a[i].RuleData.Name < a[j].RuleData.Name
	}
	// For deletes.
	if a[i].RuleID != nil && a[j].RuleID != nil {
		return *a[i].RuleID < *a[j].RuleID
	}
	// This should not happen
	return false
}