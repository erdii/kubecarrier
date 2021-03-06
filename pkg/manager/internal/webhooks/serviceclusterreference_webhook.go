/*
Copyright 2019 The KubeCarrier Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package webhooks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-logr/logr"
	adminv1beta1 "k8s.io/api/admission/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	catalogv1alpha1 "github.com/kubermatic/kubecarrier/pkg/apis/catalog/v1alpha1"
)

// ServiceClusterReferenceWebhookHandler handles mutating/validating of ServiceClusterReferences.
type ServiceClusterReferenceWebhookHandler struct {
	decoder *admission.Decoder
	Log     logr.Logger
}

var _ admission.Handler = (*ServiceClusterReferenceWebhookHandler)(nil)

// +kubebuilder:webhook:path=/validate-catalog-kubecarrier-io-v1alpha1-serviceclusterreference,mutating=false,failurePolicy=fail,groups=catalog.kubecarrier.io,resources=serviceclusterreferences,verbs=create;update,versions=v1alpha1,name=vserviceclusterreference.kubecarrier.io

// Handle is the function to handle create/update requests of ServiceClusterReferences.
func (r *ServiceClusterReferenceWebhookHandler) Handle(ctx context.Context, req admission.Request) admission.Response {
	obj := &catalogv1alpha1.ServiceClusterReference{}
	if err := r.decoder.Decode(req, obj); err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	switch req.Operation {
	case adminv1beta1.Create:
		if err := r.validateCreate(obj); err != nil {
			return admission.Denied(err.Error())
		}
	case adminv1beta1.Update:
		oldObj := &catalogv1alpha1.ServiceClusterReference{}
		if err := r.decoder.DecodeRaw(req.OldObject, oldObj); err != nil {
			return admission.Errored(http.StatusBadRequest, err)
		}
		if err := r.validateUpdate(oldObj, obj); err != nil {
			return admission.Denied(err.Error())
		}
	}
	return admission.Allowed("allowed to commit the request")

}

// ServiceClusterReferenceWebhookHandler implements admission.DecoderInjector.
// A decoder will be automatically injected.

// InjectDecoder injects the decoder.
func (r *ServiceClusterReferenceWebhookHandler) InjectDecoder(d *admission.Decoder) error {
	r.decoder = d
	return nil
}

func (r *ServiceClusterReferenceWebhookHandler) validateCreate(serviceClusterReference *catalogv1alpha1.ServiceClusterReference) error {
	r.Log.Info("validate create", "name", serviceClusterReference.Name)
	if serviceClusterReference.Spec.Provider.Name == "" {
		return fmt.Errorf("the Provider of ServiceClusterReference is not specifed")
	}
	return nil
}

func (r *ServiceClusterReferenceWebhookHandler) validateUpdate(oldObj, newObj *catalogv1alpha1.ServiceClusterReference) error {
	r.Log.Info("validate update", "name", newObj.Name)
	if newObj.Spec.Provider.Name != oldObj.Spec.Provider.Name {
		return fmt.Errorf("the Provider of ServiceClusterReference is immutable")
	}
	return nil
}
