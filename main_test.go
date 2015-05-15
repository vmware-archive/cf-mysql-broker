package main_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Executable", func() {

	BeforeEach(func() {
		err := brokerCmd.Start()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		brokerCmd.Process.Kill()
		brokerCmd.Wait()
	})

	It("returns 200 on /v2/catalog", func() {
		req, _ := http.NewRequest("GET", "http://localhost:3000/v2/catalog", nil)
		req.SetBasicAuth("", "")

		var resp *http.Response
		Eventually(func() error {
			var err error
			resp, err = http.DefaultClient.Do(req)
			return err
		}).Should(Succeed())

		Expect(resp.StatusCode).To(Equal(http.StatusOK))
	})
})
