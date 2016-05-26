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

import "encoding/xml"

//
// EAD version 3 structures
// Useful definition of the data structures can be found by looking up the related
// XML element in http://www2.archivists.org/sites/all/files/gammaEAD3TagLibrary.pdf
//

// EAD3 document container
type EAD3 struct {
	XMLName         xml.Name  `xml:"ead" json:"-"`
	XMLNameSpace    string    `xml:"xmlns,attr" json:"-"`
	Audience        string    `xml:"audience,attr,omitempty" json:"audience,omitempty"`
	Control         *Control  `xml:"control" json:"control"`
	ArchDesc        *ArchDesc `xml:"archdesc" json:"archdesc"`
	RelatedEncoding string    `xml:"relatedencoding,attr,omitempty" json:"relatedencoding,omitempty"`
	DateEncoding    string    `xml:"dateencoding,attr,omitempty" json:"dateencoding,omitempty"`
	LangEncoding    string    `xml:"langencoding,attr,omitempty" json:"langencoding,omitempty"`
}

// P provides for unparsed embedded markup of paragraph elements
type P struct {
	XMLName xml.Name `xml:"p" json:"-"`
	Value   string   `xml:",innerxml" json:"text"`
}

// Control structure for initial element of an EAD
type Control struct {
	XMLName         xml.Name `xml:"control" json:"-"`
	CountryEncoding string   `xml:"countryencoding,attr,omitempty" json:"countryencoding,omitempty"`
	DateEncoding    string   `xml:"dateencoding,attr,omitempty" json:"dateencoding,omitempty"`
	LangEncoding    string   `xml:"langencoding,attr,omitempty" json:"langencoding,omitempty"`
	RelatedEncoding string   `xml:"relatedencoding,attr,omitempty" json:"relatedencoding,omitempty"`
	ScriptEncoding  string   `xml:"scriptencoding,attr,omitempty" json:"scriptencoding,omitempty"`

	RecordID              *RecordID              `xml:"recordid" json:"recordid"`
	OtherRecordID         *OtherRecordID         `xml:"otherrecordid,omitemtpy" json:"otherrecordid,omitempty"`
	Representation        *Representation        `xml:"representation,omitempty" json:"representation,omitempty"`
	FileDesc              *FileDesc              `xml:"filedesc" json:"filedesc"`
	PublicationStatus     *PublicationStatus     `xml:"publicationstatus,omitempty" json:"publicationstatus"`
	MaintenanceStatus     *MaintenanceStatus     `xml:"maintenancestatus" json:"maintenancestatus"`
	MaintenanceAgency     *MaintenanceAgency     `xml:"maintenanceagency" json:"maintenanceagency"`
	LanguageDeclaration   *LanguageDeclaration   `xml:"languagedeclaration" json:"languagedeclaration"`
	ConventionDeclaration *ConventionDeclaration `xml:"conventiondeclaration" json:"conventiondeclaration"`
	LocalTypeDeclaration  *LocalTypeDeclaration  `xml:"localtypedeclaration,omitempty" json:"localtypedeclaration,omitempty"`
	LocalControl          []*LocalControl        `xml:"localcontrol,omitempty" json:"localcontrol,omitempty"`
	MaintenanceHistory    *MaintenanceHistory    `xml:"maintenancehistory" json:"maintenancehistory"`
	Sources               *Sources               `xml:"sources,omitempty" json:"sources,omitempty"`
}

type RecordID struct {
	XMLName     xml.Name `xml:"recordid" json:"-"`
	LocalType   string   `xml:"localtype,attr,omitempty" json:"localtype,omitempty"`
	InstanceURL string   `xml:"instanceurl,attr,omitempty" json:"instanceurl,omitempty"`
	Value       string   `xml:",chardata" json:"value"`
}

type OtherRecordID struct {
	XMLName     xml.Name `xml:"otherrecordid" json:"-"`
	LocalType   string   `xml:"localtype,attr,omitempty" json:"localtype,omitempty"`
	InstanceURL string   `xml:"instanceurl,attr,omitempty" json:"instanceurl,omitempty"`
	Value       string   `xml:",chardata" json:"value"`
}

type Representation struct {
	XMLName   xml.Name `xml:"representation" json:"-"`
	HRef      string   `xml:"href,attr,omitempty" json:"href,omitempty"`
	LocalType string   `xml:"localtype,attr,omitempty" json:"localtype,omitempty"`
	LinkTitle string   `xml:"linktitle,attr,omitempty" json:"linktitle,omitempty"`
	Show      string   `xml:"show,attr,omitempty" json:"show,omitempty"`
	Value     string   `xml:",chardata" json:"value"`
}

// FileDesc describes file system contents
type FileDesc struct {
	XMLName         xml.Name         `xml:"filedesc" json:"-"`
	TitleStmt       *TitleStmt       `xml:"titlestmt" json:"titlestmt"`
	EditionStmt     *EditionStmt     `xml:"editionstmt,omitempty" json:"editionstmt,omitempty"`
	PublicationStmt *PublicationStmt `xml:"publicationstmt,omitempty" json:"publicationstmt,omitempty"`
	NoteStmt        *NoteStmt        `xml:"notestmt,omitempty" json:"notestmt,omitempty"`
	SeriesStmt      *SeriesStmt      `xml:"seriesstmt,omitempty" json:"seriesstmt,omitempty"`
}

// TitleStmt provides structure relating to titling and authorship
type TitleStmt struct {
	XMLName     xml.Name     `xml:"titlestmt" json:"-"`
	TitleProper *TitleProper `xml:"titleproper" json:"titleproper"`
	Subtitle    *Subtitle    `xml:"subtitle,omitempty" json:"subtitle,omitempty"`
	Author      *Author      `xml:"author,omitempty" json:"author,omitempty"`
	Sponsor     *Sponsor     `xml:"sponsor,omitempty" json:"sponsor,omitempty"`
}

// TitleProper
type TitleProper struct {
	XMLName        xml.Name `xml:"titleproper" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",chardata" json:"value,omitempty"`
}

// Subtitle
type Subtitle struct {
	XMLName        xml.Name `xml:"subtitle" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",chardata" json:"value,omitempty"`
}

// Author
type Author struct {
	XMLName        xml.Name `xml:"author" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",chardata" json:"value,omitempty"`
}

// Sponsor
type Sponsor struct {
	XMLName        xml.Name `xml:"sponsor" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",chardata" json:"value,omitempty"`
}

// Publisher
type Publisher struct {
	XMLName        xml.Name `xml:"publisher" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",chardata" json:"value,omitempty"`
}

// EditionStmt provides information an about specific editions of work
type EditionStmt struct {
	XMLName xml.Name `xml:"editionstmt" json:"-"`
	Edition string   `xml:"edition,omitempty" json:"edition,omitempty"`
	P       []*P     `xml:"p" json:"p"`
	Date    *Date    `xml:"date,omitempty" json:"date,omitempty"`
}

// NoteStmt provides information an about a work
type NoteStmt struct {
	XMLName     xml.Name     `xml:"notestmt" json:"-"`
	ControlNote *ControlNote `xml:"controlnote,omitempty" json:"controlnote,omitempty"`
	Date        *Date        `xml:"date,omitempty" json:"date,omitempty"`
}

// ControlNote provides specific about processing
type ControlNote struct {
	XMLName xml.Name `xml:"controlnote" json:"-"`
	P       []*P     `xml:"p" json:"p"`
}

// PublicationStmt provides information on the publication nature of content
type PublicationStmt struct {
	XMLName   xml.Name   `xml:"publicationstmt" json:"-"`
	Publisher *Publisher `xml:"publisher,omitempty" json:"publisher,omitempty"`
	P         []*P       `xml:"p,omitempty" json:"p,omitempty"`
	Date      *Date      `xml:"date,omitempty" json:"date,omitempty"`
	Address   *Address   `xml:"address,omitempty" json:"address,omitempty"`
}

type Address struct {
	XMLName     xml.Name `xml:"address" json:"-"`
	AddressLine []string `xml:"addressline,omitempty" json:"addressline,omitempty"`
}

// SeriesStmt provides informaiton on a series
type SeriesStmt struct {
	XMLName     xml.Name `xml:"seriesstmt" json:"-"`
	TitleProper string   `xml:"titleproper" json:"titleproper"`
	Num         string   `xml:"num,omitempty" json:"num,omitempty"`
}

// PublicationStatus provides an descriptive publication status
type PublicationStatus struct {
	XMLName xml.Name `xml:"publicationstatus" json:"-"`
	Value   string   `xml:"value,attr" json:"value"`
	Text    string   `xml:",chardata" json:"text"`
}

// MaintenanceStatus provides an descriptive meantenance status
type MaintenanceStatus struct {
	XMLName xml.Name `xml:"maintenancestatus" json:"-"`
	Value   string   `xml:"value,attr" json:"value"`
	Text    string   `xml:",chardata" json:"text"`
}

// MaintenanceAgency provides content related organziation performing maintenance
type MaintenanceAgency struct {
	XMLName         xml.Name         `xml:"maintenanceagency" json:"-"`
	AgencyCode      string           `xml:"agencycode,omitempty" json:"agencycode,omitempty"`
	AgencyName      string           `xml:"agencyname,	omitempty" json:"agencyname,omitempty"`
	OtherAgencyCode *OtherAgencyCode `xml:"otheragencycode,omitempty" json:"otheragencycode,omitempty"`
}

type OtherAgencyCode struct {
	XMLName   xml.Name `xml:"otheragencycode" json:"-"`
	LocalType string   `xml:"localtype,attr,omitempty" json:"localtype,omitempty"`
	Value     string   `xml:",chardata" json:"value,omitempty"`
}

// ConventionDeclaration provides ciation declarations
type ConventionDeclaration struct {
	XMLName         xml.Name         `xml:"conventiondeclaration" json:"-"`
	Citation        *Citation        `xml:"citation" json:"citation"`
	Abbr            string           `xml:"abbr,omitempty" json:"abbr,omitempty"`
	Descriptivenote *DescriptiveNote `xml:"descriptivenote,omitempty" json:"descriptivenote,omitempty"`
}

// LocalTypeDeclaration provides ciation declarations
type LocalTypeDeclaration struct {
	XMLName         xml.Name         `xml:"localtypedeclaration" json:"-"`
	Citation        *Citation        `xml:"citation" json:"citation"`
	Abbr            string           `xml:"abbr,omitempty" json:"abbr,omitempty"`
	Descriptivenote *DescriptiveNote `xml:"descriptivenote,omitempty" json:"descriptivenote,omitempty"`
}

// LocalControl a control vocabulary
type LocalControl struct {
	XMLName   xml.Name `xml:"localcontrol" json:"-"`
	LocalType string   `xml:"localtype,attr,omitempty" json:"localtype,omitempty"`
	Term      string   `xml:"term,omitempty" json:"term,omitempty"`
}

// Citation provides citation description
type Citation struct {
	XMLName              xml.Name `xml:"citation" json:"-"`
	HRef                 string   `xml:"href,attr,omitempty" json:"href,omitempty"`
	LastDatetimeVerified string   `xml:"lastdatetimeverified,attr,omitempty" json:"lastdatetimeverified,omitempty"`
	LinkTitle            string   `xml:"linktitle,attr,omitempty" json:"linktitle,omitempty"`
	Actuate              string   `xml:"actuate,attr,omitempty" json:"actuate,omitempty"`
	Show                 string   `xml:"show,attr,omitempty" json:"show,omitempty"`
	//	LocalType            string   `xml:"localtype,attr,omitempty" json:"localtype,omitempty"`
	Value string `xml:",chardata" json:"value"`
}

// LanguageDeclaration describes relevant language implications
type LanguageDeclaration struct {
	XMLName  xml.Name  `xml:"languagedeclaration" json:"-"`
	Language *Language `xml:"language,omitempty" json:"language,omitempty"`
	Script   *Script   `xml:"script,omitempty" json:"script,omitempty"`
}

// Language describes a specific language
type Language struct {
	XMLName        xml.Name `xml:"language" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	LangCode       string   `xml:"langcode,attr,omitempty" json:"langcode,omitempty"`
	Value          string   `xml:",chardata" json:"value"`
}

// Script describes the character encoding and script representation used in content
type Script struct {
	XMLName    xml.Name `xml:"script" json:"-"`
	ScriptCode string   `xml:"scriptcode,attr" json:"scriptcode,omitempty"`
	Value      string   `xml:",chardata" json:"value,omitempty"`
}

// MaintenanceHistory provides a collection of maintenance events
type MaintenanceHistory struct {
	XMLName          xml.Name            `xml:"maintenancehistory" json:"-"`
	MaintenanceEvent []*MaintenanceEvent `xml:"maintenanceevent" json:"maintenanceevent"`
}

// MaintenanceEvent describes activities related to processing content
type MaintenanceEvent struct {
	XMLName          xml.Name       `xml:"maintenanceevent" json:"-"`
	EventType        *EventType     `xml:"eventtype" json:"eventtype"`
	EventDateTime    *EventDateTime `xml:"eventdatetime" json:"eventdatetime"`
	AgentType        *AgentType     `xml:"agenttype" json:"agenttype"`
	Agent            string         `xml:"agent" json:"agent"`
	EventDescription string         `xml:"eventdescription,omitempty" json:"eventdescription,omitempty"`
}

// EventDateTime describes a date time of event
type EventDateTime struct {
	XMLName          xml.Name `xml:"eventdatetime" json:"-"`
	StandardDateTime string   `xml:"standarddatetime,attr,omitempty" json:"standarddatetime,emitempty"`
	Value            string   `xml:",chardata" json:"value"`
}

// EventType describes the type of maintenance event
type EventType struct {
	XMLName xml.Name `xml:"eventtype" json:"-"`
	Value   string   `xml:"value,attr" json:"eventtype"`
}

// AgentType describes of the acting parties in the maintenance event
type AgentType struct {
	XMLName xml.Name `xml:"agenttype" json:"-"`
	Value   string   `xml:"value,attr" json:"agenttype"`
}

// ArchDesc provides an Archival Description of the content
type ArchDesc struct {
	XMLName           xml.Name           `xml:"archdesc" json:"-"`
	XMLNameSpace      string             `xml:"xmlns,attr,omitempty", json:"-"`
	LocalType         string             `xml:"localtype,attr,omitempty" json:"localtype,omitempty"`
	Level             string             `xml:"level,attr" json:"level,omitempty"`
	RelatedEncoding   string             `xml:"relatedencoding,attr,omitempty" json:"relatedencoding,omitempty"`
	DID               []*DID             `xml:"did,omitempty" json:"did,omitempty"`
	Bibliography      *Bibliography      `xml:"bibliography,omitempty" json:"bibliography,omitempty"`
	BiogHist          *BiogHist          `xml:"bioghist" json:"bioghist"`
	ScopeContent      *ScopeContent      `xml:"scopecontent" json:"scopecontent"`
	Arrangement       *Arrangement       `xml:"arrangement" json:"arrangement"`
	ControlAccess     []*ControlAccess   `xml:"controlaccess" json:"controlaccess"`
	RelatedMaterial   *RelatedMaterial   `xml:"relatedmaterial" json:"relatedmaterial"`
	Relations         *Relations         `xml:"relations,omitempty" json:"relation,omitempty"`
	AccessRestrict    *AccessRestrict    `xml:"accessrestrict" json:"accessrestrict"`
	UseRestrict       *UseRestrict       `xml:"userestrict" json:"userestrict"`
	AcqInfo           *AcqInfo           `xml:"acqinfo" json:"acqinfo"`
	ProcessInfo       *ProcessInfo       `xml:"processinfo" json:"processinfo"`
	AltFormAvail      *AltFormAvail      `xml:"altformavail,omitempty" json:"altformavail,omitempty"`
	Appraisal         *Appraisal         `xml:"appraisal" json:"appraisal"`
	CustodHist        *CustodHist        `xml:"custodhist" json:"custodhist"`
	FilePlan          *FilePlan          `xml:"fileplan" json:"fileplan"`
	Accruals          *Accruals          `xml:"accruals" json:"accruals"`
	LegalStatus       *LegalStatus       `xml:"legalstatus" json:"legalstatus"`
	Odd               *Odd               `xml:"odd" json:"odd"`
	OriginalsLoc      *OriginalsLoc      `xml:"originalsloc,omitempty" json:"originalsloc,omitempty"`
	PreferCite        *PreferCite        `xml:"prefercite" json:"prefercite"`
	OtherFindAID      *OtherFindAID      `xml:"otherfindaid" json:"otherfindaid"`
	PhysTech          *PhysTech          `xml:"phystech" json:"phystech"`
	SeparatedMaterial *SeparatedMaterial `xml:"separatedmaterial" json:"separatedmaterial"`
	Dsc               *Dsc               `xml:"dsc" json:"dsc"`
}

// DID  Descriptive Identification of the Unit
type DID struct {
	XMLName            xml.Name              `xml:"did" json:"-"`
	XMLNameSpace       string                `xml:"xmlns,attr,omitempty", json:"-"`
	ID                 string                `xml:"id,attr,omitempty" json:"id,omitempty"`
	Lang               string                `xml:"lang,attr,omitempty" json:"lang,omitempty"`
	Script             string                `xml:"script,attr,omitempty" json:"script,omitempty"`
	Head               *Head                 `xml:"head,omitempty" json:"head,omitempty"`
	Repository         *Repository           `xml:"repository" json:"repository"`
	Origination        *Origination          `xml:"origination" json:"origination"`
	UnitTitle          *UnitTitle            `xml:"unittitle" json:"unittitle"`
	UnitDateStructured []*UnitDateStructured `xml:"unitdatestructured,omitempty" json:"unitdatestructured,omitempty"`
	UnitDate           []*UnitDate           `xml:"unitdate,omitempty" json:"unitdates,omitempty"`
	PhysDesc           *PhysDesc             `xml:"physdesc,omitempty" json:"physdesc,omitempty"`
	PhysDescSet        *PhysDescSet          `xml:"physdescset,omitempty" json:"physdescset,omitempty"`
	PhysDescStructured []*PhysDescStructured `xml:"physdescstructured" json:"physdescstructured"`
	UnitID             *UnitID               `xml:"unitid,omitempty" json:"unitid,omitempty"`
	Abstract           *Abstract             `xml:"abstract,omitempty" json:"abstract,omitempty"`
	DIDNote            *DIDNote              `xml:"didnote,omitempty" json:"didnote,omitempty"`
	MaterialSpec       *MaterialSpec         `xml:"materialspec,omitempty" json:"materialspec,omitempty"`
	LangMaterial       *LangMaterial         `xml:"langmaterial,omitempty" json:"langmaterial,omitempty"`
	PhysLoc            *PhysLoc              `xml:"physloc,omitempty" json:"physloc,omitempty"`
	Container          []*Container          `xml:"container,omitempty" json:"container,omitempty"`
	DAO                []*DAO                `xml:"dao,omitempty" json:"dao,omitempty"`
}

type Head struct {
	XMLName   xml.Name `xml:"head" json:"-"`
	AltRender string   `xml:"altrender,attr,omitempty" json:"altrender,omitempty"`
	ID        string   `xml:"id,attr,omitempty" json:"id,omitempty"`
	Value     string   `xml:",innerxml" json:"value,omitempty"`
}

type MaterialSpec struct {
	XMLName xml.Name `xml:"materialspec" json:"-"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
}

// DAO digital archival object
type DAO struct {
	XMLName         xml.Name         `xml:"dao" json:"-"`
	HRef            string           `xml:"href,attr,omitempty" json:"href,omitempty"`
	DOAType         string           `xml:"daotype,attr,omitempty" json:"daotype,omitempty"`
	Show            string           `xml:"show,attr,omitempty" json:"show,omitempty"`
	Actuate         string           `xml:"actuate,attr,omitempty" json:"actuate,omitempty"`
	DescriptiveNote *DescriptiveNote `xml:"descriptivenote,omitempty" json:"descriptivenote,omitempty"`
}

// Container describes the container holding materials
type Container struct {
	XMLName     xml.Name `xml:"container" json:"-"`
	ContainerID string   `xml:"containerid,attr,omitempty" json:"containerid,omitempty"`
	LocalType   string   `xml:"localtype,attr,omitempty" json:"localtype,omitempty"`
	Value       string   `xml:",chardata" json:"value,omitempty"`
}

// Repository describes the repository holding the materials
type Repository struct {
	XMLName        xml.Name    `xml:"repository" json:"-"`
	Label          string      `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog string      `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Persname       []*Persname `xml:"persname,omitempty" json:"persname,omitempty"`
	Famname        []*Famname  `xml:"famname,omitempty" json:"famname,omitempty"`
	CorpName       []*CorpName `xml:"corpname,omitempty" json:"corpname,omitempty"`
}

// CorpName - Corpus name
type CorpName struct {
	XMLName        xml.Name `xml:"corpname" json:"-"`
	Source         string   `xml:"source,attr,omitempty" json:"source,omitempty"`
	Rules          string   `xml:"rules,attr,omitempty" json:"rules,omitempty"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Normal         string   `xml:"normal,attr,omitempty" json:"normal,omitempty"`
	Part           []*Part  `xml:"part" json:"part"`
}

// Part describes the part of a corpus or other organization
type Part struct {
	XMLName   xml.Name `xml:"part" json:"-"`
	LocalType string   `xml:"localtype,attr,omitempty" json:"localtype,omitempty"`
	Value     string   `xml:",chardata" json:"value"`
}

// Origination describes where things came from
type Origination struct {
	XMLName        xml.Name    `xml:"origination" json:"-"`
	Label          string      `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog string      `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Persname       []*Persname `xml:"persname,omitempty" json:"persname,omitempty"`
	Famname        []*Famname  `xml:"famname,omitempty" json:"famname,omitempty"`
	CorpName       []*CorpName `xml:"corpname,omitempty" json:"corpname,omitempty"`
}

// Persname is a person's name(s)
type Persname struct {
	XMLName        xml.Name `xml:"persname" json:"-"`
	Source         string   `xml:"source,attr,omitempty" json:"source,omitempty"`
	Relator        string   `xml:"relator,attr,omitempty" json:"relator,omitempty"`
	Rules          string   `xml:"rules,attr,omitempty" json:"rules,omitempty"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Part           []*Part  `xml:"part" json:"part"`
}

// Famname is a person's family name(s)
type Famname struct {
	XMLName        xml.Name `xml:"famname" json:"-"`
	Source         string   `xml:"source,attr,omitempty" json:"source,omitempty"`
	Relator        string   `xml:"relator,attr,omitempty" json:"relator,omitempty"`
	Rules          string   `xml:"rules,attr,omitempty" json:"rules,omitempty"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Normal         string   `xml:"normal,attr,omitempty" json:"normal,omitempty"`
	Part           []*Part  `xml:"part" json:"part"`
}

// Subject describes the subject of contents
type Subject struct {
	XMLName        xml.Name `xml:"subject" json:"-"`
	Source         string   `xml:"source,attr,omitempty" json:"source,omitempty"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Normal         string   `xml:"normal,attr,omitempty" json:"normal,omitempty"`
	Rules          string   `xml:"rules,attr,omitempty" json:"rules,omitempty"`
	Part           []*Part  `xml:"part" json:"part"`
}

// GenreForm describe the genre of the contents
type GenreForm struct {
	XMLName        xml.Name `xml:"genreform" json:"-"`
	Source         string   `xml:"source,attr,omitempty" json:"source"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Normal         string   `xml:"normal,attr,omitempty" json:"normal,omitempty"`
	Rules          string   `xml:"rules,attr,omitempty" json:"rules,omitempty"`
	Part           []*Part  `xml:"part" json:"part"`
}

// GeogName geographical name
type GeogName struct {
	XMLName        xml.Name `xml:"geogname" json:"-"`
	Source         string   `xml:"source,attr,omitempty" json:"source"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Normal         string   `xml:"normal,attr,omitempty" json:"normal,omitempty"`
	Rules          string   `xml:"rules,attr,omitempty" json:"rules,omitempty"`
	Part           []*Part  `xml:"part" json:"part"`
}

// Occupation
type Occupation struct {
	XMLName        xml.Name `xml:"occupation" json:"-"`
	Source         string   `xml:"source,attr,omitempty" json:"source"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Normal         string   `xml:"normal,attr,omitempty" json:"normal,omitempty"`
	Rules          string   `xml:"rules,attr,omitempty" json:"rules,omitempty"`
	Part           []*Part  `xml:"part" json:"part"`
}

// UnitTitle is a title for a specific unit of collect contents
type UnitTitle struct {
	XMLName        xml.Name `xml:"unittitle" json:"-"`
	Label          string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",innerxml" json:"value"`
}

// UnitDateStructured provides a descriptive structure of date information related to the unit of content
type UnitDateStructured struct {
	XMLName        xml.Name     `xml:"unitdatestructured" json"-"`
	XMLNameSpace   string       `xml:"xmlns,attr,omitempty", json:"-"`
	Label          string       `xml:"label,attr,omitempty" json:"label,omitempty"`
	Era            string       `xml:"era,attr,omitempty" json:"era,omitempty"`
	Certainty      string       `xml:"certainty,attr,omitempty" json:"certainty,omitempty"`
	UnitDateType   string       `xml:"unitdatetype,attr,omitempty" json:"unitdatetype,omitempty"`
	EncodingAnalog string       `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	DateSingle     *DateSingle  `xml:"datesingle,omitempty" json:"datesingle,omitempty"`
	DateRange      []*DateRange `xml:"daterange,omitempty" json:"daterange,omitempty"`
}

// UnitDate provides a simpler date structure relating to content
type UnitDate struct {
	XMLName        xml.Name `xml:"unitdate" json:"-"`
	XMLNameSpace   string   `xml:"xmlns,attr,omitempty", json:"-"`
	Normal         string   `xml:"normal,attr,omitempty" json:"normal,omitempty"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Lang           string   `xml:"lang,attr,omitempty" json:"lang,omitempty"`
	Script         string   `xml:"script,attr,omitempty" json:"script,omitempty"`
	Certainty      string   `xml:"certainty,attr,omitempty" json:"certainty,omitempty"`
	UnitDateType   string   `xml:"unitdatetype,attr,omitempty" json:"unitdatetype"`
	Label          string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	Value          string   `xml:",chardata" json:"value"`
}

// DateRange describes the start and end dates
type DateRange struct {
	XMLName      xml.Name  `xml:"daterange" json"-"`
	XMLNameSpace string    `xml:"xmlns,attr,omitempty", json:"-"`
	FromDate     *FromDate `xml:"fromdate" json:"fromdate"`
	ToDate       *ToDate   `xml:"todate" json:"todate"`
}

// FromDate holds the start position in a range of dates
type FromDate struct {
	XMLName      xml.Name `xml:"fromdate" json:"-"`
	XMLNameSpace string   `xml:"xmlns,attr,omitempty", json:"-"`
	NotBefore    string   `xml:"notbefore,attr,omitempty" json:"notbefore"`
	NotAfter     string   `xml:"notafter,attr,omitempty" json:"notafter"`
	StandardDate string   `xml:"standarddate,attr,omitempty" json:"standarddate,omitempty"`
	Value        string   `xml:",chardata" json:"value,omitempty"`
}

// ToDate holds the end position in a range of dates
type ToDate struct {
	XMLName      xml.Name `xml:"todate" json:"-"`
	XMLNameSpace string   `xml:"xmlns,attr,omitempty", json:"-"`
	NotBefore    string   `xml:"notbefore,attr,omitempty" json:"notbefore"`
	NotAfter     string   `xml:"notafter,attr,omitempty" json:"notafter"`
	StandardDate string   `xml:"standarddate,attr,omitempty" json:"standarddate,omitempty"`
	Value        string   `xml:",chardata" json:"value,omitempty"`
}

type DateSingle struct {
	XMLName      xml.Name `xml:"datesingle" json:"-"`
	StandardDate string   `xml:"standarddate,attr,omitempty" json:"standarddate,omitempty"`
	Normal       string   `xml:"normal,attr,omitempty" json:"normal,omitempty"`
	Value        string   `xml:",chardata" json:"value,omitempty"`
}

type Date struct {
	XMLName        xml.Name `xml:"date" json:"-"`
	Normal         string   `xml:"normal,attr,omitempty" json:"normal,omitempty"`
	StandardDate   string   `xml:"standarddate,attr,omitempty" json:"standarddate,omitempty"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",chardata" json:"value,omitempty"`
}

// PhysDesc contains the physical description of contents
type PhysDesc struct {
	XMLName        xml.Name `xml:"physdesc" json:"-"`
	Label          string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",chardata" json:"value"`
}

// PhysDescSet describes a grouping of physical descriptions of content
type PhysDescSet struct {
	XMLName            xml.Name              `xml:"physdescset" json:"-"`
	PhysDescStructured []*PhysDescStructured `xml:"physdescstructured" json:"physdescstructured"`
}

// PhysDescStructured provides a structured set of physical descriptions
type PhysDescStructured struct {
	XMLName                xml.Name    `xml:"physdescstructured" json:"-"`
	Label                  string      `xml:"label,attr,omitempty" json:"label,omitempty"`
	PhysDescStructuredType string      `xml:"physdescstructuredtype,attr,omitempty" json:"physdescstructuredtype,omitempty"`
	Coverage               string      `xml:"coverage,attr,omitempty" json:"coverage,omitempty"`
	EncodingAnalog         string      `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Quantity               *Quantity   `xml:"quantity,omitempty" json:"quantity,omitempty"`
	UnitType               *UnitType   `xml:"unittype,omitempty" json:"unittype,omitempty"`
	PhysFacet              *PhysFacet  `xml:"physfacet,omitempty" json:"physfacet,omitempty"`
	Dimensions             *Dimensions `xml:"dimensions,omitempty" json:"dimensions,omitempty"`
}

// Dimensions object measurements
type Dimensions struct {
	XMLName   xml.Name `xml:"dimensions" json:"-"`
	LocalType string   `xml:"localtype,attr,omitempty" json:"localtype,omitempty"`
	Unit      string   `xml:"unit,attr,omitempty" json:"unit,omitempty"`
	Value     string   `xml:",chardata" json:"value,omitempty"`
}

// Quantity measurement in units
type Quantity struct {
	XMLName     xml.Name `xml:"quantity" json:"-"`
	Approximate string   `xml:"approximate,attr,omitempty" json:"approximate,omitempty"`
	Value       string   `xml:",chardata" json:"value,omitempty"`
}

// PhysFacet
type PhysFacet struct {
	XMLName   xml.Name `xml:"physfacet" json:"-"`
	LocalType string   `xml:"localtype,attr,omitempty" json:"localtype,omitempty"`
	Value     string   `xml:",chardata" json:"value,omitempty"`
}

// UnitType
type UnitType struct {
	XMLName xml.Name `xml:"unittype" json:"-"`
	Source  string   `xml:"source,attr,omitempty" json:"source,omitempty"`
	Value   string   `xml:",chardata" json:"value,omitempty"`
}

// UnitID provides a unit level identifier
type UnitID struct {
	XMLName        xml.Name `xml:"unitid" json:"-"`
	Label          string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	CountryCode    string   `xml:"countrycode,attr,omitempty" json:"countrycode,omitempty"`
	RepositoryCode string   `xml:"repositorycode,attr,omitempty" json:"repositorycode,omitempty"`
	Value          string   `xml:",chardata" json:"value"`
}

// Abstract provides a summary of content
type Abstract struct {
	XMLName        xml.Name `xml:"abstract" json:"-"`
	Label          string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",chardata" json:"value"`
}

// DIDNote are notes on the describe digital identifier
type DIDNote struct {
	XMLName xml.Name `xml:"didnote" json:"-"`
	Label   string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	Value   string   `xml:",chardata" json:"value"`
}

// LangMaterial describes the material's language characteristics which could differ from document relating to content
type LangMaterial struct {
	XMLName         xml.Name         `xml:"langmaterial" json:"-"`
	Label           string           `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog  string           `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Language        *Language        `xml:"language,omitempty" json:"language,omitempty"`
	DescriptiveNote *DescriptiveNote `xml:"descriptivenote,omitempty" json:"descriptivenote,omitempty"`
}

// DescriptiveNote is a descriptive note about the content
type DescriptiveNote struct {
	XMLName xml.Name `xml:"descriptivenote" json:"-"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
	Title   string   `xml:"title,omitempty" json:"title,omitempty"`
}

// PhysLoc describes the physical location of the content
type PhysLoc struct {
	XMLName xml.Name `xml:"physloc" json:"-"`
	Label   string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	Value   string   `xml:",chardata" json:"value"`
}

// Bibliography information
type Bibliography struct {
	XMLName xml.Name  `xml:"bibliography" json:"-"`
	Head    *Head     `xml:"head,omitempty" json:"head"`
	BibRef  []*BibRef `xml:"bibref,omitempty" json:"bibref,omitempty"`
}

// BibRef a specific reference
type BibRef struct {
	XMLName xml.Name `xml:"bibref" json:"-"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
	Ref     *Ref     `xml:"ref,omitempty" json:"ref,omitempty"`
}

// BiogHist provides biographical history of content
type BiogHist struct {
	XMLName        xml.Name   `xml:"bioghist" json:"-"`
	ID             string     `xml:"id,attr,omitempty" json:"id,omitempty"`
	EncodingAnalog string     `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Head           *Head      `xml:"head,omitempty" json:"head,omitempty"`
	P              []*P       `xml:"p,omitempty" json:"p,omitempty"`
	ChronList      *ChronList `xml:"chronlist,omitempty" json:"chronlist,omitempty"`
	T              *Table     `xml:"table,omitempty" json:"table,omitempty"`
}

type ChronList struct {
	XMLName   xml.Name     `xml:"chronlist" json:"-"`
	ChronItem []*ChronItem `xml:"chronitem,omitempty" json:"chronitem,omitempty"`
}

type ChronItem struct {
	XMLName    xml.Name    `xml:"chronitem" json:"-"`
	DateSingle *DateSingle `xml:"datesingle,omitempty" json:"datesingle,omitempty"`
	Event      *Event      `xml:"event,omitempty" json:"event,omitempty"`
}

type Event struct {
	XMLName xml.Name `xml:"event" json:"-"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
}

// ScopeContent provides scoping material for content
type ScopeContent struct {
	XMLName        xml.Name `xml:"scopecontent" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Head           *Head    `xml:"head,omitempty" json:"head,omitempty"`
	P              []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// Arrangement describes the orientation of material?????
type Arrangement struct {
	XMLName        xml.Name `xml:"arrangement" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Head           *Head    `xml:"head,omitempty" json:"head,omitempty"`
	P              []*P     `xml:"p,omitempty" json:"p,omitempty"`
	List           *List    `xml:"list,omitempty" json:"list,omitempty"`
}

// List provides a list of reference to an item
type List struct {
	XMLName  xml.Name `xml:"list" json:"-"`
	ListType string   `xml:"listtype,attr,omitempty" json:"listtype,omitempty"`
	Value    string   `xml:",innerxml" json:"value,omitempty"`
	// ListHead *ListHead  `xml:"listhead,omitempty" json:"listhead,omitempty"`
	// DefItem  []*DefItem `xml:"defitem,omitempty" json:"defitem,omitempty"`
	// Item     []*Item    `xml:"item,omitempty" json:"item,omitempty"`
}

// type ListHead struct {
// 	XMLName xml.Name `xml:"listtype" json:"-"`
// 	Value   string   `xml:",innerxml" json:"value"`
// }
//
// type DefItem struct {
// 	XMLName xml.Name `xml:"defitem" json:"-"`
// 	Label   string   `xml:"label,omitempty" json:"label,omitempty"`
// 	Item    *Item    `xml:"item,omitempty" json:"item,omitempty"`
// }
//
// // Item provides a reference link to an item
// type Item struct {
// 	XMLName xml.Name `xml:"item" json:"-"`
// 	Ref     *Ref     `xml:"ref,omtempty" json:"ref,omitempty"`
// 	Value   string   `xml:",chardata" json:"value,omitempty"`
// }
//
// Ref is a link to something
type Ref struct {
	XMLName      xml.Name `xml:"ref" json:"-"`
	XMLNameSpace string   `xml:"xmlns,attr" json:"-"`
	HRef         string   `xml:"href,attr,omitempty" json:"href,omitempty"`
	Show         string   `xml:"show,attr,omitempty" json:"show,omitempty"`
	Actuate      string   `xml:"actuate,attr,omitempty" json:"actuate,omitempty"`
	//InstanceURL  string   `xml:"instanceurl,attr,omitempty" json:"instanceurl,omitempty"`
	Target string `xml:"target,attr,omitempty" json:"target,omitempty"`
	Value  string `xml:",chardata" json:"value,omitempty"`
}

// ControlAccess describes who can do what
type ControlAccess struct {
	XMLName        xml.Name         `xml:"controlaccess" json:"-"`
	EncodingAnalog string           `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Head           *Head            `xml:"head,omitempty" json:"head,omitempty"`
	P              []*P             `xml:"p,omitempty" json:"p,omitempty"`
	Persname       []*Persname      `xml:"persname,omitempty" json:"persname,omitempty"`
	Famname        []*Famname       `xml:"famname,omitempty" json:"famname,omitempty"`
	CorpName       []*CorpName      `xml:"corpname,omitempty" json:"corpname,omitempty"`
	Subject        []*Subject       `xml:"subject,omitempty" json:"subject,omitempty"`
	GenreForm      []*GenreForm     `xml:"genreform,omitempty" json:"genreform,omitempty"`
	GeogName       []*GeogName      `xml:"geogname,omitempty" json:"geogname,omitempty"`
	Occupation     []*Occupation    `xml:"occupation,omitempty" json:"occupation,omitempty"`
	ControlAccess  []*ControlAccess `xml:"controlaccess,omitempty" json:"controlaccess,omitempty"`
}

// RelatedMaterial provides links out to related material
type RelatedMaterial struct {
	XMLName        xml.Name `xml:"relatedmaterial" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Head           *Head    `xml:"head,omitempty" json:"head,omitempty"`
	P              []*P     `xml:"p,omitempty" json:"p,omitempty"`
	ArchRef        *ArchRef `xml:"archref,omitempty" json:"archref,omitempty"`
	List           *List    `xml:"list,omitempty" json:"list,omitempty"`
}

// Relations
type Relations struct {
	XMLName  xml.Name    `xml:"relations" json:"-"`
	Relation []*Relation `xml:"relation,omitempty" json:"relation,omitempty"`
}

// Relation
type Relation struct {
	XMLName          xml.Name         `xml:"relation" json:"-"`
	ID               string           `xml:"id,attr,omitempty" json:"id,omitempty"`
	RelationshipType string           `xml:"relationtype,attr,omitempty" json:"relationtype,omitempty"`
	Show             string           `xml:"show,attr,omitempty" json:"show,omitempty"`
	HRef             string           `xml:"href,attr,omitempty" json:"href,omitempty"`
	Actuate          string           `xml:"actuate,attr,omitempty" json:"accuate,omitempty"`
	RelationEntry    string           `xml:"relationentry,omitempty" json:"relationentry,omitempty"`
	DescriptiveNote  *DescriptiveNote `xml:"descriptivenote,omitempty" json:"descriptivenote,omitempty"`
}

// ArchRef archival reference
type ArchRef struct {
	XMLName xml.Name `xml:"archref" json:"-"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
}

// AccessRestrict describes the containstraints under which items can be accessed
type AccessRestrict struct {
	XMLName        xml.Name `xml:"accessrestrict" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Head           string   `xml:"head,omitempty" json:"head,omitempty"`
	P              []*P     `xml:"p,omitempty" json:"p,omitempty"`
	List           *List    `xml:"list,omitempty" json:"list,omitempty"`
}

// UseRestrict describe use restrictions that access to an item allows
type UseRestrict struct {
	XMLName        xml.Name `xml:"userestrict" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Head           string   `xml:"head,omitempty" json:"head,omitempty"`
	P              []*P     `xml:"p,omitempty" json:"p,omitempty"`
	Table          *Table   `xml:"table,omitempty" json:"table,omitempty"`
}

type Table struct {
	XMLName xml.Name `xml:"table" json:"table"`
	Frame   string   `xml:"frame,attr,omitempty" json:"frame,omitempty"`
	Colsep  string   `xml:"colsep,attr,omitempty" json:"colsep,omitempty"`
	Rowsep  string   `xml:"rowsep,attr,omitempty" json:"rowsep,omitempty"`
	TGroup  *TGroup  `xml:"tgroup,omitempty" json:"tgroup,omitempty"`
}

type TGroup struct {
	XMLName xml.Name   `xml:"tgroup" json:"-"`
	Align   string     `xml:"align,attr,omitempty" json:"align,omitempty"`
	Cols    string     `xml:"cols,attr,omitempty" json:"cols,omitempty"`
	ColSpec []*ColSpec `xml:"colspec,omitempty" json:"colspec,omitempty"`
	THead   *THead     `xml:"thead,omitempty" json:"thead,omitempty"`
	TBody   *TBody     `xml:"tbody,omitempty" json:"tbody,omitempty"`
}

type ColSpec struct {
	XMLName xml.Name `xml:"colspec" json:"-"`
	ColName string   `xml:"colname,attr,omitempty" json:"colname,omitempty"`
	ColNum  string   `xml:"colnum,attr,omitempty" json:"colnum,omitempty"`
	Value   string   `xml:",chardata" json:"value,omitempty"`
}

type THead struct {
	XMLName xml.Name `xml:"thead" json:"-"`
	VAlign  string   `xml:"valign,attr,omitempty" json:"valign,omitempty"`
	Row     []*Row   `xml:"row,omitempty" json:"row,omitempty"`
}

type Row struct {
	XMLName xml.Name `xml:"row" json:"-"`
	Entry   []*Entry `xml:"entry,omitempty" json:"entry,omitempty"`
}

type Entry struct {
	XMLName xml.Name `xml:"entry" json:"-"`
	ColName string   `xml:"colname,attr,omitempty" json:"colname,omitempty"`
	Value   string   `xml:",chardata" json:"value,omitempty"`
}

type TBody struct {
	XMLName xml.Name `xml:"tbody" json:"-"`
	VAlign  string   `xml:"valign,attr,omitempty" json:"valign,omitempty"`
	Row     []*Row   `xml:"row,omitempty" json:"row,omitempty"`
}

// AcqInfo provides information about acquisition of an item
type AcqInfo struct {
	XMLName        xml.Name `xml:"acqinfo" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Head           string   `xml:"head,omitempty" json:"head,omitempty"`
	P              []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// ProcessInfo proccessing information
type ProcessInfo struct {
	XMLName        xml.Name `xml:"processinfo" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Head           string   `xml:"head,omitempty" json:"head,omitempty"`
	P              []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// AltFormAvail describes other forms of access for an item
type AltFormAvail struct {
	XMLName xml.Name `xml:"altformavail" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// Appraisal information
type Appraisal struct {
	XMLName xml.Name `xml:"appraisal" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// CustodHist provides a custodial history of content
type CustodHist struct {
	XMLName xml.Name `xml:"custodhist" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// FilePlan Information about any classification scheme used for arranging, storing, and retrieving
// the described materials by the parties originally responsible for creating or compiling
// them.
type FilePlan struct {
	XMLName xml.Name `xml:"fileplan" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// Accruals describes added values for an element
type Accruals struct {
	XMLName xml.Name `xml:"accruals" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// LegalStatus is what is sounds like
type LegalStatus struct {
	XMLName xml.Name `xml:"legalstatus" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// Odd misc.
type Odd struct {
	XMLName xml.Name `xml:"odd" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// OriginalsLoc location of original content
type OriginalsLoc struct {
	XMLName xml.Name `xml:"originalsloc" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// PreferCite preferred citation format
type PreferCite struct {
	XMLName        xml.Name `xml:"prefercite" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Head           string   `xml:"head,omitempty" json:"head,omitempty"`
	P              []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// OtherFindAID other finding aids
type OtherFindAID struct {
	XMLName xml.Name `xml:"otherfindaid" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// PhysTech Physical Characteristics and Technical Specifications
type PhysTech struct {
	XMLName xml.Name `xml:"phystech" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// SeparatedMaterial element indicates items acquired as part of a collection
// and then subsequently removed from the collection.
type SeparatedMaterial struct {
	XMLName xml.Name `xml:"separatedmaterial" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// Dsc - descovery???
type Dsc struct {
	XMLName xml.Name `xml:"dsc" json:"-"`
	DscType string   `xml:"dsctype,attr,omitempty" json:"dsctype,omitempty"`
	Head    *Head    `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
	C       []*C     `xml:"c,omitempty" json:"c,omitempty"`
	C01     []*C01   `xml:"c01,omitempty" json:"c01"`
}

// C container un-numbered level
type C struct {
	XMLName        xml.Name        `xml:"c" json:"-"`
	Level          string          `xml:"level,attr,omitempty" json:"level"`
	ID             string          `xml:"id,attr,omitempty" json:"id"`
	DID            *DID            `xml:"did,omitempty" json:"did,omitempty"`
	ScopeContent   *ScopeContent   `xml:"scopecontent,omitempty" json:"scopecontent,omitempty"`
	AccessRestrict *AccessRestrict `xml:"accessrestrict,omitempty" json:"accessrestrict,omitempty"`
	UseRestrict    *UseRestrict    `xml:"userestrict,omitempty" json:"userestrict,omitempty"`
	PhysTech       *PhysTech       `xml:"phystech,omitempty" json:"phystech,omitempty"`
	DIDNote        *DIDNote        `xml:"didnote,omitempty" json:"didnote,omitempty"`
	Container      []*Container    `xml:"container,omitempty" json:"container,omitempty"`
	C              []*C            `xml:"c,omitempty" json:"c,omitempty"`
	Odd            *Odd            `xml:"odd,omitempty" json:"odd,omitempty"`
	ControlAccess  *ControlAccess  `xml:"controlaccess,omitempty" json:"controlaccess,omitempty"`
	OriginalsLoc   *OriginalsLoc   `xml:"originalsloc,omitempty" json:"originalsloc,omitempty"`
	AltFormAvail   *AltFormAvail   `xml:"altformavail,omitempty" json:"altformavail,omitempty"`
	Index          *Index          `xml:"index,omitempty" json:"index,omitempty"`
}

type Index struct {
	XMLName xml.Name `xml:"index" json:"-"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
}

// C01 container level 1
type C01 struct {
	XMLName        xml.Name        `xml:"c01" json:"-"`
	Level          string          `xml:"level,attr,omitempty" json:"level,omitempty"`
	ID             string          `xml:"id,attr,omitempty" json:"id,omitempty"`
	DID            *DID            `xml:"did,omitempty" json:"did,omitempty"`
	ScopeContent   *ScopeContent   `xml:"scopecontent,omitempty" json:"scopecontent,omitempty"`
	AccessRestrict *AccessRestrict `xml:"accessrestrict,omitempty" json:"accessrestrict,omitempty"`
	UseRestrict    *UseRestrict    `xml:"userestrict,omitempty" json:"userestrict,omitempty"`
	PhysTech       *PhysTech       `xml:"phystech,omitempty" json:"phystech,omitempty"`
	DIDNote        *DIDNote        `xml:"didnote,omitempty" json:"didnote,omitempty"`
	Container      []*Container    `xml:"container,omitempty" json:"container,omitempty"`
	C02            []*C02          `xml:"c02,omitempty" json:"c02,omitempty"`
}

// C02 container level 2
type C02 struct {
	XMLName        xml.Name        `xml:"c02" json:"-"`
	Level          string          `xml:"level,attr,omitempty" json:"level,omitempty"`
	ID             string          `xml:"id,attr,omitempty" json:"id,omitempty"`
	DID            *DID            `xml:"did,omitempty" json:"did,omitempty"`
	ScopeContent   *ScopeContent   `xml:"scopecontent,omitempty" json:"scopecontent,omitempty"`
	AccessRestrict *AccessRestrict `xml:"accessrestrict,omitempty" json:"accessrestrict,omitempty"`
	UseRestrict    *UseRestrict    `xml:"userestrict,omitempty" json:"userestrict,omitempty"`
	PhysTech       *PhysTech       `xml:"phystech,omitempty" json:"phystech,omitempty"`
	DIDNote        *DIDNote        `xml:"didnote,omitempty" json:"didnote,omitempty"`
	Container      []*Container    `xml:"container,omitempty" json:"container,omitempty"`
	C03            []*C03          `xml:"c03,omitempty" json:"c03,omitempty"`
}

// C03 Container level 3
type C03 struct {
	XMLName        xml.Name        `xml:"c03" json:"-"`
	Level          string          `xml:"level,attr,omitempty" json:"level,omitempty"`
	ID             string          `xml:"id,attr,omitempty" json:"id,omitempty"`
	DID            *DID            `xml:"did,omitempty" json:"did,omitempty"`
	ScopeContent   *ScopeContent   `xml:"scopecontent,omitempty" json:"scopecontent,omitempty"`
	AccessRestrict *AccessRestrict `xml:"accessrestrict,omitempty" json:"accessrestrict,omitempty"`
	UseRestrict    *UseRestrict    `xml:"userestrict,omitempty" json:"userestrict,omitempty"`
	PhysTech       *PhysTech       `xml:"phystech,omitempty" json:"phystech,omitempty"`
	DIDNote        *DIDNote        `xml:"didnote,omitempty" json:"didnote,omitempty"`
	Container      []*Container    `xml:"container,omitempty" json:"container,omitempty"`
	C04            []*C04          `xml:"c04,omitempty" json:"c04,omitempty"`
}

// C04 container level 4
type C04 struct {
	XMLName        xml.Name        `xml:"c04" json:"-"`
	Level          string          `xml:"level,attr,omitempty" json:"level,omitempty"`
	ID             string          `xml:"id,attr,omitempty" json:"id,omitempty"`
	DID            *DID            `xml:"did,omitempty" json:"did,omitempty"`
	ScopeContent   *ScopeContent   `xml:"scopecontent,omitempty" json:"scopecontent,omitempty"`
	AccessRestrict *AccessRestrict `xml:"accessrestrict,omitempty" json:"accessrestrict,omitempty"`
	UseRestrict    *UseRestrict    `xml:"userestrict,omitempty" json:"userestrict,omitempty"`
	PhysTech       *PhysTech       `xml:"phystech,omitempty" json:"phystech,omitempty"`
	DIDNote        *DIDNote        `xml:"didnote,omitempty" json:"didnote,omitempty"`
	Container      []*Container    `xml:"container,omitempty" json:"container,omitempty"`
}

// Sources contains one or more source
type Sources struct {
	XMLName xml.Name  `xml:"sources" json:"-"`
	Source  []*Source `xml:"source,omitempty" json:"source,omitempty"`
}

// Source detailed information about source
type Source struct {
	XMLName         xml.Name         `xml:"source" json:"-"`
	SourceEntry     string           `xml:"sourceentry,omitempty" json:"sourceentry,omitempty"`
	ObjectXMLWrap   *ObjectXMLWrap   `xml:"objectxmlwrap,omitempty" json:"objectxmlwrap,omitempty"`
	Descriptivenote *DescriptiveNote `xml:"descriptivenote,omitempty" json:"descriptivenote,omitempty"`
}

// ObjectXMLWrap include an existing XML object as is.
type ObjectXMLWrap struct {
	XMLName xml.Name `xml:"objectxmlwrap" json:"-"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
}

// Create a new EAD document structure
func New() *EAD3 {
	obj := new(EAD3)
	obj.XMLNameSpace = "http://ead3.archivists.org/schema/undeprecated/"
	return obj
}
