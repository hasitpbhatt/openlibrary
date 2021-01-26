package olclient

// Work refers to a work
type Work struct {
	Authors           []Item           `json:"authors"`
	ByStatement       string           `json:"by_statement"`
	Classifications   *Classifications `json:"classifications,omitempty"`
	Contributors      []Contributor    `json:"contributors,omitempty"`
	CopyrightDate     string           `json:"copyright_date,omitempty"`
	Covers            []int            `json:"covers"`
	Created           KeyValue         `json:"created"`
	Description       string           `json:"description,omitempty"`
	DeweyDecimalClass []string         `json:"dewey_decimal_class,omitempty"`
	EditionName       string           `json:"edition_name"`
	IABoxID           []string         `json:"ia_box_id,omitempty"`
	IALoadedID        []string         `json:"ia_loaded_id,omitempty"`
	Identifiers       *Identifiers     `json:"identifiers,omitempty"`
	ISBN10            []string         `json:"isbn_10,omitempty"`
	ISBN13            []string         `json:"isbn_13,omitempty"`
	Key               string           `json:"key"`
	Languages         []Item           `json:"languages"`
	LastModified      KeyValue         `json:"last_modified"`
	LatestRevision    int              `json:"latest_revision"`
	LCClassifications []string         `json:"lc_classifications"`
	LCCN              []string         `json:"lccn"`
	Notes             *KeyValue        `json:"notes,omitempty"`
	NumberOfPages     int              `json:"number_of_pages"`
	OCAID             string           `json:"ocaid,omitempty"`
	OCLCNumbers       []string         `json:"oclc_numbers,omitempty"`
	Pagination        string           `json:"pagination"`
	PhysicalFormat    string           `json:"physical_format,omitempty"`
	PublishCountry    string           `json:"publish_country"`
	PublishDate       string           `json:"publish_date"`
	PublishPlaces     []string         `json:"publish_places"`
	Publishers        []string         `json:"publishers"`
	Revision          int              `json:"revision"`
	Series            []string         `json:"series"`
	SourceRecords     []string         `json:"source_records"`
	TableOfContents   []TOC            `json:"table_of_contents,omitempty"`
	Title             string           `json:"title"`
	Type              Item             `json:"type"`
	Works             []Item           `json:"works"`
}

// Item refers to Item
type Item struct {
	Key string `json:"key"`
}

// TOC refers to TOC
type TOC struct {
	Label   string `json:"label"`
	Level   int    `json:"level"`
	Pagenum string `json:"pagenum"`
	Title   string `json:"title"`
}

// KeyValue refers to KeyValue
type KeyValue struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Classifications ...
type Classifications struct {
	LCCNPermalink []string `json:"lccn_permalink"`
}

// Contributor ...
type Contributor struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type Identifiers struct {
	Goodreads    []string `json:"goodreads"`
	Librarything []string `json:"librarything"`
}
