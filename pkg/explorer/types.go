package explorer

import "os"

type remoteEntry struct {
	name string
}

func (r *remoteEntry) Name() string               { return r.name }
func (r *remoteEntry) IsDir() bool                { return false }
func (r *remoteEntry) Type() os.FileMode          { return 0 }
func (r *remoteEntry) Info() (os.FileInfo, error) { return nil, nil }
