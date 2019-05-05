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

package custominitializer_test

import (
	"testing"
	"time"

	"k8s.io/apiserver/pkg/admission"

	"github.com/programming-kubernetes/pizza-apiserver/pkg/admission/custominitializer"
	"github.com/programming-kubernetes/pizza-apiserver/pkg/generated/clientset/versioned/fake"
	informers "github.com/programming-kubernetes/pizza-apiserver/pkg/generated/informers/externalversions"
)

// TestWantsInternalRestaurantInformerFactory ensures that the informer factory is injected
// when the WantsRestaurantInformerFactory interface is implemented by a plugin.
func TestWantsInternalRestaurantInformerFactory(t *testing.T) {
	cs := &fake.Clientset{}
	sf := informers.NewSharedInformerFactory(cs, time.Duration(1)*time.Second)
	target := custominitializer.New(sf)

	wantRestaurantInformerFactory := &wantRestaurantInformerFactory{}
	target.Initialize(wantRestaurantInformerFactory)
	if wantRestaurantInformerFactory.sf != sf {
		t.Errorf("expected informer factory to be initialized")
	}
}

// wantRestaurantInformerFactory is a test stub that fulfills the WantsRestaurantInformerFactory interface
type wantRestaurantInformerFactory struct {
	sf informers.SharedInformerFactory
}

func (self *wantRestaurantInformerFactory) SetRestaurantInformerFactory(sf informers.SharedInformerFactory) {
	self.sf = sf
}
func (self *wantRestaurantInformerFactory) Admit(a admission.Attributes) error { return nil }
func (self *wantRestaurantInformerFactory) Handles(o admission.Operation) bool { return false }
func (self *wantRestaurantInformerFactory) ValidateInitialization() error      { return nil }

var _ admission.Interface = &wantRestaurantInformerFactory{}
var _ custominitializer.WantsRestaurantInformerFactory = &wantRestaurantInformerFactory{}
