package gogobosh

func (repo BoshDirectorRepository) GetStemcells() (stemcells []Stemcell, apiResponse ApiResponse) {
	stemcellsResponse := []StemcellResponse{}

	path := "/stemcells"
	username := "admin"
	password := "admin"
	apiResponse = repo.gateway.GetResource(repo.config.TargetURL+path, username, password, &stemcellsResponse)
	if apiResponse.IsNotSuccessful() {
		return
	}

	for _, resource := range stemcellsResponse {
		stemcells = append(stemcells, resource.ToModel())
	}

	return
}

type StemcellResponse struct {
	Name string    `json:"name"`
	Version string `json:"version"`
	Cid string     `json:"cid"`
}

func (resource StemcellResponse) ToModel() (stemcell Stemcell) {
	stemcell = Stemcell{}
	stemcell.Name = resource.Name
	stemcell.Version = resource.Version
	stemcell.Cid = resource.Cid

	return
}