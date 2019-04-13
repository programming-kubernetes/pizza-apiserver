/*
Copyright 2018 The Kubernetes Authors.

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

package v1beta1

func init() {
	localSchemeBuilder.Register(RegisterDefaults)
}

func SetDefaults_PizzaSpec(obj *PizzaSpec) {
	if len(obj.Toppings) == 0 {
		obj.Toppings = []PizzaTopping{
			{"salami", 1},
			{"mozzarella", 1},
			{"tomato", 1},
		}
	}

	for i := range obj.Toppings {
		if obj.Toppings[i].Quantity == 0 {
			obj.Toppings[i].Quantity = 1
		}
	}
}
