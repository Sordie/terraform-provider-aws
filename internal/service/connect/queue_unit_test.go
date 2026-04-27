// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package connect

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	awstypes "github.com/aws/aws-sdk-go-v2/service/connect/types"
)

func TestExpandOutboundCallerConfig(t *testing.T) {
	t.Parallel()

	t.Run("absent (empty list) returns nil", func(t *testing.T) {
		t.Parallel()
		result := expandOutboundCallerConfig([]any{})
		if result != nil {
			t.Errorf("expected nil, got %+v", result)
		}
	})

	t.Run("present but all-empty returns empty non-nil struct", func(t *testing.T) {
		t.Parallel()
		result := expandOutboundCallerConfig([]any{
			map[string]any{
				"outbound_caller_id_name":      "",
				"outbound_caller_id_number_id": "",
				"outbound_flow_id":             "",
			},
		})
		if result == nil {
			t.Fatal("expected non-nil *OutboundCallerConfig, got nil")
		}
		if result.OutboundCallerIdName != nil {
			t.Errorf("expected OutboundCallerIdName to be nil, got %s", aws.ToString(result.OutboundCallerIdName))
		}
		if result.OutboundCallerIdNumberId != nil {
			t.Errorf("expected OutboundCallerIdNumberId to be nil, got %s", aws.ToString(result.OutboundCallerIdNumberId))
		}
		if result.OutboundFlowId != nil {
			t.Errorf("expected OutboundFlowId to be nil, got %s", aws.ToString(result.OutboundFlowId))
		}
	})

	t.Run("present with values populates struct fields", func(t *testing.T) {
		t.Parallel()
		result := expandOutboundCallerConfig([]any{
			map[string]any{
				"outbound_caller_id_name":      "my-name",
				"outbound_caller_id_number_id": "number-id-123",
				"outbound_flow_id":             "flow-id-456",
			},
		})
		if result == nil {
			t.Fatal("expected non-nil *OutboundCallerConfig, got nil")
		}
		if aws.ToString(result.OutboundCallerIdName) != "my-name" {
			t.Errorf("expected OutboundCallerIdName 'my-name', got %s", aws.ToString(result.OutboundCallerIdName))
		}
		if aws.ToString(result.OutboundCallerIdNumberId) != "number-id-123" {
			t.Errorf("expected OutboundCallerIdNumberId 'number-id-123', got %s", aws.ToString(result.OutboundCallerIdNumberId))
		}
		if aws.ToString(result.OutboundFlowId) != "flow-id-456" {
			t.Errorf("expected OutboundFlowId 'flow-id-456', got %s", aws.ToString(result.OutboundFlowId))
		}
	})
}

func TestFlattenOutboundCallerConfig(t *testing.T) {
	t.Parallel()

	t.Run("nil input returns empty list (no drift)", func(t *testing.T) {
		t.Parallel()
		result := flattenOutboundCallerConfig(nil)
		if len(result) != 0 {
			t.Errorf("expected empty list, got %+v", result)
		}
	})

	t.Run("all-nil fields returns empty list (no drift)", func(t *testing.T) {
		t.Parallel()
		// AWS always returns a non-nil OutboundCallerConfig even when unconfigured.
		result := flattenOutboundCallerConfig(&awstypes.OutboundCallerConfig{})
		if len(result) != 0 {
			t.Errorf("expected empty list, got %+v", result)
		}
	})

	t.Run("populated fields returns single-element list", func(t *testing.T) {
		t.Parallel()
		result := flattenOutboundCallerConfig(&awstypes.OutboundCallerConfig{
			OutboundCallerIdName:    aws.String("my-name"),
			OutboundCallerIdNumberId: aws.String("number-id-123"),
			OutboundFlowId:          aws.String("flow-id-456"),
		})
		if len(result) != 1 {
			t.Fatalf("expected list of length 1, got %d", len(result))
		}
		tfMap, ok := result[0].(map[string]any)
		if !ok {
			t.Fatalf("expected map[string]any, got %T", result[0])
		}
		if tfMap["outbound_caller_id_name"] != "my-name" {
			t.Errorf("expected outbound_caller_id_name 'my-name', got %v", tfMap["outbound_caller_id_name"])
		}
		if tfMap["outbound_caller_id_number_id"] != "number-id-123" {
			t.Errorf("expected outbound_caller_id_number_id 'number-id-123', got %v", tfMap["outbound_caller_id_number_id"])
		}
		if tfMap["outbound_flow_id"] != "flow-id-456" {
			t.Errorf("expected outbound_flow_id 'flow-id-456', got %v", tfMap["outbound_flow_id"])
		}
	})

	t.Run("partial fields returns single-element list with only set fields", func(t *testing.T) {
		t.Parallel()
		result := flattenOutboundCallerConfig(&awstypes.OutboundCallerConfig{
			OutboundCallerIdName: aws.String("only-name"),
		})
		if len(result) != 1 {
			t.Fatalf("expected list of length 1, got %d", len(result))
		}
		tfMap, ok := result[0].(map[string]any)
		if !ok {
			t.Fatalf("expected map[string]any, got %T", result[0])
		}
		if tfMap["outbound_caller_id_name"] != "only-name" {
			t.Errorf("expected outbound_caller_id_name 'only-name', got %v", tfMap["outbound_caller_id_name"])
		}
		if _, exists := tfMap["outbound_caller_id_number_id"]; exists {
			t.Errorf("expected outbound_caller_id_number_id to be absent, got %v", tfMap["outbound_caller_id_number_id"])
		}
		if _, exists := tfMap["outbound_flow_id"]; exists {
			t.Errorf("expected outbound_flow_id to be absent, got %v", tfMap["outbound_flow_id"])
		}
	})
}
