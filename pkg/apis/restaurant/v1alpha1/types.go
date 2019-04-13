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

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Pizza specifies an offered pizza with toppings.
type Pizza struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   PizzaSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status PizzaStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type PizzaSpec struct {
	// +k8s:conversion-gen=false
	// toppings is a list of Topping names. They don't have to be unique. Order does not matter.
	Toppings []string `json:"toppings" protobuf:"bytes,1,rep,name=toppings"`
}

type PizzaStatus struct {
	// cost is the cost of the whole pizza including all toppings.
	Cost float64 `json:"cost,omitempty" protobuf:"bytes,1,opt,name=cost"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PizzaList is a list of Pizza objects.
type PizzaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Pizza `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Topping is a topping put onto a pizza.
type Topping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec ToppingSpec
}

type ToppingSpec struct {
	// cost is the cost of one instance of this topping.
	Cost float64 `json:"cost" protobuf:"bytes,1,name=cost"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ToppingList is a list of Topping objects.
type ToppingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Topping `json:"items" protobuf:"bytes,2,rep,name=items"`
}
