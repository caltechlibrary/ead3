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
	"encoding/xml"
	"fmt"
	"runtime"
	"testing"
)

func errorOnNil(t *testing.T, val interface{}, msg string) {
	_, _, line, _ := runtime.Caller(1)
	if val == nil {
		t.Errorf(fmt.Sprintf("Failed test at line %d: %s\n", line, msg))
	}
}

func errorOnNotNil(t *testing.T, val interface{}, msg string) {
	_, _, line, _ := runtime.Caller(1)
	if val != nil {
		t.Errorf(fmt.Sprintf("Failed test at line %d: %s\n", line, msg))
	}
}

func TestEAD3TestXML(t *testing.T) {
	doc1 := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<ead xmlns="http://ead3.archivists.org/schema/undeprecated/">
	<control>
		<recordid>us-cchs-102</recordid>
		<filedesc>
			<titlestmt>
				<titleproper>Willam Fonds Provenance </titleproper>
				<subtitle>An Inventory of His Papers at the Cupcake Corners Historical Society</subtitle>
				<author>Finding aid created by Nancy Sahli.</author>
			</titlestmt>
			<publicationstmt>
				<p>Cupcake Corners Historical Society</p>
			</publicationstmt>
		</filedesc>
		<maintenancestatus value="new"/>
		<maintenanceagency>
			<agencyname>Cupcake Corners Historical Society</agencyname>
		</maintenanceagency>
		<languagedeclaration>
			<language langcode="eng"/>
			<script scriptcode="Latn"/>
		</languagedeclaration>
		<conventiondeclaration>
			<citation>DACS</citation>
		</conventiondeclaration>
		<maintenancehistory>
			<maintenanceevent>
				<eventtype value="created"/>
				<eventdatetime>September 11, 2013</eventdatetime>
				<agenttype value="human"/>
				<agent>Michael and Kris</agent>
			</maintenanceevent>
		</maintenancehistory>
	</control>
	<archdesc level="collection">
		<did>
			<head>Collection Summary</head>
			<repository label="Repository:" encodinganalog="855">
				<corpname source="LCNAF" encodinganalog="110">
					<part localtype="a">Cupcake Corners Historical Society</part>
					<part localtype="b">Center for the Study of the Giants of Archivy</part>
				</corpname>
			</repository>
			<origination label="Creator:" encodinganalog="100">
				<persname rules="AACR" encodinganalog="100">
					<part localtype="a">Provenance, William Fonds</part>
					<part localtype="d">1897-1956</part>
				</persname>
			</origination>
			<unittitle label="Title:" encodinganalog="245">Papers of William Fonds Provenance</unittitle>
			<unitdatestructured label="Dates:" unitdatetype="inclusive" encodinganalog="245">
				<daterange>
					<fromdate>1917</fromdate>
					<todate>1955 </todate>
				</daterange>
			</unitdatestructured>

			<!-- <unitdate label="Dates:">1917-1955</unitdate>-->

			<!-- <physdesc label="Quantity:">1.2 cubic ft. (3 boxes) </physdesc>-->

			<physdescset>
				<physdescstructured label="Quantity: " physdescstructuredtype="carrier" coverage="whole"
					encodinganalog="300">
					<quantity>3</quantity>
					<unittype>boxes</unittype>
				</physdescstructured>
				<physdescstructured label="Quantity: " physdescstructuredtype="spaceoccupied"
					coverage="whole" encodinganalog="300">
					<quantity>1.2</quantity>
					<unittype>cubic feet</unittype>
				</physdescstructured>
				<physdescstructured label="Quantity: " physdescstructuredtype="materialtype"
					coverage="whole" encodinganalog="300">
					<quantity>50</quantity>
					<unittype>diaries</unittype>
				</physdescstructured>
			</physdescset>

			<unitid label="Collection Number:">Mss 2 A</unitid>

			<abstract label="Abstract:">Correspondence, diaries,and writings of an archival theorist and author documenting his experiences in World War I, his literary endeavors, and his ideas on modern archival theory, especially regarding the centrality of the fonds.</abstract>
			<didnote label="Note:">Information about the collection.</didnote>
			<langmaterial label="Language:">
				<language langcode="eng">English</language>
				<descriptivenote>
					<p>Materials in this collection are written in English.</p>
				</descriptivenote>
			</langmaterial>
			<physloc label="Location:">This collection is on off-site storage. 24 hours notice is required for its retrieval.</physloc>
		</did>

		<bioghist encodinganalog="545">
			<head>Biography of William Provenance</head>
			<p>Archivist and author William Fonds Provenance was born at Last Chance, Nevada to Fred and Mary Jones Provenance on January 4, 1897. Little is know of his early life prior to serving in World War I as an ambulance driver. After graduating from Freen College in 1924 with a degree in cryptogamic biology, he first followed a career in commercial horticulture and later worked as an itinerant archivist. Provenance also had a lifelong interest in creative writing, producing both novels and poetry. He died at Frostbite Falls, Minnesota on March 15, 1956.</p>
		</bioghist>
		<scopecontent encodinganalog="520">
			<head>Scope and Contents of the Papers</head>
			<p>The collection consists of diaries, correspondence, manuscripts, and miscellaneous materials documenting the literary and archival career of William Fonds Provenance.</p>
			<p>The bulk of the collection consists of correspondence, principally with his mother, other archivists and writers. His diaries describe his experiences as an ambulance driver in France during World Was I in vivid detail. Major correspondents represented in the collection include Ernest Hemingway, Ernst Posner, and Provenance's long-time companion Ima Gusdorf. </p>
		</scopecontent>
		<arrangement encodinganalog="351">
			<head>Arrangement of the Papers</head>
			<p>The papers are arranged into two series:</p>
			<list>
				<item>
					<ref target="series1">Correspondence, 1919-1955.</ref>
				</item>
				<item>
					<ref target="series2">Diaries, 1917-1918</ref>
				</item>
			</list>
		</arrangement>

		<controlaccess>
			<head>Index Terms</head>
			<p>This collection is indexed under the following headings in the catalog of the Cupcake
				Corners Historical Society. Researchers desiring materials about related topics, persons, or
				places should search the catalog using these headings.</p>
			<controlaccess>
				<head>Persons</head>
				<persname source="AACR" encodinganalog="700">
					<part localtype="a">Gusdorf, Ima May</part>
				</persname>
				<persname source="LCNAF" encodinganalog="700">
					<part localtype="a">Hemingway, Ernest</part>
					<part localtype="d">1899-1961</part>
				</persname>
				<persname source="LCNAF" encodinganalog="700">
					<part localtype="a">Posner, Ernst</part>
				</persname>
				<persname encodinganalog="600">
					<part localtype="a">Ball, Joseph Hurst</part>
					<part localtype="d">1905- </part>
				</persname>
				<persname encodinganalog="600">
					<part>Bowles, Chester, 1901- </part>
				</persname>
				<persname encodinganalog="600">
					<part>McCarthy, Eugene Joseph, 1916- </part>
				</persname>
				<persname encodinganalog="600">
					<part>Shipstead, Henrik, 1881-1960</part>
				</persname>
				<persname encodinganalog="600">
					<part localtype="a">Youngdahl, Oscar Ferdinand</part>
					<part localtype="d">1893-1946</part>
				</persname>
				<persname rules="AACR" encodinganalog="600">
					<part localtype="a">Provenance, William F.</part>
					<part localtype="q">William Fonds</part>
					<part localtype="d">1897-1956</part>
				</persname>
			</controlaccess>
			<controlaccess>
				<head>Topics:</head>
				<subject source="LCSH" encodinganalog="650">
					<part>Cataloging of archival materials</part>
				</subject>
				<subject source="LCSH" encodinganalog="650">
					<part localtype="a">United States</part>
					<part localtype="x">History</part>
					<part localtype="y">World War, 1914-1918</part>
					<part localtype="x">Personal narratives, American</part>
				</subject>
				<subject encodinganalog="650">
					<part localtype="a">Cooperative marketing of farm produce</part>
					<part localtype="z">Minnesota</part>
				</subject>
			</controlaccess>

			<controlaccess>
				<head>Document Types</head>
				<genreform source="LCSH" encodinganalog="655">
					<part localtype="a">Novels</part>
				</genreform>
				<genreform source="LCSH" encodinganalog="655">
					<part localtype="a">Diaries</part>
				</genreform>
			</controlaccess>
		</controlaccess>
		<relatedmaterial encodinganalog="544">
			<head>Related Records</head>
			<p>The papers of Ima Gusdorf are located in the Freen University archives.</p>
			<p>The Historical Society has several other collections that include materials on archival
				theory.</p>
		</relatedmaterial>

		<accessrestrict encodinganalog="506">
			<head>Restrictions on Access</head>
			<p>Access to the correspondence between Provenance and Ernest Hemingway is restricted until
				2025.</p>
		</accessrestrict>

		<userestrict>
			<head>Rstrictions on Use</head>
			<p>Reproduction of any of the content of the collection beyond fair use requires the written
				permisssion of the Director of the Cucake Corners Historical Society.</p>
		</userestrict>

		<acqinfo encodinganalog="541">
			<head>Acquisition Information</head>
			<p>Acquired as a gift from Ima Gusdorf, December 17, 1952</p>
		</acqinfo>

		<processinfo encodinganalog="583">
			<head>Processing Information</head>
			<p>Collection processed and cataloged by B.W. Moos, January, 1962.</p>
		</processinfo>
		<altformavail>
			<head>Other Formats</head>
			<p>This collection is also available on microfilm.</p>
		</altformavail>
		<appraisal>
			<head>Appraisal</head>
			<p>Portions of this collection were heavily weeded.</p>
		</appraisal>
		<custodhist>
			<head>Custodial History</head>
			<p>This collection was in the possession of Ima Gusdorf until it was transfered to the
				archives</p>
		</custodhist>
		<fileplan>
			<head>File Plan</head>
			<p>An index to the collection was created by Ms. Gusdof. It is appended to the finding
				aid.</p>
		</fileplan>

		<accruals>
			<head>Additions to the Collection</head>
			<p>No additions are anticipated.</p>
		</accruals>
		<legalstatus>
			<head>Legal Status:</head>
			<p>No known laws restrict access to the use of this colleciton.</p>
		</legalstatus>
		<odd>
			<head>Other Information</head>
			<p>There is no further information about Provenance's teenage years which probably were as odd
				as you might imagine.</p>
		</odd>
		<originalsloc>
			<head>Location of Originals</head>
			<p>The self-published version of Provenance's work, <title render="italic">
					<part>The Tao of Archives,</part>
				</title> is in the Frostbite Falls Public Library. </p>
		</originalsloc>
		<prefercite>
			<head>Cite As</head>
			<p>The William Fonds Provenance Papers.</p>
		</prefercite>
		<otherfindaid>
			<head>Other Finding Aids</head>
			<p>Ms. Gusdorf's preliminary finding aid is located in the accession file.</p>
		</otherfindaid>
		<phystech>
			<head>Technical Issues</head>
			<p>Much of the collection is fragile and must be handled with gloves.</p>
		</phystech>

		<separatedmaterial>
			<head>Separated Materials</head>
			<p>Photographs have been transferred to the Iconography Department for storage.</p>
		</separatedmaterial>

		<dsc>
			<head>Detailed Description of the Collection</head>
			<p>The following section contains a detailed listing of the materials in the collection.</p>

			<c01 level="series" id="series1">
				<did>
					<unittitle>Correspondence</unittitle>
					<unitdatestructured label="Dates:" unitdatetype="inclusive" encodinganalog="245">
						<daterange>
							<fromdate>1919</fromdate>
							<todate>1955</todate>
						</daterange>
					</unitdatestructured>
					<physdesc>5 folders</physdesc>
				</did>
				<scopecontent>
					<p>Incoming letters and copies of outgoing correspondence with family, business
						associates, and prominent archivists and writers. Letters are arranged alphabetically by
						the writer's or recipient's name.</p>
				</scopecontent>
				<c02>
					<did>
						<container localtype="Box">1</container>
						<container localtype="Folder">1</container>
						<unittitle>A-F</unittitle>
					</did>
				</c02>
				<c02>
					<did>
						<container localtype="Box">1</container>
						<container localtype="Folder">2</container>
						<unittitle>Gusdorf, Ima</unittitle>
						<unitdate>1942-1955</unitdate>
					</did>
					<accessrestrict>
						<p>Access to some this file is restricted due to the personal nature of the
							contents.</p>
					</accessrestrict>
				</c02>
				<c02>
					<did>
						<container localtype="Box">2</container>
						<container localtype="Folder">1</container>
						<unittitle>H-P</unittitle>
					</did>
				</c02>
				<c02>
					<did>
						<container localtype="Box">2</container>
						<container localtype="Folder">2</container>
						<unittitle>Schellenberg, Theodore</unittitle>
						<abstract>Schellenberg describes his evolving understanding of historical
							manuscripts.</abstract>
					</did>

				</c02>
				<c02>
					<did>
						<container localtype="Box">3</container>
						<container localtype="Folder">1</container>
						<unittitle>T-Z</unittitle>
					</did>
				</c02>
			</c01>
			<c01 level="series" id="series2">
				<did>
					<unittitle>Diaries</unittitle>
					<unitdate>1917-1918</unitdate>
					<physdesc>32 v. (4 folders)</physdesc>
				</did>
				<scopecontent>
					<p>Daily accounts of Provenance's experiences during his military service in France during
						World War I, primarily documenting the daily activities of camp life, weather, battles,
						and operations of the army medical service. Also contains detailed and graphic accounts
						of his own work as an ambulance driver.</p>
				</scopecontent>
				<c02>
					<did>
						<container localtype="Box">3</container>
						<container localtype="Folder"/>
						<unitdate>1917</unitdate>
					</did>
					<c03>
						<did>
							<container localtype="Box">3</container>
							<container localtype="Folder">2</container>
							<unitdate>January-March</unitdate>
						</did>
						<phystech>
							<p>Some of the materials are highly fragile.</p>
						</phystech>
						<scopecontent>
							<p>Extensive information on the winter weather in France and the shortage of good wine
								and cigarettes. </p>
						</scopecontent>
						<userestrict>
							<p>Much of this material is very fragile. Researchers must use the microfilm
								copies.</p>
						</userestrict>
					</c03>

					<c03>
						<did>
							<container localtype="Box">3</container>
							<container localtype="Folder">3</container>
							<unitdate>April-December</unitdate>
							<didnote>April is the cruelest month.</didnote>
						</did>
						<c04>
							<did>
								<unittitle>More varied stuff</unittitle>
							<unitdatestructured>
								<datesingle>June 30</datesingle>
							</unitdatestructured>
								<abstract>Discussion of whether or not April showers bring May flowers.</abstract>
							</did>
						</c04>
					</c03>
				</c02>
				<c02>
					<did>
						<container localtype="Box">3</container>
						<container localtype="Folder"/>
						<unitdate> 1918</unitdate>
					</did>
					<c03>
						<did>
							<container localtype="Box">3</container>
							<container localtype="Folder">4</container>
							<unitdate>January</unitdate>
						</did>
						<scopecontent>
							<p>This file considers a wide range of topics including herbacious borders and German
								wines.</p>
						</scopecontent>
					</c03>
					<c03>
						<did>
							<container localtype="Box">3</container>
							<container localtype="Folder">5</container>
							<unitdate>February- June</unitdate>
						</did>
					</c03>
				</c02>
			</c01>
		</dsc>
	</archdesc>
</ead>`)

	item := new(EAD3)
	errorOnNil(t, item, "Empty EAD3 not created")
	err := xml.Unmarshal(doc1, &item)
	errorOnNotNil(t, err, fmt.Sprintf("%s", err))
	src, err := xml.MarshalIndent(item, "", "\t")
	errorOnNotNil(t, err, fmt.Sprintf("%s", err))
	fmt.Printf("DEBUG src: %s\n", src)
}
