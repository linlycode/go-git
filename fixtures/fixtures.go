package fixtures

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	check "gopkg.in/check.v1"
	"gopkg.in/src-d/go-git.v4/core"
)

var RootFolder = ""

const DataFolder = "data"

var fixtures = []*Fixture{{
	Tags:         []string{"ofs-delta"},
	URL:          "https://github.com/git-fixtures/basic",
	Head:         core.NewHash("6ecf0ef2c2dffb796033e5a02219af86ec6584e5"),
	PackfileHash: core.NewHash("a3fed42da1e8189a077c0e6846c040dcf73fc9dd"),
}, {
	Tags:         []string{"ref-delta"},
	URL:          "https://github.com/git-fixtures/basic",
	Head:         core.NewHash("6ecf0ef2c2dffb796033e5a02219af86ec6584e5"),
	PackfileHash: core.NewHash("c544593473465e6315ad4182d04d366c4592b829"),
}, {
	URL:          "https://github.com/spinnaker/spinnaker",
	Head:         core.NewHash("06ce06d0fc49646c4de733c45b7788aabad98a6f"),
	PackfileHash: core.NewHash("f2e0a8889a746f7600e07d2246a2e29a72f696be"),
}}

func All() Fixtures {
	return fixtures
}

func Basic() Fixtures {
	return ByURL("https://github.com/git-fixtures/basic")
}

func ByURL(url string) Fixtures {
	r := make(Fixtures, 0)
	for _, f := range fixtures {
		if f.URL == url {
			r = append(r, f)
		}
	}

	return r
}

type Fixture struct {
	URL          string
	Tags         []string
	Head         core.Hash
	PackfileHash core.Hash
}

func (f *Fixture) Packfile() io.ReadSeeker {
	fn := filepath.Join(RootFolder, DataFolder, fmt.Sprintf("pack-%s.pack", f.PackfileHash))
	file, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	return file
}

func (f *Fixture) Idx() io.ReadSeeker {
	fn := filepath.Join(RootFolder, DataFolder, fmt.Sprintf("pack-%s.idx", f.PackfileHash))
	file, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	return file
}

type Fixtures []*Fixture

func (g Fixtures) Test(c *check.C, test func(*Fixture)) {
	for _, f := range g {
		c.Logf("executing test at %s", f.URL)
		test(f)
	}
}

func (g Fixtures) One() *Fixture {
	return g[0]
}

func (g Fixtures) ByTag(tag string) *Fixture {
	for _, f := range g {
		for _, t := range f.Tags {
			if t == tag {
				return f
			}
		}
	}

	return nil
}
