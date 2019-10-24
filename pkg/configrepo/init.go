package configrepo

import (
	"github.com/hashicorp/go-version"
	"github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type App struct {
	Name     string
	Branches map[string]*plumbing.Reference
	Versions []*version.Version
}

func NewApp(name string) *App {
	return &App{name, make(map[string]*plumbing.Reference), make([]*version.Version, 0)}
}

type ErrorHandlerFn func(err error)

type ConfigRepo struct {
	repo          *git.Repository
	Apps          map[string]*App
	pullCh        chan bool
	cloneOpts     *git.CloneOptions
	errorsCh      chan error
	errorHandlers []ErrorHandlerFn
}

func NewConfigRepo(localPath string, cloneOpts *git.CloneOptions) (*ConfigRepo, error) {
	log := logrus.WithField("localPath", localPath)
	log.Info("New Config Repo")
	repo, err := git.PlainOpen(localPath)
	if err == git.ErrRepositoryNotExists {
		log.Warn("no repo found")
		if cloneOpts != nil {
			log.Infof("cloning it from :%+v", cloneOpts)
			cloneOpts.Tags = git.AllTags
			cloneOpts.NoCheckout = true
			repo, err = git.PlainClone(localPath, true, cloneOpts)
		}
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &ConfigRepo{repo, make(map[string]*App), make(chan bool), cloneOpts, make(chan error), make([]ErrorHandlerFn, 0)}, nil
}

func (cr *ConfigRepo) Init() error {
	cr.errorHandlerManager()
	return cr.LoadApps()
}