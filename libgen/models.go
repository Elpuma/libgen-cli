// Copyright © 2019 Antoine Chiny <antoine.chiny@inria.fr>
// Copyright © 2019 Ryan Ciehanski <ryan@ciehanski.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package libgen

import (
	"net/url"
	"time"
)

const (
	Version      = "v1.0.5"
	SearchHref   = "<a href='book/index.php.+</a>"
	SearchMD5    = "[A-Z0-9]{32}"
	booksdlReg   = "http://booksdl.org/get\\.php\\?md5=\\w{32}\\&key=\\w{16}"
	bokReg       = `\/dl\/\d{6}\/\w{6}`
	nineThreeReg = `\/main\/\d{1}\/[A-Za-z0-9]{32}\/.+?(gz|pdf|rar|djvu|epub|chm)`
	//libgenPwReg     = `http://libgen.pw/item/detail/id/\d*$`
	JSONQuery         = "id,title,author,filesize,extension,md5,year,language,pages,publisher,edition,coverurl"
	TitleMaxLength    = 68
	AuthorMaxLength   = 25
	HttpClientTimeout = time.Second * 10
)

// Book is the struct of resources on Library Genesis.
type Book struct {
	ID          string
	Title       string
	Author      string
	Filesize    string
	Extension   string
	Md5         string
	Year        string
	Language    string
	Pages       string
	Publisher   string
	Edition     string
	CoverURL    string
	DownloadURL string
	PageURL     string
}

// SearchMirrors contains all valid and tested mirrors used for
// querying against Library Genesis.
var SearchMirrors = []url.URL{
	{
		Scheme: "http",
		Host:   "gen.lib.rus.ec",
	},
	{
		Scheme: "https",
		Host:   "libgen.is",
	},
	{
		Scheme: "https",
		Host:   "libgen.unblockit.red",
	},
	{
		Scheme: "http",
		Host:   "libgen.unblockall.org",
	},
}

// DownloadMirrors contains all valid and tested mirrors used for
// downloading content from Library Genesis.
var DownloadMirrors = []url.URL{
	{
		Scheme: "http",
		Host:   "booksdl.org",
	},
	{
		Scheme: "https",
		Host:   "b-ok.cc",
	},
	{
		Scheme: "http",
		Host:   "93.174.95.29",
	},
}

// SearchOptions are the optional parameters available for the Search
// function.
type SearchOptions struct {
	Query         string
	SearchMirror  url.URL
	Results       int
	Print         bool
	RequireAuthor bool
	Extension     string
	Year          int
}

// GetDetailsOptions are the optional parameters available for the GetDetails
// function.
type GetDetailsOptions struct {
	Hashes        []string
	SearchMirror  url.URL
	Print         bool
	RequireAuthor bool
	Extension     string
	Year          int
}
