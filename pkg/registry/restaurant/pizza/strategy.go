/*
Copyright 2017 The Kubernetes Authors.

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

package pizza

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
	"github.com/programming-kubernetes/pizza-apiserver/pkg/apis/restaurant/validation"

	"github.com/programming-kubernetes/pizza-apiserver/pkg/apis/restaurant"
)

// NewStrategy creates and returns a pizzaStrategy instance
func NewStrategy(typer runtime.ObjectTyper) pizzaStrategy {
	return pizzaStrategy{typer, names.SimpleNameGenerator}
}

// GetAttrs returns labels.Set, fields.Set, the presence of Initializers if any
// and error in case the given runtime.Object is not a Pizza
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*restaurant.Pizza)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a Pizza")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), nil
}

// MatchPizza is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchPizza(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *restaurant.Pizza) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type pizzaStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (pizzaStrategy) NamespaceScoped() bool {
	return true
}

func (pizzaStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (pizzaStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (pizzaStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	pizza := obj.(*restaurant.Pizza)
	return validation.ValidatePizza(pizza)
}

func (pizzaStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (pizzaStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (pizzaStrategy) Canonicalize(obj runtime.Object) {
}

func (pizzaStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
