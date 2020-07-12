package main

// Module is an application module
type Module struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

// Modules represents the list of the application modules
type Modules struct {
	Entities map[string]Module `json:"entities,omitempty"`
	IDs      []string          `json:"ids,omitempty"`
}

var modules Modules = Modules{
	Entities: map[string]Module{
		"9a84c3f2-c84b-4e44-b2b5-3ad9fa1840e4": {
			ID:    "9a84c3f2-c84b-4e44-b2b5-3ad9fa1840e4",
			Title: "Zones",
		},
		"502d3926-63a6-4cd5-a18d-01dfa6c64454": {
			ID:    "502d3926-63a6-4cd5-a18d-01dfa6c64454",
			Title: "Routes",
		},
		"12eb993e-3c0d-4c8a-b517-313b1225363f": {
			ID:    "12eb993e-3c0d-4c8a-b517-313b1225363f",
			Title: "Customers",
		},
		"7f7d075b-fb9f-46dc-b668-e561a753daed": {
			ID:    "7f7d075b-fb9f-46dc-b668-e561a753daed",
			Title: "Suppliers",
		},
		"af9ed27e-36e0-46e0-81e6-e9ab23735b9e": {
			ID:    "af9ed27e-36e0-46e0-81e6-e9ab23735b9e",
			Title: "Projects",
		},
		"a9813694-e774-41c1-8a80-612600551f91": {
			ID:    "a9813694-e774-41c1-8a80-612600551f91",
			Title: "Manifests",
		},
		"626d3414-08ef-4274-a2bb-77633834d93e": {
			ID:    "626d3414-08ef-4274-a2bb-77633834d93e",
			Title: "Quotes",
		},
		"755931c0-94d0-4625-8cc9-5b9e2baaa2f0": {
			ID:    "755931c0-94d0-4625-8cc9-5b9e2baaa2f0",
			Title: "Consignments",
		},
		"43b2919d-8221-4e9d-91a8-97a633cbb48e": {
			ID:    "43b2919d-8221-4e9d-91a8-97a633cbb48e",
			Title: "Parcels",
		},
	},
	IDs: []string{
		"9a84c3f2-c84b-4e44-b2b5-3ad9fa1840e4",
		"502d3926-63a6-4cd5-a18d-01dfa6c64454",
		"12eb993e-3c0d-4c8a-b517-313b1225363f",
		"7f7d075b-fb9f-46dc-b668-e561a753daed",
		"af9ed27e-36e0-46e0-81e6-e9ab23735b9e",
		"a9813694-e774-41c1-8a80-612600551f91",
		"626d3414-08ef-4274-a2bb-77633834d93e",
		"755931c0-94d0-4625-8cc9-5b9e2baaa2f0",
		"43b2919d-8221-4e9d-91a8-97a633cbb48e",
	},
}
