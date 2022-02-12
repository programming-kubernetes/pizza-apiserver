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

package pizzatoppings

import (
	"fmt"
	"io"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apiserver/pkg/admission"

	"github.com/programming-kubernetes/pizza-apiserver/pkg/admission/custominitializer"
	"github.com/programming-kubernetes/pizza-apiserver/pkg/apis/restaurant"
	informers "github.com/programming-kubernetes/pizza-apiserver/pkg/generated/informers/externalversions"
	listers "github.com/programming-kubernetes/pizza-apiserver/pkg/generated/listers/restaurant/v1alpha1"
)

// Register registers a plugin
func Register(plugins *admission.Plugins) {
	plugins.Register("PizzaToppings", func(config io.Reader) (admission.Interface, error) {
		return New()
	})
}

type PizzaToppingsPlugin struct {
	*admission.Handler
	toppingLister listers.ToppingLister
}

var _ = custominitializer.WantsRestaurantInformerFactory(&PizzaToppingsPlugin{})
var _ = admission.ValidationInterface(&PizzaToppingsPlugin{})

// Admit ensures that the object in-flight is of kind Pizza.
// In addition checks that the toppings are known.
func (d *PizzaToppingsPlugin) Validate(a admission.Attributes, _ admission.ObjectInterfaces) error {
	// we are only interested in pizzas
	if a.GetKind().GroupKind() != restaurant.Kind("Pizza") {
		return nil
	}

	if !d.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
	}

	obj := a.GetObject()
	pizza := obj.(*restaurant.Pizza)
	for _, top := range pizza.Spec.Toppings {
		if _, err := d.toppingLister.Get(top.Name); err != nil && errors.IsNotFound(err) {
			return admission.NewForbidden(
				a,
				fmt.Errorf("unknown topping: %s", top.Name),
			)
		}
	}

	return nil
}

// SetRestaurantInformerFactory gets Lister from SharedInformerFactory.
// The lister knows how to lists Toppings.
func (d *PizzaToppingsPlugin) SetRestaurantInformerFactory(f informers.SharedInformerFactory) {
	d.toppingLister = f.Restaurant().V1alpha1().Toppings().Lister()
	d.SetReadyFunc(f.Restaurant().V1alpha1().Toppings().Informer().HasSynced)
}

// ValidaValidateInitializationte checks whether the plugin was correctly initialized.
func (d *PizzaToppingsPlugin) ValidateInitialization() error {
	if d.toppingLister == nil {
		return fmt.Errorf("missing policy lister")
	}
	return nil
}

// New creates a new ban pizza topping admission plugin
func New() (*PizzaToppingsPlugin, error) {
	return &PizzaToppingsPlugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}, nil
}
