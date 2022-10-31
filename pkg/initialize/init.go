package initialize

import (
	"os"
	"path"

	"github.com/servusdei2018/hype/pkg/futil"
)

// Init initializes directory structure.
func Init(version string) error {
	var err error

	// $WD/.hype/hype.env
	f, err := os.Create(path.Join(futil.WD, "hype.env"))
	if err != nil {
		return err
	}
	if _, err = f.Write([]byte("HYPE_EDITOR=vi\nHYPE_PORT=80\n")); err != nil {
		return err
	}
	f.Close()

	// $WD/www
	if err = os.Mkdir(path.Join(futil.WD, "www"), 0754); err != nil {
		return err
	}

	if f, err = os.Create(path.Join(futil.WD, "www", "index.html")); err != nil {
		return err
	}
	if _, err = f.Write([]byte(INDEX)); err != nil {
		return err
	}
	f.Close()

	if err = os.Mkdir(path.Join(futil.WD, "www", "posts"), 0754); err != nil {
		return err
	}

	return nil
}

// IsInit returns whether Hype was initialized already.
func IsInit() bool {
	return futil.Exists(path.Join(futil.WD, "hype.env"))
}
