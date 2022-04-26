package cmds

var execute bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&execute, "execute", "e", false, "execute the command.")
}
