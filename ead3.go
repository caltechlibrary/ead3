package ead3

import "encoding/xml"

//
// EAD 3 structures
//
type EAD3 struct {
	XMLName  xml.Name  `xml:"ead" json:"-"`
	Control  *Control  `xml:"control" json:"control"`
	ArchDesc *ArchDesc `xml:"archdesc" json:"archdesc"`
}

type P struct {
	XMLName xml.Name `xml:"p" json:"-"`
	Value   string   `xml:",innerxml" json:"text"`
}

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

type FileDesc struct {
	XMLName         xml.Name         `xml:"filedesc" json:"-"`
	TitleStmt       *TitleStmt       `xml:"titlestmt" json:"title_stmt"`
	PublicationStmt *PublicationStmt `xml:"publicationstmt" json:"publication_stmt"`
}

type TitleStmt struct {
	XMLName     xml.Name `xml:"titlestmt" json:"-"`
	TitleProper string   `xml:"titleproper" json:"title_proper"`
	Subtitle    string   `xml:"subtitle" json:"subtitle"`
	Author      string   `xml:"author" json:"author"`
}

type PublicationStmt struct {
	XMLName xml.Name `xml:"publicationstmt" json:"-"`
	P       []*P     `xml:"p" json:"p"`
}

type MaintenanceStatus struct {
	XMLName xml.Name `xml:"maintenancestatus" json:"-"`
	Value   string   `xml:"value,attr" json:"value"`
}

type MaintenanceAgency struct {
	XMLName    xml.Name `xml:"maintenanceagency" json:"-"`
	Agencyname string   `xml:"agencyname" json:"agency_name"`
}

type ConventionDeclaration struct {
	XMLName  xml.Name `xml:"conventiondeclaration" json:"-"`
	Citation string   `xml:"citation" json:"citation"`
}

type LanguageDeclaration struct {
	XMLName  xml.Name  `xml:"languagedeclaration" json:"-"`
	Language *Language `xml:"language" json:"language"`
	Script   *Script   `xml:"script" json:"script"`
}

type Language struct {
	XMLName  xml.Name `xml:"language" json:"-"`
	LangCode string   `xml:"langcode,attr" json:"lang_code,omitempty"`
	Value    string   `xml:",chardata" json:"value"`
}

type Script struct {
	XMLName    xml.Name `xml:"script" json:"-"`
	ScriptCode string   `xml:"scriptcode,attr" json:"script_code,omitempty"`
}

type MaintenanceHistory struct {
	XMLName          xml.Name            `xml:"maintenancehistory" json:"-"`
	MaintenanceEvent []*MaintenanceEvent `xml:"maintenanceevent" json:"maintenance_event"`
}

type MaintenanceEvent struct {
	XMLName       xml.Name   `xml:"maintenanceevent" json:"-"`
	EventType     *EventType `xml:"eventtype" json:"event_type"`
	EventDateTime string     `xml:"eventdatetime" json:"event_datetime"`
	AgentType     *AgentType `xml:"agenttype" json:"agent_type"`
	Agent         string     `xml:"agent" json:"agent"`
}

type EventType struct {
	XMLName xml.Name `xml:"eventtype" json:"-"`
	Value   string   `xml:"value,attr" json:"event_type"`
}

type AgentType struct {
	XMLName xml.Name `xml:"agenttype" json:"-"`
	Value   string   `xml:"value,attr" json:"agent_type"`
}

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

type Container struct {
	XMLName   xml.Name `xml:"container" json:"-"`
	LocalType string   `xml:"localtype,attr,omitempty" json:"local_type,omitempty"`
	Value     string   `xml:",chardata" json:"value,omitempty"`
}

type Repository struct {
	XMLName        xml.Name  `xml:"repository" json:"-"`
	Label          string    `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog string    `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	CorpName       *CorpName `xml:"corpname" json:"corpname"`
}

type CorpName struct {
	XMLName        xml.Name `xml:"corpname" json:"-"`
	Source         string   `xml:"source,attr" json:"source"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Part           []*Part  `xml:"part" json:"part"`
}

type Part struct {
	XMLName   xml.Name `xml:"part" json:"-"`
	LocalType string   `xml:"localtype,attr" json:"localtype"`
	Value     string   `xml:",chardata" json:"value"`
}

type Origination struct {
	XMLName        xml.Name    `xml:"origination" json:"-"`
	Label          string      `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog string      `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Persname       []*Persname `xml:"persname" json:"pers_name"`
}

type Persname struct {
	XMLName        xml.Name `xml:"persname" json:"-"`
	Rules          string   `xml:"rules,attr" json:"rules"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Part           []*Part  `xml:"part" json:"part"`
}

type Subject struct {
	XMLName        xml.Name `xml:"subject" json:"-"`
	Source         string   `xml:"source,attr" json:"source"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Part           []*Part  `xml:"part" json:"part"`
}

type GenreForm struct {
	XMLName        xml.Name `xml:"genreform" json:"-"`
	Source         string   `xml:"source,attr" json:"source"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Part           []*Part  `xml:"part" json:"part"`
}

type UnitTitle struct {
	XMLName        xml.Name `xml:"unittitle" json:"-"`
	Label          string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Value          string   `xml:",chardata" json:"value"`
}

type UnitDateStructured struct {
	XMLName        xml.Name     `xml:"unitdatestructured" json"-"`
	Label          string       `xml:"label,attr,omitempty" json:"label,omitempty"`
	UnitDateType   string       `xml:"unitdatetype,attr" json:"unit_date_type"`
	EncodingAnalog string       `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	DateRange      []*DateRange `xml:"daterange" json:"daterange"`
}

type UnitDates struct {
	XMLName xml.Name `xml:"unitdates" json:"-"`
	Label   string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	Value   string   `xml:",chardata" json:"value"`
}

type DateRange struct {
	XMLName  xml.Name `xml:"daterange" json"-"`
	FromDate string   `xml:"fromdate" json:"from_date"`
	ToDate   string   `xml:"todate" json:"to_date"`
}

type PhysDesc struct {
	XMLName xml.Name `xml:"physdesc" json:"-"`
	Label   string   `xml:"label,attr" json:"label"`
	Value   string   `xml:",chardata" json:"value"`
}

type PhysDescSet struct {
	XMLName            xml.Name              `xml:"physdescset" json:"-"`
	PhysDescStructured []*PhysDescStructured `xml:"physdescstructured" json:"physdescstructured"`
}

type PhysDescStructured struct {
	XMLName                xml.Name `xml:"physdescstructured" json:"-"`
	Label                  string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	PhysDescStructuredType string   `xml:"physdescstructuredtype,attr" json:"physdescstructuredtype"`
	Coverage               string   `xml:"coverage,attr" json:"coverage"`
	EncodingAnalog         string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Quantity               string   `xml:"quantity" json:"quantity"`
	UnitType               string   `xml:"unittype" json:"unittype"`
}

type UnitID struct {
	XMLName xml.Name `xml:"unitid" json:"-"`
	Label   string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	Value   string   `xml:",chardata" json:"value"`
}

type Abstract struct {
	XMLName xml.Name `xml:"abstract" json:"-"`
	Label   string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	Value   string   `xml:",chardata" json:"value"`
}

type DIDNote struct {
	XMLName xml.Name `xml:"didnote" json:"-"`
	Label   string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	Value   string   `xml:",chardata" json:"value"`
}

type LangMaterial struct {
	XMLName         xml.Name         `xml:"langmaterial" json:"-"`
	Label           string           `xml:"label,attr,omitempty" json:"label,omitempty"`
	Language        *Language        `xml:"language" json:"language"`
	DescriptiveNote *DescriptiveNote `xml:"descriptivenote" json:"descriptivenote"`
}

type DescriptiveNote struct {
	XMLName xml.Name `xml:"descriptivenote" json:"-"`
	P       []*P     `xml:"p" json:"p"`
}

type PhysLoc struct {
	XMLName xml.Name `xml:"physloc" json:"-"`
	Label   string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	Value   string   `xml:",chardata" json:"value"`
}

type BiogHist struct {
	XMLName        xml.Name `xml:"bioghist" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Head           string   `xml:"head,omitempty" json:"head,omitempty"`
	P              []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type ScopeContent struct {
	XMLName        xml.Name `xml:"scopecontent" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Head           string   `xml:"head,omitempty" json:"head,omitempty"`
	P              []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type Arrangement struct {
	XMLName        xml.Name `xml:"arrangement" json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encoding_analog,omitempty"`
	Head           string   `xml:"head,omitempty" json:"head,omitempty"`
	P              []*P     `xml:"p,omitempty" json:"p,omitempty"`
	List           *List    `xml:"list" json:"list"`
}

type List struct {
	XMLName xml.Name `xml:"list" json:"-"`
	Item    []*Item  `xml:"item" json:"item"`
}

type Item struct {
	XMLName xml.Name `xml:"item" json:"-"`
	Ref     *Ref     `xml:"ref" json:"ref"`
}

type Ref struct {
	XMLName xml.Name `xml:"ref" json:"-"`
	Target  string   `xml:"target,attr" json:"target"`
	Value   string   `xml:",chardata" json:"value"`
}

type ControlAccess struct {
	XMLName       xml.Name         `xml:"controlaccess" json:"-"`
	Head          string           `xml:"head,omitempty" json:"head,omitempty"`
	P             []*P             `xml:"p,omitempty" json:"p,omitempty"`
	Persname      []*Persname      `xml:"persname,omitempty" json:"pers_name,omitempty"`
	Subject       []*Subject       `xml:"subject,omitempty" json:"subject,omitempty"`
	GenreForm     []*GenreForm     `xml:"genreform,omitempty" json:"genreform,omitempty"`
	ControlAccess []*ControlAccess `xml:"controlaccess,omitempty" json:"control_access,omitempty"`
}

type RelatedMaterial struct {
	XMLName xml.Name `xml:"relatedmaterial" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type AccessRestrict struct {
	XMLName xml.Name `xml:"accessrestrict" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type UseRestrict struct {
	XMLName xml.Name `xml:"userestrict" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type AcqInfo struct {
	XMLName xml.Name `xml:"acqinfo" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type ProcessInfo struct {
	XMLName xml.Name `xml:"processinfo" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type AltFormAvail struct {
	XMLName xml.Name `xml:"altformavail" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type Appraisal struct {
	XMLName xml.Name `xml:"appraisal" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type CustodHist struct {
	XMLName xml.Name `xml:"custodhist" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type FilePlan struct {
	XMLName xml.Name `xml:"fileplan" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type Accruals struct {
	XMLName xml.Name `xml:"accruals" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type LegalStatus struct {
	XMLName xml.Name `xml:"legalstatus" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type Odd struct {
	XMLName xml.Name `xml:"odd" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type OriginalsLoc struct {
	XMLName xml.Name `xml:"originalsloc" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type PreferCite struct {
	XMLName xml.Name `xml:"prefercite" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type OtherFindAID struct {
	XMLName xml.Name `xml:"otherfindaid" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type PhysTech struct {
	XMLName xml.Name `xml:"phystech" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type SeparatedMaterial struct {
	XMLName xml.Name `xml:"separatedmaterial" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
}

type Dsc struct {
	XMLName xml.Name `xml:"dsc" json:"-"`
	Head    string   `xml:"head,omitempty" json:"head,omitempty"`
	P       []*P     `xml:"p,omitempty" json:"p,omitempty"`
	C01     []*C01   `xml:"c01,omitempty" json:"c01"`
}

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
