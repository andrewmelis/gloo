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

// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/solo-io/glue/implemented_modules/kube/configwatcher/crd/solo.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VirtualHostLister helps list VirtualHosts.
type VirtualHostLister interface {
	// List lists all VirtualHosts in the indexer.
	List(selector labels.Selector) (ret []*v1.VirtualHost, err error)
	// VirtualHosts returns an object that can list and get VirtualHosts.
	VirtualHosts(namespace string) VirtualHostNamespaceLister
	VirtualHostListerExpansion
}

// virtualHostLister implements the VirtualHostLister interface.
type virtualHostLister struct {
	indexer cache.Indexer
}

// NewVirtualHostLister returns a new VirtualHostLister.
func NewVirtualHostLister(indexer cache.Indexer) VirtualHostLister {
	return &virtualHostLister{indexer: indexer}
}

// List lists all VirtualHosts in the indexer.
func (s *virtualHostLister) List(selector labels.Selector) (ret []*v1.VirtualHost, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VirtualHost))
	})
	return ret, err
}

// VirtualHosts returns an object that can list and get VirtualHosts.
func (s *virtualHostLister) VirtualHosts(namespace string) VirtualHostNamespaceLister {
	return virtualHostNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualHostNamespaceLister helps list and get VirtualHosts.
type VirtualHostNamespaceLister interface {
	// List lists all VirtualHosts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.VirtualHost, err error)
	// Get retrieves the VirtualHost from the indexer for a given namespace and name.
	Get(name string) (*v1.VirtualHost, error)
	VirtualHostNamespaceListerExpansion
}

// virtualHostNamespaceLister implements the VirtualHostNamespaceLister
// interface.
type virtualHostNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualHosts in the indexer for a given namespace.
func (s virtualHostNamespaceLister) List(selector labels.Selector) (ret []*v1.VirtualHost, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VirtualHost))
	})
	return ret, err
}

// Get retrieves the VirtualHost from the indexer for a given namespace and name.
func (s virtualHostNamespaceLister) Get(name string) (*v1.VirtualHost, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("virtualhost"), name)
	}
	return obj.(*v1.VirtualHost), nil
}