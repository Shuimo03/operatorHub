package kubernetesutil

const (
	Standalone  string = "standalone"
	Replication string = "replication"
	Cluster     string = "cluster"
	Sentinel    string = "sentinel"
)

// 根据redis type生成不同的label
func buildLabel(rt string, crName string) map[string]string {
	label := map[string]string{
		rt: crName,
	}
	return label
}
