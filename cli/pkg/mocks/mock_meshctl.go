package cli_mocks

import (
	"bytes"

	"github.com/golang/mock/gomock"
	"github.com/mattn/go-shellwords"
	"github.com/onsi/ginkgo"
	cli "github.com/solo-io/mesh-projects/cli/pkg"
	"github.com/solo-io/mesh-projects/cli/pkg/common"
	"github.com/spf13/cobra"
)

// Build and execute the CLI app using the given clients
type MockMeshctl struct {
	// MUST be non-nil - we always need to produce a mocked master cluster verification client
	MockController *gomock.Controller

	// safe to leave nil if not needed
	MasterVerification func(cmd *cobra.Command, args []string) (err error)

	Clients    *common.Clients
	KubeLoader common.KubeLoader
}

// call with the same string you would pass to the meshctl binary, i.e. "cluster register --service-account-name test123"
// returns the output produced by meshctl on stdout as a string
func (m MockMeshctl) Invoke(argString string) (stdout string, err error) {
	if m.MockController == nil {
		ginkgo.Fail("Must provide the ginkgo mock controller to mock meshctl")
	}

	buffer := &bytes.Buffer{}

	masterVerification := NewMockMasterKubeConfigVerifier(m.MockController)
	masterVerification.
		EXPECT().
		BuildVerificationCallback(gomock.Any(), gomock.Any()).
		Return(m.MasterVerification)

	app := cli.BuildCli(MockClientsFactory(m.Clients), buffer, masterVerification)

	splitArgs, err := shellwords.Parse(argString)
	if err != nil {
		panic("Bad arg string: " + argString)
	}

	app.SetArgs(splitArgs)

	err = app.Execute()

	return buffer.String(), err
}