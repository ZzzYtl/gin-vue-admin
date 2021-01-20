package service

import (
	"gin-vue-admin/global"
	"go.uber.org/zap"
)

type QueryClusterResult struct {
	MasterIp string `json:"master_ip"`
	Port     int    `json:"port"`
	DbName   string `json:"db_name"`
	BackupIp string `json:"backup_ip"`
	BackupDomain string `json:"backup_domain"`
}

func QueryClusterInfo(name string) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	var data QueryClusterResult
	sql := "select a.ip master_ip,port,db_name,c.ip backup_ip,c.domain backup_domain from node_info a,cluster_info b,backup_info c where a.tag_id=b.tag_id and b.backup_id=c.backup_id and a.role_id=1 and  db_name=?"

	rows, err := global.GVA_DB.Raw(sql, name).Rows() // (*sql.Rows, error)
	defer rows.Close()
	exist := false
	for rows.Next() {
		exist = true
		rows.Scan(&data.MasterIp, &data.Port, &data.DbName, &data.BackupIp, &data.BackupDomain)
	}

	if err != nil {
		global.GVA_LOG.Error("class=ClusterModel&method=GetByName&", zap.Any("result", err), zap.Any("sql", sql))
		return nil, err
	}

	if exist {
		result["list"] = data
	} else {
		result["list"] = nil
	}
	result["num"] = 2
	return result, nil
}
