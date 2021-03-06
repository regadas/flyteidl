package logs

import (
	"fmt"

	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
)

type kubernetesLogPlugin struct {
	k8sURL string
}

func (s kubernetesLogPlugin) GetTaskLog(podName, namespace, containerName, containerID, logName string) (core.TaskLog, error) {
	return core.TaskLog{
		Uri:           fmt.Sprintf("%s/#!/log/%s/%s/pod?namespace=%s", s.k8sURL, namespace, podName, namespace),
		Name:          logName,
		MessageFormat: core.TaskLog_UNKNOWN,
	}, nil
}

func NewKubernetesLogPlugin(k8sURL string) LogPlugin {
	return &kubernetesLogPlugin{
		k8sURL: k8sURL,
	}
}
