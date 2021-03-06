package mount

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"

	"koding/klient/machine"
	"koding/klient/machine/client"
	"koding/klient/machine/index"

	"github.com/koding/logging"
)

const (
	LocalIndexName  = "index.local"  // file name of local directory index.
	RemoteIndexName = "index.remote" // file name of remote directory index.
)

// Info stores information about current mount status.
type Info struct {
	ID    ID    // Mount ID.
	Mount Mount // Mount paths stored in absolute form.

	SyncCount int // Number of synced files.
	AllCount  int // Number of all files handled by mount.

	SyncDiskSize int64 // Total size of synced files.
	AllDiskSize  int64 // Size of all files handled by mount.
}

// SupervisorOpts are the options used to configure Supervisor object.
type SupervisorOpts struct {
	// ClientFunc is a factory for dynamic clients.
	ClientFunc client.DynamicClientFunc

	// WorkDir is a working directory that will be used by syncs object. The
	// directory structure for single mount with ID will look like:
	//
	//   WorkDir
	//   |-data
	//   | +-... // mounted directory cache.
	//   |-index.remote
	//   +-index.local
	//
	WorkDir string

	// Log is used for logging. If nil, default logger will be created.
	Log logging.Logger
}

// Valid checks if provided options are correct.
func (opts *SupervisorOpts) Valid() error {
	if opts == nil {
		return errors.New("mount supervisor options are nil")
	}

	if opts.ClientFunc == nil {
		return errors.New("nil dynamic client function")
	}

	if opts.WorkDir == "" {
		return errors.New("working directory is not set")
	}

	return nil
}

// Supervisor stores and synchronizes single mount. The main goal of its logic
// is to make remote and local indexes similar.
type Supervisor struct {
	opts    SupervisorOpts
	mountID ID    // identifier of synced mount.
	m       Mount // single mount with absolute paths.
	log     logging.Logger

	ridx *index.Index // known state of remote index.
	lidx *index.Index // known state of local index.
}

// NewSupervisor creates a new supervisor instance for a given mount. It ensures
// basic mount directory structure. This function is blocking.
func NewSupervisor(mountID ID, m Mount, opts SupervisorOpts) (*Supervisor, error) {
	if err := opts.Valid(); err != nil {
		return nil, err
	}

	s := &Supervisor{
		opts:    opts,
		mountID: mountID,
		m:       m,
	}

	if opts.Log != nil {
		s.log = opts.Log.New("supervisor")
	} else {
		s.log = machine.DefaultLogger.New("supervisor")
	}

	// Create directory structure if it doesn't exist.
	if err := os.MkdirAll(filepath.Join(s.opts.WorkDir, "data"), 0755); err != nil {
		return nil, err
	}

	// Fetch or load remote index.
	var err error
	if s.ridx, err = s.loadIdx(RemoteIndexName, s.fetchRemoteIdx); err != nil {
		return nil, err
	}

	// Create or load local index.
	if s.lidx, err = s.loadIdx(LocalIndexName, s.fetchLocalIdx); err != nil {
		return nil, err
	}

	// Update local index if needed.
	if err := s.updateLocal(); err != nil {
		return nil, err
	}

	return s, nil
}

// Info returns the current status of supervised indexes.
func (s *Supervisor) Info() *Info {
	return &Info{
		ID:           s.mountID,
		Mount:        s.m,
		SyncCount:    s.lidx.Count(-1),
		AllCount:     s.ridx.Count(-1),
		SyncDiskSize: s.lidx.DiskSize(-1),
		AllDiskSize:  s.ridx.DiskSize(-1),
	}
}

// Drop closes synced mount and cleans up all resources acquired by it.
func (s *Supervisor) Drop() error {
	return os.RemoveAll(s.opts.WorkDir)
}

// loadIdx reads named index from synced working directory. If index file does
// not exist, it is fetched by calling provided `fetchIdx` function and saved to
// provided path.
func (s *Supervisor) loadIdx(name string, fetchIdx idxFunc) (*index.Index, error) {
	path := filepath.Join(s.opts.WorkDir, name)
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		idx, err := fetchIdx()
		if err != nil {
			return nil, err
		}
		return idx, index.SaveIndex(idx, path)
	} else if err != nil {
		return nil, err
	}
	defer f.Close()

	idx := index.NewIndex()
	return idx, json.NewDecoder(f).Decode(idx)
}

// idxFunc is a function used to fetch or update index.
type idxFunc func() (*index.Index, error)

// fetchRemoteIdx downloads remote index.
func (s *Supervisor) fetchRemoteIdx() (*index.Index, error) {
	spv := client.NewSupervised(s.opts.ClientFunc, 30*time.Second)
	return spv.MountGetIndex(s.m.RemotePath)
}

// fetchLocalIdx always scans mount cache directory and creates new index.
func (s *Supervisor) fetchLocalIdx() (*index.Index, error) {
	return index.NewIndexFiles(filepath.Join(s.opts.WorkDir, "data"))
}

// updateLocal updates local index and saves it to cache directory.
func (s *Supervisor) updateLocal() error {
	dataPath := filepath.Join(s.opts.WorkDir, "data")
	cs := s.lidx.Compare(dataPath)

	if len(cs) == 0 {
		return nil
	}

	s.lidx.Apply(dataPath, cs)
	return index.SaveIndex(s.lidx, filepath.Join(s.opts.WorkDir, LocalIndexName))
}
