package main

import (
	"github.com/cloudfoundry-community/gogobosh"
	"fmt"
)

func main() {
	gogobosh.Logger = gogobosh.NewLogger()

	director := gogobosh.NewDirector("https://192.168.50.4:25555", "admin", "admin")
	repo := gogobosh.NewBoshDirectorRepository(&director, gogobosh.NewDirectorGateway())

	info, apiResponse := repo.GetInfo()
	if apiResponse.IsNotSuccessful() {
		fmt.Println("Could not fetch BOSH info")
		return
	}

	fmt.Println("Director")
	fmt.Printf("  Name       %s\n", info.Name)
	fmt.Printf("  URL        %s\n", info.URL)
	fmt.Printf("  Version    %s\n", info.Version)
	fmt.Printf("  User       %s\n", info.User)
	fmt.Printf("  UUID       %s\n", info.UUID)
	fmt.Printf("  CPI        %s\n", info.CPI)
	if info.DNSEnabled {
		fmt.Printf("  dns        %#v (%s)\n", info.DNSEnabled, info.DNSDomainName)
	} else {
		fmt.Printf("  dns        %#v\n", info.DNSEnabled)
	}
	if info.CompiledPackageCacheEnabled {
		fmt.Printf("  compiled_package_cache %#v (provider: %s)\n", info.CompiledPackageCacheEnabled, info.CompiledPackageCacheProvider)
	} else {
		fmt.Printf("  compiled_package_cache %#v\n", info.CompiledPackageCacheEnabled)
	}
	fmt.Printf("  snapshots  %#v\n", info.SnapshotsEnabled)
	fmt.Println("")
	fmt.Printf("%#v\n", info)
	fmt.Println("")

	stemcells, apiResponse := repo.GetStemcells()
	if apiResponse.IsNotSuccessful() {
		fmt.Println("Could not fetch BOSH stemcells")
		return
	} else {
		for _, stemcell := range stemcells {
			fmt.Printf("%#v\n", stemcell)
		}
		fmt.Println("")
	}

	releases, apiResponse := repo.GetReleases()
	if apiResponse.IsNotSuccessful() {
		fmt.Println("Could not fetch BOSH releases")
		return
	} else {
		for _, release := range releases {
			fmt.Printf("%#v\n", release)
		}
		fmt.Println("")
	}

	deployments, apiResponse := repo.GetDeployments()
	if apiResponse.IsNotSuccessful() {
		fmt.Println("Could not fetch BOSH deployments")
		return
	} else {
		for _, deployment := range deployments {
			fmt.Printf("%#v\n", deployment)
		}
		fmt.Println("")
	}

	task, apiResponse := repo.GetTaskStatus(1)
	if apiResponse.IsNotSuccessful() {
		fmt.Println("Could not fetch BOSH task 1")
		return
	} else {
		fmt.Printf("%#v\n", task)
	}

	fmt.Println("")
	fmt.Println("VMs in cf-warden deployment:")
	vmsStatuses, apiResponse := repo.FetchVMsStatus("cf-warden")
	if apiResponse.IsNotSuccessful() {
		fmt.Println("Could not fetch VMs status for cf-warden")
		return
	} else {
		for _, vmStatus := range vmsStatuses {
			fmt.Printf("%s/%d is %s, IPs %#v\n", vmStatus.JobName, vmStatus.Index, vmStatus.JobState, vmStatus.IPs)
		}
	}
}
