package host

import (
	"testing"

	"github.com/boot2podman/machine/drivers/fakedriver"
	_ "github.com/boot2podman/machine/drivers/none"
	"github.com/boot2podman/machine/libmachine/provision"
	"github.com/boot2podman/machine/libmachine/state"
)

func TestValidateHostnameValid(t *testing.T) {
	hosts := []string{
		"zomg",
		"test-ing",
		"some.h0st",
	}

	for _, v := range hosts {
		isValid := ValidateHostName(v)
		if !isValid {
			t.Fatalf("Thought a valid hostname was invalid: %s", v)
		}
	}
}

func TestValidateHostnameInvalid(t *testing.T) {
	hosts := []string{
		"zom_g",
		"test$ing",
		"some😄host",
	}

	for _, v := range hosts {
		isValid := ValidateHostName(v)
		if isValid {
			t.Fatalf("Thought an invalid hostname was valid: %s", v)
		}
	}
}

func TestStart(t *testing.T) {
	defer provision.SetDetector(&provision.StandardDetector{})
	provision.SetDetector(&provision.FakeDetector{
		Provisioner: provision.NewNetstatProvisioner(),
	})

	host := &Host{
		Driver: &fakedriver.Driver{
			MockState: state.Stopped,
		},
	}

	if err := host.Start(); err != nil {
		t.Fatalf("Expected no error but got one: %s", err)
	}
}
