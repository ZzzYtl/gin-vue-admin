import service from '@/utils/request'

// @Tags Node
// @Summary 获取SentinelCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Node true "获取SentinelClusters"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
export const getAllSentinelClusters = (data) => {
     return service({
         url: "/sentinel_cluster/getAllSentinelClusters",
         method: 'post',
         data
     })
 }
