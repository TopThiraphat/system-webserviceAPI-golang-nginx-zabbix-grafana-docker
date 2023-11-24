package domain

// Struct XML

type Place struct {
	// XMLName xml.Name      `xml:"place"`
	Place string        `xml:"name,attr"`
	Value []Value_place `xml:"value"`
}

type Value_place struct {
	// XMLName  xml.Name `xml:"value"`
	Item     string `xml:"item,attr"`
	Itemname string `xml:"itemname,attr"`
	Name     string `xml:"name,attr"`
	Value    string `xml:"value,attr"`
}

type Response struct {
	Name                string
	QueueLength         int
	PredictWaiting      int
	WaitingTime         int
	OpenChannel         int
	AvgProcessingTime   float64
	WaitingTimeOnscreen float64
}
