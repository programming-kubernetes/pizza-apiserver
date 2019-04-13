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

package restaurant

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Pizza specifies an offered pizza with toppings.
type Pizza struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   PizzaSpec
	Status PizzaStatus
}

type PizzaSpec struct {
	// toppings is a list of Topping names. They don't have to be unique. Order does not matter.
	Toppings []PizzaTopping
}

type PizzaTopping struct {
	// name is the name of a Topping object .
	Name string
	// quantity is the number of how often the topping is put onto the pizza.
	Quantity int
}

type PizzaStatus struct {
	// cost is the cost of the whole pizza including all toppings.
	Cost float64
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PizzaList is a list of Pizza objects.
type PizzaList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Pizza
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Topping is a topping put onto a pizza.
type Topping struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec ToppingSpec
}

type ToppingSpec struct {
	// cost is the cost of one instance of this topping.
	Cost float64
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ToppingList is a list of Topping objects.
type ToppingList struct {
	metav1.TypeMeta
	metav1.ListMeta

	// Items is a list of Toppings
	Items []Topping
}
