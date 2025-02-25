// DO NOT EDIT: generated by syft/internal/packagemetadata/generate/main.go

package packagemetadata

import "github.com/anchore/syft/syft/pkg"

// AllTypes returns a list of all pkg metadata types that syft supports (that are represented in the pkg.Package.Metadata field).
func AllTypes() []any {
	return []any{pkg.AlpmMetadata{}, pkg.ApkMetadata{}, pkg.BinaryMetadata{}, pkg.CargoPackageMetadata{}, pkg.CocoapodsMetadata{}, pkg.ConanLockMetadata{}, pkg.ConanMetadata{}, pkg.DartPubMetadata{}, pkg.DotnetDepsMetadata{}, pkg.DotnetPortableExecutableMetadata{}, pkg.DpkgMetadata{}, pkg.GemMetadata{}, pkg.GolangBinMetadata{}, pkg.GolangModMetadata{}, pkg.HackageMetadata{}, pkg.JavaMetadata{}, pkg.KbPackageMetadata{}, pkg.LinuxKernelMetadata{}, pkg.LinuxKernelModuleMetadata{}, pkg.MixLockMetadata{}, pkg.NixStoreMetadata{}, pkg.NpmPackageJSONMetadata{}, pkg.NpmPackageLockJSONMetadata{}, pkg.PhpComposerJSONMetadata{}, pkg.PortageMetadata{}, pkg.PythonPackageMetadata{}, pkg.PythonPipfileLockMetadata{}, pkg.PythonRequirementsMetadata{}, pkg.RDescriptionFileMetadata{}, pkg.RebarLockMetadata{}, pkg.RpmMetadata{}}
}
