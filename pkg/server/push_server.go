package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/patnaikshekhar/kubernetescitool/cmd/cli/actions"
	"github.com/patnaikshekhar/kubernetescitool/pkg/git"
)

// StartPushServer starts the Push Server
func StartPushServer(port int) {
	http.HandleFunc("/", pushHandler)

	log.Printf("Starting push server")
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
	if err != nil {
		log.Fatalf("Push Server - Error starting push server %s", err.Error())
	}
}

func pushHandler(w http.ResponseWriter, r *http.Request) {
	var pushServerExpectedHostname = os.Getenv("PS_EXPECTED_HOSTNAME")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Push Server - Error reading request body %s", err.Error())
		return
	}

	params, err := url.ParseQuery(string(data))
	if err != nil {
		log.Printf("Push Server - Could not parse query string %s", err.Error())
		return
	}

	if payload, ok := params["payload"]; ok {
		var githubHookData git.GitHubPush

		err = json.Unmarshal([]byte(payload[0]), &githubHookData)

		// Get build.xml from the repository
		buildInfo, err := git.GetBuildDefinitionFromRepo(&githubHookData)
		if err != nil {
			log.Printf("Push Server - Error getting build definition %s", err.Error())
			return
		}

		log.Printf("Push Server - Got build definition %+v", buildInfo)

		// Start the build in the background
		log.Println("Starting Build")
		go func() {
			err = actions.StartBuild(buildInfo, pushServerExpectedHostname)
			if err != nil {
				log.Fatalf("Error in build: %s", err.Error())
			}
		}()

		w.Write([]byte("OK"))
	} else {
		log.Fatalf("Push Server - Could not find payload in query string %s", string(data))
	}
}
