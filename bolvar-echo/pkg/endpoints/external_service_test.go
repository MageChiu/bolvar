package endpoints

import "testing"

func TestName(t *testing.T) {
	kubeConfFile := ""
	kubeClient := buildKubeClient(kubeConfFile)
	genService("test", "test", kubeClient, "/Users/charles/code/goHome/bolvar/output")
}
