package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"time"
)

var collectionsUrl = "https://attack-taxii.mitre.org/api/v21/collections"

type collectionsResponse struct {
	Collections []struct {
		ID          string   `json:"id"`
		Title       string   `json:"title"`
		Description string   `json:"description"`
		CanRead     bool     `json:"can_read"`
		CanWrite    bool     `json:"can_write"`
		MediaTypes  []string `json:"media_types"`
	} `json:"collections"`
}

type objectsResponse struct {
	More    bool `json:"more"`
	Objects []struct {
		ID              string    `json:"id"`
		Modified        time.Time `json:"modified"`
		Created         time.Time `json:"created"`
		Type            string    `json:"type"`
		SpecVersion     string    `json:"spec_version"`
		Name            string    `json:"name"`
		Description     string    `json:"description"`
		KillChainPhases []struct {
			KillChainName string `json:"kill_chain_name"`
			PhaseName     string `json:"phase_name"`
		} `json:"kill_chain_phases"`
		XMitreAttackSpecVersion string   `json:"x_mitre_attack_spec_version"`
		XMitreDetection         string   `json:"x_mitre_detection"`
		XMitreDomains           []string `json:"x_mitre_domains"`
		XMitreIsSubtechnique    bool     `json:"x_mitre_is_subtechnique"`
		XMitreModifiedByRef     string   `json:"x_mitre_modified_by_ref"`
		XMitrePlatforms         []string `json:"x_mitre_platforms"`
		XMitreVersion           string   `json:"x_mitre_version"`
		XMitreDataSources       []string `json:"x_mitre_data_sources"`
		CreatedByRef            string   `json:"created_by_ref"`
		ExternalReferences      []struct {
			SourceName  string `json:"source_name"`
			URL         string `json:"url"`
			ExternalID  string `json:"external_id,omitempty"`
			Description string `json:"description,omitempty"`
		} `json:"external_references"`
		ObjectMarkingRefs []string `json:"object_marking_refs"`
	} `json:"objects"`
}

type Pick struct {
	Name        string
	Description string
}

func get_taxii_collections() *collectionsResponse {

	var collections collectionsResponse

	client := http.Client{}

	req, err := http.NewRequest("GET", collectionsUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/taxii+json;version=2.1")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&collections)
	if err != nil {
		log.Fatal(err)
	}

	return &collections
}

func get_taxii_objects(collectionID string) map[string]Pick {

	var objects objectsResponse

	client := http.Client{}

	req, err := http.NewRequest("GET", (collectionsUrl + "/" + collectionID + "/objects?match[type]=attack-pattern"), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/taxii+json;version=2.1")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&objects)
	if err != nil {
		log.Fatal(err)
	}

	picks := make(map[string]Pick)
	for {
		var pn string
		pick := rand.IntN(len(objects.Objects)-1) + 1
		//fmt.Printf("Random #: %d - %s\n", pick, objects.Objects[pick].Name)

		for _, phase := range objects.Objects[pick].KillChainPhases {
			//	fmt.Printf("Phase: %s\n", phase.PhaseName)
			pn = phase.PhaseName
			break
		}
		//fmt.Println("----------------------------------")

		_, present := picks[pn]
		if !present {
			picks[pn] = Pick{
				Name:        objects.Objects[pick].Name,
				Description: objects.Objects[pick].Description,
			}
		}

		if len(picks) == 3 {
			break
		}
	}

	return picks
}

func print_attck(picks map[string]Pick, infrastructure int) {

	fmt.Printf("\nHere's Your Clues: \n")
	for key, value := range picks {
		fmt.Printf("%s -> %s\n", key, value.Name)
	}
}

func main() {

	var infrastructure int

	fmt.Println("/ \\----./ /")
	fmt.Println("\\    e e \\")
	fmt.Println(" |       / |")
	fmt.Println(" |      ,__/  Build")
	fmt.Println("  \\_______/        -A-")
	fmt.Println("   /      \\          Threat")
	fmt.Println("  / |  \\   \\_")
	fmt.Println(" ;  \\  \\_  |_)")
	fmt.Println(" |    \\___) |  ")
	fmt.Println(" |            |")
	fmt.Println(" _;     ----/ |__/")
	fmt.Println(" \\_\\__________/")
	fmt.Printf("\n\n")

	results := get_taxii_collections()
	results_length := len(results.Collections)

	for {
		for i := 0; i < results_length; i++ {
			fmt.Printf("%d). %s\n", (i + 1), results.Collections[i].Title)
		}
		fmt.Printf("Pick Your Infrastructure: ")
		fmt.Scanf("%d", &infrastructure)

		if (infrastructure >= 1) && (infrastructure <= results_length) {
			picks := get_taxii_objects(results.Collections[infrastructure-1].ID)

			print_attck(picks, infrastructure)

			break
		}
	}

}
