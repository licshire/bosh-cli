package stemcell

import (
	boshsys "github.com/cloudfoundry/bosh-agent/system"
	bmcloud "github.com/cloudfoundry/bosh-micro-cli/cloud"
	bmeventlog "github.com/cloudfoundry/bosh-micro-cli/eventlogger"
)

type ManagerFactory interface {
	NewManager(bmcloud.Cloud) Manager
}

type managerFactory struct {
	fs          boshsys.FileSystem
	reader      Reader
	repo        Repo
	eventLogger bmeventlog.EventLogger
}

func NewManagerFactory(repo Repo, eventLogger bmeventlog.EventLogger) ManagerFactory {
	return &managerFactory{
		repo:        repo,
		eventLogger: eventLogger,
	}
}

func (f *managerFactory) NewManager(cloud bmcloud.Cloud) Manager {
	return &manager{
		repo:        f.repo,
		eventLogger: f.eventLogger,
		cloud:       cloud,
	}
}