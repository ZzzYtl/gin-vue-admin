package sentinel

import (
	"context"
	"encoding/json"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	publish "gin-vue-admin/pubsub/protocal"
	"google.golang.org/grpc/peer"
	"net"
)

type RpcServer struct {
}

func (s *RpcServer) GetConfig(c context.Context, req *publish.GetConfigReq) (*publish.GetConfigAck, error) {
	var addr string
	if pr, ok := peer.FromContext(c); ok {
		if tcpAddr, ok := pr.Addr.(*net.TCPAddr); ok {
			addr = tcpAddr.IP.String()
		} else {
			addr = pr.Addr.String()
		}
	}
	var sentinel model.SentinelInfo
	err := global.GVA_DB.Where("ip = ? AND port = ?", addr, req.Port).First(&sentinel).Error
	if err != nil {
		return &publish.GetConfigAck{
			Result: 1,
			Config: nil,
		}, nil
	}

	var config publish.Config
	config.ConfigEpoch = uint64(sentinel.ConfigEpoch)
	config.CurrentEpoch = uint64(sentinel.CurrentEpoch)
	config.SentinelClusterId = uint32(sentinel.SentinelClusterID)
	sentinelClusterID := sentinel.SentinelClusterID

	var sentinels []model.SentinelInfo
	err = global.GVA_DB.Where("sentinel_cluster_id = ?", sentinelClusterID).Find(&sentinels).Error
	if err != nil {
		return &publish.GetConfigAck{
			Result: 1,
			Config: nil,
		}, nil
	}

	var sentinelCfgs []*publish.SentinelConfig

	for _, v := range sentinels {
		sentinelCfgs = append(sentinelCfgs, &publish.SentinelConfig{
			RunId: v.RunID,
			Ip:    v.IP,
			Port:  uint64(v.Port),
		})
	}

	var dbClusters []model.SentinelDBClusterInfo
	err = global.GVA_DB.Where("sentinel_cluster_id = ?", sentinelClusterID).Find(&dbClusters).Error
	if err != nil {
		return &publish.GetConfigAck{
			Result: 1,
			Config: nil,
		}, nil
	}

	var dbClusterCfgs []*publish.DBClusterConfig
	for _, v := range dbClusters {
		dbClusterCfg := publish.DBClusterConfig{
			Name:        "",
			User:        "",
			Pw:          "",
			RlpcUser:    v.RlpcUser,
			RlpcPw:      v.RlpcPWD,
			LeaderEpoch: uint64(v.LeaderEpoch),
		}
		dbClusterInfo := model.DataBaseInfo{}
		err = global.GVA_DB.Where("cluster_id = ?", v.ClusterID).First(&dbClusterInfo).Error
		if err != nil {
			return &publish.GetConfigAck{
				Result: 1,
				Config: nil,
			}, nil
		}
		dbClusterCfg.Name = dbClusterInfo.DBName
		dbClusterCfg.User = dbClusterInfo.DBUser
		dbClusterCfg.Pw = dbClusterInfo.DBPWD

		var dbs []model.Node
		if json.Unmarshal([]byte(v.Dbs), &dbs) != nil {
			return &publish.GetConfigAck{
				Result: 1,
				Config: nil,
			}, nil
		}
		for _, v := range dbs {
			dbClusterCfg.DbConfigs = append(dbClusterCfg.DbConfigs, &publish.DBConfig{
				Ip:   v.IP,
				Port: uint64(dbClusterInfo.Port),
				Type: publish.DBType(v.RoleID),
			})
		}
		dbClusterCfgs = append(dbClusterCfgs, &dbClusterCfg)
	}

	config.SentinelConfigs = sentinelCfgs
	config.DbClusterConfigs = dbClusterCfgs
	return &publish.GetConfigAck{
		Result: 0,
		Config: &config,
	}, nil
}

func (s *RpcServer) SwitchDBCluster(c context.Context,
		req *publish.SwitchDBClusterReq) (*publish.SwitchDBClusterAck, error) {
	clusterName := req.GetClusterName()
	db := model.DataBaseInfo{}
	err := global.GVA_DB.Where("db_name = ?", clusterName).First(&db).Error
	if err != nil {
		return &publish.SwitchDBClusterAck{
			Result: 1,
		}, err
	}

	sentinelCluster := model.SentinelDBClusterInfo{}
	err = global.GVA_DB.Where("cluster_id = ?", db.ClusterID).First(&sentinelCluster).Error
	if err != nil {
		return &publish.SwitchDBClusterAck{
			Result: 1,
		}, err
	}

	if len(req.Dbs) != 0 {
		data, err := json.Marshal(req.Dbs)
		if err != nil {
			return &publish.SwitchDBClusterAck{
				Result: 1,
			}, err
		}
		sentinelCluster.Dbs = string(data)
	}

	if req.LeaderEpoch != 0 {
		sentinelCluster.LeaderEpoch = int(req.LeaderEpoch)
	}

	err = global.GVA_DB.Save(&sentinelCluster).Error
	if err != nil {
		return &publish.SwitchDBClusterAck{
			Result: 1,
		}, err
	}
	return &publish.SwitchDBClusterAck{
		Result: 0,
	}, err
}

func (s *RpcServer) SyncSentinelConfig(c context.Context,
	req *publish.SyncSentinelConfigReq) (*publish.SyncSentinelConfigAck, error) {
	sentinel := model.SentinelInfo{}
	err := global.GVA_DB.Where("ip = ? AND port = ?", req.Ip, req.Port).First(&sentinel).Error
	if err != nil {
		return &publish.SyncSentinelConfigAck{
			Result: 1,
		}, err
	}
	//update
	sentinel.RunID = req.RunId
	sentinel.CurrentEpoch = int(req.CurrentEpoch)
	sentinel.ConfigEpoch = int(req.ConfigEpoch)
	err = global.GVA_DB.Save(&sentinel).Error
	if err != nil {
		return &publish.SyncSentinelConfigAck{
			Result: 1,
		}, err
	}
	return &publish.SyncSentinelConfigAck{
		Result: 0,
	}, err
}
