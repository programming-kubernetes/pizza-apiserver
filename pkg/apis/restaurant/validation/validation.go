/*
Copyright 2016 The Kubernetes Authors.

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

package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"
	"github.com/programming-kubernetes/pizza-apiserver/pkg/apis/restaurant"
)

// ValidatePizza validates a Pizza.
func ValidatePizza(f *restaurant.Pizza) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs, ValidatePizzaSpec(&f.Spec, field.NewPath("spec"))...)

	return allErrs
}

// ValidatePizzaSpec validates a PizzaSpec.
func ValidatePizzaSpec(s *restaurant.PizzaSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	prevNames := map[string]bool{}
	for i := range s.Toppings {
		if s.Toppings[i].Quantity <= 0 {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("toppings").Index(i).Child("quantity"), s.Toppings[i].Quantity, "cannot be negative or zero"))
		}
		if len(s.Toppings[i].Name) == 0 {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("toppings").Index(i).Child("name"), s.Toppings[i].Name, "cannot be empty"))
		} else {
			if prevNames[s.Toppings[i].Name] {
				allErrs = append(allErrs, field.Invalid(fldPath.Child("toppings").Index(i).Child("name"), s.Toppings[i].Name, "must be unique"))
			}
			prevNames[s.Toppings[i].Name] = true
		}
	}

	return allErrs
}
