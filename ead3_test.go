//
// Package ead3 provides limited wrapper around EAD version 3 XML objects making it
// easier to work with EAD documents in golang.
//
// @author: R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2016, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// * Neither the name of epgo nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package ead3

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strings"
	"testing"
)

var (
	testEAD3Files = []string{
		"samples/ead3/EAD3test.xml",
		"samples/ead3/NCSU/mc00003.xml",
		"samples/ead3/NCSU/mc00019.xml",
		"samples/ead3/NCSU/mc00022.xml",
		"samples/ead3/NCSU/mc00042.xml",
		"samples/ead3/NCSU/mc00062.xml",
		"samples/ead3/NCSU/mc00084.xml",
		"samples/ead3/NCSU/mc00092.xml",
		"samples/ead3/NCSU/mc00156.xml",
		"samples/ead3/NCSU/mc00185.xml",
		"samples/ead3/NCSU/mc00192.xml",
		"samples/ead3/NCSU/mc00212.xml",
		"samples/ead3/NCSU/mc00222.xml",
		"samples/ead3/NCSU/mc00240.xml",
		"samples/ead3/NCSU/mc00246.xml",
		"samples/ead3/NCSU/mc00261.xml",
		"samples/ead3/NCSU/mc00285.xml",
		"samples/ead3/NCSU/mc00312.xml",
		"samples/ead3/NCSU/mc00313.xml",
		"samples/ead3/NCSU/mc00325.xml",
		"samples/ead3/NCSU/mc00338.xml",
		"samples/ead3/NCSU/mc00348.xml",
		"samples/ead3/NCSU/mc00353.xml",
		"samples/ead3/NCSU/mc00384.xml",
		"samples/ead3/NCSU/mc00432.xml",
		"samples/ead3/NCSU/mc00462.xml",
		"samples/ead3/NCSU/mc00480.xml",
		"samples/ead3/NCSU/mc00492.xml",
		"samples/ead3/NCSU/mc00496.xml",
		"samples/ead3/NCSU/rbc00001.xml",
		"samples/ead3/NCSU/rbc00007.xml",
		"samples/ead3/NCSU/rbc00008.xml",
		"samples/ead3/NCSU/ua012_004.xml",
		"samples/ead3/NCSU/ua015_010.xml",
		"samples/ead3/NCSU/ua015_401.xml",
		"samples/ead3/NCSU/ua015_402.xml",
		"samples/ead3/NCSU/ua016_035.xml",
		"samples/ead3/S.0001_valid.xml",
		"samples/ead3/UMN/CLRC-2155.xml",
		"samples/ead3/UMN/mss060.xml",
		"samples/ead3/UMN/naa213.xml",
		"samples/ead3/UMN/sw0116-ead3.xml",
		"samples/ead3/UMN/uarc01180.xml",
		"samples/ead3/UMN/yusa0008-ead3.xml",
		"samples/ead3/UMN/yusa0009x2x16-ead3.xml",
		"samples/ead3/schematron_test_ead3.xml",
		"samples/ead3/untitled.xml",
	}
)

func expandElements(val []byte) []byte {
	var (
		results []string
		elName  string
	)
	cnt := bytes.Count(val, []byte("/>"))
	data := bytes.Split(val, []byte("/>"))
	for _, p := range data {
		if cnt > 0 {
			start := bytes.LastIndex(p, []byte("<")) + 1
			end := bytes.Index(p[start:], []byte(" ")) + start
			if end > start {
				elName = fmt.Sprintf("%s", p[start:end])
			} else {
				//log.Printf("DEBUG epandElements(%q) start: %d, end %d\n", p, start, end)
				elName = fmt.Sprintf("%s", p[1:])
			}
			results = append(results, fmt.Sprintf("%s></%s>", p, elName))
			cnt--
		} else {
			results = append(results, fmt.Sprintf("%s", p))
		}
	}
	return []byte(strings.Join(results, ""))
}

func targetPhrase(no int, line []byte) ([]byte, bool) {
	val := bytes.TrimSpace(line[:])
	if bytes.Contains(val, []byte("<?xml")) == true {
		return []byte(""), false
	}
	if bytes.HasPrefix(val, []byte("<!-- ")) == true {
		return []byte(""), false
	}
	if bytes.Count(val, []byte("/>")) > 0 {
		return expandElements(val), true
	}
	return bytes.TrimSpace(val), (bytes.Contains(val, []byte("<")) && bytes.Contains(val, []byte(">")))
}

func errorOnNil(t *testing.T, val interface{}, msg string) {
	_, _, line, _ := runtime.Caller(1)
	if val == nil {
		t.Errorf(fmt.Sprintf("error on nil at line %d: %s\n", line, msg))
	}
}

func errorOnNotNil(t *testing.T, val interface{}, msg string) {
	_, _, line, _ := runtime.Caller(1)
	if val != nil {
		t.Errorf(fmt.Sprintf("error on not nil at line %d: %s\n", line, msg))
	}
}

func anyMatch(src, val []byte) string {
	txt := fmt.Sprintf("%s", bytes.Replace(src, []byte("\n"), []byte(" "), -1))
	for s := fmt.Sprintf("%s", val); s != ""; {
		if strings.Contains(txt, s) == true {
			return s
		}
		s = s[:len(s)-1]
	}
	return ""
}

func contains(t *testing.T, src, val []byte, msg string) error {
	if bytes.Contains(src, val) == false {
		m := anyMatch(src, val)
		if strings.Compare(m, string(val)) == 0 {
			return nil
		}
		ioutil.WriteFile("debug.xml", src, 0664)
		return fmt.Errorf(`
expected:
%s
   found:
%s`, val, anyMatch(src, val))
	}
	return nil
}

func normalize(src []byte) []byte {
	// NOTE: cut comments out
	openComment := "<!-- "
	closeComment := " -->"
	buf := strings.Replace(string(src), "\n", " ", -1)
	for strings.Index(buf, openComment) >= 0 {
		first := ""
		rest := ""
		startCut := strings.Index(buf, openComment)
		if startCut >= 0 {
			first = fmt.Sprintf("%s", buf[:startCut])
		}
		endCut := strings.Index(buf, closeComment)
		if endCut >= 0 {
			rest = fmt.Sprintf("%s", buf[:endCut])
		}
		buf = fmt.Sprintf("%s%s", first, rest)
	}
	// NOTE: rebreak on end of element
	return []byte(strings.Replace(strings.Replace(buf, ">", ">\n", -1), "<", "\n<", -1))
}

func XMLDecodeEncodeTest(t *testing.T, fname string) {
	input, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Printf("Cannot read %s, %s\n", fname, err)
		fmt.Printf("Copy %s from https://github.com/SAA-SDT/EAD3/%s\n", fname, fname)
		t.FailNow()
	}

	record := New()
	errorOnNil(t, record, "Empty EAD3 not created")
	err = xml.Unmarshal(input, &record)
	errorOnNotNil(t, err, fmt.Sprintf("%s", err))
	output, err := xml.MarshalIndent(record, " ", "    ")
	errorOnNotNil(t, err, fmt.Sprintf("%s", err))

	expectedData := normalize(input)
	header := []byte(`<?xml version="1.0" encoding="UTF-8"?>`)
	doc := append(header[:], output[:]...)
	outputData := normalize(doc)
	log.Printf("Scanning %d lines in %s\n", bytes.Count(expectedData, []byte("\n")), fname)
	for no, line := range bytes.Split(expectedData, []byte("\n")) {
		if val, ok := targetPhrase(no, line[:]); ok == true {
			err = contains(t, doc, val, fmt.Sprintf("around line %d of `%s`", no, fname))
			if err != nil {
				out := bytes.Split(outputData, []byte("\n"))
				for i, j := (no - 10), (no + 25); i <= j && i < len(out); i++ {
					if i >= 0 {
						fmt.Printf("%d %s\n", i, out[i])
					}
				}
				log.Fatalf("%s", err)
			}
		}
	}
}

func TestStructureCompleteness(t *testing.T) {
	for _, fname := range testEAD3Files {
		XMLDecodeEncodeTest(t, fname)
	}
}
