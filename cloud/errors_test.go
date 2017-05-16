package cloud_test

import (
	"github.com/cloudfoundry/bosh-cli/cloud"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Error", func() {

	var (
		cmdError cloud.CmdError
		message  string
	)

	Context("when error type is 'Bosh::Clouds::NotSupported'", func() {

		JustBeforeEach(func() {
			cmdError = cloud.CmdError{
				"Bosh::Clouds::NotSupported",
				message,
				true,
			}
		})

		Context("when the message matches '/^Method .+ not supported in photon CPI/", func() {

			BeforeEach(func() {
				message = "Method 'not-implemeted-method' is not supported in photon CPI"
			})

			It("converts the error into a 'Bosh::Clouds::NotImplemented' error", func() {
				error := cloud.NewCPIError("not-implemented-method", cmdError)

				Expect(error.Type()).To(Equal(cloud.NotImplementedError))
				Expect(error.Message()).To(Equal("CPI error 'Bosh::Clouds::NotSupported' with message 'Method 'not-implemeted-method' is not supported in photon CPI' in 'not-implemented-method' CPI method"))
				Expect(error.OkToRetry()).To(Equal(true))
			})
		})

		Context("when the message does not match '/^Method .+ not supported in photon CPI/", func() {

			BeforeEach(func() {
				message = "some-message"
			})

			It("does not convert the error", func() {
				error := cloud.NewCPIError("not-implemented-method", cmdError)

				Expect(error.Type()).To(Equal(cmdError.Type))
				Expect(error.Message()).To(Equal(cmdError.Message))
				Expect(error.OkToRetry()).To(Equal(cmdError.OkToRetry))
			})
		})
	})

	Context("when error type is 'Bosh::Clouds::CloudError'", func() {

		JustBeforeEach(func() {
			cmdError = cloud.CmdError{
				"Bosh::Clouds::CloudError",
				message,
				false,
			}
		})

		Context("when the message matches '/^Invalid Method:/", func() {

			BeforeEach(func() {
				message = "Invalid Method: 'not-implemented-method'"
			})

			It("converts the error into a 'Bosh::Clouds::NotImplemented' error", func() {
				error := cloud.NewCPIError("not-implemented-method", cmdError)

				Expect(error.Type()).To(Equal(cloud.NotImplementedError))
				Expect(error.Message()).To(Equal("CPI error 'Bosh::Clouds::CloudError' with message 'Invalid Method: 'not-implemented-method'' in 'not-implemented-method' CPI method"))
				Expect(error.OkToRetry()).To(Equal(false))
			})
		})

		Context("when the message does not match '/Invalid Method:/", func() {

			BeforeEach(func() {
				message = "some-message"
			})

			It("does not convert the error", func() {
				error := cloud.NewCPIError("not-implemented-method", cmdError)

				Expect(error.Type()).To(Equal(cmdError.Type))
				Expect(error.Message()).To(Equal(cmdError.Message))
				Expect(error.OkToRetry()).To(Equal(cmdError.OkToRetry))
			})
		})
	})

	Context("when error type is 'InvalidCall'", func() {

		JustBeforeEach(func() {
			cmdError = cloud.CmdError{
				"InvalidCall",
				message,
				false,
			}
		})

		Context("when the message matches '/^Method is not known, got/", func() {

			BeforeEach(func() {
				message = "Method is not known, got 'not-implemented-method'"
			})

			It("converts the error into a 'Bosh::Clouds::NotImplemented' error", func() {
				error := cloud.NewCPIError("not-implemented-method", cmdError)

				Expect(error.Type()).To(Equal(cloud.NotImplementedError))
				Expect(error.Message()).To(Equal("CPI error 'InvalidCall' with message 'Method is not known, got 'not-implemented-method'' in 'not-implemented-method' CPI method"))
				Expect(error.OkToRetry()).To(Equal(false))
			})
		})

		Context("when the message does not match '/^Method is not known, got/", func() {

			BeforeEach(func() {
				message = "some-message"
			})

			It("does not convert the error", func() {
				error := cloud.NewCPIError("not-implemented-method", cmdError)

				Expect(error.Type()).To(Equal(cmdError.Type))
				Expect(error.Message()).To(Equal(cmdError.Message))
				Expect(error.OkToRetry()).To(Equal(cmdError.OkToRetry))
			})
		})
	})
})