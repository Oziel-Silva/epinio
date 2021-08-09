package acceptance_test

import (
	"strings"

	"github.com/epinio/epinio/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Epinio Installation", func() {
	It("has linkerd sidecars", func() {
		out, err := helpers.Kubectl("get", "pods",
			"--namespace", "epinio",
			"-l", "app.kubernetes.io/component=epinio-server",
			"-o", "jsonpath={.items[0].spec.containers[*].name}")
		Expect(err).ToNot(HaveOccurred())
		containers := strings.Split(out, " ")
		Expect(containers).To(ContainElement("linkerd-proxy"))
	})
	It("has linkerd control plane components running", func() {
		out, err := helpers.Kubectl("get", "pods", "--namespace", "linkerd", "-o", "name")
		Expect(err).ToNot(HaveOccurred())
		Expect(out).To(MatchRegexp("linkerd-identity"))
		Expect(out).To(MatchRegexp("linkerd-proxy-injector"))
		Expect(out).To(MatchRegexp("linkerd-controller"))
		Expect(out).To(MatchRegexp("linkerd-sp-validator"))
		Expect(out).To(MatchRegexp("linkerd-destination"))
	})
})
