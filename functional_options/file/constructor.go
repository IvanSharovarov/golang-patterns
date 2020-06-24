package file

import "os"

func New(filepath string, setters ...Option) error {
	// Default Options
	args := &Options{
		UID:         os.Getuid(),
		GID:         os.Getegid(),
		Flags:       os.O_CREATE | os.O_EXCL | os.O_WRONLY,
		Contents:    "",
		Permissions: 0666,
	}

	for _, setter := range setters {
		setter(args)
	}

	f, err := os.OpenFile(filepath, args.Flags, args.Permissions)
	if err != nil {
		return err
	} else {
		defer f.Close()
	}

	if _, err := f.WriteString(args.Contents); err != nil {
		return err
	}

	return f.Chown(args.UID, args.GID)
}
