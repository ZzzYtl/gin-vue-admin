package service

import (
	"encoding/json"
	"errors"
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
	var total int64
	err = global.GVA_DB.Model(SentinelDBCluster).Where("cluster_id = ? and logic_cluster_id = ?", SentinelDBCluster.ClusterID, SentinelDBCluster.LogicClusterID).Count(&total).Error
	if err != nil {
		return err
	}
	if total != 0 {
		return errors.New("this dbcluster is already exist")
	}

	logicCluster := model.LogicCluster{
		LogicClusterID: SentinelDBCluster.LogicClusterID,
	}
	err = global.GVA_DB.First(&logicCluster).Error
	if err != nil {
		return err
	}
	dbClusterInfo := model.DataBaseInfo{}
	err = global.GVA_DB.Where("tag_id = ? and cluster_name = ?", SentinelDBCluster.ClusterID, logicCluster.Name).First(&dbClusterInfo).Error
	if err != nil {
		return err
	}

	var dbs []publish.DBConfig
	var nodes []model.Node
	err = global.GVA_DB.Where("tag_id = ?", SentinelDBCluster.ClusterID).Find(&nodes).Error
	if err != nil {
		return err
	}
	for _, v := range nodes {
		dbs = append(dbs, publish.DBConfig{
			Ip:   v.IP,
			Type: publish.DBType(v.RoleID),
		})
	}

	backUp := model.BackUpDB{
		BackUpID: dbClusterInfo.BackUpID,
	}
	err = global.GVA_DB.Where("backup_id = ?", dbClusterInfo.BackUpID).Find(&backUp).Error
	if err != nil {
		return err
	}

	if len(backUp.IP) > 0 {
		dbs = append(dbs, publish.DBConfig{
			Ip:   backUp.IP,
			Type: publish.DBType_OBServer,
		})
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
				Name:        SentinelDBCluster.GetDBClusterName(),
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
				Ip:   db.Ip,
				Port: uint64(dbClusterInfo.Port),
				Type: db.Type,
			})
		}
		global.GVA_PUBER.Pub(uint32(SentinelDBCluster.SentinelClusterID),
			&publish.PubsubMessage{AddDbCluster: msg})
	}
	return err
}

type OldMasterForm struct {
	ID int    `json:"id" form:"id"`
	IP string `json:"ip" form:"ip"`
}

func ChangeOldMaster2Slave(changeInfo OldMasterForm) (err error) {
	sentinelDBClu := model.SentinelDBClusterInfo{
		ID:                changeInfo.ID,
	}
	err = global.GVA_DB.First(&sentinelDBClu).Error
	if err != nil {
		return err
	}

	var dbs []publish.DBConfig
	if json.Unmarshal([]byte(sentinelDBClu.Dbs), &dbs) != nil {
		return errors.New("cant find this old master")
	}

	find := false
	for i := range dbs {
		if dbs[i].Ip == changeInfo.IP && dbs[i].Type == publish.DBType_OldMaster {
			find = true
			dbs[i].Type = publish.DBType_Slave
			break
		}
	}

	if find == false {
		return errors.New("cant find this old master...")
	}
	data, err := json.Marshal(dbs)
	if err != nil {
		return err
	}
	sentinelDBClu.Dbs = string(data)

	msg := &publish.ChangeOldMaster2Slave{
		DbClusterName: sentinelDBClu.GetDBClusterName(),
		OldMasterIp:   changeInfo.IP,
	}

	_, err = global.GVA_PUBER.Pub(uint32(sentinelDBClu.SentinelClusterID),
		&publish.PubsubMessage{ChangeOldMaster_2Slave: msg})
	if err == nil {
		err = global.GVA_DB.Save(&sentinelDBClu).Error
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSentinelDBClusterInfo
//@description: 删除SentinelDBClusterInfo记录
//@param: SentinelDBCluster model.SentinelDBClusterInfo
//@return: err error

func DeleteSentinelDBClusterInfo(SentinelDBCluster model.SentinelDBClusterInfo) (err error) {
	global.GVA_DB.First(&SentinelDBCluster)
	err = global.GVA_DB.Delete(&SentinelDBCluster).Error
	if err == nil && len(SentinelDBCluster.GetDBClusterName()) != 0 {
		global.GVA_PUBER.Pub(uint32(SentinelDBCluster.SentinelClusterID),
			&publish.PubsubMessage{
				DelDbCluster: &publish.DelDBCluster{
					DbClusterName: SentinelDBCluster.GetDBClusterName(),
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
