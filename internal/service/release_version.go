package service

import (
	"fmt"
	"strconv"
	"strings"

	"tirelease/internal/entity"
	"tirelease/internal/repository"

	"github.com/deckarep/golang-set"
)

func CreateReleaseVersion(releaseVersion *entity.ReleaseVersion) error {
	releaseVersion.Name = ComposeVersionName(releaseVersion)
	releaseVersion.Type = ComposeVersionType(releaseVersion)
	releaseVersion.Status = ComposeVersionStatus(releaseVersion.Type)
	releaseVersion.ReleaseBranch = ComposeVersionBranch(releaseVersion)
	return repository.CreateReleaseVersion(releaseVersion)
}

func UpdateReleaseVersion(releaseVersion *entity.ReleaseVersion) error {
	releaseVersion.Name = ComposeVersionName(releaseVersion)
	releaseVersion.Type = ComposeVersionType(releaseVersion)
	err := repository.UpdateReleaseVersion(releaseVersion)
	if nil != err {
		return err
	}

	if releaseVersion.Type == entity.ReleaseVersionTypeHotfix {
		return nil
	}
	if releaseVersion.Status == entity.ReleaseVersionStatusReleased {
		// 版本发布操作
		option := &entity.ReleaseVersionOption{
			Major: releaseVersion.Major,
			Minor: releaseVersion.Minor,
			Patch: releaseVersion.Patch + 1,
		}
		lastVersion, err := repository.SelectReleaseVersionLatest(option)
		if nil != err || nil == lastVersion {
			return nil
		}
		lastVersion.Status = entity.ReleaseVersionStatusUpcoming
		err = repository.UpdateReleaseVersion(lastVersion)
		if nil != err {
			return err
		}
	}
	return nil
}

func SelectReleaseVersion(option *entity.ReleaseVersionOption) (*[]entity.ReleaseVersion, error) {
	return repository.SelectReleaseVersion(option)
}

func SelectReleaseVersionMaintained() (*[]string, error) {
	option := &entity.ReleaseVersionOption{
		Status: entity.ReleaseVersionStatusUpcoming,
	}
	versions, err := repository.SelectReleaseVersion(option)
	if nil != err {
		return nil, err
	}
	set := mapset.NewSet()
	for _, version := range *versions {
		set.Add(ComposeVersionMinorName(&version))
	}
	var res []string
	for _, v := range set.ToSlice() {
		res = append(res, v.(string))
	}
	return &res, nil
}

// ====================================================
// ==================================================== Compose Function
func ComposeVersionName(version *entity.ReleaseVersion) string {
	if version.Addition == "" {
		return fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)
	} else {
		return fmt.Sprintf("%d.%d.%d-%s", version.Major, version.Minor, version.Patch, version.Addition)
	}
}

func ComposeVersionMinorName(version *entity.ReleaseVersion) string {
	return fmt.Sprintf("%d.%d", version.Major, version.Minor)
}

func ComposeVersionBranch(version *entity.ReleaseVersion) string {
	return fmt.Sprintf("release-%d.%d", version.Major, version.Minor)
}

func ComposeVersionType(version *entity.ReleaseVersion) entity.ReleaseVersionType {
	if version.Addition != "" {
		return entity.ReleaseVersionTypeHotfix
	} else {
		if version.Patch != 0 {
			return entity.ReleaseVersionTypePatch
		} else {
			if version.Minor != 0 {
				return entity.ReleaseVersionTypeMinor
			} else {
				return entity.ReleaseVersionTypeMajor
			}
		}
	}
}

func ComposeVersionShortType(version string) entity.ReleaseVersionShortType {
	// todo: regexp later
	slice := strings.Split(version, "-")
	if len(slice) >= 2 {
		return entity.ReleaseVersionShortTypeHotfix
	}

	slice = strings.Split(slice[0], ".")
	if len(slice) == 3 {
		return entity.ReleaseVersionShortTypePatch
	}
	if len(slice) == 2 {
		return entity.ReleaseVersionShortTypeMinor
	}
	if len(slice) == 1 {
		return entity.ReleaseVersionShortTypeMajor
	}
	return entity.ReleaseVersionShortTypeUnKnown
}

func ComposeVersionAtom(version string) (major, minor, patch int, addition string) {
	major = 0
	minor = 0
	patch = 0
	addition = ""

	slice := strings.Split(version, "-")
	if len(slice) >= 2 {
		addition = slice[1]
	}

	slice = strings.Split(slice[0], ".")
	if len(slice) >= 1 {
		major, _ = strconv.Atoi(slice[0])
	}
	if len(slice) >= 2 {
		minor, _ = strconv.Atoi(slice[1])
	}
	if len(slice) >= 3 {
		patch, _ = strconv.Atoi(slice[2])
	}

	return major, minor, patch, addition
}

func ComposeVersionStatus(vt entity.ReleaseVersionType) entity.ReleaseVersionStatus {
	if entity.ReleaseVersionTypePatch == vt {
		return entity.ReleaseVersionStatusPlanned
	} else {
		return entity.ReleaseVersionStatusUpcoming
	}
}
