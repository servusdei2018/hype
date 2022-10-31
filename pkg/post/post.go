package post

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/servusdei2018/hype/pkg/config"
	"github.com/servusdei2018/hype/pkg/futil"
)

// NewPost creates a new post.
func NewPost(conf *config.Conf) (error, string) {
	err, data := EditFile(conf)
	if err != nil {
		return err, ""
	}
	if len(data) == 0 {
		return fmt.Errorf("cancelled"), ""
	}

	// Directory structure:
	// ./www/posts/{year}/{month}/{day}-{hour}{minute}{second}
	now := time.Now()
	year, month := fmt.Sprint(now.Year()), fmt.Sprint(now.Month())
	fn := fmt.Sprintf("%d-%d%d%d", now.Day(), now.Hour(), now.Minute(), now.Second())

	// Create directories as necessary.
	if err := futil.DirMustExist(path.Join(futil.WD, "www"), "posts"); err != nil {
		return err, ""
	}
	if err := futil.DirMustExist(path.Join(futil.WD, "www", "posts"), year); err != nil {
		return err, ""
	}
	if err := futil.DirMustExist(path.Join(futil.WD, "www", "posts", year), month); err != nil {
		return err, ""
	}

	// Open file for writing.
	f, err := os.Create(path.Join(futil.WD, "www", "posts", year, month, fn))
	if err != nil {
		return err, ""
	}
	defer f.Close()

	_, err = f.Write(data)
	return err, f.Name()
}

// EditFile creates a temporary file and allows the user to edit it using a specified editor, and
// returns its contents.
func EditFile(conf *config.Conf) (error, []byte) {
	// Create temporary file
	f, err := os.CreateTemp("", "post")
	if err != nil {
		return err, nil
	}
	f.Close()

	// Edit the file with the selected editor
	cmd := exec.Command(conf.Editor, f.Name())
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	err = cmd.Run()
	if err != nil {
		return err, nil
	}

	// Read file contents
	data, err := os.ReadFile(f.Name())

	return err, data
}
