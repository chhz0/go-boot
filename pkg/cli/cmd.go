package cli

import "github.com/spf13/cobra"

type Commander interface {
	Named() string
	Short() string

	Commands() []Commander
}

type CommandOptions struct {
	Usage string
	Short string
	Long  string

	PreRunE  func(cmd *Command, args []string) error
	RunE     func(cmd *Command, args []string) error
	PostRunE func(cmd *Command, args []string) error
	Init     func(cmd *Command)
}

func New(commander Commander) (*Exec, error) {

}

type Command struct {
	Cobra *cobra.Command
	Self  Commander

	Root   *Command
	Parent *Command
	Childs []*Command
}

func (c *Command) compile() error {

}
