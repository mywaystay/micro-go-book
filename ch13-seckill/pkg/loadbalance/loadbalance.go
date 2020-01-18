package loadbalance

import (
	"errors"
	"github.com/keets2012/Micro-Go-Pracrise/ch13-seckill/pkg/discover"
	"math/rand"
)

// 负载均衡器
type LoadBalance interface {
	SelectService(service []*discover.ServiceInstance) (*discover.ServiceInstance, error)
}

type RandomLoadBalance struct {
}

// 随机负载均衡
func (loadBalance *RandomLoadBalance) SelectService(services []*discover.ServiceInstance) (*discover.ServiceInstance, error) {

	if services == nil || len(services) == 0 {
		return nil, errors.New("service instances are not exist")
	}

	return services[rand.Intn(len(services))], nil
}

type WeightRoundRobinLoadBalance struct {
}

// 权重平滑负载均衡
func (loadBalance *WeightRoundRobinLoadBalance) SelectService(services []*discover.ServiceInstance) (best *discover.ServiceInstance, err error) {

	if services == nil || len(services) == 0 {
		return nil, errors.New("service instances are not exist")
	}

	total := 0
	for i := 0; i < len(services); i++ {
		w := services[i]
		if w == nil {
			continue
		}

		w.CurWeight += w.Weight

		total += w.Weight
		if w.Weight < w.Weight {
			w.Weight++
		}
		if best == nil || w.CurWeight > best.CurWeight {
			best = w
		}
	}

	if best == nil {
		return nil, nil
	}

	best.CurWeight -= total
	return best, nil
}