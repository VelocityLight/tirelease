package service

import (
	"fmt"
	"strings"
	"strconv"

	"tirelease/internal/entity"
	"tirelease/internal/repository"
)

func CreateReleaseVersion(releaseVersion *entity.ReleaseVersion) error {
	releaseVersion.Name = ComposeVersionName(releaseVersion.Major, releaseVersion.Minor, releaseVersion.Patch, releaseVersion.Addition)
	releaseVersion.Type = ComposeVersionType(releaseVersion.Major, releaseVersion.Minor, releaseVersion.Patch, releaseVersion.Addition)
	releaseVersion.Status = ComposeVersionStatus(releaseVersion.Type)
	return repository.CreateReleaseVersion(releaseVersion)
}

func UpdateReleaseVersion(releaseVersion *entity.ReleaseVersion) error {
	releaseVersion.Name = ComposeVersionName(releaseVersion.Major, releaseVersion.Minor, releaseVersion.Patch, releaseVersion.Addition)
	releaseVersion.Type = ComposeVersionType(releaseVersion.Major, releaseVersion.Minor, releaseVersion.Patch, releaseVersion.Addition)
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

// ====================================================
// ==================================================== Compose Function
func ComposeVersionName(major, minor, patch int, addition string) string {
	if "" == addition {
		return fmt.Sprintf("%d.%d.%d", major, minor, patch)
	} else {
		return fmt.Sprintf("%d.%d.%d-%s", major, minor, patch, addition)
	}
}

func ComposeVersionMinorName(version *entity.ReleaseVersion) string {
	return fmt.Sprintf("%d.%d", version.Major, version.Minor)
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

func ComposeVersionType(major, minor, patch int, addition string) entity.ReleaseVersionType {
	if "" != addition {
		return entity.ReleaseVersionTypeHotfix
	} else {
		if patch != 0 {
			return entity.ReleaseVersionTypePatch
		} else {
			if minor != 0 {
				return entity.ReleaseVersionTypeMinor
			} else {
				return entity.ReleaseVersionTypeMajor
			}
		}
	}
}

func ComposeVersionStatus(vt entity.ReleaseVersionType)  entity.ReleaseVersionStatus {
	if entity.ReleaseVersionTypePatch == vt {
		return entity.ReleaseVersionStatusPlanning
	} else {
		return entity.ReleaseVersionStatusUpcoming
	}
}
