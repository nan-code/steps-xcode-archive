package main

import (
	"testing"

	"github.com/bitrise-io/go-xcode/exportoptions"
	"github.com/bitrise-io/go-xcode/xcarchive"
)

func Test_getExportOptions(t *testing.T) {
	type arguments struct {
		cfgUploadBitcode              string
		cfgCompileBitcode             string
		cfgICloudContainerEnvironment string
		cfgTeamID                     string
		cfgExportMethod               string
		cfgProjectPath                string
		cfgScheme                     string
		cfgConfiguration              string
		archiveExportMethod           exportoptions.Method
		mainApplication               xcarchive.IosApplication
		xcodeMajorVersion             int64
		archiveCodeSignIsXcodeManaged bool
	}

	archive, err := xcarchive.NewIosArchive("/Users/tamaspapik/Downloads/debgrunrunit/Runrunit.xcarchive")
	if err != nil {
		fail("Failed to parse archive, error: %s", err)
	}

	args := arguments{
		cfgUploadBitcode:              "no",
		cfgCompileBitcode:             "no",
		cfgICloudContainerEnvironment: "",
		cfgTeamID:                     "",
		cfgExportMethod:               "app-store",
		cfgProjectPath:                "/Users/tamaspapik/Downloads/debgrunrunit/Runrunit.xcworkspace",
		cfgScheme:                     "Runrunit",
		cfgConfiguration:              "",
		archiveExportMethod:           exportoptions.MethodAppStore,
		xcodeMajorVersion:             10,
		mainApplication:               archive.Application,
		archiveCodeSignIsXcodeManaged: false,
	}

	got, err := getExportOptions(args.cfgUploadBitcode, args.cfgCompileBitcode, args.cfgICloudContainerEnvironment, args.cfgTeamID, args.cfgExportMethod, args.cfgProjectPath, args.cfgScheme, args.cfgConfiguration, args.archiveExportMethod, args.mainApplication, args.xcodeMajorVersion, args.archiveCodeSignIsXcodeManaged)
	if err != nil {
		t.Fatal(err)
	}

	str, err := got.String()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(str)

	t.FailNow()

	// if (err != nil) != wantErr {
	// 	t.Errorf("getExportOptions() error = %v, wantErr %v", err, wantErr)
	// 	return
	// }
	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("getExportOptions() = %v, want %v", got, want)
	// }

}
