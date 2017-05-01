package main

type Tooltip struct {
	Shared    bool   `json:"shared"`
	ValueType string `json:"value_type"`
}

type Legend struct {
	AlignAsTable bool `json:"alignAsTable"`
	Avg          bool `json:"avg"`
	Current      bool `json:"current"`
	Max          bool `json:"max"`
	Min          bool `json:"min"`
	Show         bool `json:"show"`
	Total        bool `json:"total"`
	Values       bool `json:"values"`
}

type Grid struct {
	LeftLogBase     int         `json:"leftLogBase"`
	LeftMax         interface{} `json:"leftMax"`
	LeftMin         int         `json:"leftMin"`
	RightLogBase    int         `json:"rightLogBase"`
	RightMax        interface{} `json:"rightMax"`
	RightMin        interface{} `json:"rightMin"`
	Threshold1      interface{} `json:"threshold1"`
	Threshold1Color string      `json:"threshold1Color"`
	Threshold2      interface{} `json:"threshold2"`
	Threshold2Color string      `json:"threshold2Color"`
}

type TemplateOptions struct {
	AllFormat string `json:"allFormat"`
	Current   struct {
		Tags  []interface{} `json:"tags"`
		Text  string        `json:"text"`
		Value string        `json:"value"`
	} `json:"current"`
	Datasource  interface{} `json:"datasource"`
	IncludeAll  bool        `json:"includeAll"`
	Multi       bool        `json:"multi"`
	MultiFormat string      `json:"multiFormat"`
	Name        string      `json:"name"`
	Options     []struct {
		Selected bool   `json:"selected"`
		Text     string `json:"text"`
		Value    string `json:"value"`
	} `json:"options"`
	Query   string `json:"query"`
	Refresh bool   `json:"refresh"`
	Type    string `json:"type"`
}

type TimePicker struct {
	RefreshIntervals []string `json:"refresh_intervals"`
	TimeOptions      []string `json:"time_options"`
}

type Time struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Target struct {
	Key    float64 `json:"key"`
	RefID  string  `json:"refId"`
	Target string  `json:"target"`
}

type Panel struct {
	AliasColors     map[string]string `json:"aliasColors"`
	Bars            bool              `json:"bars"`
	Datasource      string            `json:"datasource"`
	Editable        bool              `json:"editable"`
	Error           bool              `json:"error"`
	Fill            int               `json:"fill"`
	Grid            Grid              `json:"grid"`
	ID              int               `json:"id"`
	IsNew           bool              `json:"isNew"`
	Legend          Legend            `json:"legend"`
	Lines           bool              `json:"lines"`
	Linewidth       int               `json:"linewidth"`
	Links           []interface{}     `json:"links"`
	NullPointMode   string            `json:"nullPointMode"`
	Percentage      bool              `json:"percentage"`
	Pointradius     int               `json:"pointradius"`
	Points          bool              `json:"points"`
	Renderer        string            `json:"renderer"`
	SeriesOverrides []interface{}     `json:"seriesOverrides"`
	Span            int               `json:"span"`
	Stack           bool              `json:"stack"`
	SteppedLine     bool              `json:"steppedLine"`
	Targets         []Target          `json:"targets"`
	TimeFrom        interface{}       `json:"timeFrom"`
	TimeShift       interface{}       `json:"timeShift"`
	Title           string            `json:"title"`
	Tooltip         Tooltip           `json:"tooltip"`
	Type            string            `json:"type"`
	XAxis           bool              `json:"x-axis"`
	YAxis           bool              `json:"y-axis"`
	YFormats        []string          `json:"y_formats"`
}

type Row struct {
	Collapse bool    `json:"collapse"`
	Editable bool    `json:"editable"`
	Height   string  `json:"height"`
	Panels   []Panel `json:"panels"`
	Title    string  `json:"title"`
}

type Dashboard struct {
	ID              int           `json:"id"`
	Title           string        `json:"title"`
	OriginalTitle   string        `json:"originalTitle"`
	Tags            []interface{} `json:"tags"`
	Style           string        `json:"style"`
	Timezone        string        `json:"timezone"`
	Editable        bool          `json:"editable"`
	HideControls    bool          `json:"hideControls"`
	SharedCrosshair bool          `json:"sharedCrosshair"`
	Rows            []Row         `json:"rows"`
	Time            Time          `json:"time"`
	TimePicker      TimePicker    `json:"timepicker"`
	Templating      struct {
		Options []TemplateOptions `json:"list"`
	} `json:"templating"`
	Annotations struct {
		List []interface{} `json:"list"`
	} `json:"annotations"`
	SchemaVersion int           `json:"schemaVersion"`
	Version       int           `json:"version"`
	Links         []interface{} `json:"links"`
}
