package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/log"
	"github.com/restream/reindexer"

	config "myTest/configs"
	"myTest/pkg"
	"myTest/pkg/repositories"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	serviceConfig, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(serviceConfig)

	reindexerDb := reindexer.NewReindex(
		fmt.Sprintf("cproto://%s:%s@%s/%s",
			serviceConfig.Database.Username,
			serviceConfig.Database.Password,
			serviceConfig.Database.Address,
			serviceConfig.Database.DBName,
		),
		reindexer.WithCreateDBIfMissing(),
	)
	repository, errDb := repositories.NewReindexerDocumentRepo(reindexerDb)
	if errDb != nil {
		fmt.Println(errDb)
		os.Exit(-1)
	}

	router := pkg.CreateRouter(repository)

	logger.Log("msg", "HTTP", "addr",
		fmt.Sprintf("%s:%s",
			serviceConfig.Server.Host,
			serviceConfig.Server.Port,
		),
	)
	logger.Log("err", http.ListenAndServe(
		fmt.Sprintf("%s:%s",
			serviceConfig.Server.Host,
			serviceConfig.Server.Port,
		), router),
	)
}
