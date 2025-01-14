//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package impl

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	service "github.com/alibaba/ioc-golang/example/autowire/autowire_allimpls/service"
	allimpls "github.com/alibaba/ioc-golang/extension/autowire/allimpls"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceImpl1_{}
		},
	})
	serviceImpl1StructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceImpl1{}
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(service.Service),
					},
				},
			},
		},
	}
	allimpls.RegisterStructDescriptor(serviceImpl1StructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceImpl2_{}
		},
	})
	serviceImpl2StructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceImpl2{}
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(service.Service),
					},
				},
			},
		},
	}
	allimpls.RegisterStructDescriptor(serviceImpl2StructDescriptor)
}

type serviceImpl1_ struct {
	GetHelloString_ func(name string) string
}

func (s *serviceImpl1_) GetHelloString(name string) string {
	return s.GetHelloString_(name)
}

type serviceImpl2_ struct {
	GetHelloString_ func(name string) string
}

func (s *serviceImpl2_) GetHelloString(name string) string {
	return s.GetHelloString_(name)
}

type serviceImpl1IOCInterface interface {
	GetHelloString(name string) string
}

type serviceImpl2IOCInterface interface {
	GetHelloString(name string) string
}

var _serviceImpl1SDID string
var _serviceImpl2SDID string
