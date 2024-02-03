package standalone

import (
	v1 "cola.redis/api/v1"
	"cola.redis/kubernetesutil"
	"github.com/sirupsen/logrus"
)

func CreateStandaloneRedis(cr *v1.Standalone) error {
	//这里考虑链式调用
	if err := kubernetesutil.CreateVolume(cr); err != nil {
		logrus.Error(err, "Cannot create standalone Volume for Redis ")
	}
	//根据configmap中信息填充到StatefulSet，之后创建对应SVC
	if err := kubernetesutil.CreateStatefulSet(cr); err != nil {
		logrus.Error(err, "Cannot create standalone statefulset for Redis")
		return err
	}
	//if err := kubernetesutil.CreateSVC(); err != nil {
	//
	//}
	return nil
}
