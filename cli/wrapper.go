package cli

import (
	"crypto/tls"

	"github.com/codegangsta/cli"
	"github.com/docker/swarm/api"
	"github.com/docker/swarm/cluster"
	"github.com/docker/swarm/discovery"
	"github.com/docker/swarm/leadership"
)

func Manage(c *cli.Context) {
	manage(c)
}

func RunWithCommand(newCommands []cli.Command) {
	commands = newCommands
	Run()
}

func GetDiscovery(c *cli.Context) string {
	return getDiscovery(c)
}

func CheckAddrFormat(addr string) bool {
	return checkAddrFormat(addr)
}

func LoadTLSConfig(ca, cert, key string, verify bool) (*tls.Config, error) {
	return loadTLSConfig(ca, cert, key, verify)
}

func NewStatusHandler(cluster cluster.Cluster, candidate *leadership.Candidate, follower *leadership.Follower) api.StatusHandler {

	handler := statusHandler{
		cluster:   cluster,
		candidate: candidate,
		follower:  follower,
	}
	return &handler
}

func CreateDiscovery(uri string, c *cli.Context) discovery.Discovery {
	return createDiscovery(uri, c)
}

func Follow(follower *leadership.Follower, replica *api.Replica, addr string) {
	follow(follower, replica, addr)
}
