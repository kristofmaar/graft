package tensorflow

import "testing"

func TestLoadPluggableDevice(t *testing.T) {
	err := LoadPluggableDevice("libmetal_plugin.dylib")
	if err != nil {
		t.Fatal(err)
	}
}
