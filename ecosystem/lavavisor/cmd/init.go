package lavavisor

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	lvstatetracker "github.com/lavanet/lava/ecosystem/lavavisor/pkg"
	"github.com/lavanet/lava/utils"
	protocoltypes "github.com/lavanet/lava/x/protocol/types"
	"github.com/spf13/cobra"
)

var cmdLavavisorInit = &cobra.Command{
	Use:   "init",
	Short: "initializes the environment for LavaVisor",
	Long: `Prepares the local environment for the operation of LavaVisor.
	config.yml should be located in the ./lavavisor/ directory.`,
	Args: cobra.ExactArgs(0),
	Example: `optional flags: --directory | --auto-download 
		lavavisor init <flags>
		lavavisor init --directory ./custom/lavavisor/path 
		lavavisor init --directory ./custom/lavavisor/path --auto-download true`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, _ := cmd.Flags().GetString("directory")
		dir, err := expandTilde(dir)
		if err != nil {
			return fmt.Errorf("unable to expand directory path: %w", err)
		}
		fmt.Println("dir: ", dir)
		// Build path to ./lavavisor
		lavavisorPath := filepath.Join(dir, "./.lavavisor")

		fmt.Println("lavavisorPath: ", lavavisorPath)

		// Check if ./lavavisor directory exists
		if _, err := os.Stat(lavavisorPath); os.IsNotExist(err) {
			// ToDo: handle case where user didn't set up the file
			return fmt.Errorf("lavavisor directory does not exist at path: %s", lavavisorPath)
		}

		// 1- check lava-protocol version from consensus
		// ...GetProtocolVersion()
		// handle flags, pass necessary fields
		ctx := context.Background()

		clientCtx, err := client.GetClientQueryContext(cmd)
		if err != nil {
			return err
		}

		sq := lvstatetracker.NewStateQuery(clientCtx)
		protoVer, err := sq.GetProtocolVersion(ctx)
		if err != nil {
			return utils.LavaFormatError("protcol version cannot be fetched from consensus", err)
		}
		version := protocoltypes.Version{
			ConsumerTarget: protoVer.ConsumerTarget,
			ConsumerMin:    protoVer.ConsumerMin,
			ProviderTarget: protoVer.ProviderTarget,
			ProviderMin:    protoVer.ProviderMin,
		}
		fmt.Println("version.ConsumerTarget: ", version.ConsumerTarget)
		fmt.Println("version.ConsumerMin: ", version.ConsumerMin)
		fmt.Println("version.ProviderTarget: ", version.ProviderTarget)
		fmt.Println("version.ProviderMin: ", version.ProviderMin)

		// 2- search extracted directory inside ./lavad/upgrades/<fetched_version>
		// first check target version, then check min version
		versionDir := filepath.Join(lavavisorPath, "upgrades", "v"+version.ConsumerMin)
		binaryPath := filepath.Join(versionDir, "lava-protocol")
		fmt.Println("verisonDir: ", versionDir)
		fmt.Println("binaryPath: ", binaryPath)

		// check if version directory exists
		if _, err := os.Stat(versionDir); os.IsNotExist(err) {
			// directory doesn't exist, check auto-download flag
			autoDownload, err := cmd.Flags().GetBool("auto-download")
			if err != nil {
				return err
			}
			if autoDownload {
				utils.LavaFormatInfo("Version directory does not exist, but auto-download is enabled. Attempting to download binary from GitHub...")
				os.MkdirAll(versionDir, os.ModePerm) // before downloading, ensure version directory exists
				err = fetchAndBuildFromGithub(version.ConsumerMin, versionDir)
				if err != nil {
					utils.LavaFormatError("Failed to auto-download binary from GitHub\n ", err)
					os.Exit(1)
				}
				utils.LavaFormatInfo("Auto-download successful.")
			} else {
				utils.LavaFormatError("Sub-directory for version not found in lavavisor", nil, utils.Attribute{Key: "version", Value: version.ConsumerMin})
				os.Exit(1)
			}
		} else {
			err = checkLavaProtocolVersion(version.ConsumerMin, binaryPath)
			if err != nil {
				// binary not found or version mismatch, check auto-download flag
				autoDownload, err := cmd.Flags().GetBool("auto-download")
				if err != nil {
					return err
				}
				if autoDownload {
					utils.LavaFormatInfo("Version mismatch or binary not found, but auto-download is enabled. Attempting to download binary from GitHub...")
					err = fetchAndBuildFromGithub(version.ConsumerMin, versionDir)
					if err != nil {
						utils.LavaFormatError("Failed to auto-download binary from GitHub\n ", err)
						os.Exit(1)
					}
					fmt.Println("Auto-download successful.")
				} else {
					utils.LavaFormatError("Protocol version mismatch or binary not found in lavavisor directory\n ", err)
					os.Exit(1)
				}
			}
		}

		fmt.Println("Version check passed")

		// 3- if found: create a link from that binary to $(which lava-protocol)
		out, err := exec.Command("which", "lava-protocol").Output()
		if err != nil {
			// if "which" command fails, copy binary to system path
			gobin, err := exec.Command("go", "env", "GOPATH").Output()
			if err != nil {
				utils.LavaFormatFatal("couldn't determine Go binary path", err)
			}

			goBinPath := strings.TrimSpace(string(gobin)) + "/bin/"
			fmt.Println("goBinPath: ", goBinPath)

			// Check if the fetched binary is executable
			// ToDo: change flag to "--version" once relased binaries support the flag
			_, err = exec.Command(binaryPath, "--help").Output()
			if err != nil {
				utils.LavaFormatFatal("binary is not a valid executable: ", err)
			}

			// Check if the link already exists and remove it
			lavaLinkPath := goBinPath + "lava-protocol"
			if _, err := os.Lstat(lavaLinkPath); err == nil {
				utils.LavaFormatInfo("Discovered an existing link. Attempting to refresh.")
				err = os.Remove(lavaLinkPath)
				if err != nil {
					utils.LavaFormatFatal("couldn't remove existing link", err)
				}
			} else if !os.IsNotExist(err) {
				// other error
				utils.LavaFormatFatal("unexpected error when checking for existing link", err)
			}
			utils.LavaFormatInfo("Old binary link successfully removed. Attempting to create the updated link.")

			err = Copy(binaryPath, goBinPath+"lava-protocol")
			if err != nil {
				utils.LavaFormatFatal("couldn't copy binary to system path", err)
			}

			// try "which" command again
			out, err = exec.Command("which", "lava-protocol").Output()
			if err != nil {
				utils.LavaFormatFatal("couldn't extract binary at the system path", err)
			}
		}
		dest := strings.TrimSpace(string(out))

		if _, err := os.Lstat(dest); err == nil {
			// if destination file exists, remove it
			err = os.Remove(dest)
			if err != nil {
				utils.LavaFormatFatal("couldn't remove existing link", err)
			}
		}

		err = os.Symlink(binaryPath, dest)
		if err != nil {
			utils.LavaFormatFatal("couldn't create symbolic link", err)
		}

		// check that the link has been established
		link, err := os.Readlink(dest)
		if err != nil || link != binaryPath {
			utils.LavaFormatFatal("failed to verify symbolic link", err)
		}

		fmt.Println("Symbolic link created successfully")

		//		if autodownload false: alert user that binary is not exist, monitor directory constantly!,

		return nil
	},
}

func expandTilde(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot get user home directory: %w", err)
	}
	return filepath.Join(home, path[1:]), nil
}

func init() {
	// Use the log directory in flags
	flags.AddQueryFlagsToCmd(cmdLavavisorInit)
	cmdLavavisorInit.MarkFlagRequired(flags.FlagFrom)
	cmdLavavisorInit.Flags().String("directory", os.ExpandEnv("~/"), "Protocol Flags Directory")
	cmdLavavisorInit.Flags().Bool("auto-download", false, "Automatically download missing binaries")
	cmdLavavisorInit.Flags().String("from", "", "Name or address of private key with which to sign")
	rootCmd.AddCommand(cmdLavavisorInit)
}

func getBinaryVersion(binaryPath string) (string, error) {
	cmd := exec.Command(binaryPath, "-v")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %w", err)
	}

	// Assume the output format is "lava-protocol version x.x.x"
	version := strings.Split(string(output), " ")[2]
	fmt.Println("getBinaryVersion - version: ", version)
	return version, nil
}

func checkLavaProtocolVersion(targetVersion, binaryPath string) error {
	binaryVersion, err := getBinaryVersion(binaryPath)
	if err != nil {
		return fmt.Errorf("failed to get binary version: %w", err)
	}

	if strings.TrimSpace(binaryVersion) != strings.TrimSpace(targetVersion) {
		return fmt.Errorf("version mismatch, expected %s but got %s", targetVersion, binaryVersion)
	}

	return nil
}

func Copy(src, dest string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return fmt.Errorf("couldn't read source file: %v", err)
	}

	err = ioutil.WriteFile(dest, input, 0755)
	if err != nil {
		return fmt.Errorf("couldn't write destination file: %v", err)
	}
	return nil
}

func fetchAndBuildFromGithub(version, versionDir string) error {
	// Clean up the binary directory if it exists
	err := os.RemoveAll(versionDir)
	if err != nil {
		return fmt.Errorf("failed to clean up binary directory: %v", err)
	}
	// URL might need to be updated based on the actual GitHub repository
	url := fmt.Sprintf("https://github.com/lavanet/lava/archive/refs/tags/v%s.zip", version)
	utils.LavaFormatInfo("Fetching the source from: ", utils.Attribute{Key: "URL", Value: url})

	// Send the request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad HTTP status: %s", resp.Status)
	}

	// Prepare the path for downloaded zip
	zipPath := filepath.Join(versionDir, version+".zip")

	// Make sure the directory exists
	dir := filepath.Dir(zipPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	// Write the body to file
	out, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	// Unzip the source
	_, err = unzip(zipPath, versionDir)
	if err != nil {
		return err
	}
	utils.LavaFormatInfo("Unzipping...")

	// Build the binary
	srcPath := versionDir + "/lava-" + version
	protocolPath := srcPath + "/protocol"
	utils.LavaFormatInfo("building protocol", utils.Attribute{Key: "protocol-path", Value: protocolPath})

	cmd := exec.Command("go", "build", "-o", "lava-protocol")
	cmd.Dir = protocolPath
	err = cmd.Run()
	if err != nil {
		return err
	}

	// Move the binary to binaryPath
	err = os.Rename(filepath.Join(protocolPath, "lava-protocol"), filepath.Join(versionDir, "lava-protocol"))
	if err != nil {
		return fmt.Errorf("failed to move compiled binary: %v", err)
	}

	// Verify the compiled binary
	versionDir += "/lava-protocol"
	binaryInfo, err := os.Stat(versionDir)
	if err != nil {
		return fmt.Errorf("failed to verify compiled binary: %v", err)
	}
	binaryMode := binaryInfo.Mode()
	if binaryMode.Perm()&0111 == 0 {
		return fmt.Errorf("compiled binary is not executable")
	}
	utils.LavaFormatInfo("lava-protocol binary is successfully verified!")

	// Remove the source files and zip file
	err = os.RemoveAll(srcPath)
	if err != nil {
		return fmt.Errorf("failed to remove source files: %v", err)
	}

	err = os.Remove(zipPath)
	if err != nil {
		return fmt.Errorf("failed to remove zip file: %v", err)
	}
	utils.LavaFormatInfo("Source and zip files removed from directory.")

	return nil
}

func unzip(src string, dest string) ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {
		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}
