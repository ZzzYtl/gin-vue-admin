package service

import (
	"encoding/json"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	publish "gin-vue-admin/pubsub/protocal"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSentinelDBClusterInfo
//@description: 创建SentinelDBClusterInfo记录
//@param: SentinelDBCluster model.SentinelDBClusterInfo
//@return: err error

func CreateSentinelDBClusterInfo(SentinelDBCluster model.SentinelDBClusterInfo) (err error) {
	dbClusterInfo := model.DataBaseInfo{}
	err = global.GVA_DB.Where("cluster_id = ?", SentinelDBCluster.ClusterID).First(&dbClusterInfo).Error
	if err != nil {
		return err
	}

	var dbs []model.Node
	err = global.GVA_DB.Where("tag_id = ?", dbClusterInfo.TagID).Find(&dbs).Error
	if err != nil {
		return err
	}
	data, err := json.Marshal(dbs)
	if err != nil {
		return err
	}
	SentinelDBCluster.Dbs = string(data)
	err = global.GVA_DB.Create(&SentinelDBCluster).Error
	if err == nil {
		msg := &publish.AddDBCluster{
			DbClusterConfig: &publish.DBClusterConfig{
				Name:        dbClusterInfo.DBName,
				User:        dbClusterInfo.DBUser,
				Pw:          dbClusterInfo.DBPWD,
				RlpcUser:    SentinelDBCluster.RlpcUser,
				RlpcPw:      SentinelDBCluster.RlpcPWD,
				LeaderEpoch: uint64(SentinelDBCluster.LeaderEpoch),
				DbConfigs:   []*publish.DBConfig{},
			},
		}
		for _, db := range dbs {
			msg.DbClusterConfig.DbConfigs = append(msg.DbClusterConfig.DbConfigs, &publish.DBConfig{
				Ip:   db.IP,
				Port: uint64(dbClusterInfo.Port),
				Type: publish.DBType(db.RoleID),
			})
		}
		global.GVA_PUBER.Pub(uint32(SentinelDBCluster.SentinelClusterID),
			&publish.PubsubMessage{AddDbCluster: msg})
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSentinelDBClusterInfo
//@description: 删除SentinelDBClusterInfo记录
//@param: SentinelDBCluster model.SentinelDBClusterInfo
//@return: err error

func DeleteSentinelDBClusterInfo(SentinelDBCluster model.SentinelDBClusterInfo) (err error) {
	cluster := &model.DataBaseInfo{}
	global.GVA_DB.Where("cluster_id = ?", SentinelDBCluster.ClusterID).First(cluster)
	global.GVA_DB.First(&SentinelDBCluster)
	err = global.GVA_DB.Delete(&SentinelDBCluster).Error
	if err == nil && len(cluster.DBName) != 0 {
		global.GVA_PUBER.Pub(uint32(SentinelDBCluster.SentinelClusterID),
			&publish.PubsubMessage{
				DelDbCluster: &publish.DelDBCluster{
					DbClusterName: cluster.DBName,
				},
			})
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSentinelDBClusterInfoByIds
//@description: 批量删除SentinelDBClusterInfo记录
//@param: ids request.IdsReq
//@return: err error

func DeleteSentinelDBClusterInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SentinelDBClusterInfo{}, "cluster_id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSentinelDBClusterInfo
//@description: 更新SentinelDBClusterInfo记录
//@param: SentinelDBCluster *model.SentinelDBClusterInfo
//@return: err error

func UpdateSentinelDBClusterInfo(SentinelDBCluster model.SentinelDBClusterInfo) (err error) {
	err = global.GVA_DB.Save(&SentinelDBCluster).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSentinelDBClusterInfo
//@description: 根据id获取SentinelDBClusterInfo记录
//@param: id uint
//@return: err error, SentinelDBCluster model.SentinelDBClusterInfo

func GetSentinelDBClusterInfo(id int) (err error, SentinelDBCluster model.SentinelDBClusterInfo) {
	err = global.GVA_DB.Where("cluster_id = ?", id).First(&SentinelDBCluster).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSentinelDBClusterInfoInfoList
//@description: 分页获取SentinelDBClusterInfo记录
//@param: info request.SentinelDBClusterInfoSearch
//@return: err error, list interface{}, total int64

func GetSentinelDBClusterInfoInfoList(info request.SentinelDBClusterInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SentinelDBClusterInfo{})
	var SentinelDBClusters []model.SentinelDBClusterInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.SentinelClusterID != 0 {
		db = db.Where("`sentinel_cluster_id` = ?", info.SentinelClusterID)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&SentinelDBClusters).Error
	return err, SentinelDBClusters, total
}
