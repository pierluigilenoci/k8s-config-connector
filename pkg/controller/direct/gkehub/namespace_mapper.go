// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gkehub

import (
	gkehubapi "google.golang.org/api/gkehub/v1"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func GKEHubNamespaceSpec_ToAPI(mapCtx *direct.MapContext, r *krm.GKEHubNamespaceSpec) *gkehubapi.Namespace {
	if r == nil {
		return nil
	}
	out := &gkehubapi.Namespace{}
	return out
}

func GKEHubNamespaceObservedState_FromProto(mapCtx *direct.MapContext, r *gkehubapi.Namespace) *krm.GKEHubNamespaceObservedState {
	if r == nil {
		return nil
	}
	out := &krm.GKEHubNamespaceObservedState{}
	out.CreateTime = direct.LazyPtr(r.CreateTime)
	out.UpdateTime = direct.LazyPtr(r.UpdateTime)
	out.DeleteTime = direct.LazyPtr(r.DeleteTime)
	out.Uid = direct.LazyPtr(r.Uid)
	if r.State != nil {
		out.State = &krm.GKEHubNamespaceStateStatus{
			Code: direct.LazyPtr(r.State.Code),
		}
	}
	return out
}

func GKEHubNamespaceStatus_FromProto(mapCtx *direct.MapContext, r *gkehubapi.Namespace) *krm.GKEHubNamespaceStatus {
	if r == nil {
		return nil
	}
	out := &krm.GKEHubNamespaceStatus{}
	out.ObservedState = GKEHubNamespaceObservedState_FromProto(mapCtx, r)
	return out
}
