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

package fuzzer

import (
	fuzz "github.com/google/gofuzz"

	"github.com/programming-kubernetes/pizza-apiserver/pkg/apis/restaurant"

	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
)

// Funcs returns the fuzzer functions for the restaurant api group.
var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{
		func(s *restaurant.PizzaSpec, c fuzz.Continue) {
			c.FuzzNoCustom(s) // fuzz first without calling this function again

			// avoid empty Toppings because that is defaulted
			if len(s.Toppings) == 0 {
				s.Toppings = []restaurant.PizzaTopping{
					{"salami", 1},
					{"mozzarella", 1},
					{"tomato", 1},
				}
			}

			seen := map[string]bool{}
			for i := range s.Toppings {
				// make quantity strictly positive and of reasonable size
				s.Toppings[i].Quantity = 1 + c.Intn(10)

				// remove duplicates
				for {
					if !seen[s.Toppings[i].Name] {
						break
					}
					s.Toppings[i].Name = c.RandString()
				}
				seen[s.Toppings[i].Name] = true
			}
		},
	}
}
