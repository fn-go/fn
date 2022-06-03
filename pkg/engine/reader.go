package engine

import (
	"fmt"
	"io/fs"
	"path"

	hfsos "github.com/hack-pad/hackpadfs/os"
	"github.com/hashicorp/go-multierror"
	"sigs.k8s.io/yaml"

	"github.com/go-fn/fn/pkg/fnfile"
)

type Reader func() (fnfile.FnFile, error)

type YamlFileReaderOptions struct {
	FS         StatReadFS
	FileFinder FileFinder
}

type StatReadFS interface {
	fs.StatFS
	fs.ReadFileFS
}

type FileFinder func(dir string, fs StatReadFS) ([]byte, error)

func AmbiguousExtensionFileFinder(name string, extensions ...string) FileFinder {
	return func(dir string, fs StatReadFS) ([]byte, error) {
		var attempts *multierror.Error
		for _, ext := range extensions {
			filePath := path.Join(dir, name+"."+ext)
			var fi, err = fs.Stat(filePath)
			attempts = multierror.Append(attempts, fmt.Errorf("trying file: %s: %w", filePath, err))
			if err == nil && !fi.IsDir() {
				return fs.ReadFile(filePath)
			}
		}

		return nil, attempts.ErrorOrNil()
	}
}

func YamlFileReader(options ...func(readerOptions *YamlFileReaderOptions)) Reader {
	opts := &YamlFileReaderOptions{
		FS:         hfsos.NewFS(),
		FileFinder: AmbiguousExtensionFileFinder("fnfile", "yml", "yaml"),
	}

	for _, o := range options {
		o(opts)
	}

	return func() (file fnfile.FnFile, err error) {
		var fnFileBytes []byte
		fnFileBytes, err = opts.FileFinder(".", opts.FS)
		if err != nil {
			err = fmt.Errorf("finding/reading file: %w", err)
			return
		}

		err = yaml.Unmarshal(fnFileBytes, &file)
		return
	}
}
