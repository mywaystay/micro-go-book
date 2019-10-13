package setup

import (
	"github.com/coreos/etcd/clientv3"
	conf "github.com/keets2012/Micro-Go-Pracrise/ch13-seckill/common/config"
	"github.com/keets2012/Micro-Go-Pracrise/ch13-seckill/sk-core/logic"
	"log"
	"time"
)

//初始化Etcd
func InitEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"39.98.179.73:2379"}, // conf.Etcd.Host
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("Connect etcd failed. Error : %v", err)
	}
	log.Print("Connect etcd sucess")
	conf.Etcd.EtcdConn = cli
	logic.LoadProductFromEtcd()
}
