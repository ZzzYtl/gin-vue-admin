import service from '@/utils/request'

export const getAllLogicClusters = (data) => {
     return service({
         url: "/logicCluster/getAllLogicClusters",
         method: 'post',
         data
     })
 }
