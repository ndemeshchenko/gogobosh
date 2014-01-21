package gogobosh_test

import (
	gogobosh "github.com/cloudfoundry-community/gogobosh"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("get list of releases", func() {
	It("GET /releases to return []DirectorRelease{}", func() {
		request := gogobosh.NewDirectorTestRequest(gogobosh.TestRequest{
			Method: "GET",
			Path:   "/releases",
			Response: gogobosh.TestResponse{
				Status: http.StatusOK,
				Body: `[
				  {
				    "name": "cf",
				    "release_versions": [
				      {
				        "version": "153",
				        "commit_hash": "009fdd9a",
				        "uncommitted_changes": true,
				        "currently_deployed": true,
				        "job_names": [
				          "cloud_controller_ng",
				          "nats",
				          "dea_next",
				          "login",
				          "health_manager_next",
				          "uaa",
				          "debian_nfs_server",
				          "loggregator",
				          "postgres",
				          "dea_logging_agent",
				          "syslog_aggregator",
				          "narc",
				          "haproxy",
				          "hm9000",
				          "saml_login",
				          "nats_stream_forwarder",
				          "collector",
				          "pivotal_login",
				          "loggregator_trafficcontroller",
				          "etcd",
				          "gorouter"
				        ]
				      }
				    ]
				  }
				]`}})
		ts, handler, repo := createDirectorRepo(request)
		defer ts.Close()

		releases, apiResponse := repo.GetReleases()

		release := releases[0]
		Expect(release.Name).To(Equal("cf"))

		release_version := release.Versions[0]
		Expect(release_version.Version).To(Equal("153"))
		Expect(release_version.CommitHash).To(Equal("009fdd9a"))
		Expect(release_version.UncommittedChanges).To(Equal(true))
		Expect(release_version.CurrentlyDeployed).To(Equal(true))

		Expect(apiResponse.IsSuccessful()).To(Equal(true))
		Expect(handler.AllRequestsCalled()).To(Equal(true))
	})

	It("DeleteReleases(name)", func() {
		request := gogobosh.NewDirectorTestRequest(gogobosh.TestRequest{
			Method: "DELETE",
			Path:   "/releases/cf?force=true",
			Response: gogobosh.TestResponse{
				Status: http.StatusFound,
				Header: http.Header{
					"Location":{"https://some.host/tasks/25"},
				},
			}})
		ts, handler, repo := createDirectorRepo(
			request,
			taskTestRequest(25, "queued"),
			taskTestRequest(25, "processing"),
			taskTestRequest(25, "done"),
		)
		defer ts.Close()

		apiResponse := repo.DeleteReleases("cf")

		Expect(apiResponse.IsSuccessful()).To(Equal(true))
		Expect(handler.AllRequestsCalled()).To(Equal(true))
	})

	It("DeleteRelease(name, version)", func() {
		request := gogobosh.NewDirectorTestRequest(gogobosh.TestRequest{
			Method: "DELETE",
			Path:   "/releases/cf?force=true&version=144",
			Response: gogobosh.TestResponse{
				Status: http.StatusFound,
				Header: http.Header{
					"Location":{"https://some.host/tasks/26"},
				},
			}})
		ts, handler, repo := createDirectorRepo(
			request,
			taskTestRequest(26, "queued"),
			taskTestRequest(26, "processing"),
			taskTestRequest(26, "done"),
		)
		defer ts.Close()

		apiResponse := repo.DeleteRelease("cf", "144")

		Expect(apiResponse.IsSuccessful()).To(Equal(true))
		Expect(handler.AllRequestsCalled()).To(Equal(true))
	})
})
