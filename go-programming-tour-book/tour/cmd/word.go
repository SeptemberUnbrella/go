var wordCmd = &cobra.Command{
	Use: "word",
	Short: "单词转换"，
	Long: "支持多种单词格式的转换"，
	Run: func(cmd, *cobar.Command,args []string) {},
}
func init()  {
	rootCmd.AddCommand(wordCmd)

}
