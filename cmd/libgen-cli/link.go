package libgen_cli

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/spf13/cobra"

	"github.com/ciehanski/libgen-cli/libgen"
)

var linkCmd = &cobra.Command{
	Use:     "link",
	Short:   "Retrieves and displays the direct download link for a specific resource.",
	Long:    `Retrieves and displays the direct download link for a specific resource.`,
	Example: "libgen link 2F2DBA2A621B693BB95601C16ED680F8",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			if err := cmd.Help(); err != nil {
				fmt.Printf("error displaying CLI help: %v\n", err)
			}
			os.Exit(1)
		}
		// Ensure provided entry is valid MD5 hash
		re := regexp.MustCompile(libgen.SearchMD5)
		if !re.MatchString(args[0]) {
			fmt.Printf("\nPlease provide a valid MD5 hash\n")
			os.Exit(1)
		}

		fmt.Printf("++ Retrieving donwload link for: %s\n", args[0])

		bookDetails, err := libgen.GetDetails(&libgen.GetDetailsOptions{
			Hashes:       args,
			SearchMirror: libgen.GetWorkingMirror(libgen.SearchMirrors),
			Print:        false,
		})
		if err != nil {
			log.Fatalf("error retrieving results from LibGen API: %v", err)
		}
		book := bookDetails[0]

		if err := libgen.GetDownloadURL(book); err != nil {
			fmt.Printf("error getting download URL: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\n%v\n", book.DownloadURL)
	},
}