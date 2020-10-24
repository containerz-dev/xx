// Copyright 2020 The containerz Authors.
// SPDX-License-Identifier: BSD-3-Clause

package xcoderelease

type Xcodereleases []*Xcoderelease

type Xcoderelease struct {
	Compilers Compilers `json:"compilers"`
	Requires  string    `json:"requires"`
	Date      Date      `json:"date"`
	Links     Links     `json:"links"`
	Version   Version   `json:"version"`
	SDKs      SDKs      `json:"sdks"`
	Name      string    `json:"name"`
	Checksums Checksums `json:"checksums"`
}

type Compilers struct {
	LLVM    []LLVM    `json:"llvm"`
	Clang   []Clang   `json:"clang"`
	Swift   []Swift   `json:"swift"`
	GCC     []GCC     `json:"gcc"`
	LLVMGCC []LLVMGCC `json:"llvm_gcc"`
}

type Clang struct {
	Number  string  `json:"number"`
	Build   string  `json:"build"`
	Release Release `json:"release"`
}

type GCC struct {
	Number  string  `json:"number"`
	Build   string  `json:"build"`
	Release Release `json:"release"`
}

type LLVM struct {
	Build   string  `json:"build"`
	Release Release `json:"release"`
}

type LLVMGCC struct {
	Number  string  `json:"number"`
	Build   string  `json:"build"`
	Release Release `json:"release"`
}

type Swift struct {
	Number  string  `json:"number"`
	Build   string  `json:"build"`
	Release Release `json:"release"`
}

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type Links struct {
	Notes    Notes    `json:"notes"`
	Download Download `json:"download"`
}

type Download struct {
	URL string `json:"url"`
}

type Notes struct {
	URL string `json:"url"`
}

type Version struct {
	Number  string  `json:"number"`
	Build   string  `json:"build"`
	Release Release `json:"release"`
}

type SDKs struct {
	MacOS   []MacOS   `json:"macOS"`
	TvOS    []TvOS    `json:"tvOS"`
	IOS     []IOS     `json:"iOS"`
	WatchOS []WatchOS `json:"watchOS"`
}

type MacOS struct {
	Number  string  `json:"number"`
	Build   string  `json:"build"`
	Release Release `json:"release"`
}

type IOS struct {
	Number  string  `json:"number"`
	Build   string  `json:"build"`
	Release Release `json:"release"`
}

type TvOS struct {
	Number  string  `json:"number"`
	Build   string  `json:"build"`
	Release Release `json:"release"`
}

type WatchOS struct {
	Number  string  `json:"number"`
	Build   string  `json:"build"`
	Release Release `json:"release"`
}

type Release struct {
	Beta   int  `json:"beta,omitempty"`
	DP     int  `json:"dp,omitempty"`
	GM     bool `json:"gm"`
	GMSeed int  `json:"gmSeed,omitempty"`
}

type Checksums struct {
	Sha1 string `json:"sha1"`
}
