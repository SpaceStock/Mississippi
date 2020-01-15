package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	restify "github.com/SpaceStock/go-restify"
	"github.com/SpaceStock/go-restify/scenario"

	"github.com/gobuffalo/packr"
)

func main() {
	domain := flag.String("domain", "", "Select a domain to be tested")
	flag.Parse()

	if domain == nil || *domain == "" {
		log.Panicln("please specify a domain e.g: apartment, housing, etc")
	}

	scenarios := map[string]restify.Scenario{}
	virtualbox := packr.NewBox("../pkg/" + *domain)
	err := virtualbox.Walk(func(path string, f packr.File) error {
		paths := strings.Split(path, "/")
		if len(paths) < 2 {
			//Skip non grouped files
			return nil
		}

		//Parse file into json bytes
		jsb, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}

		//Find scenario from cache, init if not exists
		scenarioName := paths[0]
		testScenario := scenarios[scenarioName]
		if testScenario == nil {
			testScenario = scenario.New().Set().
				Name(scenarioName).
				End()
		}

		//Parse scenario definition
		if strings.HasSuffix(path, "scenario.json") {
			temp := scenario.New()
			json.Unmarshal(jsb, &temp)
			if err != nil {
				return err
			}

			testScenario.Set().
				Name(temp.Get().Name()).
				Description(temp.Get().Description())
			return nil
		}

		//Parse individual test case
		testCase := restify.TestCase{}
		err = json.Unmarshal(jsb, &testCase)
		if err != nil {
			return err
		}

		//Add testcase to existing scenario
		testScenario.Set().AddCase(testCase)
		f.Close()

		scenarios[scenarioName] = testScenario
		return nil
	})
	fmt.Println("ERR", err)

	for _, scn := range scenarios {
		log.Println("SCENARIOS", len(scn.Get().Cases()))
		log.Println("RUNNING Scenario", scn.Get().Name())
		scn.Run(os.Stdout)
	}
}
