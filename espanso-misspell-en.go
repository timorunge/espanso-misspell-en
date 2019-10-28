package main

import (
	"fmt"
	"os"

	"github.com/client9/misspell"
	"github.com/timorunge/espanso"
)

const (
	author  = "Timo Runge"
	repo    = "https://github.com/timorunge/espanso-misspell-en"
	version = "0.1.1"
)

func main() {
	genEn()
	genEnUK()
	genEnUS()
}

func genEn() {
	p := espanso.NewPackage()
	p.SetName("misspell-en")
	p.SetParent("default")
	p.SetMatches(espanso.DictToMatches(misspell.DictMain).SetWord(true))
	p.SetVersion(version)
	if err := p.Write(); err != nil {
		panic(err)
	}

	r := espanso.NewReadme()
	r.SetAuthor(author)
	r.SetLongDesc(genLongDesc(p.Name(), "commonly misspelled english words", "yuo"))
	r.SetName(p.Name())
	r.SetRepo(repo)
	r.SetShortDesc("Replace commonly misspelled english words.")
	r.SetTitle("Misspell EN")
	r.SetVersion(version)
	if err := r.Write(p.Name()); err != nil {
		panic(err)
	}

	if err := genLicenseFile(p.Name()); err != nil {
		panic(err)
	}
}

func genEnUK() {
	p := espanso.NewPackage()
	p.SetName("misspell-en_UK")
	p.SetParent("default")
	p.SetMatches(espanso.DictToMatches(misspell.DictBritish).SetWord(true))
	p.SetVersion(version)
	if err := p.Write(); err != nil {
		panic(err)
	}

	r := espanso.NewReadme()
	r.SetAuthor(author)
	r.SetLongDesc(genLongDesc(p.Name(), "american english with british english", "color"))
	r.SetName(p.Name())
	r.SetRepo(repo)
	r.SetShortDesc("Replace american english with british english.")
	r.SetTitle("Misspell en_UK")
	r.SetVersion(version)
	if err := r.Write(p.Name()); err != nil {
		panic(err)
	}

	if err := genLicenseFile(p.Name()); err != nil {
		panic(err)
	}
}

func genEnUS() {
	p := espanso.NewPackage()
	p.SetName("misspell-en_US")
	p.SetParent("default")
	p.SetMatches(espanso.DictToMatches(misspell.DictAmerican).SetWord(true))
	p.SetVersion(version)
	if err := p.Write(); err != nil {
		panic(err)
	}

	r := espanso.NewReadme()
	r.SetAuthor(author)
	r.SetLongDesc(genLongDesc(p.Name(), "british english with american english", "tyre"))
	r.SetName(p.Name())
	r.SetRepo(repo)
	r.SetShortDesc("Replace british english with american english.")
	r.SetTitle("Misspell en_US")
	r.SetVersion(version)
	if err := r.Write(p.Name()); err != nil {
		panic(err)
	}

	if err := genLicenseFile(p.Name()); err != nil {
		panic(err)
	}
}

func genLicenseFile(d string) error {
	l := `Copyright (c) 2019 Timo Runge <me@timorunge.com>
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors
   may be used to endorse or promote products derived from this software
   without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.`
	path := fmt.Sprintf("%s/LICENSE", d)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write([]byte(l)); err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	return nil
}

func genLongDesc(pkg string, teaser string, example string) string {
	s := fmt.Sprintf("# %s\n\n", pkg)
	s += fmt.Sprintf("%s is a espanso package which is replacing %s.\n", pkg, teaser)
	s += `The package is based on [github.com/client9/misspell](https://github.com/client9/misspell).

## Installation

Install the package with:`
	s += fmt.Sprintf("\n\n```\nespanso install %s\nespanso restart\n```", pkg)
	s += fmt.Sprintf("\n\n## Usage\n\nType `%s` and see what's happening.", example)
	s += fmt.Sprintf("\n\n## License\n\n[BSD 3-Clause \"New\" or \"Revised\" License](LICENSE)\n\nMisspell is [MIT](https://github.com/client9/misspell/blob/master/LICENSE).")
	return s
}
