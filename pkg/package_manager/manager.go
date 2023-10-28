package packagemanager

import (
	"context"

	contextInternal "github.com/gameap/gameapctl/internal/context"
)

type PackageInfo struct {
	Name            string
	Architecture    string
	Version         string
	Size            string
	Description     string
	InstalledSizeKB int
}

type PackageManager interface {
	Search(ctx context.Context, name string) ([]PackageInfo, error)
	Install(ctx context.Context, packs ...string) error
	CheckForUpdates(ctx context.Context) error
	Remove(ctx context.Context, packs ...string) error
	Purge(ctx context.Context, packs ...string) error
}

//nolint:ireturn,nolintlint
func Load(ctx context.Context) (PackageManager, error) {
	osInfo := contextInternal.OSInfoFromContext(ctx)

	switch osInfo.Distribution {
	case "debian", "ubuntu":
		return NewExtendedAPT(&APT{}), nil
	case "windows":
		return NewWindowsPackageManager(), nil
	}

	return nil, NewErrUnsupportedDistribution(osInfo.Distribution)
}
