package cmd

import (

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/nyudlts/go-aspace"
	"log"
)

var (

	rootCmd = &cobra.Command{
		Use: "fa-random",
		Run: func(cmd *cobra.Command, args []string) {
			generateEADXML()
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	go_aspace.Seed()
	rootCmd.PersistentFlags().StringP("output-dir", "o", ".", "output directory for ead-xml files")
	rootCmd.PersistentFlags().IntP("count", "c", 0, "number of finding aids to generate")
	viper.BindPFlag("output-dir", rootCmd.PersistentFlags().Lookup("output-dir"))
	viper.BindPFlag("count", rootCmd.PersistentFlags().Lookup("count"))
}

func er(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func generateEADXML() {
	ASpaceClient, err := go_aspace.NewClient(20)
	er(err)

	repositories, err := ASpaceClient.GetRepositoryList()
	er(err)

	for i := 0; i < viper.GetInt("count"); i++ {
		j := go_aspace.RandInt(0, len(repositories))
		repositoryId := repositories[j]
		resources, err := ASpaceClient.GetResourceIDsByRepository(repositoryId)
		er(err)
		resourceId := resources[go_aspace.RandInt(0,len(resources))]
		err = ASpaceClient.SerializeEAD(repositoryId, resourceId, viper.GetString("output-dir"))
		er(err)
	}



}



