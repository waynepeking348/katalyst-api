/*
Copyright 2022 The Katalyst Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"
	time "time"

	autoscalingv1alpha2 "github.com/kubewharf/katalyst-api/pkg/apis/autoscaling/v1alpha2"
	versioned "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubewharf/katalyst-api/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha2 "github.com/kubewharf/katalyst-api/pkg/client/listers/autoscaling/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VirtualWorkloadInformer provides access to a shared informer and lister for
// VirtualWorkloads.
type VirtualWorkloadInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.VirtualWorkloadLister
}

type virtualWorkloadInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVirtualWorkloadInformer constructs a new informer for VirtualWorkload type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVirtualWorkloadInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVirtualWorkloadInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVirtualWorkloadInformer constructs a new informer for VirtualWorkload type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVirtualWorkloadInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV1alpha2().VirtualWorkloads(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV1alpha2().VirtualWorkloads(namespace).Watch(context.TODO(), options)
			},
		},
		&autoscalingv1alpha2.VirtualWorkload{},
		resyncPeriod,
		indexers,
	)
}

func (f *virtualWorkloadInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVirtualWorkloadInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *virtualWorkloadInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&autoscalingv1alpha2.VirtualWorkload{}, f.defaultInformer)
}

func (f *virtualWorkloadInformer) Lister() v1alpha2.VirtualWorkloadLister {
	return v1alpha2.NewVirtualWorkloadLister(f.Informer().GetIndexer())
}
