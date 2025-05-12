package command

import (
	"fmt"
	"strings"

	"github.com/cgi-fr/posimap/internal/appli/charsets"
	"github.com/spf13/cobra"
)

type Charsets struct {
	cmd *cobra.Command
}

func NewCharsetsCommand(rootname string, groupid string) *cobra.Command {
	charsets := &Charsets{
		cmd: &cobra.Command{ //nolint:exhaustruct
			Use:     "charsets",
			Short:   "List all supported charsets",
			Long:    "List all supported charsets",
			Example: "  " + rootname + " charsets",
			Args:    cobra.NoArgs,
			GroupID: groupid,
		},
	}

	charsets.cmd.Run = charsets.execute

	return charsets.cmd
}

func (c *Charsets) execute(cmd *cobra.Command, _ []string) {
	const (
		nameWidth   = 20
		descWidth   = 40
		nameHeader  = "CHARSET NAME"
		descHeader  = "DESCRIPTION"
		headerUnder = "-"
	)

	// Format pour les données et les en-têtes
	format := fmt.Sprintf("%%-%ds%%-%ds\n", nameWidth, descWidth)

	// Affichage des en-têtes
	cmd.Printf(format, nameHeader, descHeader)
	cmd.Printf(format,
		strings.Repeat(headerUnder, nameWidth-1),
		strings.Repeat(headerUnder, descWidth-1))

	// Affichage des données
	for _, charset := range charsets.List() {
		c, err := charsets.Get(charset)
		if err == nil {
			cmd.Printf(format, charset, c.String())
		}
	}
}
