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
	XMLName  xml.Name  `xml:"ead" json:"-"`
	Control  *Control  `xml:"control" json:"control"`
	ArchDesc *ArchDesc `xml:"archdesc" json:"archdesc"`
}

// P provides for unparsed embedded markup of paragraph elements
type P struct {
	XMLName xml.Name `xml:"p" json:"-"`
	Value   string   `xml:",innerxml" json:"text"`
}

// Control structure for initial element of an EAD
type Control struct {
	XMLName               xml.Name               `xml:"control" json:"-"`
	RecordID              string                 `xml:"recordid" json:"record_id"`
	FileDesc              *FileDesc              `xml:"filedesc" json:"file_desc"`
	MaintenanceStatus     *MaintenanceStatus     `xml:"maintenancestatus" json:"maintenance_status"`
	MaintenanceAgency     *MaintenanceAgency     `xml:"maintenanceagency" json:"maintenance_agency"`
	LanguageDeclaration   *LanguageDeclaration   `xml:"languagedeclaration" json:"language_declaration"`
	ConventionDeclaration *ConventionDeclaration `xml:"conventiondeclaration" json:"convention_declaration"`
	MaintenanceHistory    *MaintenanceHistory    `xml:"maintenancehistory" json:"maintenance_history"`
}

// FileDesc describes file system contents
type FileDesc struct {
	XMLName         xml.Name         `xml:"filedesc" json:"-"`
	TitleStmt       *TitleStmt       `xml:"titlestmt" json:"title_stmt"`
	PublicationStmt *PublicationStmt `xml:"publicationstmt" json:"publication_stmt"`
}

// TitleStmt provides structure relating to titling and authorship
type TitleStmt struct {
	XMLName     xml.Name `xml:"titlestmt" json:"-"`
	TitleProper string   `xml:"titleproper" json:"title_proper"`
	Subtitle    string   `xml:"subtitle" json:"subtitle"`
	Author      string   `xml:"author" json:"author"`
}

// PublicationStmt provides information on the publication nature of content
type PublicationStmt struct {
	XMLName xml.Name `xml:"publicationstmt" json:"-"`
	P       []*P     `xml:"p" json:"p"`
}

// MaintenanceStatus provides an descriptive meantenance status
type MaintenanceStatus struct {
	XMLName xml.Name `xml:"maintenancestatus" json:"-"`
	Value   string   `xml:"value,attr" json:"value"`
}

// MaintenanceAgency provides content related organziation performing maintenance
type MaintenanceAgency struct {
	XMLName    xml.Name `xml:"maintenanceagency" json:"-"`
	Agencyname string   `xml:"agencyname" json:"agency_name"`
}

// ConventionDeclaration provides ciation declarations
type ConventionDeclaration struct {
	XMLName  xml.Name `xml:"conventiondeclaration" json:"-"`
	Citation string   `xml:"citation" json:"citation"`
}

// LanguageDeclaration describes relevant language implications
type LanguageDeclaration struct {
	XMLName  xml.Name  `xml:"languagedeclaration" json:"-"`
	Language *Language `xml:"language" json:"language"`
	Script   *Script   `xml:"script" json:"script"`
}

// Language describes a specific language
type Language struct {
	XMLName  xml.Name `xml:"language" json:"-"`
	LangCode string   `xml:"langcode,attr" json:"lang_code,omitempty"`
	Value    string   `xml:",chardata" json:"value"`
}

// Script describes the character encoding and script representation used in content
type Script struct {
	XMLName    xml.Name `xml:"script" json:"-"`
	ScriptCode string   `xml:"scriptcode,attr" json:"script_code,omitempty"`
}

// MaintenanceHistory provides a collection of maintenance events
type MaintenanceHistory struct {
	XMLName          xml.Name            `xml:"maintenancehistory" json:"-"`
	MaintenanceEvent []*MaintenanceEvent `xml:"maintenanceevent" json:"maintenance_event"`
}

// MaintenanceEvent describes activities related to processing content
type MaintenanceEvent struct {
	XMLName       xml.Name   `xml:"maintenanceevent" json:"-"`
	EventType     *EventType `xml:"eventtype" json:"event_type"`
	EventDateTime string     `xml:"eventdatetime" json:"event_datetime"`
	AgentType     *AgentType `xml:"agenttype" json:"agent_type"`
	Agent         string     `xml:"agent" json:"agent"`
}

// EventType describes the type of maintenance event
type EventType struct {
	XMLName xml.Name `xml:"eventtype" json:"-"`
	Value   string   `xml:"value,attr" json:"event_type"`
}

// AgentType describes of the acting parties in the maintenance event
type AgentType struct {
	XMLName xml.Name `xml:"agenttype" json:"-"`
	Value   string   `xml:"value,attr" json:"agent_type"`
}

// ArchDesc provides an Archival Description of the content
type ArchDesc struct {
	XMLName           xml.Name           `xml:"archdesc" json:"-"`
	Level             string             `xml:"level,attr" json:"level"`
	DID               []*DID             `xml:"did" json:"did"`
	BiogHist          *BiogHist          `xml:"bioghist" json:"bioghist"`
	ScopeContent      *ScopeContent      `xml:"scopecontent" json:"scopecontent"`
	Arrangement       *Arrangement       `xml:"arrangement" json:"arrangement"`
	ControlAccess     *ControlAccess     `xml:"controlaccess" json:"controlaccess"`
	RelatedMaterial   *RelatedMaterial   `xml:"relatedmaterial" json:"related_material"`
	AccessRestrict    *AccessRestrict    `xml:"accessrestrict" json:"access_restrict"`
	UseRestrict       *UseRestrict       `xml:"userestrict" json:"use_restrict"`
	AcqInfo           *AcqInfo           `xml:"acqinfo" json:"acqinfo"`
	ProcessInfo       *ProcessInfo       `xml:"processinfo" json:"processinfo"`
	AltFormAvail      *AltFormAvail      `xml:"altformavail" json:"altformavail"`
	Appraisal         *Appraisal         `xml:"appraisal" json:"appraisal"`
	CustodHist        *CustodHist        `xml:"custodhist" json:"custod_hist"`
	FilePlan          *FilePlan          `xml:"fileplan" json:"fileplan"`
	Accruals          *Accruals          `xml:"accruals" json:"accruals"`
	LegalStatus       *LegalStatus       `xml:"legalstatus" json:"legalstatus"`
	Odd               *Odd               `xml:"odd" json:"odd"`
	OriginalsLoc      *OriginalsLoc      `xml:"originalsloc" json:"originals_loc"`
	PreferCite        *PreferCite        `xml:"prefercite" json:"prefer_cite"`
	OtherFindAID      *OtherFindAID      `xml:"otherfindaid" json:"other_find_aid"`
	PhysTech          *PhysTech          `xml:"phystech" json:"phys_tech"`
	SeparatedMaterial *SeparatedMaterial `xml:"separatedmaterial" json:"separatedmaterial"`
	Dsc               *Dsc               `xml:"dsc" json:"dsc"`
}

// DID  Descriptive Identification of the Unit
type DID struct {
	XMLName            xml.Name            `xml:"did" json:"-"`
	Head               string              `xml:"head,omitempty" json:"head,omitempty"`
	Repository         *Repository         `xml:"repository" json:"repository"`
	Origination        *Origination        `xml:"origination" json:"origination"`
	UnitTitle          *UnitTitle          `xml:"unittitle" json:"unit_title"`
	UnitDateStructured *UnitDateStructured `xml:"unitdatestructured,omitempty" json:"unit_date_structured,omitempty"`
	UnitDates          *UnitDates          `xml:"unitdates,omitempty" json:"unit_dates,omitempty"`
	PhysDesc           *PhysDesc           `xml:"physdesc,omitempty" json:"phys_desc,omitempty"`
	PhysDescSet        *PhysDescSet        `xml:"physdescset,omitempty" json:"phys_desc_set,omitempty"`
	UnitID             *UnitID             `xml:"unitid,omitempty" json:"unit_id,omitempty"`
	Abstract           *Abstract           `xml:"abstract,omitempty" json:"abstract,omitempty"`
	DIDNote            *DIDNote            `xml:"didnote,omitempty" json:"did_note,omitempty"`
	LangMaterial       *LangMaterial       `xml:"langmaterial,omitempty" json:"lang_material,omitempty"`
	PhysLoc            *PhysLoc            `xml:"physloc,omitempty" json:"phys_loc,omitempty"`
	Container          []*Container        `xml:"container,omitempty" json:"container,omitempty"`
}

// Container describes the container holding materials
type Container struct {
	XMLName   xml.Name `xml:"container" json:"-"`
	LocalType string   `xml:"localtype,attr,omitempty" json:"local_type,omitempty"`
	Value     string   `xml:",chardata" json:"value,omitempty"`
}

// Repository describes the repository holding the materials
type Repository struct {
	XMLName        xml.Name  `xml:"repository" json:"-"`
	Label          string    `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog string    `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	CorpName       *CorpName `xml:"corpname" json:"corpname"`
}

// CorpName - Corpus name
type CorpName struct {
	XMLName        xml.Name `xml:"corpname" json:"-"`
	Source         string   `xml:"source,attr" json:"source"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Part           []*Part  `xml:"part" json:"part"`
}

// Part describes the part of a corpus or other organization
type Part struct {
	XMLName   xml.Name `xml:"part" json:"-"`
	LocalType string   `xml:"localtype,attr" json:"localtype"`
	Value     string   `xml:",chardata" json:"value"`
}

// Origination describes where things came from
type Origination struct {
	XMLName        xml.Name    `xml:"origination" json:"-"`
	Label          string      `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog string      `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Persname       []*Persname `xml:"persname" json:"pers_name"`
}

// Persname is a person's name(s)
type Persname struct {
	XMLName        xml.Name `xml:"persname" json:"-"`
	Rules          string   `xml:"rules,attr" json:"rules"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Part           []*Part  `xml:"part" json:"part"`
}

// Subject describes the subject of contents
type Subject struct {
	XMLName        xml.Name `xml:"subject" json:"-"`
	Source         string   `xml:"source,attr" json:"source"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Part           []*Part  `xml:"part" json:"part"`
}

// GenreForm describe the genre of the contents
type GenreForm struct {
	XMLName        xml.Name `xml:"genreform" json:"-"`
	Source         string   `xml:"source,attr" json:"source"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Part           []*Part  `xml:"part" json:"part"`
}

// UnitTitle is a title for a specific unit of collect contents
type UnitTitle struct {
	XMLName        xml.Name `xml:"unittitle" json:"-"`
	Label          string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Value          string   `xml:",chardata" json:"value"`
}

// UnitDateStructured provides a descriptive structure of date information related to the unit of content
type UnitDateStructured struct {
	XMLName        xml.Name     `xml:"unitdatestructured" json"-"`
	Label          string       `xml:"label,attr,omitempty" json:"label,omitempty"`
	UnitDateType   string       `xml:"unitdatetype,attr" json:"unit_date_type"`
	EncodingAnalog string       `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	DateRange      []*DateRange `xml:"daterange" json:"daterange"`
}

// UnitDates provides a simpler date structure relating to content
type UnitDates struct {
	XMLName xml.Name `xml:"unitdates" json:"-"`
	Label   string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	Value   string   `xml:",chardata" json:"value"`
}

// DateRange describes the start and end dates
type DateRange struct {
	XMLName  xml.Name `xml:"daterange" json"-"`
	FromDate string   `xml:"fromdate" json:"from_date"`
	ToDate   string   `xml:"todate" json:"to_date"`
}

// PhysDesc contains the physical description of contents
type PhysDesc struct {
	XMLName xml.Name `xml:"physdesc" json:"-"`
	Label   string   `xml:"label,attr" json:"label"`
	Value   string   `xml:",chardata" json:"value"`
}

// PhysDescSet describes a grouping of physical descriptions of content
type PhysDescSet struct {
	XMLName            xml.Name              `xml:"physdescset" json:"-"`
	PhysDescStructured []*PhysDescStructured `xml:"physdescstructured" json:"physdescstructured"`
}

// PhysDescStructured provides a structured set of physical descriptions
type PhysDescStructured struct {
	XMLName                xml.Name `xml:"physdescstructured" json:"-"`
	Label                  string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	PhysDescStructuredType string   `xml:"physdescstructuredtype,attr" json:"physdescstructuredtype"`
	Coverage               string   `xml:"coverage,attr" json:"coverage"`
	EncodingAnalog         string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Quantity               string   `xml:"quantity" json:"quantity"`
	UnitType               string   `xml:"unittype" json:"unittype"`
}

// UnitID provides a unit level identifier
type UnitID struct {
	XMLName xml.Name `xml:"unitid" json:"-"`
	Label   string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	Value   string   `xml:",chardata" json:"value"`
}

// Abstract provides a summary of content
type Abstract struct {
	XMLName xml.Name `xml:"abstract" json:"-"`
	Label   string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	Value   string   `xml:",chardata" json:"value"`
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
	Language        *Language        `xml:"language" json:"language"`
	DescriptiveNote *DescriptiveNote `xml:"descriptivenote" json:"descriptivenote"`
}

// DescriptiveNote is a descriptive note about the content
type DescriptiveNote struct {
	XMLName xml.Name `xml:"descriptivenote" json:"-"`
	P       []*P     `xml:"p" json:"p"`
}

// PhysLoc describes the physical location of the content
type PhysLoc struct {
	XMLName xml.Name `xml:"physloc" json:"-"`
	Label   string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	Value   string   `xml:",chardata" json:"value"`
}

// BiogHist provides biographical history of content
type BiogHist struct {
	XMLName        xml.Name `xml:"bioghist" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Head           string   `xml:"head,omitempty" json:"head,omitempty"`
	P              []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// ScopeContent provides scoping material for content
type ScopeContent struct {
	XMLName        xml.Name `xml:"scopecontent" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Head           string   `xml:"head,omitempty" json:"head,omitempty"`
	P              []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// Arrangement describes the orientation of material?????
type Arrangement struct {
	XMLName        xml.Name `xml:"arrangement" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Head           string   `xml:"head,omitempty" json:"head,omitempty"`
	P              []*P     `xml:"p,omitempty" json:"p,omitempty"`
	List           *List    `xml:"list" json:"list"`
}

// List provides a list of reference to an item
type List struct {
	XMLName xml.Name `xml:"list" json:"-"`
	Item    []*Item  `xml:"item" json:"item"`
}

// Item provides a reference link to an item
type Item struct {
	XMLName xml.Name `xml:"item" json:"-"`
	Ref     *Ref     `xml:"ref" json:"ref"`
}

// Ref is a link to something
type Ref struct {
	XMLName xml.Name `xml:"ref" json:"-"`
	Target  string   `xml:"target,attr" json:"target"`
	Value   string   `xml:",chardata" json:"value"`
}

// ControlAccess describes who can do what
type ControlAccess struct {
	XMLName       xml.Name         `xml:"controlaccess" json:"-"`
	Head          string           `xml:"head,omitempty" json:"head,omitempty"`
	P             []*P             `xml:"p,omitempty" json:"p,omitempty"`
	Persname      []*Persname      `xml:"persname,omitempty" json:"pers_name,omitempty"`
	Subject       []*Subject       `xml:"subject,omitempty" json:"subject,omitempty"`
	GenreForm     []*GenreForm     `xml:"genreform,omitempty" json:"genreform,omitempty"`
	ControlAccess []*ControlAccess `xml:"controlaccess,omitempty" json:"control_access,omitempty"`
}

// RelatedMaterial provides links out to related material
type RelatedMaterial struct {
	XMLName xml.Name `xml:"relatedmaterial" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// AccessRestrict describes the containstraints under which items can be accessed
type AccessRestrict struct {
	XMLName xml.Name `xml:"accessrestrict" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// UseRestrict describe use restrictions that access to an item allows
type UseRestrict struct {
	XMLName xml.Name `xml:"userestrict" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// AcqInfo provides information about acquisition of an item
type AcqInfo struct {
	XMLName xml.Name `xml:"acqinfo" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

// ProcessInfo proccessing information
type ProcessInfo struct {
	XMLName xml.Name `xml:"processinfo" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
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
	XMLName xml.Name `xml:"prefercite" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
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
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
	C01     []*C01   `xml:"c01,omitempty" json:"c01"`
}

// C01 container level 1
type C01 struct {
	XMLName        xml.Name        `xml:"c01" json:"-"`
	Level          string          `xml:"level,attr" json:"level"`
	ID             string          `xml:"id,attr" json:"id"`
	DID            *DID            `xml:"did,omitempty" json:"did,omitempty"`
	ScopeContent   *ScopeContent   `xml:"scopecontent,omitempty" json:"scopecontent,omitempty"`
	AccessRestrict *AccessRestrict `xml:"accessrestrict,omitempty" json:"accessrestrict,omitempty"`
	UseRestrict    *UseRestrict    `xml:"userestrict,omitempty" json:"use_restrict,omitempty"`
	PhysTech       *PhysTech       `xml:"phystech,omitempty" json:"phys_tech,omitempty"`
	DIDNote        *DIDNote        `xml:"didnote,omitempty" json:"did_note,omitempty"`
	Container      []*Container    `xml:"container,omitempty" json:"container,omitempty"`
	C02            []*C02          `xml:"c02,omitempty" json:"c02,omitempty"`
}

// C02 container level 2
type C02 struct {
	XMLName        xml.Name        `xml:"c02" json:"-"`
	DID            *DID            `xml:"did,omitempty" json:"did,omitempty"`
	ScopeContent   *ScopeContent   `xml:"scopecontent,omitempty" json:"scopecontent,omitempty"`
	AccessRestrict *AccessRestrict `xml:"accessrestrict,omitempty" json:"accessrestrict,omitempty"`
	UseRestrict    *UseRestrict    `xml:"userestrict,omitempty" json:"use_restrict,omitempty"`
	PhysTech       *PhysTech       `xml:"phystech,omitempty" json:"phys_tech,omitempty"`
	DIDNote        *DIDNote        `xml:"didnote,omitempty" json:"did_note,omitempty"`
	Container      []*Container    `xml:"container,omitempty" json:"container,omitempty"`
	C03            []*C03          `xml:"c03,omitempty" json:"c03,omitempty"`
}

// C03 Container level 3
type C03 struct {
	XMLName        xml.Name        `xml:"c03" json:"-"`
	DID            *DID            `xml:"did,omitempty" json:"did,omitempty"`
	ScopeContent   *ScopeContent   `xml:"scopecontent,omitempty" json:"scopecontent,omitempty"`
	AccessRestrict *AccessRestrict `xml:"accessrestrict,omitempty" json:"accessrestrict,omitempty"`
	UseRestrict    *UseRestrict    `xml:"userestrict,omitempty" json:"use_restrict,omitempty"`
	PhysTech       *PhysTech       `xml:"phystech,omitempty" json:"phys_tech,omitempty"`
	DIDNote        *DIDNote        `xml:"didnote,omitempty" json:"did_note,omitempty"`
	Container      []*Container    `xml:"container,omitempty" json:"container,omitempty"`
	C04            []*C04          `xml:"c04,omitempty" json:"c04,omitempty"`
}

// C04 container level 4
type C04 struct {
	XMLName        xml.Name        `xml:"c04" json:"-"`
	DID            *DID            `xml:"did,omitempty" json:"did,omitempty"`
	ScopeContent   *ScopeContent   `xml:"scopecontent,omitempty" json:"scopecontent,omitempty"`
	AccessRestrict *AccessRestrict `xml:"accessrestrict,omitempty" json:"accessrestrict,omitempty"`
	UseRestrict    *UseRestrict    `xml:"userestrict,omitempty" json:"use_restrict,omitempty"`
	PhysTech       *PhysTech       `xml:"phystech,omitempty" json:"phys_tech,omitempty"`
	DIDNote        *DIDNote        `xml:"didnote,omitempty" json:"did_note,omitempty"`
	Container      []*Container    `xml:"container,omitempty" json:"container,omitempty"`
}
