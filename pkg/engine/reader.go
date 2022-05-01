package engine

import (
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/ghodss/yaml"
	hfsos "github.com/hack-pad/hackpadfs/os"
	"github.com/hashicorp/go-multierror"

	"github.com/go-fn/fn/pkg/fnfile"
)

type Reader func() (*fnfile.Fnfile, error)

type YamlFileReaderOptions struct {
	FS         StatReadFS
	FileFinder FileFinder
	Getwd      Getwd
}

type StatReadFS interface {
	fs.StatFS
	fs.ReadFileFS
}

type FileFinder func(workingDir string, fs StatReadFS) ([]byte, error)

type Getwd func() (string, error)

func YamlFileReader(options ...func(readerOptions *YamlFileReaderOptions)) Reader {
	opts := &YamlFileReaderOptions{
		FS:         hfsos.NewFS(),
		Getwd:      os.Getwd,
		FileFinder: AmbiguousExtensionFileFinder("fnfile", "yml", "yaml"),
	}

	for _, o := range options {
		o(opts)
	}

	return func() (*fnfile.Fnfile, error) {
		var workingDir string
		{
			var err error
			workingDir, err = opts.Getwd()
			if err != nil {
				return nil, fmt.Errorf("getting working directory: %w", err)
			}
		}

		var fnFileBytes []byte
		{
			var err error
			fnFileBytes, err = opts.FileFinder(workingDir, opts.FS)
			if err != nil {
				return nil, fmt.Errorf("finding/reading file: %w", err)
			}
		}

		var fnFile fnfile.Fnfile

		{
			err := yaml.Unmarshal(fnFileBytes, &fnFile)
			if err != nil {
				return nil, fmt.Errorf("unmarshalling yaml: %v\n", err)
			}
		}

		return &fnFile, nil
	}
}

func AmbiguousExtensionFileFinder(name string, extensions ...string) FileFinder {
	return func(workingDir string, fs StatReadFS) ([]byte, error) {
		var attempts *multierror.Error
		for _, ext := range extensions {
			filePath := path.Join(workingDir, name+"."+ext)
			var fi, err = fs.Stat(filePath)
			attempts = multierror.Append(attempts, fmt.Errorf("trying file: %s: %w", filePath, err))
			if err == nil && !fi.IsDir() {
				return fs.ReadFile(filePath)
			}
		}

		return nil, attempts.ErrorOrNil()
	}
}
